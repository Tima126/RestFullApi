package models

// Модель пользователя
type User struct {
	ID        int
	Login     string
	Password  string
	User_Role string
}
