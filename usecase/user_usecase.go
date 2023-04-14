package usecase

import (
	"assignment-golang-backend/dto"
	"assignment-golang-backend/entity"
	"assignment-golang-backend/httperror"
	"assignment-golang-backend/repository"
	"assignment-golang-backend/util"
	"net/mail"
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

	isValid := checkValidEmail(newUser.Email)
	if !isValid {
		return nil, httperror.ErrInvalidRegisterEmail
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

func checkValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
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
