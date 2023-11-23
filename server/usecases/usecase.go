package usecases

import "context"

type ITODOUsecase interface {
	CreateTODO(ctx context.Context) error
	ListTODOs(ctx context.Context) error
	UpdateTODO(ctx context.Context) error
	DeleteTODO(ctx context.Context) error
}
