package jwt

import (
	"app/logger"
	"errors"
	"net/http"
	"strings"
)

func ExtractTokenFromHeader(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")

	if token == "" {
		logger.Log.Warn("Нет токена в заголовке Authorization")
		return "", errors.New("Нет токена в заголовке")
	}

	token = strings.TrimPrefix(token, "Bearer ")
	token = strings.TrimSpace(token)

	return token, nil

}
