package repository

import (
	"auth-service/internal/model"
	"auth-service/pkg/database/postgres"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
)

type AuthRepository struct {
	db *postgres.Postgres
}

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
)

func (a *AuthRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	sql, args, err := a.db.Builder.Insert("users").
		Columns("uuid", "name", "email", "password", "createdAt", "updatedAt", "status", "type_user").
		Values(user.UUID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.Status, user.TypeUser).
		ToSql()

	if err != nil {
		return nil, errors.Wrapf(err, "repository/auth/CreateUser: fail to create sql query")
	}
	_, err = a.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return nil, ErrAlreadyExists
			}
		}
		return nil, fmt.Errorf("repository/auth/CreateUser: fail to create user in database: %s", err.Error())
	}
	return user, nil
}

func (a *AuthRepository) GetUser(ctx context.Context, password, email string) (*model.User, error) {
	sql, args, err := a.db.Builder.Select("*").
		From("users").
		Where(squirrel.Eq{"email": email, "password": password}).
		ToSql()

	var user model.User
	err = a.db.Pool.QueryRow(ctx, sql, args...).Scan(&user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Status, &user.TypeUser)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("repository/auth/GetUser: fail to select user in database: %s", err.Error())
	}

	return &user, nil
}

func (a *AuthRepository) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthRepository) UpdateStatusUser(ctx context.Context, uuid string) error {
	sql, args, err := a.db.Builder.Update("users").
		Set("status", true).
		Where(squirrel.Eq{"uuid": uuid}).
		ToSql()
	if err != nil {
		return fmt.Errorf("repository/auth/UpdateStatusUser: fail to create sql query: %s", err.Error())
	}

	_, err = a.db.Pool.Exec(ctx, sql, args)
	if err != nil {
		return fmt.Errorf("repository/auth/UpdateStatusUser: fail to update user in database: %s", err.Error())
	}
	return nil
}

func NewAuthRepository(db *postgres.Postgres) *AuthRepository {
	return &AuthRepository{db: db}
}
