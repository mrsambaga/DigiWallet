package handler

import (
	"assignment-golang-backend/usecase"
)

type Handler struct {
	walletUsecase usecase.WalletUsecase
	userUsecase   usecase.UsersUsecase
}

type Config struct {
	WalletUsecase usecase.WalletUsecase
	UserUsecase   usecase.UsersUsecase
}

func New(cfg *Config) *Handler {
	return &Handler{
		walletUsecase: cfg.WalletUsecase,
		userUsecase:   cfg.UserUsecase,
	}
}
