package entities

import "time"

// Author represents an author with a first name, last name, birth date, and books.
type Author struct {
	Id          string    `json:"_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	BirthDate   time.Time `json:"birth_date"`
	Books       []Book    `json:"books"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
	CreatedBy   int       `json:"created_by"`
	UpdatedBy   int       `json:"updated_by"`
	IsActive    bool      `json:"isactive"`
}
