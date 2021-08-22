package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/kirillov6/todo-rest-api"
)

type TodoListSql struct {
	db *sqlx.DB
}

func NewTodoListSql(db *sqlx.DB) *TodoListSql {
	return &TodoListSql{db}
}

func (r *TodoListSql) Create(userId int, list todo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, nil
	}

	var id int
	todoListsQuery := fmt.Sprintf("INSERT INTO %s (title) VALUES ($1) RETURNING id", todoListsTable)
	row := tx.QueryRow(todoListsQuery, list.Title)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	usersListsQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	if _, err := tx.Exec(usersListsQuery, userId, id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TodoListSql) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	query := fmt.Sprintf("SELECT tl.* FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *TodoListSql) GetById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList
	query := fmt.Sprintf("SELECT tl.* FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2",
		todoListsTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *TodoListSql) DeleteById(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2",
		todoListsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}

func (r *TodoListSql) UpdateById(userId, listId int, input todo.UpdateListInput) error {
	updateValues := make([]string, 0)
	args := make([]interface{}, 0)
	argPos := 1

	addArg := func(valueName string, arg interface{}) {
		updateValues = append(updateValues, fmt.Sprintf("%s = $%d", valueName, argPos))
		args = append(args, arg)
		argPos++
	}

	if input.Title != nil {
		addArg("title", input.Title)
	}

	updateValuesStr := strings.Join(updateValues, ",")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.user_id = $%d AND ul.list_id = $%d",
		todoListsTable, updateValuesStr, usersListsTable, argPos, argPos+1)

	args = append(args, userId, listId)

	_, err := r.db.Exec(query, args...)

	return err
}
