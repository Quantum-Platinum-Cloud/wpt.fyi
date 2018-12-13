// Code generated by MockGen. DO NOT EDIT.
// Source: api/query/cache/index/index.go

// Package index is a generated GoMock package.
package index

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	metrics "github.com/web-platform-tests/results-analysis/metrics"
	query "github.com/web-platform-tests/wpt.fyi/api/query"
	shared "github.com/web-platform-tests/wpt.fyi/shared"
)

// MockIndex is a mock of Index interface
type MockIndex struct {
	ctrl     *gomock.Controller
	recorder *MockIndexMockRecorder
}

// MockIndexMockRecorder is the mock recorder for MockIndex
type MockIndexMockRecorder struct {
	mock *MockIndex
}

// NewMockIndex creates a new mock instance
func NewMockIndex(ctrl *gomock.Controller) *MockIndex {
	mock := &MockIndex{ctrl: ctrl}
	mock.recorder = &MockIndexMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndex) EXPECT() *MockIndexMockRecorder {
	return m.recorder
}

// Bind mocks base method
func (m *MockIndex) Bind(arg0 []shared.TestRun, arg1 query.ConcreteQuery) (query.Plan, error) {
	ret := m.ctrl.Call(m, "Bind", arg0, arg1)
	ret0, _ := ret[0].(query.Plan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bind indicates an expected call of Bind
func (mr *MockIndexMockRecorder) Bind(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bind", reflect.TypeOf((*MockIndex)(nil).Bind), arg0, arg1)
}

// Runs mocks base method
func (m *MockIndex) Runs(arg0 []RunID) ([]shared.TestRun, error) {
	ret := m.ctrl.Call(m, "Runs", arg0)
	ret0, _ := ret[0].([]shared.TestRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Runs indicates an expected call of Runs
func (mr *MockIndexMockRecorder) Runs(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Runs", reflect.TypeOf((*MockIndex)(nil).Runs), arg0)
}

// IngestRun mocks base method
func (m *MockIndex) IngestRun(arg0 shared.TestRun) error {
	ret := m.ctrl.Call(m, "IngestRun", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// IngestRun indicates an expected call of IngestRun
func (mr *MockIndexMockRecorder) IngestRun(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IngestRun", reflect.TypeOf((*MockIndex)(nil).IngestRun), arg0)
}

// EvictAnyRun mocks base method
func (m *MockIndex) EvictAnyRun() error {
	ret := m.ctrl.Call(m, "EvictAnyRun")
	ret0, _ := ret[0].(error)
	return ret0
}

// EvictAnyRun indicates an expected call of EvictAnyRun
func (mr *MockIndexMockRecorder) EvictAnyRun() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvictAnyRun", reflect.TypeOf((*MockIndex)(nil).EvictAnyRun))
}

// MockReportLoader is a mock of ReportLoader interface
type MockReportLoader struct {
	ctrl     *gomock.Controller
	recorder *MockReportLoaderMockRecorder
}

// MockReportLoaderMockRecorder is the mock recorder for MockReportLoader
type MockReportLoaderMockRecorder struct {
	mock *MockReportLoader
}

// NewMockReportLoader creates a new mock instance
func NewMockReportLoader(ctrl *gomock.Controller) *MockReportLoader {
	mock := &MockReportLoader{ctrl: ctrl}
	mock.recorder = &MockReportLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReportLoader) EXPECT() *MockReportLoaderMockRecorder {
	return m.recorder
}

// Load mocks base method
func (m *MockReportLoader) Load(arg0 shared.TestRun) (*metrics.TestResultsReport, error) {
	ret := m.ctrl.Call(m, "Load", arg0)
	ret0, _ := ret[0].(*metrics.TestResultsReport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Load indicates an expected call of Load
func (mr *MockReportLoaderMockRecorder) Load(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockReportLoader)(nil).Load), arg0)
}
