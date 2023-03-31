package usecase

import (
	"assignment-golang-backend/dto"
	"assignment-golang-backend/entity"
	"assignment-golang-backend/repository"
)

type TransactionUsecase interface {
	GetUserTransactions(userId int) []*entity.Transaction
	Topup(newTopUp *dto.TopupRequestDTO, userId int) error
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

func (u *transactionUImp) Topup(newTopUp *dto.TopupRequestDTO, userId int) error {
	topUp := &entity.Transaction{
		SourceWalletId: uint64(userId),
		TargetWalletId: uint64(userId),
		Amount:         newTopUp.Amount,
		Description:    newTopUp.Description,
		SourceId:       newTopUp.SourceId,
	}

	return u.transactionRepository.Topup(topUp)
}
