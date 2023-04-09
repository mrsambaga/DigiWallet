package repository

import (
	"assignment-golang-backend/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Get(userId int) []*entity.Transaction
	Topup(transaction *entity.Transaction) (*entity.Transaction, error)
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

func (r *transactionRepoImp) Topup(transaction *entity.Transaction) (*entity.Transaction, error) {

	if err := r.db.Transaction(func(tx *gorm.DB) error {
		var wallet *entity.Wallet

		topup := transaction

		if err := tx.Model(&wallet).Where("wallet_id = ?", transaction.TargetWalletId).Update("balance", gorm.Expr("balance + ?", topup.Amount)).Error; err != nil {
			return err
		}

		if err := tx.Create(&topup).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return transaction, nil
}
