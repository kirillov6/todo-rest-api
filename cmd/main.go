package main

import (
	"log"

	"github.com/kirillov6/todo-rest-api"
	"github.com/kirillov6/todo-rest-api/pkg/handler"
)

func main() {
	handler := new(handler.Handler)
	serv := new(todo.Server)
	if err := serv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}

}
