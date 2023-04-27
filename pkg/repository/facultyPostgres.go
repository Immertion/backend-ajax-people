package repository

import (
	user "backend_ajax-people"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type FacultyActionPostgres struct {
	db *sqlx.DB
}

func NewFacultyActionPostgres(db *sqlx.DB) *FacultyActionPostgres {
	return &FacultyActionPostgres{db: db}
}

func (r *FacultyActionPostgres) GetAll() ([]user.Faculty, error) {
	var faculties []user.Faculty

	query := fmt.Sprintf("SELECT * FROM %s", facultiesTable)
	if err := r.db.Select(&faculties, query); err != nil {
		return nil, err
	}

	return faculties, nil
}
