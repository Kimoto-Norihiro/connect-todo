package todorepository

import (
	"context"
	"database/sql"
	"errors"

	todoservicev1 "github.com/Kimoto-Norihiro/connect-todo/server/api/todoservice/v1"
)

type TODORepository struct {
	db *sql.DB
}

func NewTODORepository(db *sql.DB) *TODORepository {
	return &TODORepository{
		db: db,
	}
}

func(r TODORepository) CreateTODO(ctx context.Context, tx *sql.Tx, title string) error {
	return errors.New("not implemented")
}

func(r TODORepository) ListTODOs(ctx context.Context, db *sql.DB) ([]*todoservicev1.TODO ,error) {
	return nil, errors.New("not implemented")
}

func(r TODORepository) UpdateTODO(ctx context.Context, tx *sql.Tx, todo *todoservicev1.TODO) error {
	return errors.New("not implemented")
}

func(r TODORepository) DeleteTODO(ctx context.Context, db *sql.DB, id int32) error {
	return errors.New("not implemented")
}
