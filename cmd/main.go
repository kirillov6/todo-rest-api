package main

import (
	"log"

	"github.com/kirillov6/todo-rest-api"
	"github.com/kirillov6/todo-rest-api/pkg/handler"
	"github.com/kirillov6/todo-rest-api/pkg/repository"
	"github.com/kirillov6/todo-rest-api/pkg/services"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initCfg(); err != nil {
		log.Fatalf("error while init config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		DBName:   viper.GetString("DB_NAME"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
		SSLMode:  viper.GetString("DB_SSLMODE"),
	})

	if err != nil {
		log.Fatalf("error while init db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	servs := services.NewServices(repo)
	handler := handler.NewHandler(servs)

	serv := new(todo.Server)
	if err := serv.Run(viper.GetString("PORT"), handler.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}

func initCfg() error {
	viper.SetConfigFile(".env")
	return viper.ReadInConfig()
}
