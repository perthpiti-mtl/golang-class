package service

import (
	"context"

	"github.com/golang-class/lab/connector"
	"github.com/golang-class/lab/model"
	"github.com/golang-class/lab/repository"
)

type RealFavoriteService struct {
	favoriteRepository repository.FavoriteRepository
	movieAPIConnector  connector.MovieAPIConnector
}

func (r *RealFavoriteService) GetFavorite(c context.Context) ([]model.FavoriteMovie, error) {
	favorite, err := r.favoriteRepository.GetFavorite(c)
	if err != nil {
		return nil, err
	}
	return favorite, nil
}

func (r *RealFavoriteService) AddFavorite(c context.Context, movieId string) error {
	movieDetail, err := r.movieAPIConnector.GetMovieDetail(c, movieId)
	if err != nil {
		return err
	}
	favoriteMovie := model.FavoriteMovie{
		MovieID: movieDetail.MovieID,
		Title:   movieDetail.Title,
		Year:    movieDetail.Year,
		Rating:  movieDetail.Rating,
	}
	err = r.favoriteRepository.AddFavorite(c, favoriteMovie)
	if err != nil {
		return err
	}
	return nil
}

func (r *RealFavoriteService) DeleteFavorite(c context.Context, movieId string) error {
	err := r.favoriteRepository.DeleteFavorite(c, movieId)
	if err != nil {
		return err
	}
	return nil
}

func NewRealFavoriteService(favoriteRepository repository.FavoriteRepository, movieAPIConnector connector.MovieAPIConnector) FavoriteService {
	return &RealFavoriteService{
		favoriteRepository: favoriteRepository,
		movieAPIConnector:  movieAPIConnector,
	}
}
