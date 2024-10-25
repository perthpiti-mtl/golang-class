package handlers

import (
	"github.com/golang-class/mocking/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService services.UserService
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.UserService.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
