// Code generated by MockGen. DO NOT EDIT.
// Source: internal/users/repository/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	entities "github.com/DoWithLogic/golang-clean-architecture/internal/users/entities"
	repository "github.com/DoWithLogic/golang-clean-architecture/internal/users/repository"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Atomic mocks base method.
func (m *MockRepository) Atomic(ctx context.Context, opt *sql.TxOptions, repo func(repository.Repository) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Atomic", ctx, opt, repo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Atomic indicates an expected call of Atomic.
func (mr *MockRepositoryMockRecorder) Atomic(ctx, opt, repo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Atomic", reflect.TypeOf((*MockRepository)(nil).Atomic), ctx, opt, repo)
}

// GetUserByEmail mocks base method.
func (m *MockRepository) GetUserByEmail(arg0 context.Context, arg1 string) (entities.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(entities.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockRepositoryMockRecorder) GetUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockRepository)(nil).GetUserByEmail), arg0, arg1)
}

// GetUserByID mocks base method.
func (m *MockRepository) GetUserByID(arg0 context.Context, arg1 int64, arg2 ...entities.LockingOpt) (entities.Users, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserByID", varargs...)
	ret0, _ := ret[0].(entities.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockRepositoryMockRecorder) GetUserByID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockRepository)(nil).GetUserByID), varargs...)
}

// IsUserExist mocks base method.
func (m *MockRepository) IsUserExist(ctx context.Context, email string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserExist", ctx, email)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUserExist indicates an expected call of IsUserExist.
func (mr *MockRepositoryMockRecorder) IsUserExist(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserExist", reflect.TypeOf((*MockRepository)(nil).IsUserExist), ctx, email)
}

// SaveNewUser mocks base method.
func (m *MockRepository) SaveNewUser(arg0 context.Context, arg1 entities.Users) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveNewUser", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveNewUser indicates an expected call of SaveNewUser.
func (mr *MockRepositoryMockRecorder) SaveNewUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveNewUser", reflect.TypeOf((*MockRepository)(nil).SaveNewUser), arg0, arg1)
}

// UpdateUserByID mocks base method.
func (m *MockRepository) UpdateUserByID(arg0 context.Context, arg1 entities.UpdateUsers) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserByID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserByID indicates an expected call of UpdateUserByID.
func (mr *MockRepositoryMockRecorder) UpdateUserByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserByID", reflect.TypeOf((*MockRepository)(nil).UpdateUserByID), arg0, arg1)
}

// UpdateUserStatusByID mocks base method.
func (m *MockRepository) UpdateUserStatusByID(arg0 context.Context, arg1 entities.UpdateUserStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserStatusByID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserStatusByID indicates an expected call of UpdateUserStatusByID.
func (mr *MockRepositoryMockRecorder) UpdateUserStatusByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserStatusByID", reflect.TypeOf((*MockRepository)(nil).UpdateUserStatusByID), arg0, arg1)
}
