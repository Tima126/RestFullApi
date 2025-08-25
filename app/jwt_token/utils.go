package jwt

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func ExtractTokenFromHeader(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	fmt.Println("Token from header:", token)
	if token == "" {
		return "", errors.New("Нет токена в заголовке")
	}

	token = strings.TrimPrefix(token, "Bearer ")
	token = strings.TrimSpace(token)

	if token == "" {
		return "", errors.New("пустой токен после Bearer")
	}

	return token, nil

}
