package login

import (
	"app/db"
	jwt "app/jwt_token"
	"app/logger"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func LoginHandler(jwtService *jwt.JWTService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var err error
		var req LoginRequest
		var UserID int
		var Login string
		var hashpasswordDb string

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		logger.Log.Infof("Пользователь %s пытается войти в систему", req.Login)

		err = db.Pool.QueryRow(context.Background(),
			"SELECT id, login, password FROM users WHERE login=$1",
			req.Login).Scan(&UserID, &Login, &hashpasswordDb)

		if err != nil {
			logger.Log.Errorf("Пользователь %s не найден: %v", req.Login, err)
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		logger.Log.Infof("Пользователь найденн с логином: %s, ID: %d ", Login, UserID)

		err = bcrypt.CompareHashAndPassword([]byte(hashpasswordDb), []byte(req.Password))
		if err != nil {
			logger.Log.Warnf("Неверный пароль для пользователя %s", req.Login)
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		logger.Log.Infof("Пароль подтверждён для пользователя %s", req.Login)

		tokenTTL := 24 * time.Hour

		token, err := jwtService.GenerateToken(UserID, "user", tokenTTL)

		if err != nil {
			logger.Log.Errorf("Ошибка генерации токена для пользователя %s: %v", req.Login, err)
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		logger.Log.Infof("JWT токен сгенерирован для пользователя %s", req.Login)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": token})

	}
}
