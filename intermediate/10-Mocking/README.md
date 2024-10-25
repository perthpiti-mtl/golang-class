# Mocking

## Install GoMock
```bash
go get go.uber.org/mock/gomock
```
Install mockgen

Install the mockgen tool globally:
```bash
go install github.com/golang/mock/mockgen@latest
```
Make sure your $GOPATH/bin is in your system’s PATH so that you can run mockgen from the command line.

## Creating an Interface for Dependency Injection

Suppose you have a service that interacts with a data store, and your Gin handler depends on this service.

### Example Service Interface
```go
// services/user_service.go
package services

type UserService interface {
    GetUser(id string) (User, error)
}

type User struct {
    ID   string
    Name string
}
```
Your Gin handler will depend on this UserService interface rather than a concrete implementation.

### Generating Mocks with mockgen

Use mockgen to generate a mock implementation of the UserService interface.

### Generate Mocks

Run the following command in your terminal:
```bash
mockgen -source=services/user_service.go -destination=mocks/mock_user_service.go -package=mocks
```
This command tells mockgen to:
- Read the source file services/user_service.go.
- Generate mocks in mocks/mock_user_service.go.
- Use the package name mocks.

## Writing Tests with Mocks

Now that you have a mock of UserService, you can write tests for your Gin handlers.

### Gin Handler Example

```go
// handlers/user_handler.go
package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "your_project/services"
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
```
### Test for the Handler
```go
// handlers/user_handler_test.go
package handlers

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "go.uber.org/mock/gomock"
    "github.com/stretchr/testify/assert"
    "your_project/mocks"
    "your_project/services"
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
```
------------------------------------------------------------
## Best Practices
- Interface Segregation: Define small interfaces that match the needs of your handlers to make mocking simpler.
- Dependency Injection: Pass dependencies (like services) into your handlers to facilitate testing.
- Avoid Global State: Global variables make testing harder; prefer passing dependencies explicitly.
- Use Test Files: Keep your test code in _test.go files to ensure they are only built during testing.
- Mock Only External Dependencies: Don’t mock the code you own unless necessary; test real implementations when possible.
