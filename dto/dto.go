package dto

import "github.com/golang-jwt/jwt/v4"

type WalletDetailDTO struct {
	Id           uint64  `json:"wallet_id"`
	UserId       uint64  `json:"user_id"`
	UserName     string  `json:"user_name"`
	Email        string  `json:"email"`
	WalletNumber uint64  `json:"wallet_number"`
	Balance      float64 `json:"balance"`
}

type OtherWalletDetailDTO struct {
	UserName     string `json:"user_name"`
	WalletNumber uint64 `json:"wallet_number"`
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

type IdTokenClaims struct {
	UserID uint64 `json:"id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

type TokenResponse struct {
	Token string `json:"token"`
}

type TopupRequestDTO struct {
	Amount   float64 `json:"amount" binding:"required"`
	SourceId uint64  `json:"source_of_funds_id" binding:"required"`
}

type TopupResponseDTO struct {
	Amount             float64 `json:"amount"`
	TransactionId      uint64  `json:"transaction_id"`
	SourceId           uint64  `json:"source_of_funds"`
	TargetWalletNumber uint64  `json:"target_wallet_id"`
	Description        string  `json:"description"`
}

type TransferRequestDTO struct {
	Amount             float64 `json:"amount" binding:"required"`
	TargetWalletNumber uint64  `json:"target_wallet_number" binding:"required"`
	Description        string  `json:"description"`
}

type TransferResponseDTO struct {
	Amount             float64 `json:"amount"`
	TransactionId      uint64  `json:"transaction_id"`
	SourceWalletNumber uint64  `json:"source_wallet_id"`
	TargetWalletNumber uint64  `json:"target_wallet_id"`
	Description        string  `json:"description"`
}
