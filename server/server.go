package server

import (
	"assignment-golang-backend/handler"
	"assignment-golang-backend/usecase"
	"log"
	"net/http"

	"assignment-golang-backend/middleware"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	WalletUsecase      usecase.WalletUsecase
	UserUsecase        usecase.UsersUsecase
	TransactionUsecase usecase.TransactionUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.New()

	h := handler.New(&handler.Config{
		WalletUsecase:      cfg.WalletUsecase,
		UserUsecase:        cfg.UserUsecase,
		TransactionUsecase: cfg.TransactionUsecase,
	})

	router.GET("/users/:user-id", middleware.AuthorizeJWT, h.GetDetail)
	router.GET("/users/transaction", middleware.AuthorizeJWT, h.GetUserTransactions)
	router.POST("/users/topup", middleware.AuthorizeJWT, h.Topup)
	router.POST("/users/transfer", middleware.AuthorizeJWT, h.Transfer)
	router.POST("/register", h.Register)
	router.POST("/login", h.Login)

	log.Fatal(http.ListenAndServe(":8000", router))
	return router
}
