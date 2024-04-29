// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repository/transaction.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTransactionRepositoryImpl is a mock of TransactionRepositoryImpl interface.
type MockTransactionRepositoryImpl struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoryImplMockRecorder
}

// MockTransactionRepositoryImplMockRecorder is the mock recorder for MockTransactionRepositoryImpl.
type MockTransactionRepositoryImplMockRecorder struct {
	mock *MockTransactionRepositoryImpl
}

// NewMockTransactionRepositoryImpl creates a new mock instance.
func NewMockTransactionRepositoryImpl(ctrl *gomock.Controller) *MockTransactionRepositoryImpl {
	mock := &MockTransactionRepositoryImpl{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoryImplMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepositoryImpl) EXPECT() *MockTransactionRepositoryImplMockRecorder {
	return m.recorder
}

// WithTransaction mocks base method.
func (m *MockTransactionRepositoryImpl) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTransaction", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithTransaction indicates an expected call of WithTransaction.
func (mr *MockTransactionRepositoryImplMockRecorder) WithTransaction(ctx, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTransaction", reflect.TypeOf((*MockTransactionRepositoryImpl)(nil).WithTransaction), ctx, fn)
}