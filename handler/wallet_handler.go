package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetDetail(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": "Invalid user ID parameter",
		})
		return
	}

	wallet, err := h.walletUsecase.GetDetail(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wallet})
}
