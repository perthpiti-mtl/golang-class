package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Port         int      `envconfig:"PORT" default:"8080"`
	Host         string   `envconfig:"HOST" default:"localhost"`
	Debug        bool     `envconfig:"DEBUG" default:"false"`
	AllowedHosts []string `envconfig:"ALLOWED_HOSTS" default:"localhost"`
	DatabaseURL  string   `envconfig:"DATABASE_URL" required:"true"`
}

func main() {
	// Load variables from .env into the environment
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var cfg Config
	err = envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Error processing env variables: %v", err)
	}

	fmt.Printf("Config: %+v\n", cfg)
}
