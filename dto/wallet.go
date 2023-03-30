package dto

type WalletDetailDTO struct {
	Id           uint64  `json:"wallet_id"`
	UserId       uint64  `json:"user_id"`
	UserName     string  `json:"user_name"`
	Email        string  `json:"email"`
	WalletNumber uint64  `json:"wallet_number"`
	Balance      float64 `json:"balance"`
}

type OtherWalletDetailDTO struct {
	UserName     string  `json:"user_name"`
	WalletNumber uint64  `json:"wallet_number"`
}

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

type LoginRequestDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
