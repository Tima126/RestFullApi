package auth

/*
***********************************************
тестовоя реализациия регистрации пользователья
***********************************************
*/

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

// структура для мапинга json с полями структуры
type RegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// основной метод дял регистрации
func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var req RegisterRequest // инициализация сруктуры для мапинга json

	// декодирование json в структуру

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Log.Warnf("Invalid request payload: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	logger.Log.Infof("Регистрация нового пользователя: %s", req.Login)

	// генерациия хеш паролья для полученного пароля
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Log.Errorf("Ошибка хэширования пароля для %s: %v", req.Login, err)
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// зополнение структуры пользователя
	user := models.User{
		Login:     req.Login,
		Password:  string(hash),
		User_Role: "user",
	}

	// добавление в бд пользователя
	err = db.Pool.QueryRow(context.Background(),
		"INSERT INTO users (login, password, user_role) VALUES ($1, $2, $3) RETURNING id",
		user.Login, user.Password, user.User_Role).Scan(&user.ID)

	if err != nil {
		logger.Log.Errorf("Ошибка создания пользователя %s: %v", req.Login, err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	logger.Log.Infof("Пользователь %s успешно создан с ID %d", req.Login, user.ID)

	// генерация JWT токена для нового пользователя
	tokenTTL := 24 * time.Hour
	token, err := jwt.GenerateToken(user.ID, user.User_Role, tokenTTL)
	if err != nil {
		logger.Log.Errorf("Ошибка генерации токена для пользователя %s: %v", req.Login, err)
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	logger.Log.Infof("JWT токен сгенерирован для пользователя %s", req.Login)

	// возврат токена в ответе
	resp := map[string]string{"token": token}
	// установка заголовков и код статуса
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// кодирование ответа в json и отправка
	json.NewEncoder(w).Encode(resp)
}
