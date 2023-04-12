package usecase

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/repository"
)

type GamesUsecase interface {
	ProcessGames(userId uint64, box *entity.Boxes) ([]*entity.Boxes, error)
	GetChance(userId uint64) (*entity.Chance, error)
	GetLeaderboard() []*entity.Leaderboard
}

type gamesUImp struct {
	gamesRepository repository.GamesRepository
}

type GamesUConfig struct {
	GamesRepository repository.GamesRepository
}

func NewGamesUsecase(cfg *GamesUConfig) GamesUsecase {
	return &gamesUImp{
		gamesRepository: cfg.GamesRepository,
	}
}

func (u *gamesUImp) ProcessGames(userId uint64, box *entity.Boxes) ([]*entity.Boxes, error) {
	boxes, err := u.gamesRepository.ProcessGames(userId, box.BoxId)
	if err != nil {
		return nil, err
	}

	return boxes, nil
}

func (u *gamesUImp) GetChance(userId uint64) (*entity.Chance, error) {
	chance, err := u.gamesRepository.GetChance(userId)
	if err != nil {
		return nil, err
	}

	return chance, nil
}

func (u *gamesUImp) GetLeaderboard() []*entity.Leaderboard {
	leaderboard := u.gamesRepository.GetLeaderboard()

	return leaderboard
}
