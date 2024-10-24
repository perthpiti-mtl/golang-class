package service

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-class/http-client-api/connector"
	"github.com/golang-class/http-client-api/model"
)

type RealCatService struct {
	catImageAPIClient connector.CatImageAPIClient
}

func (r *RealCatService) FetchImage(ctx *gin.Context) ([]model.CatImage, error) {
	return r.catImageAPIClient.Search(ctx, 10)
}

func NewRealCatService(catImageAPIClient connector.CatImageAPIClient) CatService {
	return &RealCatService{
		catImageAPIClient: catImageAPIClient,
	}
}
