package model

type Film struct {
	Id           string  `json:"id" db:"id"`
	Title        string  `json:"title" db:"title"`
	Description  string  `json:"description" db:"description"`
	ReleasedDate string  `json:"released_at" db:"released_at"`
	Rating       float32 `json:"rating" db:"rating"`
	Actors       []Actor `json:"actors"`
}
