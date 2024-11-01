package model

import "time"

type FavoriteMovie struct {
	MovieID   string    `json:"movie_id"`
	Title     string    `json:"title"`
	Year      int       `json:"year"`
	Rating    float32   `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}

type AddFavoriteMovieRequest struct {
	MovieID string `json:"movie_id"`
}
