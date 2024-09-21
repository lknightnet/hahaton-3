package http

import (
	"auth-service/internal/model"
	"auth-service/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type authController struct {
	authUC service.Auth
}

func newAuthController(authUC service.Auth) *authController {
	return &authController{authUC: authUC}
}

func NewAuthRoutes(authUC service.Auth, r *mux.Router) {
	authCtrl := newAuthController(authUC)

	r.HandleFunc("/auth/signup", authCtrl.signup).Methods(http.MethodPost)
	r.HandleFunc("/auth/login", authCtrl.login).Methods(http.MethodPost)
	r.HandleFunc("/auth/update-user/{uuid}", authCtrl.updateUser).Methods(http.MethodGet)
}

func (a *authController) signup(w http.ResponseWriter, r *http.Request) {
	var body model.User

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.authUC.Signup(r.Context(), &body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *authController) login(w http.ResponseWriter, r *http.Request) {
	var body model.User

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := a.authUC.Login(r.Context(), &body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	marshal, err := json.Marshal(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write(marshal)
	if err != nil {
		log.Println(err)
		return
	}
}

func (a *authController) updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars["uuid"])
	err := a.authUC.UpdateStatusUser(r.Context(), vars["uuid"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
