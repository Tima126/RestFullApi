package handlers

import (
	"app/middleware"
	"fmt"
	"net/http"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	login := r.Context().Value(middleware.UserCtxKey)

	if login == nil {
		http.Error(w, "пользователь не найден в контексте", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Привет, %s! Это твой профиль ", login)
}
