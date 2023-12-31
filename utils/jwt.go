package utils

import (
	"megachasma/config"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJwt(userID string) (string, error) {
	jwtSecret := config.LoadEnv().JwtEnv.JwtSecret

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"sub": userID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	token, err := t.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateJwt(token string) (*jwt.Token, error) {
	jwtSecret := config.LoadEnv().JwtEnv.JwtSecret
	return jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
}
