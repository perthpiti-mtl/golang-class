package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func NewDatabasePool() *pgxpool.Pool {
	connStr := "postgres://admin_user:admin_password@localhost:5432/database"

	// Parse the connection string into a configuration
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		panic(fmt.Errorf("unable to parse connection string: %v", err))
	}

	// Optionally configure the pool settings
	config.MaxConns = 10 // maximum number of connections in the pool
	config.MinConns = 2  // minimum number of connections in the pool
	config.MaxConnIdleTime = 5 * time.Minute

	// Create the connection pool
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic(fmt.Errorf("unable to create connection pool: %v", err))
	}

	return pool
}
