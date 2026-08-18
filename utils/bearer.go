package utils

import (
	"errors"
	"strings"
	"time"

	"dnj-backend/config"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateBearerToken(userID uint, email, role string) (string, error) {
	secret := config.LoadEnv().AuthConfig.JWTSecret
	if secret == "" {
		return "", errors.New("jwt secret is empty")
	}

	claims := TokenClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return "Bearer " + signed, nil
}

func ValidateBearerToken(bearerToken string) (*TokenClaims, error) {
	secret := config.LoadEnv().AuthConfig.JWTSecret
	if secret == "" {
		return nil, errors.New("jwt secret is empty")
	}

	tokenString := strings.TrimSpace(bearerToken)
	if tokenString == "" {
		return nil, errors.New("token is empty")
	}

	if after, ok :=strings.CutPrefix(tokenString, "Bearer "); ok  {
		tokenString = after
	}

	claims := &TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method == nil || token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
