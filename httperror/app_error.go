package httperror

import (
	"errors"
)

var (
	ErrEmailAlreadyRegistered = errors.New("email already registered")
	ErrGenerateHash           = errors.New("failed to generate hash")
	ErrCreateUser             = errors.New("failed to create user")
	ErrInvalidEmailPassword   = errors.New("invalid email or password")
	ErrUserNotExist           = errors.New("user not exist")
	ErrFailedCreateToken      = errors.New("failed to creae token")
	ErrInvalidRegisterEmail   = errors.New("invalid email")
	ErrSourceOfFundsNotExist  = errors.New("source of funds not exist")
	ErrInvalidTopupAmount     = errors.New("invalid topup amount")
	ErrInvalidTransferAmount  = errors.New("invalid transfer amount")
	ErrInvalidUserWalletId    = errors.New("invalid user wallet id")
	ErrInvalidTargetWallet    = errors.New("invalid target wallet for transaction")
	ErrInvalidSourceWallet    = errors.New("invalid source wallet for transaction")
	ErrInvalidTransfer        = errors.New("invalid transfer")
	ErrWalletNotFound         = errors.New("wallet not found")
	ErrInvalidLimit           = errors.New("invalid limit")
)
