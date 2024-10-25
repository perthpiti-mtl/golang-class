package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Favorite struct {
	ID        int       `json:"id"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	// Connect to the database using a connection pool
	pool, err := connectDB()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	defer pool.Close()

	// Create a new favorite
	newFavorite := Favorite{
		ImageUrl:  "https://example.com/image.png",
		CreatedAt: time.Now(),
	}

	// Insert the favorite into the database
	id, err := insertFavorite(pool, newFavorite)
	if err != nil {
		log.Fatalf("Insert error: %v", err)
	}
	fmt.Printf("Inserted favorite with ID: %d\n", id)

	// Retrieve all favorites
	favorites, err := getFavorites(pool)
	if err != nil {
		log.Fatalf("Retrieval error: %v", err)
	}

	// Print retrieved favorites
	for _, fav := range favorites {
		fmt.Printf("ID: %d, ImageURL: %s, CreatedAt: %s\n", fav.ID, fav.ImageUrl, fav.CreatedAt)
	}
}

func connectDB() (*pgxpool.Pool, error) {
	connStr := "postgres://admin_user:admin_password@localhost:5432/database"

	// Parse the connection string into a configuration
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %v", err)
	}

	// Optionally configure the pool settings
	config.MaxConns = 10 // maximum number of connections in the pool
	config.MinConns = 2  // minimum number of connections in the pool
	config.MaxConnIdleTime = 5 * time.Minute

	// Create the connection pool
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	return pool, nil
}

func insertFavorite(pool *pgxpool.Pool, favorite Favorite) (int, error) {
	var id int
	err := pool.QueryRow(
		context.Background(),
		"INSERT INTO favorites (image_url, created_at) VALUES ($1, $2) RETURNING id",
		favorite.ImageUrl,
		favorite.CreatedAt,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("insert failed: %v", err)
	}
	return id, nil
}

func getFavorites(pool *pgxpool.Pool) ([]Favorite, error) {
	rows, err := pool.Query(context.Background(), "SELECT id, image_url, created_at FROM favorites")
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}
	defer rows.Close()

	var favorites []Favorite
	for rows.Next() {
		var fav Favorite
		err := rows.Scan(&fav.ID, &fav.ImageUrl, &fav.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %v", err)
		}
		favorites = append(favorites, fav)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return favorites, nil
}
