package usecase

import (
	"assignment-golang-backend/dto"
	"assignment-golang-backend/entity"
	"assignment-golang-backend/httperror"
	"assignment-golang-backend/repository"
	"regexp"
)

type UsersUsecase interface {
	Register(newUserDTO *dto.RegisterRequestDTO) (*dto.RegisterResponseDTO, error)
	Login(loginUserDTO *dto.LoginRequestDTO) (string, error)
}

type usersUsecaseImp struct {
	usersRepository repository.UsersRepository
}

type UsersUsecaseConfig struct {
	UsersRepository repository.UsersRepository
}

func NewUsersUsecase(cfg *UsersUsecaseConfig) UsersUsecase {
	return &usersUsecaseImp{
		usersRepository: cfg.UsersRepository,
	}
}

func (u *usersUsecaseImp) Register(newUserDTO *dto.RegisterRequestDTO) (*dto.RegisterResponseDTO, error) {
	newUser := &entity.User{
		Name:     newUserDTO.Name,
		Email:    newUserDTO.Email,
		Password: newUserDTO.Password,
	}

	validEmail := checkValidEmail(newUser.Email)
	if !validEmail {
		return nil, httperror.ErrInvalidRegisterEmail
	}

	newWallet, err := u.usersRepository.Register(newUser)
	if err != nil {
		return nil, err
	}

	out := &dto.RegisterResponseDTO{}
	out.UserId = newWallet.UserId
	out.Name = newWallet.User.Name
	out.Email = newWallet.User.Email
	out.Password = newUser.Password
	out.WalletNumber = newWallet.WalletNumber
	out.Balance = newWallet.Balance

	return out, nil
}

func checkValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9]+@gmail\.com$`
	valid, _ := regexp.MatchString(pattern, email)
	return valid
}

func (u *usersUsecaseImp) Login(loginUserDTO *dto.LoginRequestDTO) (string, error) {
	loginUser := &entity.User{
		Email:    loginUserDTO.Email,
		Password: loginUserDTO.Password,
	}

	tokenString, err := u.usersRepository.Login(loginUser)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
