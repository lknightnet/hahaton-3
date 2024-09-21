package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"student-resume/internal/model"
	"student-resume/internal/service"
)

type studentController struct {
	studentUC service.Student
}

func newStudentController(studentUC service.Student) *studentController {
	return &studentController{studentUC: studentUC}
}

func NewStudentRoutes(studentUC service.Student, r *mux.Router) {
	studentCtrl := newStudentController(studentUC)

	r.HandleFunc("/student/get/{id:[0-9]+}", studentCtrl.createStudent).Methods(http.MethodPost)
	r.HandleFunc("/student/create", studentCtrl.getStudent).Methods(http.MethodGet)
}

func (a *studentController) createStudent(w http.ResponseWriter, r *http.Request) {
	var body model.Student

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = a.studentUC.CreateStudent(r.Context(), &body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *studentController) getStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	student, err := a.studentUC.GetStudentByID(r.Context(), vars["id"])
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
