package main

import (
	"log"

	"github.com/kirillov6/todo-rest-api"
	"github.com/kirillov6/todo-rest-api/pkg/handler"
	"github.com/kirillov6/todo-rest-api/pkg/repository"
	"github.com/kirillov6/todo-rest-api/pkg/services"
)

func main() {
	repo := repository.NewRepository()
	servs := services.NewServices(repo)
	handler := handler.NewHandler(servs)

	serv := new(todo.Server)
	if err := serv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}
