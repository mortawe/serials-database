package models

type Genre struct {
	GenreID int    `json:"genre_id" db:"genre_id"`
	Name    string `json:"name" db:"name"`
}
