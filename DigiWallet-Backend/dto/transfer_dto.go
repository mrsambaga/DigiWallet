package dto

type TransferRequestDTO struct {
	Amount             float64 `json:"amount" binding:"required"`
	TargetWalletNumber uint64  `json:"target_wallet_number" binding:"required"`
	Description        string  `json:"description"`
}

type TransferResponseDTO struct {
	Amount             float64 `json:"amount"`
	TransactionId      uint64  `json:"transaction_id"`
	SourceWalletNumber uint64  `json:"source_wallet_number"`
	TargetWalletNumber uint64  `json:"target_wallet_number"`
	Description        string  `json:"description"`
}
