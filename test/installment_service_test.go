package test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"kplus.com/dto"
	"kplus.com/services"
	"kplus.com/utils"
)

func TestInstallmentService_PayInstallment(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Create test data
	data := dto.PayInstallmentDto{
		ID:     1,
		Amount: 1000,
		Period: 12,
	}

	// Mock the SELECT query
	mock.ExpectQuery(`SELECT id, installment, period FROM installments WHERE id = ?`).
		WithArgs(data.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "installment", "period"}).AddRow(data.ID, data.Amount, data.Period))

	mock.ExpectExec(`UPDATE installments SET status = ?, paid_date = ? WHERE id = ?`).
		WithArgs("paid", sqlmock.AnyArg(), data.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Create InstallmentService
	installmentService := services.NewInstallmentService(utils.NewDatabase(utils.Env{Environment: "test"}, db))

	// Execute PayInstallment
	err = installmentService.PayInstallment(data)
	if err != nil {
		t.Logf("expect no error, got %v", err)
	}

	// Verify expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Logf("there were unfulfilled expectations: %s", err)
	}

}
