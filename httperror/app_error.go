package httperror

import (
	"errors"
)

var (
	ErrEmailAlreadyRegistered = errors.New("email already registered")
	ErrGenerateHash = errors.New("failed to generate hash")
	ErrCreateUser = errors.New("failed to create user")
	ErrEmailNotRegistered = errors.New("email not registered")
	ErrUserNotExist = errors.New("user not exist")
)