// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock_config is a generated GoMock package.
package mock_config

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLevel is a mock of Level interface.
type MockLevel struct {
	ctrl     *gomock.Controller
	recorder *MockLevelMockRecorder
}

// MockLevelMockRecorder is the mock recorder for MockLevel.
type MockLevelMockRecorder struct {
	mock *MockLevel
}

// NewMockLevel creates a new mock instance.
func NewMockLevel(ctrl *gomock.Controller) *MockLevel {
	mock := &MockLevel{ctrl: ctrl}
	mock.recorder = &MockLevelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLevel) EXPECT() *MockLevelMockRecorder {
	return m.recorder
}

// UnmarshalText mocks base method.
func (m *MockLevel) UnmarshalText(text []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnmarshalText", text)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnmarshalText indicates an expected call of UnmarshalText.
func (mr *MockLevelMockRecorder) UnmarshalText(text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmarshalText", reflect.TypeOf((*MockLevel)(nil).UnmarshalText), text)
}