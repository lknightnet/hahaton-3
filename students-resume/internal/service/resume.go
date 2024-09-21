package service

import (
	"context"
	"student-resume/internal/model"
	"student-resume/internal/repository"
)

type ResumeService struct {
	DBResume repository.Resume
}

func (r *ResumeService) CreateResume(ctx context.Context, resume *model.Resume) (*model.Resume, error) {
	return r.DBResume.CreateResume(ctx, resume)
}

func (r *ResumeService) GetResumeByID(ctx context.Context, resumeID string) (*model.Resume, error) {
	return r.DBResume.GetResumeByID(ctx, resumeID)
}

func (r *ResumeService) GetResumes(ctx context.Context) ([]*model.Resume, error) {
	return r.DBResume.GetResumes(ctx)
}

func (r *ResumeService) DeleteResumeByID(ctx context.Context, resumeID string) error {
	return r.DBResume.DeleteResumeByID(ctx, resumeID)
}

func (r *ResumeService) UpdateResumeByID(ctx context.Context, resumeID string, resume *model.Resume) (*model.Resume, error) {
	return r.DBResume.UpdateResumeByID(ctx, resumeID, resume)
}

func NewResumeService(DBResume repository.Resume) *ResumeService {
	return &ResumeService{DBResume: DBResume}
}
