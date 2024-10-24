package service

import (
	"github.com/golang-class/di-api/model"
	"github.com/golang-class/di-api/repository"
)

type RealFavoriteService struct {
	database repository.Database
}

func (r *RealFavoriteService) Add(favoriteData model.Favorite) ([]model.Favorite, error) {
	//TODO to be implement
	return nil, nil
}

func (r *RealFavoriteService) Delete(id string) ([]model.Favorite, error) {
	//TODO to be implement
	return nil, nil
}

func NewRealFavoriteService(database repository.Database) FavoriteService {
	return &RealFavoriteService{
		database: database,
	}
}
