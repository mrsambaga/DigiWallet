package usecase

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/repository"
)

type ChanceUsecase interface {
	GetChance(userId uint64) (*entity.Chance, error)
}

type chanceUImp struct {
	chanceRepository repository.ChanceRepository
}

type ChanceUConfig struct {
	ChanceRepository repository.ChanceRepository
}

func NewChanceUsecase(cfg *ChanceUConfig) ChanceUsecase {
	return &chanceUImp{
		chanceRepository: cfg.ChanceRepository,
	}
}

func (u *chanceUImp) GetChance(userId uint64) (*entity.Chance, error) {
	chance, err := u.chanceRepository.GetChance(userId)
	if err != nil {
		return nil, err
	}

	return chance, nil
}
