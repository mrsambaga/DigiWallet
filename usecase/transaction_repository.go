package usecase

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/repository"
)

type TransactionUsecase interface {
	GetUserTransactions(userId int) []*entity.Transaction
}

type transactionUImp struct {
	transactionRepository repository.TransactionRepository
}

type TransactionUConfig struct {
	TransactionRepository repository.TransactionRepository
}

func NewTransactionUsecase(cfg *TransactionUConfig) TransactionUsecase {
	return &transactionUImp{
		transactionRepository: cfg.TransactionRepository,
	}
}

func (u *transactionUImp) GetUserTransactions(userId int) []*entity.Transaction {
	return u.transactionRepository.Get(userId)
}
