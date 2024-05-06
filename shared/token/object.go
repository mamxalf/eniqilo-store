package token

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type UserData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}

type GenerateTokenParams struct {
	AccessTokenSecret string
	AccessTokenExpiry time.Duration
}

type Token struct {
	AccessToken string `json:"accessToken"`
}

// JWTToken is
type JWTToken struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	OwnerID string `json:"ownerID"`
	jwt.StandardClaims
}

// JWTVerifyEmail is
type JWTVerifyEmail struct {
	jwt.StandardClaims
	Email string `json:"email"`
}
