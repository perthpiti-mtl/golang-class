// providers.go
//go:build wireinject
// +build wireinject

package main

import (
	"github.com/golang-class/di-api/app"
	"github.com/golang-class/di-api/connector"
	"github.com/golang-class/di-api/handler"
	"github.com/golang-class/di-api/repository"
	"github.com/golang-class/di-api/service"
	"github.com/google/wire"
)

func InitializeApp() *app.App {
	wire.Build(
		service.NewRealCatService,
		service.NewRealFavoriteService,
		handler.NewHandler,
		connector.NewRealHTTPClient,
		repository.NewRealDatabase,
		app.NewApp,
	)
	return nil
}
