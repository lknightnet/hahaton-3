package service

import (
	"email-service/internal/model"
	"net/smtp"
)

type Mail interface {
	SendMail(m model.Mail) error
}

type Services struct {
	Mail Mail
}

type Dependencies struct {
	SmtpHost string
	SmtpPort string
	Auth     smtp.Auth
	From     string
}

func NewServices(deps *Dependencies) *Services {
	return &Services{Mail: newMailService(deps.SmtpHost, deps.SmtpPort, deps.Auth, deps.From)}
}
