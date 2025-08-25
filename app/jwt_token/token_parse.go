package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func ParseToken(tokenString string) (*CustomClaims, error) {

	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return Secret_key, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid || claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("Invalid Token")
	}

	return claims, nil

}
