package connector

import (
	"context"
	"github.com/golang-class/lab/model"
)

type MovieAPIConnector interface {
	SearchMovie(c context.Context, keyword string) ([]model.Movie, error)
	GetMovieDetail(c context.Context, movieId string) (*model.Movie, error)
}
