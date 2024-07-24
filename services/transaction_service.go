package services

import (
	"context"
	"errors"
	"fmt"
	"math"
	"sync"
	"time"

	"kplus.com/dto"
	"kplus.com/utils"
)

type TransactionService struct {
	db utils.Database
}

func (t TransactionService) GetTransaction(id string) (*dto.TransactionDto, error) {
	data := dto.TransactionDto{}
	if err := t.db.QueryRow(`
		SELECT id, contract_number, user_id, otr, fee, installment, interest, status, asset_name, created_at
	 	FROM transactions WHERE id = ?`, id).Scan(
		&data.ID, &data.ContractNumber, &data.UserID, &data.OTR, &data.Fee, &data.Installment, &data.Interest, &data.Status, &data.AssetName, &data.CreatedAt,
	); err != nil {
		return nil, err
	}
	rows, err := t.db.QueryContext(context.Background(), `
		SELECT id, transaction_id, installment, due_date, paid_date, period, status, created_at FROM installments
		WHERE transaction_id = ?`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var n dto.InstallmentDto
		if err := rows.Scan(
			&n.ID,
			&n.TransactionID,
			&n.Installment,
			&n.DueDate,
			&n.PaidDate,
			&n.Period,
			&n.Status,
			&n.CreatedAt,
		); err != nil {
			return nil, err
		}
		data.Installments = append(data.Installments, n)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &data, nil
}

func (t TransactionService) GetTransactions(userID int) ([]dto.TransactionDto, error) {

	var result []dto.TransactionDto
	rows, err := t.db.QueryContext(context.Background(), `
		SELECT id, contract_number, user_id, otr, fee, installment, interest, status, asset_name, created_at
		FROM transactions
		WHERE user_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data dto.TransactionDto
		if err := rows.Scan(
			&data.ID,
			&data.ContractNumber,
			&data.UserID,
			&data.OTR,
			&data.Fee,
			&data.Installment,
			&data.Interest,
			&data.Status,
			&data.AssetName,
			&data.CreatedAt,
		); err != nil {
			return nil, err
		}
		result = append(result, data)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (t TransactionService) CreateTransaction(data *dto.CreateTransactionDto, userId int) error {
	var wg sync.WaitGroup
	perfomtransaction := func(data *dto.CreateTransactionDto, userId int) error {
		defer wg.Done()
		trx, err := t.db.Begin()
		if err != nil {
			return err
		}

		contractNumber := utils.RandomInt(48, 50)
		fee := 0.0
		interest := 0.0

		if err := t.db.QueryRow(`SELECT fee FROM fees WHERE tenor = ?`, data.Tenor).Scan(&fee); err != nil {
			return err
		}

		if err := t.db.QueryRow(`SELECT interest FROM interests WHERE tenor = ?`, data.Tenor).Scan(&interest); err != nil {
			return err
		}

		if fee == 0.0 || interest == 0.0 {
			return errors.New("tenor not found")
		}

		installment := t.calculateInterset(data.Amount, interest, data.Tenor)
		r, err := trx.ExecContext(context.Background(), `
		INSERT INTO transactions
		(contract_number, user_id, otr, fee, installment, interest, status, asset_name)
		VALUES
		(?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			contractNumber,
			userId,
			data.Amount,
			fee,
			installment+fee,
			interest,
			"pending",
			data.AssetName,
		)
		if err != nil {
			trx.Rollback()
			return err
		}

		id, _ := r.LastInsertId()

		for i := 1; i <= data.Tenor; i++ {
			in := fmt.Sprintf("%.2f", math.Floor(installment))
			_, err := trx.ExecContext(context.Background(), `
			INSERT INTO installments
			(transaction_id, installment, due_date, period, status)
			VALUES
			(?, ?, ?, ?, ?)`,
				id,
				utils.StringToFloat(in),
				time.Now().Add(24*30*time.Hour*time.Duration(i)),
				i,
				"pending",
			)
			if err != nil {
				trx.Rollback()
				return err
			}
		}

		return trx.Commit()
	}
	wg.Add(1)
	go perfomtransaction(data, userId)
	wg.Wait()
	return nil
}

func (t TransactionService) calculateInterset(loan float64, interest float64, tenor int) float64 {
	interestMonth := interest / 100
	totalInterest := loan * interestMonth * float64(tenor)
	totalPay := loan + totalInterest
	installment := totalPay / float64(tenor)
	return installment
}

func NewTransactionService(db utils.Database) TransactionService {
	return TransactionService{
		db: db,
	}
}
