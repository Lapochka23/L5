package users

import (
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Service struct {
}

func (s *Service) CreateUser( /* Create user params */ ) (*User, error) {

}

func (s *Service) ListUsers() ([]User, error) {
	// ...
}

func (s *Service) GetUser(userID string) (*User, error) {
	// ...
}

func (s *Service) DeleteUser(userID string) error {
	// ...
}
