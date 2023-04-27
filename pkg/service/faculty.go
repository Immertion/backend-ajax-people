package service

import (
	user "backend_ajax-people"
	"backend_ajax-people/pkg/repository"
)

type FacultyActionService struct {
	repo repository.Faculty
}

func NewFacultyActionService(repo repository.Faculty) *FacultyActionService {
	return &FacultyActionService{repo: repo}
}

func (s *FacultyActionService) GetAll() ([]user.Faculty, error) {
	return s.repo.GetAll()
}
