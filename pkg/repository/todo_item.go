package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/kirillov6/todo-rest-api"
)

type TodoItemSql struct {
	db *sqlx.DB
}

func NewTodoItemSql(db *sqlx.DB) *TodoItemSql {
	return &TodoItemSql{db}
}

func (r *TodoItemSql) Create(listId int, item todo.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, nil
	}

	var id int
	todoItemsQuery := fmt.Sprintf("INSERT INTO %s (title, note) VALUES ($1, $2) RETURNING id", todoItemsTable)
	row := tx.QueryRow(todoItemsQuery, item.Title, item.Note)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	listsItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listsItemsTable)
	if _, err := tx.Exec(listsItemsQuery, listId, id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TodoItemSql) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem
	query := fmt.Sprintf("SELECT ti.* FROM %s ti INNER JOIN %s li ON ti.id = li.item_id INNER JOIN %s ul ON li.list_id = ul.list_id WHERE ul.user_id = $1 AND li.list_id = $2",
		todoItemsTable, listsItemsTable, usersListsTable)
	err := r.db.Select(&items, query, userId, listId)

	return items, err
}

func (r *TodoItemSql) GetById(userId, itemId int) (todo.TodoItem, error) {
	var item todo.TodoItem
	query := fmt.Sprintf("SELECT ti.* FROM %s ti INNER JOIN %s li ON ti.id = li.item_id INNER JOIN %s ul ON li.list_id = ul.list_id WHERE ul.user_id = $1 AND ti.id = $2",
		todoItemsTable, listsItemsTable, usersListsTable)
	err := r.db.Get(&item, query, userId, itemId)

	return item, err
}

func (r *TodoItemSql) DeleteById(userId, itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s ti USING %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_Id AND ul.user_id = $1 AND ti.id = $2",
		todoItemsTable, listsItemsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, itemId)

	return err
}

func (r *TodoItemSql) UpdateById(userId, itemId int, input todo.UpdateItemInput) error {
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

	if input.Note != nil {
		addArg("note", input.Note)
	}

	if input.Done != nil {
		addArg("done", input.Done)
	}

	updateValuesStr := strings.Join(updateValues, ",")

	query := fmt.Sprintf("UPDATE %s ti SET %s FROM %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d",
		todoItemsTable, updateValuesStr, listsItemsTable, usersListsTable, argPos, argPos+1)

	args = append(args, userId, itemId)

	_, err := r.db.Exec(query, args...)

	return err
}
