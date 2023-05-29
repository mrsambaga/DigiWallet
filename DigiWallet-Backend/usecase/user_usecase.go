package usecase

import (
	"assignment-golang-backend/dto"
	"assignment-golang-backend/entity"
	"assignment-golang-backend/httperror"
	"assignment-golang-backend/repository"
	"assignment-golang-backend/util"
	"errors"
	"fmt"
	"regexp"
)

type UsersUsecase interface {
	Register(newUserDTO *dto.RegisterRequestDTO) (*dto.RegisterResponseDTO, error)
	Login(loginUserDTO *dto.LoginRequestDTO) (*dto.TokenResponse, error)
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

	err := checkValidEmail2(newUser.Email)
	if err != nil {
		return nil, httperror.ErrInvalidRegisterEmail
	}

	err = checkValidPassword(newUser.Password)
	if err != nil {
		return nil, errors.New("invalid password : must be more than 8 char")
	}

	hashedPassword, err := util.HashPassword(newUser.Password)
	if err != nil {
		return nil, err
	}
	newUser.Password = hashedPassword

	err = u.usersRepository.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	newWallet, err := u.usersRepository.CreateNewWallet(newUser)
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

// func checkValidEmail(email string) error {
// 	_, err := mail.ParseAddress(email)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func checkValidEmail2(email string) error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

func checkValidPassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}
	return nil
}

func (u *usersUsecaseImp) Login(loginUserDTO *dto.LoginRequestDTO) (*dto.TokenResponse, error) {
	loginUser := &entity.User{
		Email:    loginUserDTO.Email,
		Password: loginUserDTO.Password,
	}

	registeredUser, err := u.usersRepository.GetUserByEmail(loginUser.Email)
	if err != nil {
		return nil, err
	}

	ok := util.ComparePassword(registeredUser.Password, loginUser.Password)
	if !ok {
		return nil, httperror.ErrInvalidEmailPassword
	}

	loginUser.UserId = registeredUser.UserId
	token, err := util.GenerateAccessToken(loginUser)
	if err != nil {
		return nil, err
	}

	return token, nil
}
