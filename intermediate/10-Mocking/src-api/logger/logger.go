package logger

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

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
			"status_code":   statusCode,
			"latency_time":  latency,
			"client_ip":     c.ClientIP(),
			"method":        c.Request.Method,
			"path":          c.Request.URL.Path,
			"error_message": c.Errors.ByType(gin.ErrorTypePrivate).String(),
		}).Info("Incoming request")
	}
}
