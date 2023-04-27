package service

import (
	user "backend_ajax-people"
	"backend_ajax-people/pkg/repository"
	"container/list"
)

type Authorization interface {
	CreateUser(user user.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type UserAction interface {
	CreateUser(user user.User) (int, error)
	GetUserById(id int) (user.User, error)
	DeleteUser(id int) error
	UpdateUser(id int, user user.User) (int, error)
	GetAllUsers() (*list.List, error)
}

type Faculty interface {
	GetAll() ([]user.Faculty, error)
}

type Mail interface {
	SendMessage(message string, jwt string) (int, string, error)
}

type Service struct {
	Authorization
	UserAction
	Mail
	Faculty
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		NewAuthService(repos.Authorization),
		NewUserActionService(repos.UserAction),
		NewSendMessageService(repos.Mail),
		NewFacultyActionService(repos.Faculty),
	}
}

// NewSendMessageService Заглушка
func NewSendMessageService(mail repository.Mail) Mail {
	return nil
}
