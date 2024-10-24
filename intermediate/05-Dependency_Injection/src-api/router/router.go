package router

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-class/di-api/handler"
)

func Router(handler handler.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/cat", handler.GetCatList)
	router.GET("/favorite", handler.GetFavoriteList)
	router.POST("/favorite", handler.AddFavorite)
	router.DELETE("/favorite", handler.DeleteFavorite)
	return router
}
