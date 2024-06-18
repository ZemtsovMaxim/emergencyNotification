package models

type Contact struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Telegram string `json:"telegram"`
	Email    string `json:"email"`
}
