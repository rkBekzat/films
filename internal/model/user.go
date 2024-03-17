package model

type User struct {
	Id       string `jsong:"id" db:"id"`
	Username string `json:"name" db:"name"`
	Role     string `json:"role" db:"role"`
	Email    string `json:"email" db:"email"`
}
