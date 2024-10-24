package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type ServerConfig struct {
	Port int    `envconfig:"PORT" default:"8080"`
	Host string `envconfig:"HOST" default:"localhost"`
}

type DatabaseConfig struct {
	URL      string `envconfig:"URL" required:"true"`
	Username string `envconfig:"USERNAME" default:"root"`
	Password string `envconfig:"PASSWORD"`
}

type Config struct {
	Server   ServerConfig   `envconfig:"SERVER"`
	Database DatabaseConfig `envconfig:"DATABASE"`
	Debug    bool           `envconfig:"DEBUG" default:"false"`
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
