// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/core/port/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/sifer169966/device-interactions/internal/core/domain"
)

// MockDeviceInteractionsRepository is a mock of DeviceInteractionsRepository interface.
type MockDeviceInteractionsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceInteractionsRepositoryMockRecorder
}

// MockDeviceInteractionsRepositoryMockRecorder is the mock recorder for MockDeviceInteractionsRepository.
type MockDeviceInteractionsRepositoryMockRecorder struct {
	mock *MockDeviceInteractionsRepository
}

// NewMockDeviceInteractionsRepository creates a new mock instance.
func NewMockDeviceInteractionsRepository(ctrl *gomock.Controller) *MockDeviceInteractionsRepository {
	mock := &MockDeviceInteractionsRepository{ctrl: ctrl}
	mock.recorder = &MockDeviceInteractionsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeviceInteractionsRepository) EXPECT() *MockDeviceInteractionsRepositoryMockRecorder {
	return m.recorder
}

// CreateMany mocks base method.
func (m *MockDeviceInteractionsRepository) CreateMany(d domain.DeviceSubmission) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMany", d)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMany indicates an expected call of CreateMany.
func (mr *MockDeviceInteractionsRepositoryMockRecorder) CreateMany(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMany", reflect.TypeOf((*MockDeviceInteractionsRepository)(nil).CreateMany), d)
}
