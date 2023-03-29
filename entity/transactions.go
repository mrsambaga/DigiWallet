package entity

import "time"

type Transaction struct {
	TransactionId       uint64 `gorm:"PrimaryKey"`
	SourceWalletId      uint64
	SourceWallet        Wallet
	DestinationWalletId uint64
	DestinationWallet   Wallet
	Amount              float64
	SourceId            uint64
	CreatedAt           time.Time
}
