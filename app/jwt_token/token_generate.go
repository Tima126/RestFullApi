package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var Secret_key = []byte("12345qwert")

func GenerateToken(login, role string, ttl time.Duration) (string, error) {
	now := time.Now()

	return generateTokenClaims(CustomClaims{
		Login: login,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	})

}

func generateTokenClaims(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Secret_key)
}
