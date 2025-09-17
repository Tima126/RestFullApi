package auth

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestPasswordHash(t *testing.T) {
	password := "qweasd124!!"

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		t.Fatalf("Ошибка хэширования пароля: %v", err)
	}

	if string(hash) == password {
		t.Fatalf("Хеш совпадает с паролем, это неправильно")
	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		t.Fatalf("Хеш не проходит проверку: %v", err)
	}

}
