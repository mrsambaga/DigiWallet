package usecase

import (
	"assignment-golang-backend/dto"
	"assignment-golang-backend/entity"
	"assignment-golang-backend/httperror"
	"assignment-golang-backend/repository"
)

type GamesUsecase interface {
	ProcessGames(userId uint64, boxDTO *dto.BoxRequestDTO) ([]*entity.Boxes, error)
	GetLeaderboard() []*entity.Leaderboard
}

type gamesUImp struct {
	gamesRepository  repository.GamesRepository
	chanceRepository repository.ChanceRepository
}

type GamesUConfig struct {
	GamesRepository  repository.GamesRepository
	ChanceRepository repository.ChanceRepository
}

func NewGamesUsecase(cfg *GamesUConfig) GamesUsecase {
	return &gamesUImp{
		gamesRepository:  cfg.GamesRepository,
		chanceRepository: cfg.ChanceRepository,
	}
}

func (u *gamesUImp) ProcessGames(userId uint64, boxDTO *dto.BoxRequestDTO) ([]*entity.Boxes, error) {
	chance, err := u.chanceRepository.GetChance(userId)
	if err != nil {
		return nil, err
	}
	if chance.Chance == 0 {
		return nil, httperror.ErrNoChance
	}

	boxes, err := u.gamesRepository.ProcessGames(userId, boxDTO.BoxId)
	if err != nil {
		return nil, err
	}

	return boxes, nil
}

func (u *gamesUImp) GetLeaderboard() []*entity.Leaderboard {
	leaderboard := u.gamesRepository.GetLeaderboard()

	return leaderboard
}
