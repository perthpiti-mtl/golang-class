package service

import (
	"context"
	"github.com/golang-class/lab/model"
)

type FavoriteService interface {
	GetFavorite(ctx context.Context) ([]model.FavoriteMovie, error)
}
