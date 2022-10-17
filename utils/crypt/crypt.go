package crypt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	math "math/rand"
	"os"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// HashText generates hased password
func HashText(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

// CheckTextHash compares hash and user plain password
func CheckTextHash(text, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
	if err != nil {
		return err == nil
	}
	return true
}

// Jwt generates json web tokens for stateless authentication
func Jwt(claim interface{}) string {
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": claim,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		logrus.Error(err)
	}
	return tokenString
}

// Jwt generates json web tokens for stateless authentication
func JwtWithExpr(claim interface{}, expr time.Duration) string {
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": claim,
		"exp":  time.Now().Add(time.Minute * expr).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		logrus.Error(err)
	}
	return tokenString
}

// GenerateRandomString generates a random string
func GenerateRandomString(length int) string {
	SeededRand := math.New(math.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[SeededRand.Intn(len(charset))]
	}
	return string(b)
}
