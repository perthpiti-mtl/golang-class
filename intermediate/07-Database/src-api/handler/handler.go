package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-class/api/model"
	"github.com/golang-class/api/service"
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
		return
	}
	ctx.JSON(http.StatusOK, imageList)
}

func (a *Handler) GetFavoriteList(ctx *gin.Context) {
	list, err := a.favoriteService.GetFavoriteList(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, list)
}

func (a *Handler) AddFavorite(ctx *gin.Context) {
	var favoriteRequest model.FavoriteAddRequest
	if err := ctx.ShouldBindJSON(&favoriteRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	favorite, err := a.favoriteService.Add(ctx, favoriteRequest.ImageUrl)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, favorite)
}

func (a *Handler) DeleteFavorite(ctx *gin.Context) {
	id := ctx.Param("id")
	favorite, err := a.favoriteService.Delete(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, favorite)
}
