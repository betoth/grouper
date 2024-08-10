package secutiry

import (
	"golang.org/x/crypto/bcrypt"
)

func HashSHA256(input string) (string, error) {
	output, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	return string(output), err
}

func VerifyPassword(passwordString, passwordHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordString))
}
