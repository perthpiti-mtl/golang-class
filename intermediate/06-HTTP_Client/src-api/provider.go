// providers.go
//go:build wireinject
// +build wireinject

package main

import (
	"github.com/golang-class/http-client-api/app"
	"github.com/golang-class/http-client-api/connector"
	"github.com/golang-class/http-client-api/handler"
	"github.com/golang-class/http-client-api/repository"
	"github.com/golang-class/http-client-api/service"
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
