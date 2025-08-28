package jwt

import "github.com/golang-jwt/jwt/v5"

var Secret_key = []byte("12345qwert")

// Общий тип Claims -> jwt
type CustomClaims struct {
	UserID int `json:"user_id"`
	Role   int `json:"role,omitempty"`
	jwt.RegisteredClaims
}
