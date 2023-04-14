package repository

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/httperror"

	"gorm.io/gorm"
)

type GamesRepository interface {
	ProcessGames(userId uint64, boxIdx uint64) ([]*entity.Boxes, error)
	GetLeaderboard() []*entity.Leaderboard
}

type gamesRepoImp struct {
	db         *gorm.DB
	chanceRepo chanceRepositoryImp
}

type GamesRConfig struct {
	DB         *gorm.DB
	ChanceRepo ChanceRConfig
}

func NewGamesRepository(cfg *GamesRConfig) GamesRepository {
	return &gamesRepoImp{
		db: cfg.DB,
	}
}

func (r *gamesRepoImp) ProcessGames(userId uint64, boxIdx uint64) ([]*entity.Boxes, error) {
	var Boxes []*entity.Boxes
	r.db.Find(&Boxes)

	if err := r.db.Transaction(func(tx *gorm.DB) error {

		//Subtract chance
		err := r.chanceRepo.SubtractChance(tx, userId)
		if err != nil {
			return err
		}

		// Get selected box prize
		var selectedBox *entity.Boxes
		if err := tx.Model(&Boxes).Where("box_id = ?", boxIdx).First(&selectedBox).Error; err != nil {
			return err
		}

		// Update wallet balance
		var wallet *entity.Wallet
		if err := tx.Model(&wallet).Where("user_id = ?", userId).Update("balance", gorm.Expr("balance + ?", selectedBox.Prize)).Error; err != nil {
			return err
		}
		if err := tx.Where("user_id = ?", userId).First(&wallet).Error; err != nil {
			return httperror.ErrInvalidUserWalletId
		}

		// Create record in leaderboard / update if already exist
		var leaderboard entity.Leaderboard
		if err := tx.Where("user_id = ?", userId).First(&leaderboard).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				var user *entity.User
				if err := tx.Where("user_id = ?", userId).First(&user).Error; err != nil {
					return err
				}

				leaderboard.UserId = userId
				leaderboard.Name = user.Name
				leaderboard.TotalWinning = selectedBox.Prize
				if err := tx.Create(&leaderboard).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}
		if err := tx.Model(&leaderboard).Where("user_id = ?", userId).Update("total_winning", gorm.Expr("total_winning + ?", selectedBox.Prize)).Error; err != nil {
			return err
		}

		// Create record in transaction as reward
		gamesSourceId := uint64(1004)
		var gamesSourceIdPtr *uint64 = &gamesSourceId
		newTransaction := &entity.Transaction{
			TargetWalletNumber: wallet.WalletNumber,
			Amount:             selectedBox.Prize,
			Description:        "Topup from Reward",
			SourceId:           gamesSourceIdPtr,
		}
		if err := tx.Create(&newTransaction).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}
	return Boxes, nil
}

func (r *gamesRepoImp) GetLeaderboard() []*entity.Leaderboard {
	var leaderboard []*entity.Leaderboard
	r.db.Order("total_winning DESC").Find(&leaderboard)

	return leaderboard
}
