package utils

import (
    "github.com/golang-jwt/jwt/v4"
    "time"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
    UserID uint
    jwt.RegisteredClaims
}

func GenerateToken(userID uint) (string, error) {
    claims := Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}