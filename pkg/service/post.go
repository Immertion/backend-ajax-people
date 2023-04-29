package service

import (
	user "backend_ajax-people"
	"backend_ajax-people/pkg/repository"
	"fmt"
	"time"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

const (
	timeFormat = "2006-01-02T15:04:05"
)

func (s *PostService) CreatePost(text string, userId int) (int, error) {
	pblTime := fmt.Sprintln(time.Now().Format(timeFormat))
	newPost := user.Post{UserId: userId, Text: text, PublicationTime: pblTime}
	return s.repo.CreatePost(newPost)
}

func (s *PostService) GetPostById(id int) (user.Post, error) {
	return s.repo.GetPostById(id)
}

func (s *PostService) GetAllPosts() ([]user.Post, error) {
	return s.repo.GetAllPosts()
}

func (s *PostService) UpdatePost(id int, isModerated bool) error {
	return s.repo.UpdatePost(id, isModerated)
}

func (s *PostService) DeletePost(id int) error {
	return s.repo.DeletePost(id)
}
