package todo

import (
	"errors"
	"reflect"
)

type (
	TodoList struct {
		Id    int    `json:"id" db:"id"`
		Title string `json:"title" db:"title" binding:"required"`
	}

	TodoItem struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
		Note  string `json:"note"`
		Done  bool   `json:"done"`
	}

	ListsItems struct {
		Id     int
		ListId int
		ItemId int
	}

	UpdateListInput struct {
		Title *string `json:"title" db:"title"`
	}
)

func (i *UpdateListInput) Validate() error {
	value := reflect.ValueOf(i).Elem()

	for j := 0; j < value.NumField(); j++ {
		if !value.Field(j).IsNil() {
			return nil
		}
	}

	return errors.New("update structure is empty")
}
