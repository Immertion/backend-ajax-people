package service

import (
	user "backend_ajax-people"
	"backend_ajax-people/pkg/repository"
	"time"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
) // тоже на будещее

type AuthService struct {
	repo repository.Authorization
}

func (s *AuthService) CreateUser(user user.User) (int, error) {
	//user.Password = generatePasswordHash(user.Password) - на будущее для хэширования
	return s.repo.CreateUser(user)
}
