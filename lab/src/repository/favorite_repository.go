package repository

import (
	"context"
	"github.com/golang-class/lab/model"
)

type FavoriteRepository interface {
	GetFavorite(c context.Context) ([]model.FavoriteMovie, error)
}
