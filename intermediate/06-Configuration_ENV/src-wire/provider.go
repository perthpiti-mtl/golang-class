// providers.go
//go:build wireinject
// +build wireinject

package main

import (
	"example.com/05-di/m/app"
	"example.com/05-di/m/config"
	"example.com/05-di/m/connector"
	"example.com/05-di/m/repository"
	"github.com/google/wire"
)

func InitializeApp() *app.App {
	wire.Build(
		config.NewConfig,
		connector.NewRealHTTPClient,
		repository.NewRealDatabase,
		app.NewApp,
	)
	return nil
}
