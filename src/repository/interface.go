package repository

import (
	"context"

	"shows/src/models"
	"shows/src/models/search"
)

type IShowR interface {
	Create(ctx context.Context, show *models.Show) (*models.Show, error)
	Update(ctx context.Context, show *models.Show) (*models.Show, error)
	Find(ctx context.Context, name string, sort search.Sort) ([]models.Show, error)
	ShowsByPersonID(ctx context.Context, id int) ([]models.Show, error)
	AddPerson(ctx context.Context, personID int, showID int) error
	Get(ctx context.Context, id int) (*models.Show, error)
	DeletePersonsFromShow(ctx context.Context, id int) error
	GetByGenre(ctx context.Context, genreID int) ([]models.Show, error)
}

type IPersonR interface {
	Create(ctx context.Context, show *models.Person) (*models.Person, error)
	Update(ctx context.Context, show *models.Person) (*models.Person, error)
	Find(ctx context.Context, name string, sort search.Sort) ([]models.Person, error)
	Get(ctx context.Context, id int) (*models.Person, error)
	GetByShowID(ctx context.Context, id int) ([]models.Person, error)
}

type IGenreR interface {
	Create(ctx context.Context, name string) (*models.Genre, error)
	GetAll(ctx context.Context) ([]models.Genre, error)
}
