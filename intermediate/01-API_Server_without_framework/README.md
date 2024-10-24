# API Server without framework

### 1. Creating Handler Functions

In Go, an HTTP handler is any object that implements the http.Handler interface, which requires a ServeHTTP method, or you can use http.HandlerFunc, which is a type that adapts a function to the http.Handler interface.

```go
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}
```

### 2. Routing Requests

You can route incoming HTTP requests to handler functions using http.HandleFunc or by creating a custom ServeMux.

*Using DefaultServeMux*
```go
func main() {
    http.HandleFunc("/hello", helloHandler)
    http.ListenAndServe(":8080", nil)
}
```
*Using a Custom ServeMux*
```go
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/hello", helloHandler)
    http.ListenAndServe(":8080", mux)
}
```

### 3. Handling HTTP Methods

Within your handler function, you can check the HTTP method and respond accordingly.

```go
func helloHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        fmt.Fprintf(w, "GET Hello, World!")
    case http.MethodPost:
        fmt.Fprintf(w, "POST Hello, World!")
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
```

### 4. Parsing Request Data

*Query Parameters*
```go
func queryHandler(w http.ResponseWriter, r *http.Request) {
    values := r.URL.Query()
    name := values.Get("name")
    fmt.Fprintf(w, "Hello, %s!", name)
}
```

*Form Data*
```go
func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }
    name := r.FormValue("name")
    fmt.Fprintf(w, "Hello, %s!", name)
}
```


*JSON Payloads*
```go
import (
    "encoding/json"
)

type Person struct {
    Name string `json:"name"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
    var person Person
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&person); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    fmt.Fprintf(w, "Hello, %s!", person.Name)
}
```

### 5. Setting Response Headers and Status Codes
```go
func customResponseHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, `{"status":"created"}`)
}
```

### 6. Starting the Server

Start the server by calling http.ListenAndServe:
```go
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/hello", helloHandler)
    // Add more routes here
    fmt.Println("Server starting on port 8080...")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal(err)
    }
}
```

