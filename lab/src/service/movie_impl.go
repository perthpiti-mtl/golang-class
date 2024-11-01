package service

import (
	"context"
	"github.com/golang-class/lab/connector"
	"github.com/golang-class/lab/model"
)

type RealMovieService struct {
	MovieAPIConnector connector.MovieAPIConnector
}

func (r *RealMovieService) ListMovie(c context.Context) ([]model.Movie, error) {
	movie, err := r.MovieAPIConnector.ListMovie(c)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (r *RealMovieService) GetMovieDetail(c context.Context, movieId string) (*model.Movie, error) {
	movie, err := r.MovieAPIConnector.GetMovieDetail(c, movieId)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func NewRealMovieService(movieAPIConnector connector.MovieAPIConnector) MovieService {
	return &RealMovieService{
		MovieAPIConnector: movieAPIConnector,
	}
}
