// app.go
package app

import (
	"fmt"
	"github.com/golang-class/api/config"
	"github.com/golang-class/api/handler"
	"github.com/golang-class/api/router"
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
	fmt.Printf("Server starting on %s:%d...\n", a.config.Server.Host, a.config.Server.Port)
	return router.Router(a.handler).Run(fmt.Sprintf("%s:%d", a.config.Server.Host, a.config.Server.Port))
}
