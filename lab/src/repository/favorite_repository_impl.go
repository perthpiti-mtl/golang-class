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

	rows, err := r.db.Query(c, "SELECT movie_id, title, year, image, created_at FROM favorite_movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var movie model.FavoriteMovie
		err := rows.Scan(&movie.MovieID, &movie.Title, &movie.Year, &movie.Image, &movie.CreatedAt)
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

func (r *RealFavoriteRepository) AddFavorite(c context.Context, favoriteMovie model.FavoriteMovie) error {
	_, err := r.db.Exec(c, "INSERT INTO favorite_movies (movie_id, title, year, image) VALUES ($1, $2, $3, $4)",
		favoriteMovie.MovieID, favoriteMovie.Title, favoriteMovie.Year, favoriteMovie.Image)
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

func (r *RealFavoriteRepository) RemoveFavorite(c context.Context, movieId string) error {
	result, err := r.db.Exec(c, "DELETE FROM favorite_movies WHERE movie_id = $1", movieId)
	if err != nil {
		return err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("movie not found")
	}
	return nil
}

func NewRealFavoriteRepository(db *pgxpool.Pool) FavoriteRepository {
	return &RealFavoriteRepository{
		db: db,
	}
}
