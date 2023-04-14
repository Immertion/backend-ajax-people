package main

import (
	user "backend_ajax-people"
	"backend_ajax-people/pkg/handler"
	"backend_ajax-people/pkg/repository"
	"backend_ajax-people/pkg/service"
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "32769",
		Username: "postgres",
		DBName:   "userdb",
		Password: "postgrespw",
		SSLMode:  "disable",
	})
	if err != nil {
		fmt.Printf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(user.Server)
	if err := srv.Run(("8080"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while ruunning http server: %s", err.Error())
	}

	logrus.Print("TodoApp Started")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

}
