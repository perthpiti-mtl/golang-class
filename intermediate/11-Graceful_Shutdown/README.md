# Graceful Shutdown

Implementing a graceful shutdown in a Go application using the Gin framework involves properly handling system signals to allow the server to finish processing ongoing requests before shutting down. This ensures a smooth termination without abruptly dropping active connections or losing data.

## 1. Import Necessary Packages

You’ll need to import the following packages:

```go
import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/gin-gonic/gin"
)
```
- context: To manage cancellation and timeouts.
- os/signal: To capture OS signals like SIGINT or SIGTERM.
- syscall: Provides constants for system call interrupts.
- time: To set timeouts for the shutdown process.
- github.com/gin-gonic/gin: The Gin web framework.

## 2. Set Up Your Gin Router and HTTP Server

Initialize your Gin router and wrap it in an http.Server:

```go
router := gin.Default()

// Define your routes here
router.GET("/", func(c *gin.Context) {
    c.String(http.StatusOK, "Hello, World!")
})

// Create the HTTP server
srv := &http.Server{
    Addr:    ":8080",
    Handler: router,
}
```

## 3. Start the Server in a Goroutine

Run the server in a separate goroutine so that it doesn’t block the graceful shutdown handling code:
```go
go func() {
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatalf("listen: %s\n", err)
    }
}()
log.Println("Server started on port 8080")
```

## 4. Capture OS Signals for Shutdown

Set up a channel to listen for interrupt signals:
```go
quit := make(chan os.Signal, 1)
// SIGINT (Ctrl+C) and SIGTERM (Docker and Kubernetes stop signals)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
<-quit
log.Println("Shutdown signal received")
```

## 5. Implement the Graceful Shutdown

When a shutdown signal is received, create a context with a timeout and call srv.Shutdown:
```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

if err := srv.Shutdown(ctx); err != nil {
    log.Fatalf("Server forced to shutdown: %v", err)
}

log.Println("Server exiting")
```
- The Shutdown method will stop accepting new connections and wait for the existing ones to finish within the timeout period.
- If the server doesn’t shut down within the specified timeout, it will forcefully exit.

## 6. Full Example Code

Putting it all together:

```go
package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, World!")
    })

    srv := &http.Server{
        Addr:    ":8080",
        Handler: router,
    }

    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %s\n", err)
        }
    }()
    log.Println("Server started on port 8080")

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Shutdown signal received")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("Server forced to shutdown: %v", err)
    }

    log.Println("Server exiting")
}
```

## 7. Testing the Graceful Shutdown

To test the graceful shutdown:

1. Run your application:
```bash
go run main.go
```

2. In another terminal, simulate a long-running request:
```bash
curl localhost:8080 -v
```

3. While the request is in progress, send a SIGINT signal by pressing Ctrl+C in the terminal where the server is running.

4. Observe the logs to ensure that the server waits for the request to complete before shutting down.