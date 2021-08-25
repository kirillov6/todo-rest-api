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
		Create(userId, listId int, item todo.TodoItem) (int, error)
		GetAll(userId, listId int) ([]todo.TodoItem, error)
		GetById(userId, itemId int) (todo.TodoItem, error)
		DeleteById(userId, itemId int) error
		UpdateById(userId, itemId int, input todo.UpdateItemInput) error
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
		TodoItem:      NewTodoItemService(repo.TodoItem, repo.TodoList),
	}
}
