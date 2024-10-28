package service

import (
	"context"
	"github.com/golang-class/lab/model"
)

type FavoriteService interface {
	AddFavorite(ctx context.Context, movieId string) error
	RemoveFavorite(ctx context.Context, movieId string) error
	GetFavorite(ctx context.Context) ([]model.FavoriteMovie, error)
}
