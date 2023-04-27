package repository

import (
	user "backend_ajax-people"
	"github.com/jmoiron/sqlx"
)

const (
	facultiesTable = "faculty"
)

type Authorization interface {
	CreateUser(user user.User) (int, error)
	GetUser(username, password string) (user.User, error)
}

type UserAction interface {
	CreateUser(user user.User) (int, error)
	GetUser(id int) (user.User, error)
	GetAllUsers() ([]user.User, error)
	DeleteUser(id int) error
	UpdateUser(id int, user user.UpdateUserInput) error
}

type Mail interface {
	SendCodeActivation(id int, rdmKey string) (string, error)
	CheckCodeActivation(id int, rdmKey string) (bool, error)
}

type Faculty interface {
	GetAll() ([]user.Faculty, error)
}

type Repository struct {
	Authorization
	UserAction
	Mail
	Faculty
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		UserAction:    NewUserActionPostgres(db),
		Mail:          NewMailPostgres(db),
		Faculty:       NewFacultyActionPostgres(db),
	}
}
