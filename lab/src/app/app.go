package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-class/lab/config"
	"github.com/golang-class/lab/handler"
	"github.com/golang-class/lab/router"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	Config  *config.Config
	Handler *handler.Handler
}

func (a *App) Run() error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.Config.Server.Port),
		Handler: router.Router(a.Handler),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Could not listen on %d: %v\n", a.Config.Server.Port)
		}
	}()

	fmt.Printf("Server starting on %d...\n", a.Config.Server.Port)

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

func NewApp(config *config.Config, handler *handler.Handler) *App {
	return &App{
		Config:  config,
		Handler: handler,
	}
}
