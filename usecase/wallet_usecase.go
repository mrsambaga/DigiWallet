package usecase

import (
	"assignment-golang-backend/dto"
	"assignment-golang-backend/repository"
)

type WalletUsecase interface {
	GetSelfDetail(userId int) (*dto.WalletDetailDTO, error)
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

func (u *walletUsecaseImp) GetSelfDetail(userId int) (*dto.WalletDetailDTO, error) {
	return u.walletRepository.GetSelfDetail(userId)
}
