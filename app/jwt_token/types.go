package jwt

import (
	"app/config"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	Secret string
}

func NewJWTService(cfg *config.Config) *JWTService {
	return &JWTService{Secret: cfg.JWT_Secret}
}

// структура для пользовательских клеймов
type CustomClaims struct {
	UserID int    `json:"user_id"`
	Login  string `json:"login"`
	Role   string `json:"role,omitempty"`

	jwt.RegisteredClaims
}
