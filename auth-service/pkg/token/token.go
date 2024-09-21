package token

import (
	"auth-service/internal/model"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTDependencies struct {
	signKey string
}

type AuthClaims struct {
	jwt.RegisteredClaims
	User *model.User
}

func NewJWTDependencies(signKey string) *JWTDependencies {
	return &JWTDependencies{
		signKey: signKey,
	}
}

func (j *JWTDependencies) GenerateTokens(user *model.User) (*model.Token, error) {
	aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &AuthClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		User: user,
	})

	accessToken, err := aToken.SignedString([]byte(j.signKey))
	if err != nil {
		return nil, fmt.Errorf("generateToken/accessToken: fail convert token to string: %s", err.Error())
	}

	token := model.Token{
		Token:  accessToken,
		UserID: user.ID,
	}

	return &token, nil
}
