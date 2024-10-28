// app.go
package app

import (
	"context"
	"fmt"
	"github.com/golang-class/api/config"
	"github.com/golang-class/api/handler"
	"github.com/golang-class/api/router"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	handler handler.Handler
	config  config.Config
}

func NewApp(handler *handler.Handler, config *config.Config) *App {
	return &App{
		handler: *handler,
		config:  *config,
	}
}

func (a *App) Run() error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.Server.Port),
		Handler: router.Router(a.handler),
	}

	// Start server in a goroutine
	go func() {
		fmt.Printf("Server starting on %d...\n", a.config.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %d: %v\n", a.config.Server.Port, err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("Server exiting")
	return nil
}
