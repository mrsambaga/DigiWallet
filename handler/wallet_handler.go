package handler

import (
	"assignment-golang-backend/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetDetail(c *gin.Context) {
	userIdStr := c.Param("user-id")
	loggedUserId := c.GetInt("id")

	userIdParam, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": "Invalid user ID parameter",
		})
		return
	}

	if loggedUserId == userIdParam {
		wallet, err := h.walletUsecase.GetSelfDetail(userIdParam)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    "INTERNAL_SERVER_ERROR",
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": wallet})
		return
	}

	wallet, err := h.walletUsecase.GetOtherUserDetail(userIdParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": wallet})
}

func (h *Handler) GetSelfDetail(id int) (*dto.WalletDetailDTO, error) {
	wallet, err := h.walletUsecase.GetSelfDetail(id)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (h *Handler) GetOtherUserDetail(userIdStr string) (*dto.OtherWalletDetailDTO, error) {

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return nil, err
	}

	wallet, err := h.walletUsecase.GetOtherUserDetail(userId)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}
