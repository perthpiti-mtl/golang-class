package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-class/lab/handler"
	"github.com/golang-class/lab/router"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	Handler *handler.Handler
}

func (a *App) Run() error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: router.Router(a.Handler),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Could not listen on %d: %v\n", 8080)
		}
	}()

	fmt.Printf("Server starting on %d...\n", 8080)

	// Graceful shutdown
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

func NewApp(handler *handler.Handler) *App {
	return &App{
		Handler: handler,
	}
}
