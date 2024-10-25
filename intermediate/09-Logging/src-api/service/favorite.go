package service

import (
	"context"
	"github.com/golang-class/api/model"
)

type FavoriteService interface {
	GetFavoriteList(ctx context.Context) ([]model.Favorite, error)
	Add(ctx context.Context, url string) (*model.Favorite, error)
	Delete(ctx context.Context, id string) (*model.Favorite, error)
}
