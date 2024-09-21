package mail

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"
)

type Mail struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Topic   string   `json:"topic"`
	Message string   `json:"message"`
}

type EmailMailService struct {
	ApiEmailSender string
}

func NewMailSender(apiEmailSender string) *EmailMailService {
	return &EmailMailService{ApiEmailSender: apiEmailSender}
}

func (s *EmailMailService) SendMail(m Mail) error {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(m)
	if err != nil {
		log.Println("pkg.mail.SendMail: " + err.Error())
		return errors.Wrap(err, "service.Mail.SendMail")
	}

	req, err := http.NewRequest(http.MethodPost, s.ApiEmailSender, &buf)
	if err != nil {
		log.Println("pkg.mail.SendMail: " + err.Error())
		return errors.Wrap(err, "service.Mail.SendMail")
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("pkg.mail.SendMail: " + err.Error())
		return errors.Wrap(err, "service.Mail.SendMail")
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("pkg.mail.SendMail")
		return errors.Errorf("pkg.mail.SendMail: unknown error, bad status code: %d", resp.StatusCode)
	}
	return nil
}
