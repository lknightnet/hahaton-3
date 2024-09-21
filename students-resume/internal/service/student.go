package service

import (
	"context"
	"student-resume/internal/model"
	"student-resume/internal/repository"
)

type StudentService struct {
	DBStudent repository.Student
}

func (s *StudentService) CreateStudent(ctx context.Context, student *model.Student) (*model.Student, error) {
	return s.DBStudent.CreateStudent(ctx, student)
}

func (s *StudentService) GetStudentByID(ctx context.Context, studentID string) (*model.Student, error) {
	return s.DBStudent.GetStudentByID(ctx, studentID)
}

func (s *StudentService) DeleteStudentByID(ctx context.Context, studentID string) error {
	return s.DBStudent.DeleteStudentByID(ctx, studentID)
}

func NewStudentService(DBStudent repository.Student) *StudentService {
	return &StudentService{DBStudent: DBStudent}
}
