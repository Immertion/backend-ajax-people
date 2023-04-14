package main

import (
	user "backend_ajax-people"
	"backend_ajax-people/pkg/handler"
	"backend_ajax-people/pkg/repository"
	"backend_ajax-people/pkg/service"
	"log"
)

func main() {

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(user.Server)
	if err := srv.Run(("8080"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while ruunning http server: %s", err.Error())
	}
}
