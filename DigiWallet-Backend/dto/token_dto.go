package dto

import "github.com/golang-jwt/jwt/v4"

type IdTokenClaims struct {
	UserID uint64 `json:"id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

type TokenResponse struct {
	Token string `json:"token"`
}
