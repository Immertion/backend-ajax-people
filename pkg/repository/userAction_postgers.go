package repository

import (
	user "backend_ajax-people"
	"container/list"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserActionPostgres struct {
	db *sqlx.DB
}

func NewUserActionPostgres(db *sqlx.DB) *UserActionPostgres {
	return &UserActionPostgres{db: db}
}

func (r *UserActionPostgres) CreateUser(user user.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (firstname, lastname, password) values ($1, $2, $3) RETURNING id", "users")

	row := r.db.QueryRow(query, user.FirstName, user.LastName, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserActionPostgres) GetUser(id int) (user.User, error) {
	var user user.User

	query := fmt.Sprintf("SELECT firstname, lastname, password  FROM %s WHERE id=$1", "users")
	err := r.db.Get(&user, query, id)
	fmt.Println(user)
	return user, err
}

func (r *UserActionPostgres) GetAllUsers() (*list.List, error) {
	userList := list.New()

	query := fmt.Sprintf("SELECT * FROM %s", "users")
	err := r.db.Get(&userList, query)
	fmt.Println(userList)

	return userList, err

}

func (r *UserActionPostgres) DeleteUser(id int) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", "users")
	_, err := r.db.Exec(query, id)
	return err
}

func (r *UserActionPostgres) UpdateUser(id int, user user.User) (int, error) {

	query := fmt.Sprintf("UPDATE %s SET firstname = $1, lastname = $2 WHERE id=$3;", "users")
	_, err := r.db.Exec(query, user.FirstName, user.LastName, id)

	return id, err
}
