package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/phongloihong/user_framework/db/types"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password []byte) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedByte), nil
}

func CompareHashWithPassword(hashedPassord, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassord, password)
}

func GenerateToken(user types.TokenClaims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
