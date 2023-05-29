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
	sort := c.Query("sort")
	sortBy := c.Query("sortBy")
	limit := c.Query("limit")
	search := c.Query("search")
	sortComplete := ""
	if sort != "" && sortBy != "" {
		sortComplete = sortBy + " " + sort
	}

	transactions, err := h.transactionUsecase.GetUserTransactions(uint64(userId), sortComplete, limit, search)
	if err != nil {
		if errors.Is(err, httperror.ErrWalletNotFound) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "BAD_REQUEST",
				"message": "Wallet not found !",
				"data":    nil,
			})
			return
		}
		if errors.Is(err, httperror.ErrInvalidLimit) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "BAD_REQUEST",
				"message": "Invalid limit : please enter a number",
				"data":    nil,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESS_CREATED",
		"message": "Get transaction history successful !",
		"data":    transactions,
	})
}

func (h *Handler) Topup(c *gin.Context) {
	var newTopupRequest *dto.TopupRequestDTO
	userId := c.GetInt("id")

	if err := c.ShouldBindJSON(&newTopupRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	newTopupResponse, err := h.transactionUsecase.Topup(newTopupRequest, uint64(userId))
	if err != nil {
		if errors.Is(err, httperror.ErrInvalidTopupAmount) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "BAD_REQUEST",
				"message": "Invalid topup amount : please enter amount between 50.000 & 10.000.000",
				"data":    nil,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": err.Error(),
			"data":    nil,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESS_CREATED",
		"message": "Topup successful !",
		"data":    newTopupResponse,
	})
}

func (h *Handler) Transfer(c *gin.Context) {
	var newTransferRequest *dto.TransferRequestDTO
	userId := c.GetInt("id")

	if err := c.ShouldBindJSON(&newTransferRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	newTransferResponse, err := h.transactionUsecase.Transfer(newTransferRequest, uint64(userId))
	if err != nil {
		if errors.Is(err, httperror.ErrInvalidTransferAmount) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "BAD_REQUEST",
				"message": "Invalid transfer amount : please enter amount between 1000 & 50.000.000",
				"data":    nil,
			})
			return
		} else if errors.Is(err, httperror.ErrInvalidTransfer) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "BAD_REQUEST",
				"message": "Invalid transfer : cannot transfer to self wallet",
				"data":    nil,
			})
			return
		} else if errors.Is(err, httperror.ErrInvalidTargetWallet) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "BAD_REQUEST",
				"message": "Invalid transfer : target wallet not found",
				"data":    nil,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESS_CREATED",
		"message": "Transfer successful !",
		"data":    newTransferResponse,
	})
}
