package database

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-class/lab/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDatabasePool(c *config.Config) *pgxpool.Pool {
	// Initialize the configuration using a structured approach
	configData, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s",
		c.Database.Host, c.Database.Port, c.Database.DatabaseName, c.Database.Username, c.Database.Password,
	))
	if err != nil {
		panic(fmt.Errorf("unable to parse config: %v", err))
	}

	// Customize pool settings
	configData.MaxConns = c.Database.MaxConnection // Maximum number of connections
	configData.MinConns = c.Database.MinConnection // Minimum number of connections
	configData.MaxConnIdleTime = time.Duration(c.Database.MinConnectionIdleMinute) * time.Minute

	// Create the connection pool
	pool, err := pgxpool.NewWithConfig(context.Background(), configData)
	if err != nil {
		panic(fmt.Errorf("unable to create connection pool: %v", err))
	}

	return pool
}
