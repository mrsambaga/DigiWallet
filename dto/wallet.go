package dto

type WalletDetail struct {
	Id           uint64  `json:"wallet_id"`
	UserId       uint64  `json:"user_id"`
	UserName     string  `json:"user_name"`
	Email        string  `json:"email"`
	WalletNumber uint64  `json:"wallet_number"`
	Balance      float64 `json:"balance"`
}
