package repository

import (
	"context"
	"github.com/golang-class/lab/model"
)

type FavoriteRepository interface {
	GetFavorite(c context.Context) ([]model.FavoriteMovie, error)
	AddFavorite(c context.Context, movie model.FavoriteMovie) error
}
