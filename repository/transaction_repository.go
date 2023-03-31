package repository

import (
	"assignment-golang-backend/entity"
	"errors"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Get(userId int) []*entity.Transaction
	Topup(transaction *entity.Transaction) error
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

func (r *transactionRepoImp) Topup(transaction *entity.Transaction) error {
	var sourceFunds *entity.Sources

	err := r.db.Where("source_id = ?", transaction.SourceId).First(&sourceFunds).Error
	if err != nil {
		return errors.New("source of funds invalid")
	}

	if err := r.db.Transaction(func(tx *gorm.DB) error {
		var wallet *entity.Wallet
		err := tx.Where("wallet_id = ?", transaction.SourceWalletId).First(&wallet).Error
		if err != nil {
			return err
		}

		transactions := transaction

		if err := tx.Model(&wallet).Update("balance", gorm.Expr("balance + ?", transactions.Amount)).Error; err != nil {
			return err
		}

		if err := tx.Create(&transactions).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
