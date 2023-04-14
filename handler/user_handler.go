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
			"data":    nil,
		})
		return
	}

	result, err := h.userUsecase.Register(newUser)
	if err != nil {
		if errors.Is(err, httperror.ErrEmailAlreadyRegistered) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "BAD_REQUEST",
				"message": "Email already registered !",
				"data":    nil,
			})
			return
		} else if errors.Is(err, httperror.ErrCreateUser) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "BAD_REQUEST",
				"message": "Failed to create user !",
				"data":    nil,
			})
			return
		} else if errors.Is(err, httperror.ErrGenerateHash) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "BAD_REQUEST",
				"message": "Failed to generate hash password !",
				"data":    nil,
			})
			return
		} else if errors.Is(err, httperror.ErrInvalidRegisterEmail) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "BAD_REQUEST",
				"message": "Invalid email, please enter this format : 'xxx@xxx.com'",
				"data":    nil,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "BAD_REQUEST",
			"message": err.Error(),
			"data":    nil,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESS_CREATED",
		"message": "Register successful !",
		"data":    result,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var loginUserDTO *dto.LoginRequestDTO

	if err := c.ShouldBindJSON(&loginUserDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "BAD_REQUEST",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	token, err := h.userUsecase.Login(loginUserDTO)
	if err != nil {
		if errors.Is(err, httperror.ErrInvalidEmailPassword) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "BAD_REQUEST",
				"message": "Wrong email or password !",
				"data":    nil,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "BAD_REQUEST",
			"message": err.Error(),
			"data":    nil,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESS_CREATED",
		"message": "Login successful !",
		"data":    token,
	})
}
