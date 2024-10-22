// providers.go
//go:build wireinject
// +build wireinject

package main

import (
	"example.com/04-di/m/app"
	"example.com/04-di/m/connector"
	"example.com/04-di/m/repository"
	"github.com/google/wire"
)

func InitializeApp() *app.App {
	wire.Build(
		connector.NewRealHTTPClient,
		repository.NewRealDatabase,
		app.NewApp,
	)
	return nil
}
