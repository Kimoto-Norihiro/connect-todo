package todorepository

import (
	"database/sql"
	"errors"
)

type TODORepository struct {
	db *sql.DB
}

func NewTODORepository(db *sql.DB) *TODORepository {
	return &TODORepository{
		db: db,
	}
}

func(r TODORepository) CreateTODO() error {
	return errors.New("not implemented")
}

func(r TODORepository) ListTODOs() error {
	return errors.New("not implemented")
}

func(r TODORepository) UpdateTODO() error {
	return errors.New("not implemented")
}

func(r TODORepository) DeleteTODO() error {
	return errors.New("not implemented")
}
