package service

import "github.com/golang-class/http-client-api/model"

type FavoriteService interface {
	Add(favoriteData model.Favorite) ([]model.Favorite, error)
	Delete(id string) ([]model.Favorite, error)
}
