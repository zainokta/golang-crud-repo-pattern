package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId int64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	token, err := sign.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
