package connector

import (
	"context"
	"github.com/golang-class/lab/model"
)

type MovieAPIConnector interface {
	ListMovie(c context.Context) ([]model.Movie, error)
	GetMovieDetail(c context.Context, movieId string) (*model.Movie, error)
}
