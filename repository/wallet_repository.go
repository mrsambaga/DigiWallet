package repository

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/httperror"
	"strings"

	"gorm.io/gorm"
)

type WalletRepository interface {
	GetSelfDetail(userId int) (*entity.Wallet, error)
	GetOtherUserDetail(userId int) (*entity.Wallet, error)
}

type walletRepositoryImp struct {
	db *gorm.DB
}

type WalletRConfig struct {
	DB *gorm.DB
}

func NewWalletRepository(cfg *WalletRConfig) WalletRepository {
	return &walletRepositoryImp{
		db: cfg.DB,
	}
}

func (r *walletRepositoryImp) GetSelfDetail(userId int) (*entity.Wallet, error) {
	var wallet *entity.Wallet
	err := r.db.Where("user_id = ?", userId).Preload("User").Find(&wallet).Error
	if err != nil {
		return nil, httperror.ErrUserNotExist
	}

	return wallet, nil
}

func (r *walletRepositoryImp) GetOtherUserDetail(userId int) (*entity.Wallet, error) {
	var wallet *entity.Wallet
	err := r.db.Where("user_id = ?", userId).Preload("User").Find(&wallet).Error
	if err != nil {
		return nil, httperror.ErrUserNotExist
	}

	hiddenName := r.hideName(wallet.User.Name)

	wallet.User.Name = hiddenName

	return wallet, nil
}

func (r *walletRepositoryImp) hideName(name string) string {
	var names []string

	nameSplit := strings.Split(name, " ")
	for _, name := range nameSplit {
		hiddenName := name[:1] + "* "
		names = append(names, hiddenName)
	}

	joinedName := strings.Join(names, "")
	return joinedName[:len(joinedName)-1]
}
