package repository

import (
	"auth-service/internal/model"
	"auth-service/pkg/database/postgres"
	"context"
)

type Auth interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUser(ctx context.Context, password, email string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	UpdateStatusUser(ctx context.Context, uuid string) error
}

type Token interface {
	AddToken(ctx context.Context, token *model.Token) (*model.Token, error)
	GetToken(ctx context.Context, token string) (*model.Token, error)
}

type Repositories struct {
	Auth  Auth
	Token Token
}

func NewRepositories(db *postgres.Postgres) *Repositories {
	return &Repositories{
		Auth:  NewAuthRepository(db),
		Token: NewTokenRepository(db),
	}
}
