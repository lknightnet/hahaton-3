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

type StudentRepository struct {
	db *postgres.Postgres
}

func (s *StudentRepository) CreateStudent(ctx context.Context, student *model.Student) (*model.Student, error) {
	sql, args, err := s.db.Builder.Insert("students").
		Columns("userID", "avatar_url", "email").
		Values(student.UserID, student.AvatarUrl, student.Email).
		ToSql()

	if err != nil {
		return nil, errors.Wrapf(err, "repository/resume/CreateStudent: fail to create sql query")
	}
	_, err = s.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return nil, ErrAlreadyExists
			}
		}
		return nil, fmt.Errorf("repository/resume/CreateStudent: fail to create resume in database: %s", err.Error())
	}
	return student, nil
}

func (s *StudentRepository) GetStudentByID(ctx context.Context, studentID string) (*model.Student, error) {
	sql, args, err := s.db.Builder.Select("*").
		From("students").
		Where(squirrel.Eq{"id": studentID}).
		ToSql()

	var student model.Student
	err = s.db.Pool.QueryRow(ctx, sql, args...).Scan(&student.ID, &student.UserID, &student.AvatarUrl, &student.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("repository/resume/GetStudentByID: fail to select resume in database: %s", err.Error())
	}

	return &student, nil
}

func (s *StudentRepository) DeleteStudentByID(ctx context.Context, studentID string) error {
	sql, args, err := s.db.Builder.Delete("students").
		Where(squirrel.Eq{"id": studentID}).
		ToSql()

	if err != nil {
		return errors.Wrapf(err, "repository/resume/DeleteStudentByID: fail to create sql query")
	}

	_, err = s.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return ErrAlreadyExists
			}
		}
		return fmt.Errorf("repository/resume/DeleteStudentByID: fail to create student in database: %s", err.Error())
	}

	return nil
}

func NewStudentRepository(db *postgres.Postgres) *StudentRepository {
	return &StudentRepository{db}
}
