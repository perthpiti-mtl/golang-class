package main

import "github.com/golang-class/api/di"

func main() {
	appRunner := di.InitializeApp()
	if err := appRunner.Run(); err != nil {
		panic(err)
	}
}
