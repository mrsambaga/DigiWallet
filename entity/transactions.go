package entity

import "time"

type Transaction struct {
	TransactionId  uint64 `gorm:"PrimaryKey"`
	SourceWalletId *uint64
	TargetWalletId uint64
	Amount         float64
	Description    string
	SourceId       *uint64
	CreatedAt      time.Time
}
