package main

import (
	user "backend_ajax-people"
	"backend_ajax-people/pkg/handler"
	"backend_ajax-people/pkg/repository"
	"backend_ajax-people/pkg/service"
	"fmt"
	"log"
)

func main() {

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "32768",
		Username: "postgres",
		DBName:   "userdb",
		SSLMode:  "disable",
		Password: "postgrespw",
	})
	if err != nil {
		fmt.Printf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	const port = "8080"

	srv := new(user.Server)
	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while ruunning http server: %s", err.Error())
	}

	log.Printf("Server is running on port %s", port)
}
