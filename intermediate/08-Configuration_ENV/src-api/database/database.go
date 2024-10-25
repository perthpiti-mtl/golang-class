package database

import (
	"context"
	"fmt"
	"github.com/golang-class/api/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func NewDatabasePool(cfg *config.Config) *pgxpool.Pool {
	// Initialize the configuration using a structured approach
	configData, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.DatabaseName, cfg.Database.Username, cfg.Database.Password,
	))
	if err != nil {
		panic(fmt.Errorf("unable to parse config: %v", err))
	}

	// Customize pool settings
	configData.MaxConns = cfg.Database.MaxConnection // Maximum number of connections
	configData.MinConns = cfg.Database.MinConnection // Minimum number of connections
	configData.MaxConnIdleTime = time.Duration(cfg.Database.MinConnectionIdleMinute) * time.Minute

	// Create the connection pool
	pool, err := pgxpool.NewWithConfig(context.Background(), configData)
	if err != nil {
		panic(fmt.Errorf("unable to create connection pool: %v", err))
	}

	return pool
}
