package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-class/lab/model"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RealFavoriteRepository struct {
	db *pgxpool.Pool
}

func (r *RealFavoriteRepository) GetFavorite(c context.Context) ([]model.FavoriteMovie, error) {
	var favoriteMovies []model.FavoriteMovie

	rows, err := r.db.Query(c, "SELECT movie_id, title, year, rating, created_at FROM favorite_movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var movie model.FavoriteMovie
		err := rows.Scan(&movie.MovieID, &movie.Title, &movie.Year, &movie.Rating, &movie.CreatedAt)
		if err != nil {
			return nil, err
		}
		favoriteMovies = append(favoriteMovies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return favoriteMovies, nil
}

func (r *RealFavoriteRepository) AddFavorite(c context.Context, movie model.FavoriteMovie) error {
	_, err := r.db.Exec(c, "INSERT INTO favorite_movies (movie_id, title, year, rating) VALUES ($1, $2, $3, $4)", movie.MovieID, movie.Title, movie.Year, movie.Rating)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" { // Unique violation error code
				return fmt.Errorf("movie already in favorite list")
			}
		}
		return err
	}
	return nil
}

func (r *RealFavoriteRepository) DeleteFavorite(c context.Context, movieID string) error {
	result, err := r.db.Exec(c, "DELETE FROM favorite_movies WHERE movie_id = $1", movieID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("movie not found in favorite list")
	}
	return nil
}

func NewRealFavoriteRepository(db *pgxpool.Pool) FavoriteRepository {
	return &RealFavoriteRepository{
		db: db,
	}
}
