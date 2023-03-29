package repository

import (
	"assignment-golang-backend/dto"
	"assignment-golang-backend/entity"

	"gorm.io/gorm"
)

type WalletRepository interface {
	GetDetail(userId int) (*dto.WalletDetail, error)
}

type walletRepositoryImp struct {
	db *gorm.DB
}

type WalletRConfig struct {
	DB *gorm.DB
}

func NewWalletRepository(cfg *WalletRConfig) WalletRepository {
	return &walletRepositoryImp{
		db: cfg.DB,
	}
}

func (r *walletRepositoryImp) GetDetail(userId int) (*dto.WalletDetail, error) {
	var wallet *entity.Wallet
	err := r.db.Where("user_id = ?", userId).Preload("User").Find(&wallet).Error
	if err != nil {
		return nil, err
	}

	dtoResponse := &dto.WalletDetail{
		Id:           wallet.WalletId,
		UserId:       wallet.UserId,
		UserName:     wallet.User.Name,
		Email:        wallet.User.Email,
		WalletNumber: wallet.WalletNumber,
		Balance:      wallet.Balance,
	}

	return dtoResponse, nil
}
