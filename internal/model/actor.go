package model

type Actor struct {
	Name      string `json:"name" db:"name"`
	Gender    string `json:"gender" db:"gender"`
	BirthDate string `json:"birth_date" db:"birth_date"`
}
