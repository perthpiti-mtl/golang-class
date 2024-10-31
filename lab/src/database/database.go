package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func NewDatabasePool() *pgxpool.Pool {
	// Initialize the configuration using a structured approach
	configData, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s",
		"localhost", 5432, "database", "admin_user", "admin_password",
	))
	if err != nil {
		panic(fmt.Errorf("unable to parse config: %v", err))
	}

	// Customize pool settings
	configData.MaxConns = 10 // Maximum number of connections
	configData.MinConns = 2  // Minimum number of connections
	configData.MaxConnIdleTime = time.Duration(5) * time.Minute

	// Create the connection pool
	pool, err := pgxpool.NewWithConfig(context.Background(), configData)
	if err != nil {
		panic(fmt.Errorf("unable to create connection pool: %v", err))
	}

	return pool
}
