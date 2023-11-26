package todousecase

import (
	"context"
	"database/sql"

	todoservicev1 "github.com/Kimoto-Norihiro/connect-todo/server/api/todoservice/v1"
	"github.com/Kimoto-Norihiro/connect-todo/server/repositories"
)

type TODOUsecase struct {
	db   *sql.DB
	repo repositories.ITODORepository
}

func NewTODOUsecase(db *sql.DB, repo repositories.ITODORepository) *TODOUsecase {
	return &TODOUsecase{
		db:   db,
		repo: repo,
	}
}

func (u TODOUsecase) CreateTODO(ctx context.Context, title string) (err error) {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = u.repo.CreateTODO(ctx, tx, title)
	if err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
	}
	return nil
}

func (u TODOUsecase) ListTODOs(ctx context.Context) ([]*todoservicev1.TODO, error) {
	todos, err := u.repo.ListTODOs(ctx, u.db)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (u TODOUsecase) UpdateTODO(ctx context.Context, todo *todoservicev1.TODO) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = u.repo.UpdateTODO(ctx, tx, todo)
	if err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
	}
	return nil
}

func (u TODOUsecase) DeleteTODO(ctx context.Context, id int32) error {
	err := u.repo.DeleteTODO(ctx, u.db, id)
	if err != nil {
		return err
	}
	return nil
}
