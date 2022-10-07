// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/digitalocean/client (interfaces: MonitoringService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

// MockMonitoringService is a mock of MonitoringService interface.
type MockMonitoringService struct {
	ctrl     *gomock.Controller
	recorder *MockMonitoringServiceMockRecorder
}

// MockMonitoringServiceMockRecorder is the mock recorder for MockMonitoringService.
type MockMonitoringServiceMockRecorder struct {
	mock *MockMonitoringService
}

// NewMockMonitoringService creates a new mock instance.
func NewMockMonitoringService(ctrl *gomock.Controller) *MockMonitoringService {
	mock := &MockMonitoringService{ctrl: ctrl}
	mock.recorder = &MockMonitoringServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMonitoringService) EXPECT() *MockMonitoringServiceMockRecorder {
	return m.recorder
}

// ListAlertPolicies mocks base method.
func (m *MockMonitoringService) ListAlertPolicies(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.AlertPolicy, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAlertPolicies", arg0, arg1)
	ret0, _ := ret[0].([]godo.AlertPolicy)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListAlertPolicies indicates an expected call of ListAlertPolicies.
func (mr *MockMonitoringServiceMockRecorder) ListAlertPolicies(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAlertPolicies", reflect.TypeOf((*MockMonitoringService)(nil).ListAlertPolicies), arg0, arg1)
}
