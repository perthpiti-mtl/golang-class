package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type ServerConfig struct {
	Port int `envconfig:"PORT" default:"8080"`
}

type MovieAPIConfig struct {
	Url string `envconfig:"URL" default:"https://imdb.iamidiotareyoutoo.com"`
}

type DatabaseConfig struct {
	Host                    string `envconfig:"HOST" required:"true"`
	DatabaseName            string `envconfig:"DATABASE_NAME" default:"database"`
	Port                    uint16 `envconfig:"PORT" default:"5432"`
	Username                string `envconfig:"USERNAME" required:"true"`
	Password                string `envconfig:"PASSWORD" required:"true"`
	MaxConnection           int32  `envconfig:"MAX_CONNECTION" default:"10"`
	MinConnection           int32  `envconfig:"MIN_CONNECTION" default:"2"`
	MinConnectionIdleMinute int32  `envconfig:"MIN_CONNECTION_IDLE_MINUTE" default:"5"`
}

type Config struct {
	Server   ServerConfig   `envconfig:"SERVER"`
	MovieAPI MovieAPIConfig `envconfig:"MOVIE_API"`
	Database DatabaseConfig `envconfig:"DATABASE"`
}

func NewConfig() *Config {
	// Load variables from .env into the environment
	err := godotenv.Load()
	if err != nil {
		log.Warnf("Error loading .env file: %v", err)
	}

	var cfg Config
	err = envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Error processing env variables: %v", err)
	}

	return &cfg
}
