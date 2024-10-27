package service

import (
	"github.com/golang-class/api/connector"
	"github.com/golang-class/api/model"
)

type RealCatService struct {
	httpClient connector.HTTPClient
}

func (r *RealCatService) FetchImage() ([]model.CatImage, error) {
	//TODO to be implement
	return nil, nil
}

func NewRealCatService(httpClient connector.HTTPClient) CatService {
	return &RealCatService{
		httpClient: httpClient,
	}
}
