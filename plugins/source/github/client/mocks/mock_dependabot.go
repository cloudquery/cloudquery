// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/github/client (interfaces: DependabotService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v48/github"
)

// MockDependabotService is a mock of DependabotService interface.
type MockDependabotService struct {
	ctrl     *gomock.Controller
	recorder *MockDependabotServiceMockRecorder
}

// MockDependabotServiceMockRecorder is the mock recorder for MockDependabotService.
type MockDependabotServiceMockRecorder struct {
	mock *MockDependabotService
}

// NewMockDependabotService creates a new mock instance.
func NewMockDependabotService(ctrl *gomock.Controller) *MockDependabotService {
	mock := &MockDependabotService{ctrl: ctrl}
	mock.recorder = &MockDependabotServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDependabotService) EXPECT() *MockDependabotServiceMockRecorder {
	return m.recorder
}

// ListOrgAlerts mocks base method.
func (m *MockDependabotService) ListOrgAlerts(arg0 context.Context, arg1 string, arg2 *github.ListAlertsOptions) ([]*github.DependabotAlert, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrgAlerts", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*github.DependabotAlert)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListOrgAlerts indicates an expected call of ListOrgAlerts.
func (mr *MockDependabotServiceMockRecorder) ListOrgAlerts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrgAlerts", reflect.TypeOf((*MockDependabotService)(nil).ListOrgAlerts), arg0, arg1, arg2)
}

// ListOrgSecrets mocks base method.
func (m *MockDependabotService) ListOrgSecrets(arg0 context.Context, arg1 string, arg2 *github.ListOptions) (*github.Secrets, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrgSecrets", arg0, arg1, arg2)
	ret0, _ := ret[0].(*github.Secrets)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListOrgSecrets indicates an expected call of ListOrgSecrets.
func (mr *MockDependabotServiceMockRecorder) ListOrgSecrets(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrgSecrets", reflect.TypeOf((*MockDependabotService)(nil).ListOrgSecrets), arg0, arg1, arg2)
}

// ListRepoAlerts mocks base method.
func (m *MockDependabotService) ListRepoAlerts(arg0 context.Context, arg1, arg2 string, arg3 *github.ListAlertsOptions) ([]*github.DependabotAlert, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRepoAlerts", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*github.DependabotAlert)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListRepoAlerts indicates an expected call of ListRepoAlerts.
func (mr *MockDependabotServiceMockRecorder) ListRepoAlerts(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRepoAlerts", reflect.TypeOf((*MockDependabotService)(nil).ListRepoAlerts), arg0, arg1, arg2, arg3)
}

// ListRepoSecrets mocks base method.
func (m *MockDependabotService) ListRepoSecrets(arg0 context.Context, arg1, arg2 string, arg3 *github.ListOptions) (*github.Secrets, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRepoSecrets", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*github.Secrets)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListRepoSecrets indicates an expected call of ListRepoSecrets.
func (mr *MockDependabotServiceMockRecorder) ListRepoSecrets(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRepoSecrets", reflect.TypeOf((*MockDependabotService)(nil).ListRepoSecrets), arg0, arg1, arg2, arg3)
}
