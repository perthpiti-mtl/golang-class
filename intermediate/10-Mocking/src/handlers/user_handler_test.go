// handlers/user_handler_test.go
package handlers

import (
	"errors"
	"github.com/golang-class/mocking/mocks"
	"github.com/golang-class/mocking/services"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUser_Success(t *testing.T) {
	// Create a Gin router with the handler
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)

	// Set up expected calls and return values
	expectedUser := services.User{ID: "123", Name: "John Doe"}
	mockUserService.
		EXPECT().
		GetUser("123").
		Return(expectedUser, nil)

	userHandler := &UserHandler{
		UserService: mockUserService,
	}

	router.GET("/users/:id", userHandler.GetUser)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/users/123", nil)
	resp := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "John Doe")
}

func TestGetUser_NotFound(t *testing.T) {
	// Similar setup as above
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)

	// Set up expected calls and return values
	mockUserService.
		EXPECT().
		GetUser("999").
		Return(services.User{}, errors.New("user not found"))

	userHandler := &UserHandler{
		UserService: mockUserService,
	}

	router.GET("/users/:id", userHandler.GetUser)

	req, _ := http.NewRequest("GET", "/users/999", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
	assert.Contains(t, resp.Body.String(), "User not found")
}
