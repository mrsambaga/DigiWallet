package repository

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/httperror"
	"strings"
	"time"

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
	err := r.db.Joins("User").Where("wallets.user_id = ?", userId).First(&wallet).Error
	if err != nil {
		return nil, httperror.ErrUserNotExist
	}

	return wallet, nil
}

func (r *walletRepositoryImp) GetOtherUserDetail(userId int) (*entity.Wallet, error) {
	var wallet *entity.Wallet
	err := r.db.Joins("User").Where("wallets.user_id = ?", userId).First(&wallet).Error
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

func (r *userRepositoryImp) DeleteUserSoft(userID uint64) error {
	user := &entity.User{}
	if err := r.db.Model(user).Where("user_id = ?", userID).Update("DeletedAt", time.Now()).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepositoryImp) DeleteUserSoft2(userID uint64) error {
	user := &entity.User{}
	if err := r.db.Where("user_id = ?", userID).Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepositoryImp) DeleteUserPermanent(userID uint64) error {
	user := &entity.User{}
	if err := r.db.Unscoped().Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *walletRepositoryImp) UpdateWallet(userId uint64, balance float64) error {
	wallet := &entity.Wallet{}
	err := r.db.Where("user_id = ?", userId).First(wallet).Error
	if err != nil {
		return err
	}
	wallet.Balance = balance
	wallet.UpdatedAt = time.Now()
	err = r.db.Save(wallet).Error
	if err != nil {
		return err
	}
	return nil
}
