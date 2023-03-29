package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserId    uint64         `gorm:"PrimaryKey" json:"user_id"`
	Name      string         
	Email     string         
	Password  string         
	CreatedAt time.Time      
	UpdatedAt time.Time      
	DeletedAt gorm.DeletedAt 
}
