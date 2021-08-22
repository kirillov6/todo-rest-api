package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kirillov6/todo-rest-api"
)

type (
	Authorization interface {
		CreateUser(user todo.User) (int, error)
		GetUser(username, password string) (todo.User, error)
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

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSql(db),
		TodoList:      NewTodoListSql(db),
	}
}
