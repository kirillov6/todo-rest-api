package services

import (
	"github.com/kirillov6/todo-rest-api"
	"github.com/kirillov6/todo-rest-api/pkg/repository"
)

type (
	Authorization interface {
		CreateUser(user todo.User) (int, error)
		GenerateToken(username, password string) (string, error)
		ParseToken(token string) (int, error)
	}

	TodoList interface {
	}

	TodoItem interface {
	}
)

type Services struct {
	Authorization
	TodoList
	TodoItem
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Authorization: NewAuthService(repo),
	}
}
