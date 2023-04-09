package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func ConnectDB() error {
	dsn := "host=localhost user=radjasa.dzar password=radjasa.dzar dbname=wallet_starter_db port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func GetError() error {
	return err
}
