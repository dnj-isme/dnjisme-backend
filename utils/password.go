package utils

import (
	"dnj-backend/config"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const DefaultCost = bcrypt.DefaultCost

func passwordWithSalt(raw string) string {
	return raw + config.LoadEnv().AuthConfig.Salt
}

func HashPassword(raw string) (string, error) {
	if len(raw) == 0 {
		return "", errors.New("password cannot be empty")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(passwordWithSalt(raw)), DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func VerifyPassword(raw, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(passwordWithSalt(raw)))
}

func IsValidPassword(raw, hashed string) bool {
	return VerifyPassword(raw, hashed) == nil
}

func PasswordValidator(raw string) error {
	if len(raw) < 8 {
		return errors.New("Password must be at least 8 characters long")
	}

	return nil
}