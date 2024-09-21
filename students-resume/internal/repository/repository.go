package repository

import (
	"context"
	"student-resume/internal/model"
	"student-resume/pkg/database/postgres"
)

type Student interface {
	CreateStudent(ctx context.Context, student *model.Student) (*model.Student, error)
	GetStudentByID(ctx context.Context, studentID string) (*model.Student, error)
	DeleteStudentByID(ctx context.Context, studentID string) error
}

type Resume interface {
	CreateResume(ctx context.Context, resume *model.Resume) (*model.Resume, error)
	GetResumeByID(ctx context.Context, resumeID string) (*model.Resume, error)
	GetResumes(ctx context.Context) ([]*model.Resume, error)
	DeleteResumeByID(ctx context.Context, resumeID string) error
	UpdateResumeByID(ctx context.Context, resumeID string, resume *model.Resume) (*model.Resume, error)
}

type Repositories struct {
	Student Student
	Resume  Resume
}

func NewRepositories(db *postgres.Postgres) *Repositories {
	return &Repositories{
		Student: NewStudentRepository(db),
		Resume:  NewResumeRepository(db),
	}
}
