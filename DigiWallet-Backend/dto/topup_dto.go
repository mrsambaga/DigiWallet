package dto

type TopupRequestDTO struct {
	Amount   float64 `json:"amount" binding:"required"`
	SourceId uint64  `json:"source_of_funds_id" binding:"required"`
}

type TopupResponseDTO struct {
	Amount             float64 `json:"amount"`
	TransactionId      uint64  `json:"transaction_id"`
	SourceId           uint64  `json:"source_of_funds"`
	TargetWalletNumber uint64  `json:"target_wallet_number"`
	Description        string  `json:"description"`
}
