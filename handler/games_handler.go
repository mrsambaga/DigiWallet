package handler

import (
	"assignment-golang-backend/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ProcessGames(c *gin.Context) {
	var boxDTO dto.BoxRequestDTO
	userId := c.GetInt("id")

	if err := c.ShouldBindJSON(&boxDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	boxes, err := h.gamesUsecase.ProcessGames(uint64(userId), &boxDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESS_CREATED",
		"message": "Games successfully processed !",
		"data":    boxes,
	})
}

func (h *Handler) GetChance(c *gin.Context) {
	userId := c.GetInt("id")

	chance, err := h.gamesUsecase.GetChance(uint64(userId))
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
		"message": "Success get self chance",
		"data":    chance,
	})
}

func (h *Handler) GetLeaderboard(c *gin.Context) {
	leaderboard := h.gamesUsecase.GetLeaderboard()

	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESS_CREATED",
		"message": "Success get self chance",
		"data":    leaderboard,
	})
}
