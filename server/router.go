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
		DB: db.GetDB(),
	})
	walletUsecase := usecase.NewWalletUsecase(&usecase.WalletUConfig{
		WalletRepository: walletRepo,
	})
	userRepo := repository.NewUserRepository(&repository.UserRConfig{
		DB: db.GetDB(),
	})
	userUsecase := usecase.NewUsersUsecase(&usecase.UsersUsecaseConfig{
		UsersRepository: userRepo,
	})

	return NewRouter(&RouterConfig{
		WalletUsecase: walletUsecase,
		UserUsecase:   userUsecase,
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
