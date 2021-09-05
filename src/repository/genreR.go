package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"shows/src/models"
)

type GenreRepo struct {
	db *sqlx.DB
}

func NewGenreRepo(db *sqlx.DB) *GenreRepo {
	return &GenreRepo{db: db}
}

const createGenreQ = `INSERT  INTO genre (name) VALUES (:name) RETURNING *`

func (r *GenreRepo) Create(ctx context.Context, name string) (*models.Genre, error) {
	genre := models.Genre{}
	query, args, err := r.db.BindNamed(createGenreQ, map[string]interface{}{"name": name})
	if err != nil {
		return nil, err
	}
	err = r.db.GetContext(ctx, &genre, query, args...)
	return &genre, err
}

const getAllGenreQ = `SELECT * FROM genre`

func (r *GenreRepo) GetAll(ctx context.Context) ([]models.Genre, error) {
	genre := []models.Genre{}
	err := r.db.SelectContext(ctx, &genre, getAllGenreQ)
	return genre, err
}

