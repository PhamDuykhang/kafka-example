// Code generated by MockGen. DO NOT EDIT.
// Source: handler/greetinhandler/ping.go

// Package mock is a generated GoMock package.
package mock

import (
	types "github.com/PhamDuyKhang/kafkaexamples/testdemo/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockGateKeeper is a mock of GateKeeper interface
type MockGateKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockGateKeeperMockRecorder
}

// MockGateKeeperMockRecorder is the mock recorder for MockGateKeeper
type MockGateKeeperMockRecorder struct {
	mock *MockGateKeeper
}

// NewMockGateKeeper creates a new mock instance
func NewMockGateKeeper(ctrl *gomock.Controller) *MockGateKeeper {
	mock := &MockGateKeeper{ctrl: ctrl}
	mock.recorder = &MockGateKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGateKeeper) EXPECT() *MockGateKeeperMockRecorder {
	return m.recorder
}

// GetGreeting mocks base method
func (m *MockGateKeeper) GetGreeting() types.Greeting {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGreeting")
	ret0, _ := ret[0].(types.Greeting)
	return ret0
}

// GetGreeting indicates an expected call of GetGreeting
func (mr *MockGateKeeperMockRecorder) GetGreeting() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGreeting", reflect.TypeOf((*MockGateKeeper)(nil).GetGreeting))
}

// SaySimply mocks base method
func (m *MockGateKeeper) SaySimply() types.Greeting {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaySimply")
	ret0, _ := ret[0].(types.Greeting)
	return ret0
}

// SaySimply indicates an expected call of SaySimply
func (mr *MockGateKeeperMockRecorder) SaySimply() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaySimply", reflect.TypeOf((*MockGateKeeper)(nil).SaySimply))
}

// SayWithName mocks base method
func (m *MockGateKeeper) SayWithName(name string) types.Greeting {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SayWithName", name)
	ret0, _ := ret[0].(types.Greeting)
	return ret0
}

// SayWithName indicates an expected call of SayWithName
func (mr *MockGateKeeperMockRecorder) SayWithName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SayWithName", reflect.TypeOf((*MockGateKeeper)(nil).SayWithName), name)
}
