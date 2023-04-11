package handler

import (
	"assignment-golang-backend/dto"
	"assignment-golang-backend/httperror"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserTransactions(c *gin.Context) {
	userId := c.GetInt("id")

	transactions, err := h.transactionUsecase.GetUserTransactions(uint64(userId))
	if err != nil {
		if errors.Is(err, httperror.ErrWalletNotFound) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "BAD_REQUEST",
				"message": "Wallet not found !",
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "BAD_REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *Handler) Topup(c *gin.Context) {
	var newTopupRequest *dto.TopupRequestDTO
	userId := c.GetInt("id")

	if err := c.ShouldBindJSON(&newTopupRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "BAD_REQUEST",
			"message": err.Error(),
		})
		return
	}

	newTopupResponse, err := h.transactionUsecase.Topup(newTopupRequest, uint64(userId))
	if err != nil {
		if errors.Is(err, httperror.ErrInvalidTopupAmount) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "BAD_REQUEST",
				"message": "Invalid topup amount : please enter amount between 50.000 & 10.000.000",
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "BAD_REQUEST",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": newTopupResponse})
}

func (h *Handler) Transfer(c *gin.Context) {
	var newTransferRequest *dto.TransferRequestDTO
	userId := c.GetInt("id")

	if err := c.ShouldBindJSON(&newTransferRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "BAD_REQUEST",
			"message": err.Error(),
		})
		return
	}

	newTransferResponse, err := h.transactionUsecase.Transfer(newTransferRequest, uint64(userId))
	if err != nil {
		if errors.Is(err, httperror.ErrInvalidTopupAmount) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "BAD_REQUEST",
				"message": "Invalid topup amount : please enter amount between 1000 & 50.000.000",
			})
			return
		} else if errors.Is(err, httperror.ErrInvalidTransfer) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "BAD_REQUEST",
				"message": "Invalid transfer : cannot transfer to self wallet",
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "BAD_REQUEST",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": newTransferResponse})
}
