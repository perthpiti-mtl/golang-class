package main

func main() {
	appRunner := InitializeApp()
	if err := appRunner.Run(); err != nil {
		panic(err)
	}
}
