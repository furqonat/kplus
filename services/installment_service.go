package services

import (
	"context"
	"database/sql"
	"log"
	"time"

	"kplus.com/dto"
	"kplus.com/utils"
)

type InstallmentService struct {
	db utils.Database
}

func (i InstallmentService) PayInstallment(data dto.PayInstallmentDto) error {
	trx, err := i.db.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return err
	}
	var id int
	var amount float64
	var period int
	var transactionID int

	if err := trx.QueryRow(`
		SELECT id, installment, transaction_id, period FROM installments
		WHERE id = ?`, data.ID).Scan(&id, &amount, &transactionID, &period); err != nil {
		trx.Rollback()
		return err
	}
	status := "unpaid"
	if amount == data.Amount && period == data.Period {
		status = "paid"
	}
	if _, err := trx.Exec(`UPDATE installments
		SET status = ?, paid_date = ?
		WHERE id = ?`, status, time.Now(), data.ID); err != nil {
		trx.Rollback()
		return err
	}
	if _, err := i.db.Exec(`
		UPDATE installments
		SET status = ?, paid_date = ?
		WHERE id = ?`, status, time.Now(), data.ID); err != nil {
		trx.Rollback()
		return err
	}
	var userID int
	var otr float64
	var tenor int
	if err := i.db.QueryRow(`SELECT user_id, otr, tenor FROM transactions WHERE id = ?`, transactionID).Scan(&userID, &otr, &tenor); err != nil {
		trx.Rollback()
		return err
	}
	if _, err := i.db.Exec(`INSERT INTO payments(installment_id, user_id, amount) VALUES (?, ?, ?)`, data.ID, userID, data.Amount); err != nil {
		trx.Rollback()
		return err
	}

	if err := i.increaseUserLoanLimit(userID, amount, tenor); err != nil {
		trx.Rollback()
		return err
	}
	return trx.Commit()
}

func (i InstallmentService) increaseUserLoanLimit(userID int, total float64, tenor int) error {
	var currentLimit float64
	if err := i.db.QueryRow("SELECT `limit` FROM loans WHERE user_id = ? AND tenor = ?", userID, tenor).Scan(&currentLimit); err != nil {
		log.Println(err.Error())
		return err
	}
	_, err := i.db.Exec("UPDATE loans SET `limit` = `limit` + (`limit` * ?) WHERE user_id = ?", (total / currentLimit), userID)
	if err != nil {
		return err
	}
	return nil
}
func NewInstallmentService(db utils.Database) InstallmentService {
	return InstallmentService{
		db: db,
	}
}
