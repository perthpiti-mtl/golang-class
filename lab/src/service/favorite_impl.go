package service

import (
	"context"
	"github.com/golang-class/lab/model"
	"github.com/golang-class/lab/repository"
)

type RealFavoriteService struct {
	favoriteRepository repository.FavoriteRepository
}

func (r *RealFavoriteService) GetFavorite(c context.Context) ([]model.FavoriteMovie, error) {
	favorite, err := r.favoriteRepository.GetFavorite(c)
	if err != nil {
		return nil, err
	}
	return favorite, nil
}

func NewRealFavoriteService(favoriteRepository repository.FavoriteRepository) FavoriteService {
	return &RealFavoriteService{
		favoriteRepository: favoriteRepository,
	}
}
