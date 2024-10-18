package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type User struct {
	Name      string    `json:"name"`
	Age       int       `json:"age,omitempty"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	user := User{
		Name:      "Eve",
		Age:       29,
		Email:     "eve@example.com",
		CreatedAt: time.Now(),
	}

	// Marshal to JSON
	jsonData, err := json.MarshalIndent(user, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Marshaled JSON:")
	fmt.Println(string(jsonData))

	// Unmarshal back to struct
	var newUser User
	err = json.Unmarshal(jsonData, &newUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nUnmarshaled Struct:")
	fmt.Printf("%+v\n", newUser)
}
