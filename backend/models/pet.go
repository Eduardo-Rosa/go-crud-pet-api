package models

import "time"

type Pet struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Species   string    `json:"species"`
	Breed     string    `json:"breed"`
	Age       int       `json:"age"`
	BirthDate time.Time `json:"birth_date"`
	OwnerName string    `json:"owner_name"`
}
