package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
    accessSecret  = []byte("ACCESS_SECRET_KEY")
    refreshSecret = []byte("REFRESH_SECRET_KEY")
)


type AccessClaims struct {
    UserID string `json:"sub"`
    Email  string `json:"email"`
    jwt.RegisteredClaims
}



func GenerateAccessToken(userID, email, role string) (string, error) {
    claims := AccessClaims{
        UserID: userID,
        Email:  email,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(accessSecret)
}


func GenerateRefreshToken(userID string, tokenID string) (string, error) {
    claims := jwt.RegisteredClaims{
        Subject:   userID,
        ID:        tokenID,
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
        IssuedAt:  jwt.NewNumericDate(time.Now()),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(refreshSecret)
}

