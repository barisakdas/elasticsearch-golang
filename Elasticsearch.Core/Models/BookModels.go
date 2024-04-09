package models

import "time"

type CreateBookModel struct {
	Title       string    `json:"title"`
	Abstract    string    `json:"abstract"`
	PublishDate time.Time `json:"publish_date"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Categories  []string  `json:"categories"`
}

type UpdateBookModel struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Abstract    string    `json:"abstract"`
	PublishDate time.Time `json:"publish_date"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Categories  []string  `json:"categories"`
}

type SearchBookModel struct {
	Title            string    `json:"title"`
	Abstract         string    `json:"abstract"`
	PublishDateStart time.Time `json:"publish_date_start"`
	MinPrice         float64   `json:"min_price"`
	MinStock         int       `json:"min_stock"`
}
