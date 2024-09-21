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

type TokenRepository struct {
	db *postgres.Postgres
}

func (t *TokenRepository) AddToken(ctx context.Context, token *model.Token) (*model.Token, error) {
	sql, args, err := t.db.Builder.Insert("tokens").
		Columns("user_id", "token").
		Values(token.UserID, token.Token).
		ToSql()

	if err != nil {
		return nil, errors.Wrapf(err, "repository/token/AddToken: fail to create sql query")
	}
	_, err = t.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return nil, ErrAlreadyExists
			}
		}
		return nil, fmt.Errorf("repository/token/AddToken: fail to create user in database: %s", err.Error())
	}
	return token, nil
}

func (t *TokenRepository) GetToken(ctx context.Context, token string) (*model.Token, error) {
	sql, args, err := t.db.Builder.Select("*").
		From("tokens").
		Where(squirrel.Eq{"token": token}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("repository/token/GetToken: fail to create sql query: %s", err.Error())
	}
	var tokenModel model.Token
	err = t.db.Pool.QueryRow(ctx, sql, args...).Scan(&tokenModel.UserID, &tokenModel.Token)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("repository/token/GetToken: fail to select user in database: %s", err.Error())
	}

	return &tokenModel, nil
}

func NewTokenRepository(db *postgres.Postgres) *TokenRepository {
	return &TokenRepository{db: db}
}
