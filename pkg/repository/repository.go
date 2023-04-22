package repository

import (
	user "backend_ajax-people"
	"container/list"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user user.User) (int, error)
	GetUser(username, password string) (user.User, error)
}

type UserAction interface {
	CreateUser(user user.User) (int, error)
	GetUser(id int) (user.User, error)
	GetAllUsers() (*list.List, error)
	DeleteUser(id int) error
	UpdateUser(id int, user user.User) (int, error)
}

type Mail interface {
	SendMessage(message string, jwt string) (int, string)
}

type Repository struct {
	Authorization
	UserAction
	Mail
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		UserAction:    NewUserActionPostgres(db),
		Mail:          NewMailPostgres(db),
	}
}
