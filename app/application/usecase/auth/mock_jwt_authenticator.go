// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface_jwt_authenticator.go
//
// Generated by this command:
//
//	mockgen -package=auth -source=./interface_jwt_authenticator.go -destination=./mock_jwt_authenticator.go
//

// Package auth is a generated GoMock package.
package auth

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockJwtAuthenticator is a mock of JwtAuthenticator interface.
type MockJwtAuthenticator struct {
	ctrl     *gomock.Controller
	recorder *MockJwtAuthenticatorMockRecorder
	isgomock struct{}
}

// MockJwtAuthenticatorMockRecorder is the mock recorder for MockJwtAuthenticator.
type MockJwtAuthenticatorMockRecorder struct {
	mock *MockJwtAuthenticator
}

// NewMockJwtAuthenticator creates a new mock instance.
func NewMockJwtAuthenticator(ctrl *gomock.Controller) *MockJwtAuthenticator {
	mock := &MockJwtAuthenticator{ctrl: ctrl}
	mock.recorder = &MockJwtAuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJwtAuthenticator) EXPECT() *MockJwtAuthenticatorMockRecorder {
	return m.recorder
}

// GenerateJwtToken mocks base method.
func (m *MockJwtAuthenticator) GenerateJwtToken(sub, jti string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateJwtToken", sub, jti)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateJwtToken indicates an expected call of GenerateJwtToken.
func (mr *MockJwtAuthenticatorMockRecorder) GenerateJwtToken(sub, jti any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateJwtToken", reflect.TypeOf((*MockJwtAuthenticator)(nil).GenerateJwtToken), sub, jti)
}

// VerifyJwtToken mocks base method.
func (m *MockJwtAuthenticator) VerifyJwtToken(signedToken string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyJwtToken", signedToken)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// VerifyJwtToken indicates an expected call of VerifyJwtToken.
func (mr *MockJwtAuthenticatorMockRecorder) VerifyJwtToken(signedToken any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyJwtToken", reflect.TypeOf((*MockJwtAuthenticator)(nil).VerifyJwtToken), signedToken)
}
