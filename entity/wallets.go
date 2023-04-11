package entity

import (
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	WalletNumber uint64 `gorm:"PrimaryKey" json:"wallet_number"`
	UserId       uint64
	Balance      float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	User         User
}
