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
	UserName     string `json:"user_name"`
	WalletNumber uint64 `json:"wallet_number"`
}
