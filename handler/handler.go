package handler

import (
	"assignment-golang-backend/usecase"
)

type Handler struct {
	walletUsecase usecase.WalletUsecase
}

type Config struct {
	WalletUsecase usecase.WalletUsecase
}

func New(cfg *Config) *Handler {
	return &Handler{
		walletUsecase: cfg.WalletUsecase,
	}
}
