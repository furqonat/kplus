package services

import (
	"time"

	"kplus.com/dto"
	"kplus.com/utils"
)

type InstallmentService struct {
	db utils.Database
}

func (i InstallmentService) PayInstallment(data dto.PayInstallmentDto) error {
	var id int
	var amount float64
	var period int
	if err := i.db.QueryRow(`
		SELECT id, installment, period FROM installments
		WHERE id = ?`, data.ID).Scan(&id, &amount, &period); err != nil {
		return err
	}
	if amount != data.Amount {
		return nil
	}
	if period != data.Period {
		return nil
	}
	if err := i.db.QueryRow(`
		UPDATE installments
		SET status = ?, paid_date = ?
		WHERE id = ?`, "paid", time.Now(), data.ID).Scan(&id); err != nil {
		return err
	}
	return nil
}
func NewInstallmentService(db utils.Database) InstallmentService {
	return InstallmentService{
		db: db,
	}
}
