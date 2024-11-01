package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type ServerConfig struct {
	Port int `envconfig:"PORT" default:"8080"`
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

type CatAPIConfig struct {
	Url           string `envconfig:"URL" default:"https://distribution-uat.dev.muangthai.co.th/mtl-node-red/golang-course/cat-api"`
	TimeoutSecond int    `envconfig:"TIMEOUT" default:"10"`
}

type Config struct {
	Server   ServerConfig   `envconfig:"SERVER"`
	Database DatabaseConfig `envconfig:"DATABASE"`
	CatAPI   CatAPIConfig   `envconfig:"CAT_API"`
}

func NewConfig() *Config {
	// Load variables from .env into the environment
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: .env file not found, using environment variables instead")
	}

	var cfg Config
	err = envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Error processing env variables: %v", err)
	}

	return &cfg
}
