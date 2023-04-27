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

	query := fmt.Sprintf("INSERT INTO %s (firstname, lastname, password, mail) values ($1, $2, $3, $4) RETURNING id", userTable)

	row := r.db.QueryRow(query, user.FirstName, user.LastName, user.Password, user.Mail)
	fmt.Println(user.Mail)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(firstname, password string) (user.User, error) {
	var user user.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE firstname=$1 AND password=$2", userTable)
	err := r.db.Get(&user, query, firstname, password)

	return user, err
}
