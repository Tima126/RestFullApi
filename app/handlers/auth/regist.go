package auth

import (
	"app/db"
	jwt "app/jwt_token"
	"app/logger"
	"app/models"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Log.Warnf("Invalid request payload: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	logger.Log.Infof("Регистрация нового пользователя: %s", req.Login)

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Log.Errorf("Ошибка хэширования пароля для %s: %v", req.Login, err)
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Login:     req.Login,
		Password:  string(hash),
		User_Role: "user",
	}

	err = db.Pool.QueryRow(context.Background(),
		"INSERT INTO users (login, password, user_role) VALUES ($1, $2, $3) RETURNING id",
		user.Login, user.Password, user.User_Role).Scan(&user.ID)
	if err != nil {
		logger.Log.Errorf("Ошибка создания пользователя %s: %v", req.Login, err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	logger.Log.Infof("Пользователь %s успешно создан с ID %d", req.Login, user.ID)

	tokenTTL := 24 * time.Hour
	token, err := jwt.GenerateToken(user.ID, user.User_Role, tokenTTL)
	if err != nil {
		logger.Log.Errorf("Ошибка генерации токена для пользователя %s: %v", req.Login, err)
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	logger.Log.Infof("JWT токен сгенерирован для пользователя %s", req.Login)

	resp := map[string]string{"token": token}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
