package router

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-class/api/handler"
	"github.com/golang-class/api/logger"
)

func Router(handler handler.Handler) *gin.Engine {
	router := gin.Default()
	router.Use(logger.LogrusLogger())
	router.GET("/cat", handler.GetCatList)
	router.GET("/favorite", handler.GetFavoriteList)
	router.POST("/favorite", handler.AddFavorite)
	router.DELETE("/favorite/:id", handler.DeleteFavorite)
	return router
}
