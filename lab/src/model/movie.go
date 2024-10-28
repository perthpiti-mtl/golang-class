package model

type Movie struct {
	MovieID string `json:"movie_id"`
	Title   string `json:"title"`
	Year    int    `json:"year"`
	Image   string `json:"image"`
}
