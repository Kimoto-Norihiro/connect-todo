package todoservice

import (
	"context"
	"database/sql"

	"connectrpc.com/connect"
	todoservicev1 "github.com/Kimoto-Norihiro/connect-todo/server/api/todoservice/v1"
	"github.com/Kimoto-Norihiro/connect-todo/server/repositories/todorepository"
	"github.com/Kimoto-Norihiro/connect-todo/server/usecases"
	"github.com/Kimoto-Norihiro/connect-todo/server/usecases/todousecase"
)

type TODOService struct {
	usecase usecases.ITODOUsecase
}

func NewTODOService(db *sql.DB) *TODOService {
	repo := todorepository.NewTODORepository(db)
	usecase := todousecase.NewTODOUsecase(db, repo)

	return &TODOService{
		usecase: usecase,
	}
}

func (s *TODOService) CreateTODO(
	ctx context.Context,
	req *connect.Request[todoservicev1.CreateTODORequest],
) (*connect.Response[todoservicev1.CreateTODOResponse], error) {
	res := connect.NewResponse(&todoservicev1.CreateTODOResponse{})
	return res, nil
}

func (s *TODOService) ListTODOs(
	ctx context.Context,
	req *connect.Request[todoservicev1.ListTODOsRequest],
) (*connect.Response[todoservicev1.ListTODOsResponse], error) {
	res := connect.NewResponse(&todoservicev1.ListTODOsResponse{})
	return res, nil
}

func (s *TODOService) UpdateTODO(
	ctx context.Context,
	req *connect.Request[todoservicev1.UpdateTODORequest],
) (*connect.Response[todoservicev1.UpdateTODOResponse], error) {
	res := connect.NewResponse(&todoservicev1.UpdateTODOResponse{})
	return res, nil
}

func (s *TODOService) DeleteTODO(
	ctx context.Context,
	req *connect.Request[todoservicev1.DeleteTODORequest],
) (*connect.Response[todoservicev1.DeleteTODOResponse], error) {
	res := connect.NewResponse(&todoservicev1.DeleteTODOResponse{})
	return res, nil
}
