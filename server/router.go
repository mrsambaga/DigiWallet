package server

import (
	"assignment-golang-backend/db"
	"assignment-golang-backend/repository"
	"assignment-golang-backend/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	walletRepo := repository.NewWalletRepository(&repository.WalletRConfig{
		DB: db.Get(),
	})
	walletUsecase := usecase.NewWalletUsecase(&usecase.WalletUConfig{
		WalletRepository: walletRepo,
	})
	userRepo := repository.NewUserRepository(&repository.UserRConfig{
		DB: db.Get(),
	})
	userUsecase := usecase.NewUsersUsecase(&usecase.UsersUsecaseConfig{
		UsersRepository: userRepo,
	})
	transactionRepo := repository.NewTransactionRepository(&repository.TransactionRConfig{
		DB: db.Get(),
	})
	transactionUsecase := usecase.NewTransactionUsecase(&usecase.TransactionUConfig{
		TransactionRepository: transactionRepo,
	})
	gamesRepo := repository.NewGamesRepository(&repository.GamesRConfig{
		DB: db.Get(),
	})
	chanceRepo := repository.NewChanceRepository(&repository.ChanceRConfig{
		DB: db.Get(),
	})
	gamesUsecase := usecase.NewGamesUsecase(&usecase.GamesUConfig{
		GamesRepository:  gamesRepo,
		ChanceRepository: chanceRepo,
	})
	chanceUsecase := usecase.NewChanceUsecase(&usecase.ChanceUConfig{
		ChanceRepository: chanceRepo,
	})

	return NewRouter(&RouterConfig{
		WalletUsecase:      walletUsecase,
		UserUsecase:        userUsecase,
		TransactionUsecase: transactionUsecase,
		GamesUsecase:       gamesUsecase,
		ChanceUsecase:      chanceUsecase,
	})
}

func Init() {
	r := createRouter()
	err := r.Run()
	if err != nil {
		log.Println("error while running server", err)
		return
	}
}
