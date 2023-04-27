package repository

import (
	user "backend_ajax-people"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type UserActionPostgres struct {
	db *sqlx.DB
}

func NewUserActionPostgres(db *sqlx.DB) *UserActionPostgres {
	return &UserActionPostgres{db: db}
}

func (r *UserActionPostgres) CreateUser(user user.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (firstname, lastname, password, mail) values ($1, $2, $3, $4) RETURNING id", userTable)

	row := r.db.QueryRow(query, user.FirstName, user.LastName, user.Password, user.Mail)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserActionPostgres) GetUser(id int) (user.User, error) {
	var user user.User

	query := fmt.Sprintf("SELECT firstname, lastname, mail, is_admin FROM %s WHERE id=$1", userTable)
	err := r.db.Get(&user, query, id)

	return user, err
}

func (r *UserActionPostgres) GetAllUsers() ([]user.User, error) {
	var userList []user.User

	query := fmt.Sprintf("SELECT firstname, lastname FROM %s", userTable)
	if err := r.db.Select(&userList, query); err != nil {
		return nil, err
	}

	return userList, nil
}

func (r *UserActionPostgres) DeleteUser(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", userTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *UserActionPostgres) UpdateUser(id int, user user.UpdateUserInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if user.FirstName != nil {
		setValues = append(setValues, fmt.Sprintf("firstname=$%d", argId))
		args = append(args, *user.FirstName)
		argId++
	}

	if user.LastName != nil {
		setValues = append(setValues, fmt.Sprintf("lastname=$%d", argId))
		args = append(args, *user.LastName)
		argId++
	}

	if user.StatusUser != nil {
		setValues = append(setValues, fmt.Sprintf("status_user=$%d", argId))
		args = append(args, *user.StatusUser)
		argId++
	}

	if user.AdmissionYear != nil {
		setValues = append(setValues, fmt.Sprintf("admission_year=$%d", argId))
		args = append(args, *user.AdmissionYear)
		argId++
	}

	if user.Age != nil {
		setValues = append(setValues, fmt.Sprintf("age=$%d", argId))
		args = append(args, *user.Age)
		argId++
	}

	if user.EducationLevel != nil {
		setValues = append(setValues, fmt.Sprintf("education_level=$%d", argId))
		args = append(args, *user.EducationLevel)
		argId++
	}

	if user.GraduationYear != nil {
		setValues = append(setValues, fmt.Sprintf("graduation_year=$%d", argId))
		args = append(args, *user.GraduationYear)
		argId++
	}

	if user.StudyProgramId != nil {
		setValues = append(setValues, fmt.Sprintf("study_program_id=$%d", argId))
		args = append(args, *user.StudyProgramId)
		argId++
	}

	if user.SchoolId != nil {
		setValues = append(setValues, fmt.Sprintf("school_id=$%d", argId))
		args = append(args, *user.SchoolId)
		argId++
	}

	if user.AvatarPath != nil {
		setValues = append(setValues, fmt.Sprintf("avatar_path=$%d", argId))
		args = append(args, *user.AvatarPath)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d;", userTable, setQuery, argId)

	args = append(args, id)
	_, err := r.db.Exec(query, args...)

	return err
}
