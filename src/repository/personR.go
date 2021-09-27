package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"shows/src/models"
	"shows/src/models/search"
)

type PersonRepo struct {
	db *sqlx.DB
}

func NewPersonRepo(db *sqlx.DB) *PersonRepo {
	return &PersonRepo{db: db}
}

const insertPersonQuery = `INSERT INTO person (name, birthdate, bio, awards) VALUES (
	:name, :birthdate, :bio, :awards) RETURNING *`

func (r *PersonRepo) Create(ctx context.Context, s *models.Person) (*models.Person, error) {
	show := models.Person{}
	query, args, err := r.db.BindNamed(insertPersonQuery, s)
	if err != nil {
		return nil, err
	}
	err = r.db.GetContext(ctx, &show, query, args...)
	return &show, err
}

const updatePersonQuery = `UPDATE person SET name=:name, birthdate=:birthdate, bio=:bio, awards=:awards WHERE person_id = :person_id`

func (r *PersonRepo) Update(ctx context.Context, s *models.Person) (*models.Person, error) {
	show := models.Person{}
	query, args, err := r.db.BindNamed(updatePersonQuery, s)
	if err != nil {
		return nil, err
	}
	_, err = r.db.ExecContext(ctx, query, args...)
	return &show, err
}

const (
	findPersonQueryBase = `SELECT * FROM person `
)

func (r *PersonRepo) Find(ctx context.Context, person search.Person, sort search.Sort) ([]models.Person, error) {
	dest := &[]models.Person{}
	person.Name = "%" + person.Name + "%"
	query, args, err := r.db.BindNamed(findPersonQueryBase+person.ToSQL()+sort.ToSQL(), person)
	if err != nil {
		return nil, err
	}
	err = r.db.SelectContext(ctx, dest, query, args...)
	return *dest, err
}

const getPersonQ = `SELECT * FROM person WHERE person_id = :id`

func (r *PersonRepo) Get(ctx context.Context, id int) (*models.Person, error) {
	person := models.Person{}
	query, args, err := r.db.BindNamed(getPersonQ, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	err = r.db.GetContext(ctx, &person, query, args...)
	return &person, err
}

const getPersonByShowID = `SELECT person.* FROM person JOIN person_show ON person.person_id = person_show.person_id WHERE show_id = :id`

func (r *PersonRepo) GetByShowID(ctx context.Context, id int) ([]models.Person, error) {
	dest := &[]models.Person{}
	query, args, err := r.db.BindNamed(getPersonByShowID, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return nil, err
	}
	err = r.db.SelectContext(ctx, dest, query, args...)
	return *dest, err
}

const deletePersonQ = `DELETE FROM person WHERE person_id = :id`

func (r *PersonRepo) Delete(ctx context.Context, id int) error {
	query, args, err := r.db.BindNamed(deletePersonQ, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query, args...)
	return err
}
