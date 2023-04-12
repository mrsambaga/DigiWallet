package handler

import (
	"assignment-golang-backend/usecase"
)

type Handler struct {
	walletUsecase      usecase.WalletUsecase
	userUsecase        usecase.UsersUsecase
	transactionUsecase usecase.TransactionUsecase
	gamesUsecase       usecase.GamesUsecase
}

type Config struct {
	WalletUsecase      usecase.WalletUsecase
	UserUsecase        usecase.UsersUsecase
	TransactionUsecase usecase.TransactionUsecase
	GamesUsecase       usecase.GamesUsecase
}

func New(cfg *Config) *Handler {
	return &Handler{
		walletUsecase:      cfg.WalletUsecase,
		userUsecase:        cfg.UserUsecase,
		transactionUsecase: cfg.TransactionUsecase,
		gamesUsecase:       cfg.GamesUsecase,
	}
}
