package usecase

import (
	"assignment-golang-backend/dto"
	"assignment-golang-backend/entity"
	"assignment-golang-backend/httperror"
	"assignment-golang-backend/repository"
)

type TransactionUsecase interface {
	GetUserTransactions(userId int) []*entity.Transaction
	Topup(newTopUp *dto.TopupRequestDTO, userId int) (*dto.TopupResponseDTO, error)
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

func (u *transactionUImp) Topup(newTopUp *dto.TopupRequestDTO, userId int) (*dto.TopupResponseDTO, error) {
	err := u.checkValidTopupAmount(newTopUp.Amount)
	if err != nil {
		return nil, err
	}

	description, err := u.generateDesc(newTopUp.SourceId)
	if err != nil {
		return nil, err
	}

	topUp := &entity.Transaction{
		TargetWalletId: uint64(userId),
		Amount:         newTopUp.Amount,
		SourceId:       &newTopUp.SourceId,
		Description:    description,
	}

	out, err := u.transactionRepository.Topup(topUp)
	if err != nil {
		return nil, err
	}

	newTransaction := &dto.TopupResponseDTO{
		TransactionId:  out.TransactionId,
		Amount:         out.Amount,
		TargetWalletId: out.TargetWalletId,
		SourceId:       *out.SourceId,
		Description:    out.Description,
	}

	return newTransaction, nil
}

func (u *transactionUImp) generateDesc(sourceId uint64) (string, error) {
	var description string
	switch sourceId {
	case 1:
		description = "Topup from Bank Transfer"
	case 2:
		description = "Topup from Credit Card"
	case 3:
		description = "Topup from Cash"
	case 4:
		description = "Topup from Reward"
	default:
		return "", httperror.ErrSourceOfFundsNotExist
	}

	return description, nil
}

func (u *transactionUImp) checkValidTopupAmount(amount float64) error {
	if 50000 <= amount && amount <= 10000000 {
		return nil
	}

	return httperror.ErrInvalidTopupAmount
}
