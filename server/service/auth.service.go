package service

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func AuthenticateAgainstHash(password string, hashedPassword string) bool {
	byteHash := []byte(hashedPassword)
	bytePassword := []byte(password)

	if err := bcrypt.CompareHashAndPassword(byteHash, bytePassword); err != nil {
		return false
	}

	return true
}
