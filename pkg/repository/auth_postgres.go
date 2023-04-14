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

	query := fmt.Sprintf("INSERT INTO %s (firstname, lastname) values ($1, $2) RETURNING id", "users")

	row := r.db.QueryRow(query, user.FirstName, user.LastName)
	fmt.Println(query)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
