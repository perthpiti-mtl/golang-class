# HTTP Client

Go’s net/http package provides a robust set of tools for making HTTP requests. Whether you need to perform simple GET
requests or handle complex interactions with web APIs, Go’s HTTP client has you covered.

## Basic Usage

**Making a GET Request**

You can make a simple GET request using http.Get, which returns *http.Response and error.

```go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.example.com/data")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
```

**Important:** Always defer resp.Body.Close() to ensure that the response body is properly closed, preventing resource
leaks.

**Making a POST Request**

To make a POST request with form data, you can use http.PostForm.

```go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	data := url.Values{}
	data.Set("username", "example")
	data.Set("password", "password123")

	resp, err := http.PostForm("https://api.example.com/login", data)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
```

**Customizing Requests**

For more control over your HTTP requests, you can create a custom http.Request object.

**Setting Headers**

```go
req, err := http.NewRequest("GET", "https://api.example.com/data", nil)
if err != nil {
log.Fatal(err)
}

req.Header.Set("User-Agent", "MyCustomUserAgent/1.0")
req.Header.Set("Accept", "application/json")
```

**Setting Query Parameters**

```go
q := req.URL.Query()
q.Add("search", "golang")
q.Add("page", "1")
req.URL.RawQuery = q.Encode()
```

**Sending JSON Data**

When sending JSON data in a POST request, you need to set the Content-Type header and provide a bytes.Buffer containing the JSON payload.
```go
import (
    "bytes"
    "encoding/json"
    // other imports
)

type Payload struct {
    Field1 string `json:"field1"`
    Field2 int    `json:"field2"`
}

func main() {
    payload := Payload{
        Field1: "value1",
        Field2: 123,
    }

    jsonData, err := json.Marshal(payload)
    if err != nil {
        log.Fatal(err)
    }

    req, err := http.NewRequest("POST", "https://api.example.com/data", bytes.NewBuffer(jsonData))
    if err != nil {
        log.Fatal(err)
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    // Process response...
}
```

## Advanced Usage

**Custom HTTP Client**

Creating a custom http.Client allows you to configure settings like timeouts and transports.
```go
client := &http.Client{
    Timeout: time.Second * 10, // Set a timeout of 10 seconds
}
```

**Timeouts and Context**

To prevent your request from hanging indefinitely, you can use a context.Context with a timeout.
```go
import (
    "context"
    "time"
)

ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
defer cancel()

req, err := http.NewRequestWithContext(ctx, "GET", "https://api.example.com/data", nil)
if err != nil {
    log.Fatal(err)
}

resp, err := client.Do(req)
// Handle response...
```

**Error Handling and Response Status**

Always check the StatusCode in the response to handle HTTP errors appropriately.
```go
if resp.StatusCode != http.StatusOK {
    log.Fatalf("HTTP error: %s", resp.Status)
}
```
You can also handle specific status codes as needed.

## Complete Example

Here’s a comprehensive example that fetches JSON data from an API and decodes it into a Go struct.
    
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type ApiResponse struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

func main() {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", "https://api.example.com/data", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("HTTP error: %s", resp.Status)
	}

	var result ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %+v\n", result)
}
```
--------------------------------------------
# Using Go’s HTTP Client with Gin Context

When building web applications with the Gin framework in Go, you might need to make HTTP requests to external services within your route handlers. Leveraging Gin’s Context (*gin.Context) can help manage request-scoped data, deadlines, and cancellation signals effectively.


### Basic Usage in a Gin Handler

Here’s how you might typically use the HTTP client inside a Gin handler:

```go
func GetExternalData(c *gin.Context) {
    resp, err := http.Get("https://api.example.com/data")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer resp.Body.Close()

    // Read and process the response body...
}
```
While this works, it doesn’t take advantage of Gin’s context, particularly for managing timeouts and cancellations.


### Using Gin Context with HTTP Client

Creating a Request with Context

Gin’s Context embeds Go’s context.Context, which carries deadlines, cancellation signals, and request-scoped values across API boundaries and between processes.

You can create an HTTP request that uses Gin’s context:
```go
func GetExternalData(c *gin.Context) {
    req, err := http.NewRequestWithContext(c.Request.Context(), "GET", "https://api.example.com/data", nil)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
        return
    }
    defer resp.Body.Close()

    // Read and process the response body...
}
```

Handling Timeouts and Cancellation

By using c.Request.Context(), your HTTP request will automatically be canceled if the client disconnects or if the request times out. This helps prevent resource leaks and unnecessary processing.

If you want to set a specific timeout for the outbound HTTP request, you can create a derived context with a timeout:
```go
import (
    "context"
    "time"
)

func GetExternalData(c *gin.Context) {
    ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
    defer cancel()

    req, err := http.NewRequestWithContext(ctx, "GET", "https://api.example.com/data", nil)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        // Check if the context was canceled due to timeout
        if ctx.Err() == context.DeadlineExceeded {
            c.JSON(http.StatusGatewayTimeout, gin.H{"error": "Request timed out"})
        } else {
            c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
        }
        return
    }
    defer resp.Body.Close()

    // Read and process the response body...
}
```