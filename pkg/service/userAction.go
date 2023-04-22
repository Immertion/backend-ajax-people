package service

import (
	user "backend_ajax-people"
	"backend_ajax-people/pkg/repository"
	"container/list"
)

type UserActionService struct {
	repo repository.UserAction
}

func NewUserActionService(repo repository.UserAction) *UserActionService {
	return &UserActionService{repo: repo}
}

func (s *UserActionService) CreateUser(user user.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *UserActionService) GetUserById(id int) (user.User, error) {
	return s.repo.GetUser(id)
}

func (s *UserActionService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}

func (s *UserActionService) UpdateUser(id int, user user.User) (int, error) {
	return s.repo.UpdateUser(id, user)
}

func (s *UserActionService) GetAllUsers() (*list.List, error) {
	return s.repo.GetAllUsers()
}
