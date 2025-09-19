package jwt

import (
	"app/logger"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// метод генерации токена
func (s *JWTService) GenerateToken(user_id int, login string, role string, ttl time.Duration) (string, error) {
	now := time.Now()

	// создание токена с пользовательскими клеймами
	token, err := s.generateTokenClaims(CustomClaims{
		UserID: user_id,
		Role:   role,
		Login:  login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	})

	if err != nil {
		logger.Log.Errorf("Ошибка при генерации токена: %v", err)
		return "", err
	}
	logger.Log.Debugf("Сгенерированный токен: %s", token)
	return token, nil

}

// вспомогательная функция для генерации токена с клеймами
func (s *JWTService) generateTokenClaims(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.Secret))
}
