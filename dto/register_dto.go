package dto

type RegisterRequestDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterResponseDTO struct {
	UserId       uint64  `json:"user_id"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Password     string  `json:"password"`
	WalletNumber uint64  `json:"wallet_number"`
	Balance      float64 `json:"balance"`
}
