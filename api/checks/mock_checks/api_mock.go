// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/web-platform-tests/wpt.fyi/api/checks (interfaces: API)

// Package mock_checks is a generated GoMock package.
package mock_checks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/github"
	shared "github.com/web-platform-tests/wpt.fyi/shared"
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

// CancelRun mocks base method
func (m *MockAPI) CancelRun(arg0, arg1, arg2 string, arg3 *github.CheckRun, arg4 *github.Installation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelRun", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelRun indicates an expected call of CancelRun
func (mr *MockAPIMockRecorder) CancelRun(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelRun", reflect.TypeOf((*MockAPI)(nil).CancelRun), arg0, arg1, arg2, arg3, arg4)
}

// Context mocks base method
func (m *MockAPI) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockAPIMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAPI)(nil).Context))
}

// CreateWPTCheckSuite mocks base method
func (m *MockAPI) CreateWPTCheckSuite(arg0, arg1 int64, arg2 string, arg3 ...int) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateWPTCheckSuite", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWPTCheckSuite indicates an expected call of CreateWPTCheckSuite
func (mr *MockAPIMockRecorder) CreateWPTCheckSuite(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWPTCheckSuite", reflect.TypeOf((*MockAPI)(nil).CreateWPTCheckSuite), varargs...)
}

// GetSuitesForSHA mocks base method
func (m *MockAPI) GetSuitesForSHA(arg0 string) ([]shared.CheckSuite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSuitesForSHA", arg0)
	ret0, _ := ret[0].([]shared.CheckSuite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSuitesForSHA indicates an expected call of GetSuitesForSHA
func (mr *MockAPIMockRecorder) GetSuitesForSHA(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSuitesForSHA", reflect.TypeOf((*MockAPI)(nil).GetSuitesForSHA), arg0)
}

// GetWPTRepoAppInstallationIDs mocks base method
func (m *MockAPI) GetWPTRepoAppInstallationIDs() (int64, int64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWPTRepoAppInstallationIDs")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(int64)
	return ret0, ret1
}

// GetWPTRepoAppInstallationIDs indicates an expected call of GetWPTRepoAppInstallationIDs
func (mr *MockAPIMockRecorder) GetWPTRepoAppInstallationIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWPTRepoAppInstallationIDs", reflect.TypeOf((*MockAPI)(nil).GetWPTRepoAppInstallationIDs))
}

// IgnoreFailure mocks base method
func (m *MockAPI) IgnoreFailure(arg0, arg1, arg2 string, arg3 *github.CheckRun, arg4 *github.Installation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IgnoreFailure", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// IgnoreFailure indicates an expected call of IgnoreFailure
func (mr *MockAPIMockRecorder) IgnoreFailure(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IgnoreFailure", reflect.TypeOf((*MockAPI)(nil).IgnoreFailure), arg0, arg1, arg2, arg3, arg4)
}

// PendingCheckRun mocks base method
func (m *MockAPI) PendingCheckRun(arg0 shared.CheckSuite, arg1 shared.ProductSpec) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingCheckRun", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PendingCheckRun indicates an expected call of PendingCheckRun
func (mr *MockAPIMockRecorder) PendingCheckRun(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingCheckRun", reflect.TypeOf((*MockAPI)(nil).PendingCheckRun), arg0, arg1)
}

// ScheduleResultsProcessing mocks base method
func (m *MockAPI) ScheduleResultsProcessing(arg0 string, arg1 shared.ProductSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScheduleResultsProcessing", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScheduleResultsProcessing indicates an expected call of ScheduleResultsProcessing
func (mr *MockAPIMockRecorder) ScheduleResultsProcessing(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleResultsProcessing", reflect.TypeOf((*MockAPI)(nil).ScheduleResultsProcessing), arg0, arg1)
}