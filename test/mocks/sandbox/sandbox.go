// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cri-o/cri-o/internal/lib/sandbox (interfaces: NamespaceIface)

// Package sandboxmock is a generated GoMock package.
package sandboxmock

import (
	sandbox "github.com/cri-o/cri-o/internal/lib/sandbox"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockNamespaceIface is a mock of NamespaceIface interface
type MockNamespaceIface struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceIfaceMockRecorder
}

// MockNamespaceIfaceMockRecorder is the mock recorder for MockNamespaceIface
type MockNamespaceIfaceMockRecorder struct {
	mock *MockNamespaceIface
}

// NewMockNamespaceIface creates a new mock instance
func NewMockNamespaceIface(ctrl *gomock.Controller) *MockNamespaceIface {
	mock := &MockNamespaceIface{ctrl: ctrl}
	mock.recorder = &MockNamespaceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNamespaceIface) EXPECT() *MockNamespaceIfaceMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockNamespaceIface) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockNamespaceIfaceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockNamespaceIface)(nil).Close))
}

// Get mocks base method
func (m *MockNamespaceIface) Get() *sandbox.Namespace {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(*sandbox.Namespace)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockNamespaceIfaceMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockNamespaceIface)(nil).Get))
}

// Initialize mocks base method
func (m *MockNamespaceIface) Initialize() sandbox.NamespaceIface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize")
	ret0, _ := ret[0].(sandbox.NamespaceIface)
	return ret0
}

// Initialize indicates an expected call of Initialize
func (mr *MockNamespaceIfaceMockRecorder) Initialize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockNamespaceIface)(nil).Initialize))
}

// Initialized mocks base method
func (m *MockNamespaceIface) Initialized() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialized")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Initialized indicates an expected call of Initialized
func (mr *MockNamespaceIfaceMockRecorder) Initialized() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialized", reflect.TypeOf((*MockNamespaceIface)(nil).Initialized))
}

// Path mocks base method
func (m *MockNamespaceIface) Path() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Path")
	ret0, _ := ret[0].(string)
	return ret0
}

// Path indicates an expected call of Path
func (mr *MockNamespaceIfaceMockRecorder) Path() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Path", reflect.TypeOf((*MockNamespaceIface)(nil).Path))
}

// Remove mocks base method
func (m *MockNamespaceIface) Remove() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove")
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockNamespaceIfaceMockRecorder) Remove() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockNamespaceIface)(nil).Remove))
}

// Type mocks base method
func (m *MockNamespaceIface) Type() sandbox.NSType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(sandbox.NSType)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockNamespaceIfaceMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockNamespaceIface)(nil).Type))
}
