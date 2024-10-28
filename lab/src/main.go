package main

import "github.com/golang-class/lab/di"

func main() {
	app := di.InitializeApp()
	err := app.Run()
	if err != nil {
		return
	}
}
