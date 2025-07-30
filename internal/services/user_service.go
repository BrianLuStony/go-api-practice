package services

import (
	"errors"
	"example/api/internal/entities"
)

type UserService interface {
	Login(email string, password string) (entities.User, string, error)
	Register(user entities.User) (entities.User, error)
	GetUserByID(id string) (entities.User, error)
}

type userServiceImpl struct{}

func NewUserService() UserService {
	return &userServiceImpl{}
}

func (s *userServiceImpl) Login(email string, password string) (entities.User, string, error) {
	if email == "" || password == "" {
		return entities.User{}, "", errors.New("email and password cannot be empty")
	}
	if email == "admin@example.com" && password == "password" {
		user := entities.User{ID: "1", Email: email}
		token := "valid-token"
		return user, token, nil
	}
	return entities.User{}, "", errors.New("invalid credentials")

}

func (s *userServiceImpl) Register(user entities.User) (entities.User, error) { return user, nil }

func (s *userServiceImpl) GetUserByID(id string) (entities.User, error) {
	if id == "" {
		return entities.User{}, errors.New("user ID cannot be empty")
	}
	if id == "1" {
		return entities.User{ID: "1", Email: "test1@gmail.com"}, nil
	}

	return entities.User{}, errors.New("user not found")
}
