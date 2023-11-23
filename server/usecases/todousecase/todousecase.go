package todousecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Kimoto-Norihiro/connect-todo/server/repositories"
)

type TODOUsecase struct {
	db   *sql.DB
	repo repositories.ITODORepository
}

func NewTODOUsecase(db *sql.DB, repo repositories.ITODORepository) *TODOUsecase {
	return &TODOUsecase{
		db: db,
		repo: repo,
	}
}

func (r TODOUsecase) CreateTODO(ctx context.Context,) error {
	return errors.New("not implemented")
}

func (r TODOUsecase) ListTODOs(ctx context.Context,) error {
	return errors.New("not implemented")
}

func (r TODOUsecase) UpdateTODO(ctx context.Context,) error {
	return errors.New("not implemented")
}

func (r TODOUsecase) DeleteTODO(ctx context.Context,) error {
	return errors.New("not implemented")
}
