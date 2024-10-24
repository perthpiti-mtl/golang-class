package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
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

func NewConfig() *Config {
	// Load .env file if it exists
	_ = godotenv.Load()

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
