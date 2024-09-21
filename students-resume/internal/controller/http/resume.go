package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"student-resume/internal/model"
	"student-resume/internal/service"
)

type resumeController struct {
	studentUC service.Resume
}

func newResumeController(studentUC service.Resume) *resumeController {
	return &resumeController{studentUC: studentUC}
}

func NewResumeRoutes(studentUC service.Resume, r *mux.Router) {
	studentCtrl := newResumeController(studentUC)

	r.HandleFunc("/student/get/{id:[0-9]+}", studentCtrl.createResume).Methods(http.MethodPost)
	r.HandleFunc("/student/create", studentCtrl.getResume).Methods(http.MethodGet)
}

func (a *resumeController) createResume(w http.ResponseWriter, r *http.Request) {
	var body model.Resume

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = a.studentUC.CreateResume(r.Context(), &body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *resumeController) getResume(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	student, err := a.studentUC.GetResumeByID(r.Context(), vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	marshal, err := json.Marshal(student)
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

func (a *resumeController) getResumes(w http.ResponseWriter, r *http.Request) {
	resumes, err := a.studentUC.GetResumes(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	marshal, err := json.Marshal(resumes)
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
