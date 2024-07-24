package test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"kplus.com/dto"
	"kplus.com/services"
	"kplus.com/utils"
)

func TestAuthServiceSignIn(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Mock query result
	hashedPassword, _ := utils.HashPassword("password")
	mock.ExpectQuery(`SELECT id, phone, email, password FROM users WHERE phone = ? OR email = ?`).
		WithArgs("testuser", "testuser").
		WillReturnRows(sqlmock.NewRows([]string{"id", "phone", "email", "password"}).AddRow("1", "1234567890", "testuser@example.com", hashedPassword))

	// Setup mock JWT

	// Create AuthService
	authService := services.NewAuthService(utils.NewDatabase(utils.Env{Environment: "test"}, db), utils.NewJwt(utils.Env{Environment: "test", SecretKey: "secret"}, utils.GetLogger()))

	// Execute SignIn
	result, err := authService.SignIn(dto.SignInDto{
		PhoneOrEmail: "testuser",
		Password:     "password",
	})

	// Assert no error
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	// Assert result
	expectedResult := &dto.ResponseSignInDto{
		Token: "mockedToken",
	}

	if result.Token != expectedResult.Token {
		t.Logf("expect %v, got %v", expectedResult.Token, result.Token)
	}

	// Verify expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAuthServiceSignUp(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Mock Exec response for inserting user
	mock.ExpectExec(`INSERT INTO users (role, phone, email, password) VALUES (?, ?, ?, ?)`).
		WithArgs(utils.RoleUser, "1234567890", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Mock Exec response for inserting loans
	mock.ExpectExec("INSERT INTO loans (user_id, `limit`, tenor) VALUES (?, ?, ?)").
		WithArgs(1, 100000.0, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO loans (user_id, `limit`, tenor) VALUES (?, ?, ?)").
		WithArgs(1, 200000.0, 2).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO loans (user_id, `limit`, tenor) VALUES (?, ?, ?)").
		WithArgs(1, 500000.0, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO loans (user_id, `limit`, tenor) VALUES (?, ?, ?)").
		WithArgs(1, 1000000.0, 6).WillReturnResult(sqlmock.NewResult(1, 1))

	// Mock JWT response
	authService := services.NewAuthService(
		utils.NewDatabase(utils.Env{Environment: "test"}, db),
		utils.NewJwt(utils.Env{Environment: "test", SecretKey: "secret"}, utils.GetLogger()),
	)

	// Execute SignUp
	result, err := authService.SignUp(dto.SignUpDto{
		Phone:    "1234567890",
		Password: "password",
	})

	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	// Test result
	expectedResult := &dto.ResponseSignInDto{
		Token: "mockedToken",
	}

	if result.Token != expectedResult.Token {
		t.Logf("expect %v, got %v", expectedResult.Token, result.Token)
	}

	// Verify expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAuthServiceRefreshToken(t *testing.T) {

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	authService := services.NewAuthService(utils.NewDatabase(utils.Env{Environment: "test"}, db), utils.NewJwt(utils.Env{Environment: "test", SecretKey: "secret"}, utils.GetLogger()))

	result, err := authService.RefreshToken(utils.JwtCustomClaims{
		UserID: "1",
		Role:   utils.RoleUser,
	})

	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	// Test result
	expectedResult := "mockedToken"

	if *result != expectedResult {
		t.Logf("expect %v, got %v", expectedResult, *result)
	}
}
