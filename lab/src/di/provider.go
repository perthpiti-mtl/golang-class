// providers.go
//go:build wireinject
// +build wireinject

package di

import (
	"github.com/golang-class/lab/app"
	"github.com/golang-class/lab/config"
	"github.com/golang-class/lab/connector"
	"github.com/golang-class/lab/database"
	"github.com/golang-class/lab/handler"
	"github.com/golang-class/lab/repository"
	"github.com/golang-class/lab/service"
	"github.com/google/wire"
)

func InitializeApp() *app.App {
	wire.Build(
		config.NewConfig,
		database.NewDatabasePool,
		repository.NewRealFavoriteRepository,
		connector.NewRealMovieAPI,
		handler.NewHandler,
		service.NewRealMovieService,
		service.NewRealFavoriteService,
		app.NewApp,
	)
	return nil
}
