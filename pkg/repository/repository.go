package repository

import (
	user "backend_ajax-people"
	"github.com/jmoiron/sqlx"
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

type RegisterData interface {
	GetAllFaculties() ([]user.Faculty, error)
	GetAllInterests() ([]user.Interest, error)
	GetAllSchools() ([]user.School, error)
}

type Post interface {
	CreatePost(post user.Post) (int, error)
	GetPostById(id int) (user.Post, error)
	GetAllPosts() ([]user.Post, error)
	UpdatePost(id int, isModerated bool) error
	DeletePost(id int) error
}

type Repository struct {
	Authorization
	UserAction
	Mail
	RegisterData
	Post
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		UserAction:    NewUserActionPostgres(db),
		Mail:          NewMailPostgres(db),
		RegisterData:  NewRegisterDataPostgres(db),
		Post:          NewPostPostgres(db),
	}
}
