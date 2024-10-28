package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-class/lab/model"
	"github.com/golang-class/lab/service"
	"net/http"
)

type Handler struct {
	MovieService    service.MovieService
	FavoriteService service.FavoriteService
}

// SearchMovie handles the search movie endpoint
func (h *Handler) SearchMovie(c *gin.Context) {
	searchQuery := c.Query("query")
	movie, err := h.MovieService.SearchMovie(c, searchQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

// GetMovieDetail handles the get movie detail endpoint
func (h *Handler) GetMovieDetail(c *gin.Context) {
	id := c.Param("id")
	detail, err := h.MovieService.GetMovieDetail(c, id)
	if err != nil {
		if err.Error() == "movie not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, detail)
}

// GetFavoriteList handles the get favorite movies endpoint
func (h *Handler) GetFavoriteList(c *gin.Context) {
	favorite, err := h.FavoriteService.GetFavorite(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, favorite)
}

// AddFavorite handles the add favorite movie endpoint
func (h *Handler) AddFavorite(c *gin.Context) {
	var req model.FavoriteMovieAddRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.FavoriteService.AddFavorite(c, req.MovieID)
	if err != nil {
		if err.Error() == "movie not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else if err.Error() == "movie already in favorite list" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Movie added to favorite list"})
}

// DeleteFavorite handles the delete favorite movie endpoint
func (h *Handler) DeleteFavorite(c *gin.Context) {
	id := c.Param("id")
	err := h.FavoriteService.RemoveFavorite(c, id)
	if err != nil {
		if err.Error() == "movie not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Movie removed from favorite list"})
}

func NewHandler(movieService service.MovieService, favoriteService service.FavoriteService) *Handler {
	return &Handler{
		MovieService:    movieService,
		FavoriteService: favoriteService,
	}
}
