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
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.New()
	h := handler.New(&handler.Config{
		WalletUsecase: cfg.WalletUsecase,
	})

	router.GET("/wallet/:user-id", h.GetDetail)

	log.Fatal(http.ListenAndServe(":8013", router))
	return router
}
