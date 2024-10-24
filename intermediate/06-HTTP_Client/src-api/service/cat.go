package service

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-class/http-client-api/model"
)

type CatService interface {
	FetchImage(ctx *gin.Context) ([]model.CatImage, error)
}
