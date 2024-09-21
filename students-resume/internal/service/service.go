package service

import (
	"context"
	"student-resume/internal/model"
	"student-resume/internal/repository"
)

type Student interface {
	CreateStudent(ctx context.Context, student *model.Student) (*model.Student, error)
	GetStudentByID(ctx context.Context, studentID string) (*model.Student, error)
	DeleteStudentByID(ctx context.Context, studentID string) error
}

type Resume interface {
	CreateResume(ctx context.Context, resume *model.Resume) (*model.Resume, error)
	GetResumeByID(ctx context.Context, resumeID string) (*model.Resume, error)
	DeleteResumeByID(ctx context.Context, resumeID string) error
	UpdateResumeByID(ctx context.Context, resumeID string, resume *model.Resume) (*model.Resume, error)
	GetResumes(ctx context.Context) ([]*model.Resume, error)
}

type Service struct {
	Student Student
	Resume  Resume
}

type Dependencies struct {
	DBStudent repository.Student
	DBResume  repository.Resume
}

func NewServices(deps *Dependencies) *Service {
	return &Service{
		Student: NewStudentService(deps.DBStudent),
		Resume:  NewResumeService(deps.DBResume),
	}
}
