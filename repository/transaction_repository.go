package repository

import (
	"assignment-golang-backend/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Get(userId int) []*entity.Transaction
}

type transactionRepoImp struct {
	db *gorm.DB
}

type TransactionRConfig struct {
	DB *gorm.DB
}

func NewTransactionRepository(cfg *TransactionRConfig) TransactionRepository {
	return &transactionRepoImp{
		db: cfg.DB,
	}
}

func (r *transactionRepoImp) Get(userId int) []*entity.Transaction {
	var transactions []*entity.Transaction

	r.db.Where("source_wallet_id = ?", userId).
		Order("created_at DESC").
		Limit(10).
		Find(&transactions)

	return transactions
}
