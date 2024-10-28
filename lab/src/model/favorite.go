package model

import "time"

type FavoriteMovie struct {
	MovieID   string    `json:"movie_id"`
	Title     string    `json:"title"`
	Year      int       `json:"year"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
}

type FavoriteMovieAddRequest struct {
	MovieID string `json:"movie_id"`
}
