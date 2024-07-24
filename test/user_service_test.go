package test

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"kplus.com/dto"
	"kplus.com/services"
	"kplus.com/utils"
)

func TestUserServiceGetUser(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Setup mock row
	row := sqlmock.NewRows([]string{"id", "phone", "email", "role", "ud.id", "full_name", "legal_name", "place_of_birth", "date_of_birth", "salary", "selfie", "selfie_with_national_id", "identity_number", "national_id_image"}).
		AddRow("1", "123456789", "test@example.com", "user", "1", "John Doe", "John Doe", "City", "2000-01-01", 5000, "selfie.jpg", "selfie_with_id.jpg", "1234567890", "id.jpg")

	// Setup mock expectations
	mock.ExpectQuery(`
		SELECT u.id, u.phone, u.email, u.role, ud.id, ud.full_name, ud.legal_name, ud.place_of_birth, ud.date_of_birth, ud.salary, ud.selfie, ud.selfie_with_national_id, ud.identity_number, ud.national_id_image FROM users u
		LEFT JOIN user_details ud ON u.id = ud.user_id
		WHERE u.id = ?`).WithArgs("1").WillReturnRows(row)

	// Create UserService
	userService := services.NewUserService(utils.NewDatabase(utils.Env{Environment: "test"}, db))

	// Execute GetUser
	result, err := userService.GetUser("1")
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	// Test result
	expectedResult := &dto.UserDto{
		ID:    1,
		Phone: "123456789",
		Email: utils.StringPtr("test@example.com"),
		Role:  "user",
		Details: dto.UserDetailsDto{
			ID:                   utils.IntPtr(1),
			FullName:             utils.StringPtr("John Doe"),
			LegalName:            utils.StringPtr("John Doe"),
			PlaceOfBirth:         utils.StringPtr("City"),
			DateOfBirth:          utils.StringPtr("2000-01-01"),
			Salary:               utils.StringPtr("5000"),
			Selfie:               utils.StringPtr("selfie.jpg"),
			SelfieWithNationalID: utils.StringPtr("selfie_with_id.jpg"),
			Nik:                  utils.StringPtr("1234567890"),
			NationalIdImage:      utils.StringPtr("id.jpg"),
		},
	}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("expect %v, got %v", expectedResult, result)
	}

	// Verify expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserServiceCreateUserDetails(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Setup mock expectations
	mock.ExpectExec(`
		INSERT INTO user_details
		(user_id, full_name, legal_name, place_of_birth, date_of_birth, salary, selfie, selfie_with_national_id, identity_number, national_id_image)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
	).WithArgs(
		1, "John Doe", "John Doe", "City", "2000-01-01", "5000", "selfie.jpg", "selfie_with_id.jpg", "1234567890", "id.jpg",
	).WillReturnResult(sqlmock.NewResult(1, 1))

	// Create UserService
	userService := services.NewUserService(utils.NewDatabase(utils.Env{Environment: "test"}, db))

	// Execute CreateUserDetails
	data := dto.UserDetailsDto{
		UserID:               utils.IntPtr(1),
		FullName:             utils.StringPtr("John Doe"),
		LegalName:            utils.StringPtr("John Doe"),
		PlaceOfBirth:         utils.StringPtr("City"),
		DateOfBirth:          utils.StringPtr("2000-01-01"),
		Salary:               utils.StringPtr("5000"),
		Selfie:               utils.StringPtr("selfie.jpg"),
		SelfieWithNationalID: utils.StringPtr("selfie_with_id.jpg"),
		Nik:                  utils.StringPtr("1234567890"),
		NationalIdImage:      utils.StringPtr("id.jpg"),
	}
	id, err := userService.CreateUserDetails(data)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	// Test result
	if id != 1 {
		t.Errorf("expect id 1, got %d", id)
	}

	// Verify expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserServiceUpdateUserDetails(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Setup mock expectations
	mock.ExpectExec(`UPDATE user_details SET full_name = \? WHERE user_id = \?`).
		WithArgs("John Doe", "1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Create UserService
	userService := services.NewUserService(utils.NewDatabase(utils.Env{Environment: "test"}, db))

	// Execute UpdateUserDetails
	data := dto.UserDetailsDto{
		FullName: utils.StringPtr("John Doe"),
	}
	err = userService.UpdateUserDetails(data, "1")
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	// Verify expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
func TestUserServiceGetLoanLimit(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Setup mock rows
	rows := sqlmock.NewRows([]string{"id", "limit", "tenor"}).
		AddRow("1", 1000, 12).
		AddRow("2", 2000, 24)

	// Setup mock expectations
	mock.ExpectQuery(`
		SELECT l.id, l.limit, l.tenor FROM loans l
		WHERE l.user_id = ? LIMIT 5`,
	).WithArgs(1).WillReturnRows(rows)

	// Create UserService
	userService := services.NewUserService(utils.NewDatabase(utils.Env{Environment: "test"}, db))

	// Execute GetLoanLimit
	result, err := userService.GetLoanLimit(1)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	// Test result
	expectedResult := []dto.LoanDto{
		{ID: 1, Limit: 1000, Tenor: 12},
		{ID: 2, Limit: 2000, Tenor: 24},
	}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("expect %v, got %v", expectedResult, result)
	}

	// Verify expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
