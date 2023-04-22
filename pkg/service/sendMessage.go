package service

import (
	"backend_ajax-people/pkg/repository"
	"net/smtp"
)

type SendMessageService struct {
	repo repository.Mail
}

func NewSendMessageService(repo repository.Mail) *SendMessageService {
	return &SendMessageService{repo: repo}
}

func (s SendMessageService) SendMessage(message string, jwt string) (int, string, error) {

	// Добавить в конфиг addr, from, to
	// SendMail("localhost:25" - addr, "from@mail.test" - from , "Subject text - subject", "Body text" - body, "to1@tomail.com - to")

	addr := "smtp.localhost:8080"
	host := "smtp.mail.ru"
	from := "ezepiz@mail.ru"
	password := "3141592654Invoker"
	to := "serbinovich.md@students.dvfu.ru"

	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(addr, auth, "ezepiz@mail.ru", []string{to}, []byte(message))

	return 0, message, err
}
