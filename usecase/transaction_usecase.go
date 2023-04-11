package usecase

import (
	"assignment-golang-backend/dto"
	"assignment-golang-backend/entity"
	"assignment-golang-backend/httperror"
	"assignment-golang-backend/repository"
)

type TransactionUsecase interface {
	GetUserTransactions(userId uint64) ([]*entity.Transaction, error)
	Topup(newTopUp *dto.TopupRequestDTO, userId uint64) (*dto.TopupResponseDTO, error)
	Transfer(newTransfer *dto.TransferRequestDTO, userId uint64) (*dto.TransferResponseDTO, error)
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

func (u *transactionUImp) GetUserTransactions(userId uint64) ([]*entity.Transaction, error) {
	transactions, err := u.transactionRepository.GetTransaction(userId)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (u *transactionUImp) Transfer(newTransfer *dto.TransferRequestDTO, userId uint64) (*dto.TransferResponseDTO, error) {
	err := u.checkValidTransferAmount(newTransfer.Amount)
	if err != nil {
		return nil, err
	}

	transfer := &entity.Transaction{
		TargetWalletNumber: newTransfer.TargetWalletNumber,
		Amount:             newTransfer.Amount,
		Description:        newTransfer.Description,
	}

	out, err := u.transactionRepository.Transaction(transfer, userId)
	if err != nil {
		return nil, err
	}

	newTransaction := &dto.TransferResponseDTO{
		TransactionId:      out.TransactionId,
		Amount:             out.Amount,
		TargetWalletNumber: out.TargetWalletNumber,
		SourceWalletNumber: *out.SourceWalletNumber,
		Description:        out.Description,
	}

	return newTransaction, nil

}

func (u *transactionUImp) Topup(newTopUp *dto.TopupRequestDTO, userId uint64) (*dto.TopupResponseDTO, error) {
	err := u.checkValidTopupAmount(newTopUp.Amount)
	if err != nil {
		return nil, err
	}

	description, err := u.generateDesc(newTopUp.SourceId)
	if err != nil {
		return nil, err
	}

	topUp := &entity.Transaction{
		Amount:      newTopUp.Amount,
		SourceId:    &newTopUp.SourceId,
		Description: description,
	}

	out, err := u.transactionRepository.Transaction(topUp, userId)
	if err != nil {
		return nil, err
	}

	newTransaction := &dto.TopupResponseDTO{
		TransactionId:      out.TransactionId,
		Amount:             out.Amount,
		TargetWalletNumber: out.TargetWalletNumber,
		SourceId:           *out.SourceId,
		Description:        out.Description,
	}

	return newTransaction, nil
}

func (u *transactionUImp) generateDesc(sourceId uint64) (string, error) {
	var description string
	switch sourceId {
	case 1001:
		description = "Topup from Bank Transfer"
	case 1002:
		description = "Topup from Credit Card"
	case 1003:
		description = "Topup from Cash"
	case 1004:
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

func (u *transactionUImp) checkValidTransferAmount(amount float64) error {
	if 1000 <= amount && amount <= 50000000 {
		return nil
	}

	return httperror.ErrInvalidTransferAmount
}
