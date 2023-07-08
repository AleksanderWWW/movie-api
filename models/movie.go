package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Year     int    `json:"year"`
}
