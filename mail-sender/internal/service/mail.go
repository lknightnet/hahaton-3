package service

import (
	"email-service/internal/model"
	"github.com/pkg/errors"
	"log"
	"net/smtp"
)

type mailService struct {
	SmtpHost string
	SmtpPort string
	Auth     smtp.Auth
	From     string
}

func (ms *mailService) SendMail(m model.Mail) error {
	msg := m.NewMessage(m.Topic, m.Message)
	log.Println(m)
	err := smtp.SendMail(ms.SmtpHost+":"+ms.SmtpPort, ms.Auth, ms.From, m.To, msg)
	if err != nil {
		log.Println("service.SendMail: " + err.Error())
		return errors.Wrap(err, "service.SendMail")
	}
	return nil
}

func newMailService(SmtpHost string, SmtpPort string, Auth smtp.Auth, From string) *mailService {
	return &mailService{SmtpHost: SmtpHost, SmtpPort: SmtpPort, Auth: Auth, From: From}
}
