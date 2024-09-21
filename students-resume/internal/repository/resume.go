package repository

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
	"student-resume/internal/model"
	"student-resume/pkg/database/postgres"
)

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
)

type ResumeRepository struct {
	db *postgres.Postgres
}

func (r *ResumeRepository) CreateResume(ctx context.Context, resume *model.Resume) (*model.Resume, error) {
	sql, args, err := r.db.Builder.Insert("resumes").
		Columns("userID", "name", "user_name", "user_age", "city", "university", "course", "course_number", "status", "description", "document_id").
		Values(resume.UserID, resume.Name, resume.UserName, resume.UserAge, resume.City, resume.University, resume.Course, resume.CourseNumber, resume.Status, resume.Description, resume.DocumentID).
		ToSql()

	if err != nil {
		return nil, errors.Wrapf(err, "repository/resume/CreateResume: fail to create sql query")
	}
	_, err = r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return nil, ErrAlreadyExists
			}
		}
		return nil, fmt.Errorf("repository/resume/CreateResume: fail to create resume in database: %s", err.Error())
	}
	return resume, nil
}

func (r *ResumeRepository) GetResumeByID(ctx context.Context, resumeID string) (*model.Resume, error) {
	sql, args, err := r.db.Builder.Select("*").
		From("resumes").
		Where(squirrel.Eq{"id": resumeID}).
		ToSql()

	var resume model.Resume
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(&resume.ID, &resume.UserID, &resume.Name, &resume.UserName, &resume.UserAge, &resume.City, &resume.University, &resume.Course, &resume.CourseNumber, &resume.Status, &resume.Description, &resume.DocumentID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("repository/resume/CreateResume: fail to select resume in database: %s", err.Error())
	}

	return &resume, nil
}

func (r *ResumeRepository) DeleteResumeByID(ctx context.Context, resumeID string) error {
	sql, args, err := r.db.Builder.Delete("resumes").
		Where(squirrel.Eq{"id": resumeID}).
		ToSql()

	if err != nil {
		return errors.Wrapf(err, "repository/resume/DeleteResumeByID: fail to create sql query")
	}

	_, err = r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return ErrAlreadyExists
			}
		}
		return fmt.Errorf("repository/resume/DeleteResumeByID: fail to create resume in database: %s", err.Error())
	}

	return nil

}

func (r *ResumeRepository) UpdateResumeByID(ctx context.Context, resumeID string, resume *model.Resume) (*model.Resume, error) {
	query, args, err := r.db.Builder.Update("resumes").
		Set("status", resume.Status).
		Where(squirrel.Eq{"id": resumeID}).
		ToSql()
	if err != nil {
		return nil, errors.Wrapf(err, "repository/resume/UpdateResumeByID: fail to create sql query")
	}

	_, err = r.db.Pool.Exec(ctx, query, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return nil, ErrAlreadyExists
			}
		}
		return nil, fmt.Errorf("repository/resume/UpdateResumeByID: fail to create resume in database: %s", err.Error())
	}
	return resume, nil
}

func (r *ResumeRepository) GetResumes(ctx context.Context) ([]*model.Resume, error) {
	sql, args, err := r.db.Builder.Select("*").
		From("resumes").
		ToSql() // Убираем условие Where, чтобы получить все записи

	if err != nil {
		return nil, fmt.Errorf("repository/resume/GetResumes: fail to build query: %s", err.Error())
	}

	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("repository/resume/GetResumes: fail to execute query: %s", err.Error())
	}
	defer rows.Close()

	var resumes []*model.Resume

	for rows.Next() {
		var resume model.Resume
		err = rows.Scan(&resume.ID, &resume.UserID, &resume.Name, &resume.UserName, &resume.UserAge, &resume.City, &resume.University, &resume.Course, &resume.CourseNumber, &resume.Status, &resume.Description, &resume.DocumentID)
		if err != nil {
			return nil, fmt.Errorf("repository/resume/GetResumes: fail to scan row: %s", err.Error())
		}
		resumes = append(resumes, &resume)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("repository/resume/GetResumes: row iteration error: %s", err.Error())
	}

	return resumes, nil
}

func NewResumeRepository(db *postgres.Postgres) *ResumeRepository {
	return &ResumeRepository{db}
}
