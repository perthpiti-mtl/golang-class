package router

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-class/lab/handler"
)

func Router(handler *handler.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/movies", handler.ListMovie)
	router.GET("/movies/:id", handler.GetMovieDetail)
	router.GET("/favorites", handler.GetFavoriteList)
	router.POST("/favorites", handler.AddFavorite)
	return router

}
