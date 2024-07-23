package services

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"kplus.com/dto"
	"kplus.com/utils"
)

type AuthService struct {
	db  utils.Database
	jwt utils.Jwt
}

func NewAuthService(db utils.Database, jwt utils.Jwt) AuthService {
	return AuthService{
		db:  db,
		jwt: jwt,
	}
}

func (a AuthService) SignIn(data dto.SignInDto) (*dto.ResponseSignInDto, error) {
	var id, password, phone string
	var email *string
	if err := a.db.QueryRow(`SELECT id, phone, email, password FROM users WHERE phone = ? OR email = ?`, data.PhoneOrEmail, data.PhoneOrEmail).Scan(
		&id, &phone, &email, &password,
	); errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("invalid phone number or email address")
	}
	if !utils.CheckPasswordHash(data.Password, password) {
		return nil, errors.New("invalid credentials, wrong password")
	}

	return a.generateToken(utils.JwtCustomClaims{
		UserID:    id,
		ExpiresAt: utils.Int64Pointer(time.Now().Add(utils.OneDay).Unix()),
		Role:      utils.RoleUser,
		TokenType: utils.AccessToken,
	})
}

func (a AuthService) RefreshToken(claims utils.JwtCustomClaims) (*string, error) {
	expiresAt := time.Now().Add(utils.OneDay)
	refreshClaims := utils.JwtCustomClaims{
		UserID:    claims.UserID,
		ExpiresAt: utils.Int64Pointer(expiresAt.Unix()),
		TokenType: utils.AccessToken,
		Role:      claims.Role,
	}
	refresh, err := a.jwt.GenerateToken(&refreshClaims)
	if err != nil {
		return nil, err
	}

	return &refresh, nil
}

func (a AuthService) SignUp(user dto.SignUpDto) (*dto.ResponseSignInDto, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	role := utils.RoleUser
	if user.Role != nil {
		role = *user.Role
	}
	if !utils.IsPhoneNumber(user.Phone) {
		return nil, errors.New("invalid phone number")
	}
	if user.Email != nil && !utils.IsGmailAddress(*user.Email) {
		return nil, errors.New("invalid email address")
	}
	r, err := a.db.Exec(`INSERT INTO users (role, phone, email, password) VALUES (?, ?, ?, ?)`, role, user.Phone, user.Email, hashedPassword)
	if err != nil {
		return nil, err
	}

	id, err := r.LastInsertId()
	if err != nil {
		return nil, err
	}
	expiresAt := time.Now().Add(utils.OneDay)
	return a.generateToken(utils.JwtCustomClaims{
		UserID:    fmt.Sprint(id),
		ExpiresAt: utils.Int64Pointer(expiresAt.Unix()),
		Role:      utils.RoleUser,
		TokenType: utils.AccessToken,
	})
}

func (a AuthService) generateToken(claims utils.JwtCustomClaims) (*dto.ResponseSignInDto, error) {

	token, err := a.jwt.GenerateToken(&claims)
	if err != nil {
		return &dto.ResponseSignInDto{}, err
	}
	expiresAt := time.Now().Add(utils.FiveYear)
	refreshClaims := utils.JwtCustomClaims{
		UserID:    claims.UserID,
		ExpiresAt: utils.Int64Pointer(expiresAt.Unix()),
		TokenType: utils.RefreshToken,
		Role:      claims.Role,
	}
	refresh, err := a.jwt.GenerateToken(&refreshClaims)
	if err != nil {
		return &dto.ResponseSignInDto{}, err
	}
	return &dto.ResponseSignInDto{
		Token:        token,
		RefreshToken: refresh,
	}, nil
}
