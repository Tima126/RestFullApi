package main

import (
	"app/jwt"
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/login", loginHandler)
	http.Handle("/profile", jwt.Middleware(http.HandlerFunc(profiler)))
	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query().Get("login")

	if login == "" {
		http.Error(w, "Укажите логин", http.StatusBadRequest)
		return
	}

	token, err := jwt.GenerateToken(login, 24*time.Hour)

	if err != nil {
		http.Error(w, "Ошибка генерации токена", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}

func profiler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Profiler endpoint"))
}
