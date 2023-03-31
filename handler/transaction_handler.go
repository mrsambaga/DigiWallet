package handler

import (
	"assignment-golang-backend/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserTransactions(c *gin.Context) {
	userId := c.GetInt("id")

	transactions := h.transactionUsecase.GetUserTransactions(userId)

	c.JSON(http.StatusOK, transactions)
}

func (h *Handler) Topup(c *gin.Context) {
	var newTopUp *dto.TopupRequestDTO
	userId := c.GetInt("id")

	if err := c.ShouldBindJSON(&newTopUp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "BAD_REQUEST",
			"message": err.Error(),
		})
		return
	}

	err := h.transactionUsecase.Topup(newTopUp, userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "BAD_REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "Top Up Successful")
}
