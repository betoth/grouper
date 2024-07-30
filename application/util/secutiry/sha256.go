package util

import "golang.org/x/crypto/bcrypt"

func HashSHA256(input string) (string, error) {

	output, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)

	return string(output), err
}

// VerifyPassword compare password hash with a string
func VerifyPassword(passwordString, passwordHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordString))
}
