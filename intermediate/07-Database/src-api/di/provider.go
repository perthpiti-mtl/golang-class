// providers.go
//go:build wireinject
// +build wireinject

package di

import (
	"github.com/golang-class/api/app"
	"github.com/golang-class/api/connector"
	"github.com/golang-class/api/database"
	"github.com/golang-class/api/handler"
	"github.com/golang-class/api/repository"
	"github.com/golang-class/api/service"
	"github.com/google/wire"
)

func InitializeApp() *app.App {
	wire.Build(
		database.NewDatabasePool,
		repository.NewRealFavoriteRepository,
		service.NewRealCatService,
		service.NewRealFavoriteService,
		handler.NewHandler,
		connector.NewRealHTTPClient,
		app.NewApp,
	)
	return nil
}
