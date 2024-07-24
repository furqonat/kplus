package services

import (
	"context"

	"kplus.com/dto"
	"kplus.com/utils"
)

type InstallmentService struct {
	db utils.Database
}

func (t InstallmentService) GetInstallment(id string) (*dto.InstallmentDto, error) {
	data := dto.InstallmentDto{}
	if err := t.db.QueryRow(`SELECT id, transaction_id, installment, due_date, paid_date, period, status, created_at FROM installments WHERE id = ?`, id).Scan(
		&data.ID, &data.TransactionID, &data.Installment, &data.DueDate, &data.PaidDate, &data.Period, &data.Status, &data.CreatedAt,
	); err != nil {
		return nil, err
	}
	return &data, nil
}

func (t InstallmentService) GetInstallments(transactionID int) ([]dto.InstallmentDto, error) {
	var result []dto.InstallmentDto
	rows, err := t.db.QueryContext(context.Background(), `
		SELECT id, transaction_id, installment, due_date, paid_date, period, status, created_at FROM installments
		WHERE transaction_id = ?`,
		transactionID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data dto.InstallmentDto
		if err := rows.Scan(
			&data.ID,
			&data.TransactionID,
			&data.Installment,
			&data.DueDate,
			&data.PaidDate,
			&data.Period,
			&data.Status,
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

func NewInstallmentService(db utils.Database) InstallmentService {
	return InstallmentService{
		db: db,
	}
}
