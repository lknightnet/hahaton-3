package http

import (
	"email-service/internal/model"
	"email-service/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

type mailController struct {
	mail service.Mail
}

func newMailService(mail service.Mail) *mailController {
	return &mailController{mail: mail}
}

func (mc *mailController) postSendMail(w http.ResponseWriter, r *http.Request) {
	var mail model.Mail

	err := json.NewDecoder(r.Body).Decode(&mail)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	log.Println(mail)

	err = mc.mail.SendMail(mail)
	if err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
