package test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"kplus.com/dto"
	"kplus.com/services"
	"kplus.com/utils"
)

type MockRandomIntGenerator struct{}

func (m MockRandomIntGenerator) RandomInt(min, max int) int {
	return 49
}

func TestTransactionService_GetTransaction(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Expected SQL and results for GetTransaction
	id := "1"
	mock.ExpectQuery(`SELECT id, contract_number, user_id, otr, fee, installment, interest, status, asset_name, created_at FROM transactions WHERE id = ?`).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "contract_number", "user_id", "otr", "fee", "installment", "interest", "status", "asset_name", "created_at"}).
			AddRow(id, "123456", 1, 1000.0, 10.0, 110.0, 5.0, "pending", "asset", time.Now()))

	mock.ExpectQuery(`SELECT id, transaction_id, installment, due_date, paid_date, period, status, created_at FROM installments WHERE transaction_id = ?`).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "transaction_id", "installment", "due_date", "paid_date", "period", "status", "created_at"}).
			AddRow(1, id, 110.0, time.Now(), time.Now(), 1, "pending", time.Now()))

	// Create TransactionService
	transactionService := services.NewTransactionService(utils.NewDatabase(utils.Env{Environment: "test"}, db), MockRandomIntGenerator{})

	// Execute GetTransaction
	transaction, err := transactionService.GetTransaction(id)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	// Test result
	if transaction.ID != id {
		t.Errorf("expect id %s, got %s", id, transaction.ID)
	}

	// Verify expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestTransactionService_GetTransactions(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Expected SQL and results for GetTransactions
	userID := 1
	mock.ExpectQuery(`SELECT id, contract_number, user_id, otr, fee, installment, interest, status, asset_name, created_at FROM transactions WHERE user_id = ?`).
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "contract_number", "user_id", "otr", "fee", "installment", "interest", "status", "asset_name", "created_at"}).
			AddRow(1, "123456", userID, 1000.0, 10.0, 110.0, 5.0, "pending", "asset", time.Now()))

	// Create TransactionService
	transactionService := services.NewTransactionService(utils.NewDatabase(utils.Env{Environment: "test"}, db), MockRandomIntGenerator{})

	// Execute GetTransactions
	transactions, err := transactionService.GetTransactions(userID)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	// Test result
	if len(transactions) != 1 {
		t.Errorf("expect 1 transaction, got %d", len(transactions))
	}

	// Verify expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestTransactionService_CreateTransaction(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Expected SQL and results for CreateTransaction
	data := &dto.CreateTransactionDto{
		Amount:    1000,
		Tenor:     12,
		AssetName: "asset",
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`SELECT fee FROM fees WHERE tenor = ?`).WithArgs(data.Tenor).WillReturnRows(sqlmock.NewRows([]string{"fee"}).AddRow(10.0))
	mock.ExpectQuery(`SELECT interest FROM interests WHERE tenor = ?`).WithArgs(data.Tenor).WillReturnRows(sqlmock.NewRows([]string{"interest"}).AddRow(5.0))

	mock.ExpectExec(`INSERT INTO transactions (contract_number, user_id, otr, fee, installment, interest, status, asset_name) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`).
		WithArgs(49, 1, data.Amount, 10, 110, 5, "pending", data.AssetName).
		WillReturnResult(sqlmock.NewResult(1, 1))

	for i := 1; i <= data.Tenor; i++ {
		mock.ExpectExec(`INSERT INTO installments (transaction_id, installment, due_date, period, status) VALUES (?, ?, ?, ?, ?)`).
			WithArgs(1, 91.67, sqlmock.AnyArg(), i, "pending").
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	mock.ExpectCommit()

	// Create TransactionService with mock RandomInt generator
	transactionService := services.NewTransactionService(utils.NewDatabase(utils.Env{Environment: "test"}, db), MockRandomIntGenerator{})

	// Execute CreateTransaction
	err = transactionService.CreateTransaction(data, 1)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

}
