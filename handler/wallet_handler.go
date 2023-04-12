package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetSelfDetailHandler(c *gin.Context) {
	loggedUserId := c.GetInt("id")

	wallet, err := h.walletUsecase.GetSelfDetail(loggedUserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESS_CREATED",
		"message": "Success get self detail",
		"data":    wallet,
	})
}

func (h *Handler) GetOtherDetailHandler(c *gin.Context) {
	userIdStr := c.Param("user-id")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "BAD_REQUEST",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	wallet, err := h.walletUsecase.GetOtherUserDetail(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESS_CREATED",
		"message": "Success get self detail",
		"data":    wallet,
	})
}
