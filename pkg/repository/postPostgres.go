package repository

import (
	user "backend_ajax-people"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PostPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}
}

func (r *PostPostgres) CreatePost(post user.Post) (int, error) {
	var postId int

	query := fmt.Sprintf("INSERT INTO %s (user_id, text, publication_time) VALUES ($1, $2, $3) RETURNING id", postsTable)

	row := r.db.QueryRow(query, post.UserId, post.Text, post.PublicationTime)
	if err := row.Scan(&postId); err != nil {
		return 0, err
	}

	return postId, nil
}

func (r *PostPostgres) GetPostById(id int) (user.Post, error) {
	var post user.Post

	query := fmt.Sprintf("SELECT user_id, text, is_moderated, publication_time FROM %s WHERE id=$1", postsTable)
	err := r.db.Get(&post, query, id)

	return post, err
}

func (r *PostPostgres) GetAllPosts() ([]user.Post, error) {
	var postsList []user.Post

	query := fmt.Sprintf("SELECT id, user_id, text, is_moderated, publication_time FROM %s", postsTable)
	if err := r.db.Select(&postsList, query); err != nil {
		return nil, err
	}

	return postsList, nil
}

func (r *PostPostgres) UpdatePost(id int, isModerated bool) error {
	query := fmt.Sprintf("UPDATE %s SET is_moderated=$1", postsTable)
	_, err := r.db.Exec(query, isModerated)
	return err
}

func (r *PostPostgres) DeletePost(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postsTable)
	_, err := r.db.Exec(query, id)
	return err
}
