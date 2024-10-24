// app.go
package app

import (
	"github.com/golang-class/di-manual/connector"
	"github.com/golang-class/di-manual/repository"
)

type App struct {
	httpClient connector.HTTPClient
	database   repository.Database
}

func NewApp(httpClient connector.HTTPClient, database repository.Database) *App {
	return &App{
		httpClient: httpClient,
		database:   database,
	}
}

func (a *App) Run(url string) error {
	data, err := a.httpClient.Get(url)
	if err != nil {
		return err
	}
	return a.database.Save(data)
}
