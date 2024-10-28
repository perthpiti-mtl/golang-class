package repository

import (
	"context"
	"github.com/golang-class/lab/model"
)

type FavoriteRepository interface {
	GetFavorite(c context.Context) ([]model.FavoriteMovie, error)
	AddFavorite(c context.Context, favoriteMovie model.FavoriteMovie) error
	RemoveFavorite(c context.Context, movieId string) error
}
