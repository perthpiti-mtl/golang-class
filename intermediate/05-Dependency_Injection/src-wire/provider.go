// providers.go
//go:build wireinject
// +build wireinject

package main

import (
	"github.com/golang-class/di-wire/app"
	"github.com/golang-class/di-wire/connector"
	"github.com/golang-class/di-wire/repository"
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
