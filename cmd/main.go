package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	server := new(todo.Server)

	go func() {
		if err := server.Run(viper.GetString("PORT"), handler.InitRoutes()); err != nil {
			log.Fatalf("error while running http server: %s", err.Error())
		}
	}()

	log.Print("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("error while shutting down server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Fatalf("error while closing db connection: %s", err.Error())
	}

	log.Print("Server shutdown")
}

func initCfg() error {
	viper.SetConfigFile(".env")
	return viper.ReadInConfig()
}
