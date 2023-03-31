package httperror

import (
	"errors"
)

var (
	ErrEmailAlreadyRegistered = errors.New("email already registered")
	ErrGenerateHash           = errors.New("failed to generate hash")
	ErrCreateUser             = errors.New("failed to create user")
	ErrInvalidEmail           = errors.New("invalid email")
	ErrUserNotExist           = errors.New("user not exist")
	ErrFailedCreateToken      = errors.New("failed to creae token")
	ErrInvalidPassword        = errors.New("invalid password")
)
