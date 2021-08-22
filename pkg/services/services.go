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
		Create(userId int, list todo.TodoList) (int, error)
		GetAll(userId int) ([]todo.TodoList, error)
		GetById(userId, listId int) (todo.TodoList, error)
		DeleteById(userId, listId int) error
		UpdateById(userId, listId int, input todo.UpdateListInput) error
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
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
	}
}
