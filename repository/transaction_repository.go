package repository

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/httperror"
	"strconv"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetTransaction(userId uint64, sort string, limit string, search string) ([]*entity.Transaction, error)
	Transaction(transaction *entity.Transaction, userId uint64) (*entity.Transaction, error)
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

func (r *transactionRepoImp) GetTransaction(userId uint64, sort string, limit string, search string) ([]*entity.Transaction, error) {
	var userWallet entity.Wallet
	if err := r.db.Where("user_id = ?", userId).First(&userWallet).Error; err != nil {
		return nil, httperror.ErrWalletNotFound
	}

	var transactions []*entity.Transaction
	transaction := r.db.Where("source_wallet_number = ?", userWallet.WalletNumber).Or("target_wallet_number", userWallet.WalletNumber).Where("description LIKE ? ", "%"+search+"%")
	if sort != " " {
		transaction = transaction.Order(sort)
	}
	if limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			return nil, httperror.ErrInvalidLimit
		}
		transaction = transaction.Limit(limitInt)
	} else {
		transaction = transaction.Limit(10)
	}

	transaction.Find(&transactions)

	return transactions, nil
}

func (r *transactionRepoImp) Transaction(transaction *entity.Transaction, userId uint64) (*entity.Transaction, error) {

	if err := r.db.Transaction(func(tx *gorm.DB) error {
		var wallet *entity.Wallet

		if err := tx.Where("user_id = ?", userId).First(&wallet).Error; err != nil {
			return httperror.ErrInvalidUserWalletId
		}

		// Topup has source of funds id but has no source wallet number
		if transaction.SourceId != nil {
			transaction.TargetWalletNumber = wallet.WalletNumber

			if err := tx.Model(&wallet).Where("wallet_number = ?", transaction.TargetWalletNumber).Update("balance", gorm.Expr("balance + ?", transaction.Amount)).Error; err != nil {
				return err
			}
		} else if transaction.SourceId == nil { //Transfer has no source of funds id but has target wallet number
			transaction.SourceWalletNumber = &wallet.WalletNumber

			var targetWallet entity.Wallet
			if err := tx.Where("wallet_number = ?", transaction.TargetWalletNumber).First(&targetWallet).Error; err != nil {
				return httperror.ErrInvalidTargetWallet
			}
			if targetWallet.WalletNumber == *transaction.SourceWalletNumber {
				return httperror.ErrInvalidTransfer
			}

			if err := tx.Model(&wallet).Where("wallet_number = ?", transaction.SourceWalletNumber).Update("balance", gorm.Expr("balance - ?", transaction.Amount)).Error; err != nil {
				return err
			}

			if err := tx.Model(&wallet).Where("wallet_number = ?", transaction.TargetWalletNumber).Update("balance", gorm.Expr("balance + ?", transaction.Amount)).Error; err != nil {
				return err
			}
		}

		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return transaction, nil
}
