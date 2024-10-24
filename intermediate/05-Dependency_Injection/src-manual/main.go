package main

import (
	app2 "github.com/golang-class/di-manual/app"
	"github.com/golang-class/di-manual/connector"
	"github.com/golang-class/di-manual/repository"
)

func main() {
	httpClient := connector.NewRealHTTPClient()
	database := repository.NewRealDatabase()
	app := app2.NewApp(httpClient, database)

	if err := app.Run("https://api.thecatapi.com/v1/images/search"); err != nil {
		panic(err)
	}
}
