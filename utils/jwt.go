package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

type JwtCustomClaims struct {
	jwt.RegisteredClaims
	Role      string
	ExpiresAt *int64
	UserID    string
	TokenType TokenType
}

type Jwt struct {
	env    Env
	logger Logger
}

func NewJwt(env Env, logger Logger) *Jwt {
	return &Jwt{
		env:    env,
		logger: logger,
	}
}

func (j *Jwt) GenerateToken(payload *JwtCustomClaims) (string, error) {
	payload.RegisteredClaims.IssuedAt = jwt.NewNumericDate(time.Now())
	payload.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	payload.RegisteredClaims.NotBefore = jwt.NewNumericDate(time.Now())
	if payload.ExpiresAt != nil {
		payload.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Unix(*payload.ExpiresAt, 0))
	}
	s := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return s.SignedString([]byte(j.env.SecretKey))
}

func (j *Jwt) VerifyToken(token string) (JwtCustomClaims, error) {
	jwtString, err := jwt.ParseWithClaims(token, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.env.SecretKey), nil
	})
	if err != nil {
		return JwtCustomClaims{}, err
	}
	claims, ok := jwtString.Claims.(*JwtCustomClaims)
	if !ok || !jwtString.Valid {
		return JwtCustomClaims{}, err
	}

	if claims.RegisteredClaims.ExpiresAt.Before(time.Now()) {
		return JwtCustomClaims{}, errors.New("token expired")
	}

	return *claims, nil
}
