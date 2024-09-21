package http

import (
	"email-service/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

func NewMailRoute(r *mux.Router, mail service.Mail) {
	ctrl := newMailService(mail)

	r.HandleFunc("/mail/send", ctrl.postSendMail).Methods(http.MethodPost)
}
