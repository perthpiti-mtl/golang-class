# Containerizing a Go Application

Suppose that this is Go application that you want to containerize. This guide will walk you through the process of
creating a Docker image for your Go application and running it in a container.

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080") // Listen and serve on port 8080
}
```

## Write a Dockerfile for the Application

We’ll create a Dockerfile to containerize the Go application using a multi-stage build to keep the final image
lightweight.

```Dockerfile
# Stage 1: Build the Go binary
FROM golang:1.23.2-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies (optimize caching)
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o gin-api

# Stage 2: Run the application
FROM alpine:3.20

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/gin-api .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./gin-api"]
```

### Explanation:

#### Builder Stage:

- Base Image: golang:1.23.2-alpine (lightweight Go image).
- Working Directory: /app.
- Dependency Management: Copies go.mod and go.sum and runs go mod download to cache dependencies.
- Source Code: Copies the rest of the code and builds the binary.

#### Runtime Stage:

- Base Image: alpine:latest (minimalist image).
- Copies Binary: From the builder stage.
- Exposes Port: 8080.
- Runs Application: Specifies the command to run the binary.

## Build the Docker Image
#### Build the Image
Run the following command in the project root directory:
```bash
docker build -t gin-api:latest .
```
- -t gin-api:latest tags the image.
- The . specifies the current directory as the build context.

#### Verify the Image
List the Docker images to confirm:
```bash
docker images
```

## Run the Docker Container
Run the Docker container using the following command:
```bash
docker run -d -p 8080:8080 --name gin-api-container gin-api:latest
```
- -d runs the container in detached mode.
- -p 8080:8080 maps the host port 8080 to the container port 8080.
- --name gin-api-container names the container.

### Test the Application Inside the Container
```bash
curl http://localhost:8080/ping
```
