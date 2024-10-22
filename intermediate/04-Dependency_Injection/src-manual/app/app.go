// app.go
package app

import (
	"example.com/04-di/m/connector"
	"example.com/04-di/m/repository"
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
