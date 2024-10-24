// app.go
package app

import (
	"fmt"
	"github.com/golang-class/http-client-api/handler"
	"github.com/golang-class/http-client-api/router"
)

type App struct {
	handler handler.Handler
}

func NewApp(handler *handler.Handler) *App {
	return &App{
		handler: *handler,
	}
}

func (a *App) Run() error {
	fmt.Println("Server starting on port 8080...")
	return router.Router(a.handler).Run(":8080")
}
