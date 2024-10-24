package main

import (
	"fmt"
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
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Config: %+v\n", cfg)
	// Your application logic here
}
