package usecases

import (
	"context"

	todoservicev1 "github.com/Kimoto-Norihiro/connect-todo/server/api/todoservice/v1"
)

//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE

type ITODOUsecase interface {
	CreateTODO(ctx context.Context, title string) error
	ListTODOs(ctx context.Context) ([]*todoservicev1.TODO, error)
	UpdateTODO(ctx context.Context, todo *todoservicev1.TODO) error
	DeleteTODO(ctx context.Context, id int32) error
}
