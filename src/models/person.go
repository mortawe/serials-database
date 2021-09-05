package models

import "time"

type Person struct {
	PersonID  int       `json:"person_id" db:"person_id"`
	Name      string    `json:"name" db:"name"`
	Birthdate time.Time `json:"birthdate" db:"birthdate"`
	Bio       string    `json:"bio" db:"bio"`
}

type ExtPerson struct {
	Person
	Shows  []Show  `json:"show" db:"shows"`
	Genres []Genre `json:"genre" db:"genre"`
}
