package repository

type Authorization interface {
}

type TodoUser interface {
}

type Repository struct {
	Authorization
	TodoUser
}

func NewRepository() *Repository {
	return &Repository{}
}
