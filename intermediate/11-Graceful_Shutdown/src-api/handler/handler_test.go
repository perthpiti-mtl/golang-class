package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-class/api/model"
	"github.com/golang-class/api/service/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteFavorite_Success(t *testing.T) {
	// Create a Gin router with the handler
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFavoriteService := mock.NewMockFavoriteService(ctrl)

	// Set up expected calls and return values
	expectedFavorite := &model.Favorite{ID: 1, ImageUrl: "http://example.com/image.jpg"}
	mockFavoriteService.
		EXPECT().
		Delete(gomock.Any(), "1").
		Return(expectedFavorite, nil)

	handler := NewHandler(nil, mockFavoriteService)

	router.DELETE("/favorites/:id", handler.DeleteFavorite)

	// Create a request to send to the above route
	req, _ := http.NewRequest("DELETE", "/favorites/1", nil)
	resp := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "http://example.com/image.jpg")
}

func TestDeleteFavorite_NotFound(t *testing.T) {
	// Create a Gin router with the handler
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFavoriteService := mock.NewMockFavoriteService(ctrl)

	// Set up expected calls and return values
	mockFavoriteService.
		EXPECT().
		Delete(gomock.Any(), "1").
		Return(nil, errors.New("favorite not found"))

	handler := NewHandler(nil, mockFavoriteService)

	router.DELETE("/favorites/:id", handler.DeleteFavorite)

	// Create a request to send to the above route
	req, _ := http.NewRequest("DELETE", "/favorites/1", nil)
	resp := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusNotFound, resp.Code)
	assert.Contains(t, resp.Body.String(), "favorite not found")
}

func TestDeleteFavorite_InternalServerError(t *testing.T) {
	// Create a Gin router with the handler
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFavoriteService := mock.NewMockFavoriteService(ctrl)

	// Set up expected calls and return values
	mockFavoriteService.
		EXPECT().
		Delete(gomock.Any(), "1").
		Return(nil, errors.New("internal server error"))

	handler := NewHandler(nil, mockFavoriteService)

	router.DELETE("/favorites/:id", handler.DeleteFavorite)

	// Create a request to send to the above route
	req, _ := http.NewRequest("DELETE", "/favorites/1", nil)
	resp := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	assert.Contains(t, resp.Body.String(), "internal server error")
}
