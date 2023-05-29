package repository

import (
	"assignment-golang-backend/entity"

	"gorm.io/gorm"
)

type ChanceRepository interface {
	GetChance(userId uint64) (*entity.Chance, error)
	SubtractChance(tx *gorm.DB, userId uint64) error
}

type chanceRepositoryImp struct {
	db *gorm.DB
}

type ChanceRConfig struct {
	DB *gorm.DB
}

func NewChanceRepository(cfg *ChanceRConfig) ChanceRepository {
	return &chanceRepositoryImp{
		db: cfg.DB,
	}
}

func (r *chanceRepositoryImp) GetChance(userId uint64) (*entity.Chance, error) {
	var chance *entity.Chance
	err := r.db.Where("user_id = ?", userId).First(&chance).Error
	if err != nil {
		return nil, err
	}
	return chance, nil
}

func (r *chanceRepositoryImp) SubtractChance(tx *gorm.DB, userId uint64) error {
	var chance *entity.Chance
	err := tx.Model(chance).Where("user_id = ?", userId).Update("chance", gorm.Expr("chance - ?", 1)).Error
	if err != nil {
		return err
	}

	return nil
}
