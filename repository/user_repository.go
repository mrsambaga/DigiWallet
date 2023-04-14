package repository

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/httperror"
	"strings"

	"gorm.io/gorm"
)

type UsersRepository interface {
	CreateUser(newUser *entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
	CreateNewWallet(user *entity.User) (*entity.Wallet, error)
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

func (r *userRepositoryImp) GetUserByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	if err := r.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, httperror.ErrInvalidEmailPassword
	}
	return user, nil
}

func (r *userRepositoryImp) CreateUser(newUser *entity.User) error {
	if err := r.db.Create(newUser).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return httperror.ErrEmailAlreadyRegistered
		}

		return httperror.ErrCreateUser
	}

	//New record in chance
	newChance := &entity.Chance{
		UserId: newUser.UserId,
		Chance: 0,
	}
	if err := r.db.Create(&newChance).Error; err != nil {
		return httperror.ErrCreateUser
	}

	return nil
}

func (r *userRepositoryImp) CreateNewWallet(user *entity.User) (*entity.Wallet, error) {
	newWallet := &entity.Wallet{
		UserId:  user.UserId,
		Balance: 0,
		User:    *user,
	}

	if err := r.db.Create(newWallet).Error; err != nil {
		return nil, httperror.ErrCreateUser
	}

	return newWallet, nil
}
