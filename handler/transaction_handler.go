package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserTransactions(c *gin.Context) {
	userId := c.GetInt("id")

	transactions := h.transactionUsecase.GetUserTransactions(userId)

	c.JSON(http.StatusOK, transactions)
}
