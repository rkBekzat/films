package model

type Film struct {
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	ReleasedAt  string  `json:"released_at" db:"released_at"`
	Rating      float32 `json:"rating" db:"rating"`
	Actors      []Actor `json:"actors"`
}
