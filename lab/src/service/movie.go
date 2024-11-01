package service

import (
	"context"
	"github.com/golang-class/lab/model"
)

type MovieService interface {
	ListMovie(ctx context.Context) ([]model.Movie, error)
	GetMovieDetail(ctx context.Context, movieId string) (*model.Movie, error)
}
