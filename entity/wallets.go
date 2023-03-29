package entity

import (
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	WalletId     uint64         `gorm:"PrimaryKey" json:"wallet_id"`
	UserId       uint64         
	WalletNumber uint64         
	Balance      float64        
	CreatedAt    time.Time      
	UpdatedAt    time.Time      
	DeletedAt    gorm.DeletedAt 
	User         User           
}
