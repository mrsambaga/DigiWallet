package repository

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/httperror"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		return nil, httperror.ErrGenerateHash
	}
	newUser.Password = string(hashedPassword)

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

	r.db.Order("wallet_number desc").First(&lastWallet)
	if lastWallet == nil {
		walletNumber = 123000000000
	}

	walletNumber = lastWallet.WalletNumber + 1

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
		return "", httperror.ErrEmailNotRegistered
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": loginUser.UserId,
		"exp": time.Now().Add(time.Minute * 20).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_API")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
