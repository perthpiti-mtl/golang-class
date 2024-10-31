package router

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-class/lab/handler"
)

func Router(handler *handler.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/movies/search", handler.SearchMovie)
	router.GET("/movies/:id", handler.GetMovieDetail)
	router.GET("/favorites", handler.GetFavoriteList)
	return router

}
