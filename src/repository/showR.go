package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"shows/src/models"
	"shows/src/models/search"
)

type ShowRepo struct {
	db *sqlx.DB
}

func NewShowRepo(db *sqlx.DB) *ShowRepo {
	return &ShowRepo{db: db}
}

const insertShowQuery = `INSERT INTO show (title, release, description, genre_id) VALUES (
	:title, :release, :description, :genre_id) RETURNING *`

func (r *ShowRepo) Create(ctx context.Context, s *models.Show) (*models.Show, error) {
	show := models.Show{}
	query, args, err := r.db.BindNamed(insertShowQuery, s)
	if err != nil {
		return nil, err
	}
	err = r.db.GetContext(ctx, &show, query, args...)
	return &show, err
}

const updateShowQuery = `UPDATE show SET title=:title, release=:release, description=:description, genre_id=:genre_id WHERE show_id = :show_id
RETURNING *`

func (r *ShowRepo) Update(ctx context.Context, s *models.Show) (*models.Show, error) {
	show := models.Show{}
	query, args, err := r.db.BindNamed(updateShowQuery, s)
	if err != nil {
		return nil, err
	}
	err = r.db.GetContext(ctx, &show, query, args...)
	return &show, err
}

const (
	findShowQueryBase = `SELECT show_id, title, release, description, episode_num, genre_id FROM show `
)

// todo maxpage
func (r *ShowRepo) Find(ctx context.Context, name string, sort search.Sort) ([]models.Show, error) {
	dest := &[]models.Show{}
	arguments := map[string]interface{}{}
	query, args, err := r.db.BindNamed(findShowQueryBase+sort.ToSQL(), arguments)
	if err != nil {
		return nil, err
	}
	err = r.db.SelectContext(ctx, dest, query, args...)
	return *dest, err
}

const (
	getShowsByPersonIDQuery = `SELECT show.* FROM person_show JOIN show ON person_show.show_id = show.show_id 
WHERE person_id = :id`
)

func (r *ShowRepo) ShowsByPersonID(ctx context.Context, id int) ([]models.Show, error) {
	dest := &[]models.Show{}
	query, args, err := r.db.BindNamed(getShowsByPersonIDQuery, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	err = r.db.SelectContext(ctx, dest, query, args...)
	return *dest, err
}

const addPersonQuery = `INSERT INTO person_show (person_id, show_id) VALUES (:person_id, :show_id)`

func (r *ShowRepo) AddPerson(ctx context.Context, personID int, showID int) error {
	query, args, err := r.db.BindNamed(addPersonQuery, map[string]interface{}{"person_id": personID,
		"show_id": showID})
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

const getShowQ = `SELECT * FROM show WHERE show_id = :id`

func (r *ShowRepo) Get(ctx context.Context, id int) (*models.Show, error) {
	show := models.Show{}
	query, args, err := r.db.BindNamed(getShowQ, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	err = r.db.GetContext(ctx, &show, query, args...)
	return &show, err
}

const deletePersonQ = `DELETE  FROM person_show WHERE show_id = :id`

func (r *ShowRepo) DeletePersonsFromShow(ctx context.Context, id int) error {
	query, args, err := r.db.BindNamed(deletePersonQ, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query, args...)
	return err
}

const getByGenreQ = `SELECT show.* FROM show  WHERE genre_id = :id`

func (r *ShowRepo) GetByGenre(ctx context.Context, genreID int) ([]models.Show, error) {
	dest := []models.Show{}
	query, args, err := r.db.BindNamed(deletePersonQ, map[string]interface{}{"id": genreID})
	if err != nil {
		return nil, err
	}
	err = r.db.SelectContext(ctx, &dest, query, args...)
	return dest, err
}
