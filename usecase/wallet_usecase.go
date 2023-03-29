package usecase

import (
	"assignment-golang-backend/dto"
	"assignment-golang-backend/repository"
)

type WalletUsecase interface {
	GetDetail(userId int) (*dto.WalletDetail, error)
}

type walletUsecaseImp struct {
	walletRepository repository.WalletRepository
}

type WalletUConfig struct {
	WalletRepository repository.WalletRepository
}

func NewWalletUsecase(cfg *WalletUConfig) WalletUsecase {
	return &walletUsecaseImp{
		walletRepository: cfg.WalletRepository,
	}
}

func (u *walletUsecaseImp) GetDetail(userId int) (*dto.WalletDetail, error) {
	return u.walletRepository.GetDetail(userId)
}
