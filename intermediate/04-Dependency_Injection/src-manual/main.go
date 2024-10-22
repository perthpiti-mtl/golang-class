package main

import (
	app2 "example.com/04-di/m/app"
	"example.com/04-di/m/connector"
	"example.com/04-di/m/repository"
)

func main() {
	httpClient := connector.NewRealHTTPClient()
	database := repository.NewRealDatabase()
	app := app2.NewApp(httpClient, database)

	if err := app.Run("https://api.thecatapi.com/v1/images/search"); err != nil {
		panic(err)
	}
}
