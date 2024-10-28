package service

import (
	"context"
	"github.com/golang-class/lab/model"
)

type MovieService interface {
	SearchMovie(ctx context.Context, keyword string) ([]model.Movie, error)
	GetMovieDetail(ctx context.Context, movieId string) (*model.Movie, error)
}
