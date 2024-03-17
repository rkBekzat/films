package model

type User struct {
	Id       string `jsong:"id" db:"id"`
	Username string `json:"name" db:"name"`
	Gender   string `json:"gender" db:"gender"`
	Role     string `json:"role" db:"role"`
	Email    string `json:"email" db:"email"`
	Password string `jsong:"password" db:"password"`
}
