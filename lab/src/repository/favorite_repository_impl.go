package repository

import (
	"context"
	"github.com/golang-class/lab/model"
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

func NewRealFavoriteRepository(db *pgxpool.Pool) FavoriteRepository {
	return &RealFavoriteRepository{
		db: db,
	}
}
