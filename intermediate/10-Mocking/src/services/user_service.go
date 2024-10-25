// services/user_service.go
package services

type UserService interface {
	GetUser(id string) (User, error)
}

type User struct {
	ID   string
	Name string
}
