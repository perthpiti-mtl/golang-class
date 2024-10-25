package connector

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-class/api/model"
)

type CatImageAPIClient interface {
	Search(ctx *gin.Context, limit int) ([]model.CatImage, error)
}
