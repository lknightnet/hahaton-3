package service

import (
	"auth-service/internal/model"
	"auth-service/internal/repository"
	"auth-service/pkg/mail"
	"auth-service/pkg/token"
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type AuthService struct {
	Salt    string
	DBUser  repository.Auth
	DBToken repository.Token
	JWT     *token.JWTDependencies
	Email   *mail.EmailMailService
	Topic   string
	Link    string
}

func (a *AuthService) Signup(ctx context.Context, user *model.User) error {
	user.UUID = uuid.NewString()
	hash := sha1.New()
	hash.Write([]byte(user.Password))
	user.Password = fmt.Sprintf("%x", hash.Sum([]byte(a.Salt)))
	user.CreatedAt = time.Now()
	user.Status = false
	user.UpdatedAt = time.Now()

	_, err := a.DBUser.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	m := mail.Mail{
		To:      []string{user.Email},
		Topic:   a.Topic,
		Message: "Перейдите по ссылке: " + a.Link + user.UUID,
	}

	err = a.Email.SendMail(m)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthService) Login(ctx context.Context, user *model.User) (*model.Token, error) {
	hash := sha1.New()
	hash.Write([]byte(user.Password))
	password := fmt.Sprintf("%x", hash.Sum([]byte(a.Salt)))

	user, err := a.DBUser.GetUser(ctx, password, user.Email)
	if err != nil {
		return nil, err
	}

	generateTokens, err := a.JWT.GenerateTokens(user)
	if err != nil {
		return nil, err
	}

	addToken, err := a.DBToken.AddToken(ctx, generateTokens)
	if err != nil {
		return nil, err
	}

	return addToken, nil
}

func (a *AuthService) UpdateUser(ctx context.Context, user *model.User) (*model.Token, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthService) UpdateStatusUser(ctx context.Context, uuid string) error {
	return a.DBUser.UpdateStatusUser(ctx, uuid)
}

func NewAuthService(Salt string, DBUser repository.Auth, DBToken repository.Token, JWT *token.JWTDependencies,
	Email *mail.EmailMailService, Topic string, Link string) *AuthService {
	return &AuthService{
		Salt:    Salt,
		DBUser:  DBUser,
		DBToken: DBToken,
		JWT:     JWT,
		Email:   Email,
		Topic:   Topic,
		Link:    Link,
	}
}
