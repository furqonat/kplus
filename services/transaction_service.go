package services

import (
	"kplus.com/dto"
	"kplus.com/utils"
)

type TransactionService struct {
	db utils.Database
}

func (t TransactionService) GetTransaction(id string) (*dto.TransactionDto, error) {
	return nil, nil
}

func NewTransactionService(db utils.Database) TransactionService {
	return TransactionService{
		db: db,
	}
}
