package handler

type signUpInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Gender   string `jsong:"gender" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type actorDto struct {
	Name      string `json:"name" db:"name"`
	Gender    string `json:"gender" db:"gender"`
	BirthDate string `json:"birth_date" db:"birth_date"`
}
