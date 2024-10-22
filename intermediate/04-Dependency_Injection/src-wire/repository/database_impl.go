package repository

import "fmt"

type RealDatabase struct{}

func (db *RealDatabase) Save(data []byte) error {
	// Imagine this saves data to a real database
	fmt.Println("Data saved to the database")
	return nil
}

func NewRealDatabase() Database {
	return &RealDatabase{}
}
