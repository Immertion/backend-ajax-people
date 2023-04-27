package service

import (
	user "backend_ajax-people"
	"backend_ajax-people/pkg/repository"
)

type Authorization interface {
	CreateUser(user user.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, bool, error)
}

type UserAction interface {
	CreateUser(user user.User) (int, error)
	GetUserById(id int) (user.User, error)
	DeleteUser(id int) error
	UpdateUser(id int, user user.UpdateUserInput) error
	GetAllUsers() ([]user.User, error)
}

type Faculty interface {
	GetAll() ([]user.Faculty, error)
}

type Mail interface {
	SendCodeActivation(id int) error
	CheckCodeActivation(id int, rdmKey string) (bool, error)
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
//func NewSendMessageService(mail repository.Mail) Mail {
//	return nil
//}
