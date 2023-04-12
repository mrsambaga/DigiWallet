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
	GamesUsecase       usecase.GamesUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.New()

	h := handler.New(&handler.Config{
		WalletUsecase:      cfg.WalletUsecase,
		UserUsecase:        cfg.UserUsecase,
		TransactionUsecase: cfg.TransactionUsecase,
		GamesUsecase:       cfg.GamesUsecase,
	})

	router.GET("/profile", middleware.AuthorizeJWT, h.GetSelfDetailHandler)
	router.GET("/inquiry/:user-id", middleware.AuthorizeJWT, h.GetOtherDetailHandler)
	router.GET("/profile/transaction", middleware.AuthorizeJWT, h.GetUserTransactions)
	router.POST("/topup", middleware.AuthorizeJWT, h.Topup)
	router.POST("/transfer", middleware.AuthorizeJWT, h.Transfer)
	router.POST("/games", middleware.AuthorizeJWT, h.ProcessGames)
	router.GET("/chance", middleware.AuthorizeJWT, h.GetChance)
	router.GET("/leaderboard", middleware.AuthorizeJWT, h.GetLeaderboard)
	router.POST("/register", h.Register)
	router.POST("/login", h.Login)

	log.Fatal(http.ListenAndServe(":8000", router))
	return router
}
