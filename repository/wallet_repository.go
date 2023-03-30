package repository

import (
	"assignment-golang-backend/dto"
	"assignment-golang-backend/entity"
	"assignment-golang-backend/httperror"

	"gorm.io/gorm"
)

type WalletRepository interface {
	GetSelfDetail(userId int) (*dto.WalletDetailDTO, error)
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

func (r *walletRepositoryImp) GetSelfDetail(userId int) (*dto.WalletDetailDTO, error) {
	var wallet *entity.Wallet
	err := r.db.Where("user_id = ?", userId).Preload("User").Find(&wallet).Error
	if err != nil {
		return nil, httperror.ErrUserNotExist
	}

	dtoResponse := &dto.WalletDetailDTO{
		Id:           wallet.WalletId,
		UserId:       wallet.UserId,
		UserName:     wallet.User.Name,
		Email:        wallet.User.Email,
		WalletNumber: wallet.WalletNumber,
		Balance:      wallet.Balance,
	}

	return dtoResponse, nil
}
