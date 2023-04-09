package httperror

import (
	"errors"
)

var (
	ErrEmailAlreadyRegistered = errors.New("email already registered")
	ErrGenerateHash           = errors.New("failed to generate hash")
	ErrCreateUser             = errors.New("failed to create user")
	ErrInvalidEmailPassword   = errors.New("invalid email or password")
	ErrUserNotExist           = errors.New("user not exist")
	ErrFailedCreateToken      = errors.New("failed to creae token")
	ErrInvalidRegisterEmail   = errors.New("invalid email")
)
