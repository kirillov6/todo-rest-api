package services

import "github.com/kirillov6/todo-rest-api/pkg/repository"

type (
	Authorization interface {
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
	return &Services{}
}
