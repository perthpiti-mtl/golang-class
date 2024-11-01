package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type MovieAPIConnectorConfig struct {
	URL string `envconfig:"URL" default:"https://distribution-uat.dev.muangthai.co.th/mtl-node-red/golang-course/movie-api"`
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

type ServerConfig struct {
	Port string
}

type Config struct {
	Server            ServerConfig            `envconfig:"SERVER"`
	Database          DatabaseConfig          `envconfig:"DATABASE"`
	MovieAPIConnector MovieAPIConnectorConfig `envconfig:"MOVIE_API_CONNECTOR"`
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Warn("Error loading .env file: %v", err)
	}
	var cfg Config
	err = envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &cfg
}
