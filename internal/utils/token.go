package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
	Email string `json:"email"`
	ID    int    `json:"id"`
}

const secret = "secret"

// GenerateToken creates a new JWT token with given email and ID.
func GenerateToken(email string, id int) (string, error) {
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Date(2024, 01, 01, 12, 0, 0, 0, time.UTC).Unix(),
		},
		Email: email,
		ID:    id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secret))
	return tokenStr, err
}

// ParseToken parses and validates the JWT token string.
func ParseToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
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

// TokenCheck validates the JWT token and returns the claims if valid.
func TokenCheck(jwtToken string) (map[string]interface{}, error) {
	claims, err := ParseToken(jwtToken)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"email": claims.Email,
		"id":    claims.ID,
	}, nil
}
