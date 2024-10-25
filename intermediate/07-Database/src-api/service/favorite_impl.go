package service

import (
	"context"
	"github.com/golang-class/api/model"
	"github.com/golang-class/api/repository"
)

type RealFavoriteService struct {
	favoriteRepo repository.FavoriteRepository
}

func (r *RealFavoriteService) GetFavoriteList(ctx context.Context) ([]model.Favorite, error) {
	favorites, err := r.favoriteRepo.GetAllFavorites(ctx)
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func (r *RealFavoriteService) Add(ctx context.Context, url string) (*model.Favorite, error) {
	favorite, err := r.favoriteRepo.InsertFavorite(ctx, url)
	if err != nil {
		return nil, err
	}
	return favorite, nil
}

func (r *RealFavoriteService) Delete(ctx context.Context, id string) (*model.Favorite, error) {
	favorite, err := r.favoriteRepo.DeleteFavoriteByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return favorite, nil
}

func NewRealFavoriteService(favoriteRepo repository.FavoriteRepository) FavoriteService {
	return &RealFavoriteService{
		favoriteRepo: favoriteRepo,
	}
}
