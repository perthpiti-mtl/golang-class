package service

import (
	"context"
	"github.com/golang-class/lab/connector"
	"github.com/golang-class/lab/model"
	"github.com/golang-class/lab/repository"
)

type RealFavoriteService struct {
	favoriteRepository repository.FavoriteRepository
	movieConnector     connector.MovieAPIConnector
}

func (r *RealFavoriteService) AddFavorite(c context.Context, movieId string) error {
	movie, err := r.movieConnector.GetMovieDetail(c, movieId)
	if err != nil {
		return err
	}
	favoriteMovie := model.FavoriteMovie{
		MovieID: movie.MovieID,
		Title:   movie.Title,
		Year:    movie.Year,
		Image:   movie.Image,
	}
	err = r.favoriteRepository.AddFavorite(c, favoriteMovie)
	if err != nil {
		return err
	}
	return nil
}

func (r *RealFavoriteService) RemoveFavorite(c context.Context, movieId string) error {
	err := r.favoriteRepository.RemoveFavorite(c, movieId)
	if err != nil {
		return err
	}
	return nil
}

func (r *RealFavoriteService) GetFavorite(c context.Context) ([]model.FavoriteMovie, error) {
	favorite, err := r.favoriteRepository.GetFavorite(c)
	if err != nil {
		return nil, err
	}
	return favorite, nil
}

func NewRealFavoriteService(movieConnector connector.MovieAPIConnector, favoriteRepository repository.FavoriteRepository) FavoriteService {
	return &RealFavoriteService{
		movieConnector:     movieConnector,
		favoriteRepository: favoriteRepository,
	}
}
