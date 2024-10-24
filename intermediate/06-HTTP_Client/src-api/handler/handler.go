package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-class/http-client-api/service"
	"net/http"
)

type Handler struct {
	catService      service.CatService
	favoriteService service.FavoriteService
}

func NewHandler(catService service.CatService, favoriteService service.FavoriteService) *Handler {
	return &Handler{
		catService:      catService,
		favoriteService: favoriteService,
	}
}

func (a *Handler) GetCatList(ctx *gin.Context) {
	imageList, err := a.catService.FetchImage(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, imageList)
}

func (a *Handler) GetFavoriteList(ctx *gin.Context) {
}

func (a *Handler) AddFavorite(ctx *gin.Context) {
}

func (a *Handler) DeleteFavorite(ctx *gin.Context) {
}
