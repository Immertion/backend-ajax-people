package repository

import "github.com/jmoiron/sqlx"

type MailPostgres struct {
	db *sqlx.DB
}

func NewMailPostgres(db *sqlx.DB) *MailPostgres {
	return &MailPostgres{db: db}
}

func (r *MailPostgres) SendMessage(message string, jwt string) (int, string) {
	// Надо по jwt токену определить id в базе данных
	return 0, message
}
