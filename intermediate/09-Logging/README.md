# Logging
[Logrus](https://github.com/sirupsen/logrus) is a structured logger for Go that provides extensive features for logging with levels, hooks, and formatting options.

## Installation
```bash
go get github.com/sirupsen/logrus
```
## Basic Usage of Logrus

Import Logrus into your Go file:
```go
import (
    log "github.com/sirupsen/logrus"
)
```
Set up Logrus in your init() function or at the start of your main() function:
```go
func init() {
    // Set the log format to JSON
    log.SetFormatter(&log.JSONFormatter{})

    // Output logs to stdout instead of the default stderr
    log.SetOutput(os.Stdout)

    // Set the log level (debug, info, warn, error, fatal, panic)
    log.SetLevel(log.InfoLevel)
}
```
Basic Logging Examples
```go
log.Info("This is an info message")
log.Warn("This is a warning message")
log.Error("This is an error message")

// Logging with fields (structured logging)
log.WithFields(log.Fields{
    "user_id":  123,
    "function": "main",
}).Info("User logged in")
```
## Integrating Logrus with Gin

Gin uses its own logging middleware by default. To integrate Logrus, you can replace Gin’s logger with a custom middleware.
```go
package main

import (
    "time"

    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"
)

func main() {
    // Initialize Logrus
    log.SetFormatter(&log.JSONFormatter{})
    log.SetOutput(os.Stdout)
    log.SetLevel(log.InfoLevel)

    // Create a new Gin router
    router := gin.New()

    // Use the custom Logrus middleware
    router.Use(LogrusLogger())

    // Define routes
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    // Start the server
    router.Run(":8080")
}

// LogrusLogger is a middleware that logs requests using Logrus
func LogrusLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        startTime := time.Now()

        // Process the request
        c.Next()

        // Calculate the latency
        latency := time.Since(startTime)

        // Get the status code
        statusCode := c.Writer.Status()

        // Log request details
        log.WithFields(log.Fields{
            "status_code":  statusCode,
            "latency_time": latency,
            "client_ip":    c.ClientIP(),
            "method":       c.Request.Method,
            "path":         c.Request.URL.Path,
            "error_message": c.Errors.ByType(gin.ErrorTypePrivate).String(),
        }).Info("Incoming request")
    }
}
```



### Using Logrus in Handlers

Within your Gin handlers, you can use Logrus for application-specific logging.
```go
func GetUser(c *gin.Context) {
    userID := c.Param("id")

    // Fetch user data (pseudo-code)
    user, err := fetchUserFromDB(userID)
    if err != nil {
        log.WithFields(log.Fields{
            "user_id": userID,
            "error":   err.Error(),
        }).Error("Failed to fetch user")
        c.JSON(500, gin.H{"error": "Internal Server Error"})
        return
    }

    log.WithFields(log.Fields{
        "user_id": userID,
    }).Info("User data retrieved")

    c.JSON(200, user)
}
```

### Setting Log Level Based on Environment

You might want different log levels for development and production.
```go
env := os.Getenv("GIN_MODE")
if env == "release" {
    log.SetLevel(log.WarnLevel)
} else {
    log.SetLevel(log.DebugLevel)
}
```