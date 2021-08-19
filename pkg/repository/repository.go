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
	}
}
