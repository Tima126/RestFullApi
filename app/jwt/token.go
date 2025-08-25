package jwt

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret_key = []byte("12345qwert")



func GenerateToken(login, role string, ttl time.Duration) (string, error) {
	now := time.now()

	return generateTokenClaims(CustomClaims{
		Login: login,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})

}


func generateTokenClaims(claims CustomClaims) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret_key)
}

