package main

func main() {
	appRunner := InitializeApp()
	if err := appRunner.Run("https://api.thecatapi.com/v1/images/search"); err != nil {
		panic(err)
	}
}
