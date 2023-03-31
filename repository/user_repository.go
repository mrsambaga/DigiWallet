package repository

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/httperror"
	"assignment-golang-backend/util"

	"gorm.io/gorm"
)

type UsersRepository interface {
	Register(newUser *entity.User) (*entity.Wallet, error)
	Login(loginUser *entity.User) (string, error)
}

type userRepositoryImp struct {
	db *gorm.DB
}

type UserRConfig struct {
	DB *gorm.DB
}

func NewUserRepository(cfg *UserRConfig) UsersRepository {
	return &userRepositoryImp{
		db: cfg.DB,
	}
}

func (r *userRepositoryImp) Register(newUser *entity.User) (*entity.Wallet, error) {
	var user entity.User

	if err := r.db.Where("email = ?", newUser.Email).First(&user).Error; err == nil {
		return nil, httperror.ErrEmailAlreadyRegistered
	}

	hashedPassword, err := util.HashPassword(newUser.Password)
	if err != nil {
		return nil, err
	}
	newUser.Password = hashedPassword

	if err := r.db.Create(&newUser).Error; err != nil {
		return nil, httperror.ErrCreateUser
	}

	newWallet := r.NewWallet(newUser)
	if err := r.db.Create(&newWallet).Error; err != nil {
		return nil, httperror.ErrCreateUser
	}

	return newWallet, nil
}

func (r *userRepositoryImp) NewWallet(user *entity.User) *entity.Wallet {
	var lastWallet *entity.Wallet
	var walletNumber uint64

	if err := r.db.Order("wallet_number desc").First(&lastWallet).Error; err != nil {
		walletNumber = 123000000000
	} else {
		walletNumber = lastWallet.WalletNumber + 1
	}

	return &entity.Wallet{
		UserId:       user.UserId,
		WalletNumber: walletNumber,
		Balance:      0,
		User:         *user,
	}
}

func (r *userRepositoryImp) Login(loginUser *entity.User) (string, error) {
	var user *entity.User

	if err := r.db.Where("email = ?", loginUser.Email).First(&user).Error; err != nil {
		return "", httperror.ErrInvalidEmail
	}

	ok := util.ComparePassword(user.Password, loginUser.Password)
	if !ok {
		return "", httperror.ErrInvalidPassword
	}

	loginUser.UserId = user.UserId
	tokenString, err := util.GenerateAccessToken(loginUser)
	if err != nil {
		return "", err
	}

	return tokenString.Token, nil
}
