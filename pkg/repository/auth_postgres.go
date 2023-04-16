package repository

import (
	user "backend_ajax-people"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user user.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (firstname, lastname, password) values ($1, $2, $3) RETURNING id", "users")

	row := r.db.QueryRow(query, user.FirstName, user.LastName, user.Password)
	fmt.Println(user.LastName)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(firstname, password string) (user.User, error) {
	var user user.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE firstname=$1 AND password=$2", "users")
	err := r.db.Get(&user, query, firstname, password)

	return user, err
}
