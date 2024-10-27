package repository

import (
	"context"
	"github.com/golang-class/api/model"
)

type FavoriteRepository interface {
	InsertFavorite(ctx context.Context, imageUrl string) (*model.Favorite, error)
	GetFavoriteByID(ctx context.Context, id string) (*model.Favorite, error)
	GetAllFavorites(ctx context.Context) ([]model.Favorite, error)
	DeleteFavoriteByID(ctx context.Context, id string) (*model.Favorite, error)
}
