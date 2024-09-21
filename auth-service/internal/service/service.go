package service

import (
	"auth-service/internal/model"
	"auth-service/internal/repository"
	"auth-service/pkg/mail"
	"auth-service/pkg/token"
	"context"
)

type Auth interface {
	Signup(ctx context.Context, user *model.User) error
	Login(ctx context.Context, user *model.User) (*model.Token, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.Token, error)
	UpdateStatusUser(ctx context.Context, uuid string) error
}

type Service struct {
	Auth Auth
}

type Dependencies struct {
	Salt    string
	DBUser  repository.Auth
	DBToken repository.Token
	JWT     *token.JWTDependencies
	Email   *mail.EmailMailService
	Topic   string
	Link    string
}

func NewServices(deps *Dependencies) *Service {
	return &Service{
		Auth: NewAuthService(deps.Salt, deps.DBUser, deps.DBToken, deps.JWT, deps.Email, deps.Topic, deps.Link),
	}
}
