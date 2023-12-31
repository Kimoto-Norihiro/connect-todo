// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock_usecases is a generated GoMock package.
package mock_usecases

import (
	context "context"
	reflect "reflect"

	todoservicev1 "github.com/Kimoto-Norihiro/connect-todo/server/api/todoservice/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockITODOUsecase is a mock of ITODOUsecase interface.
type MockITODOUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockITODOUsecaseMockRecorder
}

// MockITODOUsecaseMockRecorder is the mock recorder for MockITODOUsecase.
type MockITODOUsecaseMockRecorder struct {
	mock *MockITODOUsecase
}

// NewMockITODOUsecase creates a new mock instance.
func NewMockITODOUsecase(ctrl *gomock.Controller) *MockITODOUsecase {
	mock := &MockITODOUsecase{ctrl: ctrl}
	mock.recorder = &MockITODOUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITODOUsecase) EXPECT() *MockITODOUsecaseMockRecorder {
	return m.recorder
}

// CreateTODO mocks base method.
func (m *MockITODOUsecase) CreateTODO(ctx context.Context, title string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTODO", ctx, title)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTODO indicates an expected call of CreateTODO.
func (mr *MockITODOUsecaseMockRecorder) CreateTODO(ctx, title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTODO", reflect.TypeOf((*MockITODOUsecase)(nil).CreateTODO), ctx, title)
}

// DeleteTODO mocks base method.
func (m *MockITODOUsecase) DeleteTODO(ctx context.Context, id int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTODO", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTODO indicates an expected call of DeleteTODO.
func (mr *MockITODOUsecaseMockRecorder) DeleteTODO(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTODO", reflect.TypeOf((*MockITODOUsecase)(nil).DeleteTODO), ctx, id)
}

// ListTODOs mocks base method.
func (m *MockITODOUsecase) ListTODOs(ctx context.Context) ([]*todoservicev1.TODO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTODOs", ctx)
	ret0, _ := ret[0].([]*todoservicev1.TODO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTODOs indicates an expected call of ListTODOs.
func (mr *MockITODOUsecaseMockRecorder) ListTODOs(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTODOs", reflect.TypeOf((*MockITODOUsecase)(nil).ListTODOs), ctx)
}

// UpdateTODO mocks base method.
func (m *MockITODOUsecase) UpdateTODO(ctx context.Context, todo *todoservicev1.TODO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTODO", ctx, todo)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTODO indicates an expected call of UpdateTODO.
func (mr *MockITODOUsecaseMockRecorder) UpdateTODO(ctx, todo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTODO", reflect.TypeOf((*MockITODOUsecase)(nil).UpdateTODO), ctx, todo)
}
