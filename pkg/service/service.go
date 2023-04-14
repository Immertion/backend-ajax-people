package service

import "backend_ajax-people/pkg/repository"

type Authorization interface {
}

type TodoUser interface {
}

type Service struct {
	Authorization
	TodoUser
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
