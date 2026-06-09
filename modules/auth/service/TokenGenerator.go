package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(
	userId string,
	username string,
) (string, error) {
	claims := jwt.MapClaims{
		"sub":      userId,
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(JwtSecret)
}

func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)

	return err == nil
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (any, error) {
			return JwtSecret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}
