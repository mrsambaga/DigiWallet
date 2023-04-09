package handler

import (
	"assignment-golang-backend/dto"
	"assignment-golang-backend/httperror"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	var newUser *dto.RegisterRequestDTO

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "BAD_REQUEST",
			"message": err.Error(),
		})
		return
	}

	result, err := h.userUsecase.Register(newUser)
	if err != nil {
		if errors.Is(err, httperror.ErrEmailAlreadyRegistered) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "BAD_REQUEST",
				"message": "Email already registered !",
			})
			return
		} else if errors.Is(err, httperror.ErrCreateUser) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "BAD_REQUEST",
				"message": "Failed to create user !",
			})
			return
		} else if errors.Is(err, httperror.ErrGenerateHash) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "BAD_REQUEST",
				"message": "Failed to generate hash password !",
			})
			return
		} else if errors.Is(err, httperror.ErrInvalidRegisterEmail) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "BAD_REQUEST",
				"message": "Invalid email, please enter this format : 'xxx@shopee.com'",
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "BAD_REQUEST",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (h *Handler) Login(c *gin.Context) {
	var loginUserDTO *dto.LoginRequestDTO

	if err := c.ShouldBindJSON(&loginUserDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "BAD_REQUEST",
			"message": err.Error(),
		})
		return
	}

	token, err := h.userUsecase.Login(loginUserDTO)
	if err != nil {
		if errors.Is(err, httperror.ErrInvalidEmailPassword) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "BAD_REQUEST",
				"message": "Wrong email or password !",
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "BAD_REQUEST",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
