package main

import "github.com/golang-class/di-wire/di"

func main() {
	appRunner := di.InitializeApp()
	if err := appRunner.Run("https://api.thecatapi.com/v1/images/search"); err != nil {
		panic(err)
	}
}
