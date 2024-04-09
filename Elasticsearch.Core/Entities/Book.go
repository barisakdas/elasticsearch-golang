package entities

import "time"

type Book struct {
	Id          string    `json:"_id"`
	Title       string    `json:"title"`
	Abstract    string    `json:"abstract"`
	PublishDate time.Time `json:"publish_date"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Categories  []string  `json:"categories"`
	Author      Author    `json:"author"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
	CreatedBy   int       `json:"created_by"`
	UpdatedBy   int       `json:"updated_by"`
	IsActive    bool      `json:"isactive"`
}
