package server

import (
	"assignment-golang-backend/handler"
	"assignment-golang-backend/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	WalletUsecase usecase.WalletUsecase
	UserUsecase   usecase.UsersUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.New()
	h := handler.New(&handler.Config{
		WalletUsecase: cfg.WalletUsecase,
		UserUsecase:   cfg.UserUsecase,
	})

	router.GET("/users/:user-id", h.GetSelfDetail)

	router.POST("/register", h.Register)

	router.POST("/login", h.Login)

	log.Fatal(http.ListenAndServe(":8031", router))
	return router
}
