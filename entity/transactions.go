package entity

import "time"

type Transaction struct {
	TransactionId      uint64 `gorm:"PrimaryKey"`
	TargetWalletNumber uint64
	SourceWalletNumber *uint64
	Amount             float64
	Description        string
	SourceId           *uint64
	CreatedAt          time.Time
}

type TransactionHistory struct {
}
