package repositories

import (
	"context"
	"database/sql"

	todoservicev1 "github.com/Kimoto-Norihiro/connect-todo/server/api/todoservice/v1"
)

type ITODORepository interface {
	CreateTODO(ctx context.Context, tx *sql.Tx, title string) error
	ListTODOs(ctx context.Context, db *sql.DB) ([]*todoservicev1.TODO, error)
	UpdateTODO(ctx context.Context, tx *sql.Tx, todo *todoservicev1.TODO) error
	DeleteTODO(ctx context.Context, db *sql.DB, id int32) error
}
