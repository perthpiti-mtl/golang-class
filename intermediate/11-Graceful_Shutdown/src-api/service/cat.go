package service

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-class/api/model"
)

type CatService interface {
	FetchImage(ctx *gin.Context) ([]model.CatImage, error)
}
