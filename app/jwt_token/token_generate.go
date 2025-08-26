package jwt

import (
	"app/logger"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(login, role string, ttl time.Duration) (string, error) {
	now := time.Now()

	logger.Log.Infof("Генерация токена для пользователя: %s с ролью: %s", login, role)

	token, err := generateTokenClaims(CustomClaims{
		Login: login,
		Role:  role,
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

func generateTokenClaims(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Secret_key)
}
