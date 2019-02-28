// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/web-platform-tests/wpt.fyi/api/taskcluster (interfaces: API)

// Package mock_taskcluster is a generated GoMock package.
package mock_taskcluster

import (
	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/github"
	reflect "reflect"
)

// MockAPI is a mock of API interface
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// HandleCheckRunEvent mocks base method
func (m *MockAPI) HandleCheckRunEvent(arg0 *github.CheckRunEvent) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleCheckRunEvent", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleCheckRunEvent indicates an expected call of HandleCheckRunEvent
func (mr *MockAPIMockRecorder) HandleCheckRunEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleCheckRunEvent", reflect.TypeOf((*MockAPI)(nil).HandleCheckRunEvent), arg0)
}
