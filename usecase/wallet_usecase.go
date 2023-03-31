package usecase

import (
	"assignment-golang-backend/dto"
	"assignment-golang-backend/repository"
)

type WalletUsecase interface {
	GetSelfDetail(userId int) (*dto.WalletDetailDTO, error)
	GetOtherUserDetail(userId int) (*dto.OtherWalletDetailDTO, error)
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
	detail, err := u.walletRepository.GetSelfDetail(userId)
	if err != nil {
		return nil, err
	}

	out := &dto.WalletDetailDTO{}
	out.Id = detail.User.UserId
	out.UserId = detail.UserId
	out.UserName = detail.User.Name
	out.Email = detail.User.Email
	out.WalletNumber = detail.WalletNumber
	out.Balance = detail.Balance

	return out, nil
}

func (u *walletUsecaseImp) GetOtherUserDetail(userId int) (*dto.OtherWalletDetailDTO, error) {
	detail, err := u.walletRepository.GetOtherUserDetail(userId)
	if err != nil {
		return nil, err
	}

	out := &dto.OtherWalletDetailDTO{}
	out.UserName = detail.User.Name
	out.WalletNumber = detail.WalletNumber

	return out, nil
}
