package handlers

import (
	jwt "app/jwt_token"
	"net/http"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query().Get("login")

	if login == "" {
		http.Error(w, "Укажите логин", http.StatusBadRequest)
		return
	}

	token, err := jwt.GenerateToken(login, "user", 24*time.Hour)

	if err != nil {
		http.Error(w, "Ошибка генерации токена", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}
