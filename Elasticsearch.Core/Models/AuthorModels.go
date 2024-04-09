package models

import "time"

type CreateAuthorModel struct {
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	BirthDate time.Time `json:"birth-date"`
}

type UpdateAuthorModel struct {
	Id        string    `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	BirthDate time.Time `json:"birth-date"`
}
