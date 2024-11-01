package main

import "github.com/golang-class/di-wire/di"

func main() {
	appRunner := di.InitializeApp()
	if err := appRunner.Run("https://distribution-uat.dev.muangthai.co.th/mtl-node-red/golang-course/cat-api/list"); err != nil {
		panic(err)
	}
}
