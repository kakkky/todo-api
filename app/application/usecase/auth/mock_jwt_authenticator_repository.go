// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface_jwt_authenticator_repository.go
//
// Generated by this command:
//
//	mockgen -package=auth -source=./interface_jwt_authenticator_repository.go -destination=./mock_jwt_authenticator_repository.go
//

// Package auth is a generated GoMock package.
package auth

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockJwtAuthenticatorRepository is a mock of JwtAuthenticatorRepository interface.
type MockJwtAuthenticatorRepository struct {
	ctrl     *gomock.Controller
	recorder *MockJwtAuthenticatorRepositoryMockRecorder
	isgomock struct{}
}

// MockJwtAuthenticatorRepositoryMockRecorder is the mock recorder for MockJwtAuthenticatorRepository.
type MockJwtAuthenticatorRepositoryMockRecorder struct {
	mock *MockJwtAuthenticatorRepository
}

// NewMockJwtAuthenticatorRepository creates a new mock instance.
func NewMockJwtAuthenticatorRepository(ctrl *gomock.Controller) *MockJwtAuthenticatorRepository {
	mock := &MockJwtAuthenticatorRepository{ctrl: ctrl}
	mock.recorder = &MockJwtAuthenticatorRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJwtAuthenticatorRepository) EXPECT() *MockJwtAuthenticatorRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockJwtAuthenticatorRepository) Delete(ctx context.Context, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockJwtAuthenticatorRepositoryMockRecorder) Delete(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockJwtAuthenticatorRepository)(nil).Delete), ctx, userID)
}

// Load mocks base method.
func (m *MockJwtAuthenticatorRepository) Load(ctx context.Context, userID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", ctx, userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Load indicates an expected call of Load.
func (mr *MockJwtAuthenticatorRepositoryMockRecorder) Load(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockJwtAuthenticatorRepository)(nil).Load), ctx, userID)
}

// Save mocks base method.
func (m *MockJwtAuthenticatorRepository) Save(ctx context.Context, duration time.Duration, userID, jwtID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, duration, userID, jwtID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockJwtAuthenticatorRepositoryMockRecorder) Save(ctx, duration, userID, jwtID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockJwtAuthenticatorRepository)(nil).Save), ctx, duration, userID, jwtID)
}
