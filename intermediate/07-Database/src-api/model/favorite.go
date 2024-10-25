package model

import "time"

type FavoriteAddRequest struct {
	ImageUrl string `json:"image_url" binding:"required"`
}

type Favorite struct {
	ID        int       `json:"id"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}
