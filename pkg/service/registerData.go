package service

import (
	user "backend_ajax-people"
	"backend_ajax-people/pkg/repository"
)

type RegisterDataService struct {
	repo repository.RegisterData
}

func NewRegisterDataService(repo repository.RegisterData) *RegisterDataService {
	return &RegisterDataService{repo: repo}
}

func (s *RegisterDataService) GetAllFaculties() ([]user.Faculty, error) {
	return s.repo.GetAllFaculties()
}

func (s *RegisterDataService) GetAllInterests() ([]user.Interest, error) {
	return s.repo.GetAllInterests()
}

var statusesSlice = []user.StatusUser{
	{0, "student"},
	{1, "teacher"},
	{2, "graduated"},
	{3, "other"},
}

func (s *RegisterDataService) GetAllStatuses() []user.StatusUser {
	return statusesSlice
}

var edLevelsSlice = []user.EducationLevel{
	{0, "bachelor"},
	{1, "master"},
	{2, "specialist"},
	{3, "other"},
}

func (s *RegisterDataService) GetAllEdLevels() []user.EducationLevel {
	return edLevelsSlice
}

func (s *RegisterDataService) GetAllSchools() ([]user.School, error) {
	return s.repo.GetAllSchools()
}
