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