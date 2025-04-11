package auth

import (
	"avito-pvz/internal/config"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken generates and returns JWT-token with "email" and "role" claims
func GenerateToken(email string, role string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"role":  role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWT.SecretKey))
}

// ParseToken parse tokenString including checking for valid and returns its claims
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.AppConfig.JWT.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
