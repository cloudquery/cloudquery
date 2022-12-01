// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/github/client (interfaces: BillingService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v48/github"
)

// MockBillingService is a mock of BillingService interface.
type MockBillingService struct {
	ctrl     *gomock.Controller
	recorder *MockBillingServiceMockRecorder
}

// MockBillingServiceMockRecorder is the mock recorder for MockBillingService.
type MockBillingServiceMockRecorder struct {
	mock *MockBillingService
}

// NewMockBillingService creates a new mock instance.
func NewMockBillingService(ctrl *gomock.Controller) *MockBillingService {
	mock := &MockBillingService{ctrl: ctrl}
	mock.recorder = &MockBillingServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBillingService) EXPECT() *MockBillingServiceMockRecorder {
	return m.recorder
}

// GetActionsBillingOrg mocks base method.
func (m *MockBillingService) GetActionsBillingOrg(arg0 context.Context, arg1 string) (*github.ActionBilling, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActionsBillingOrg", arg0, arg1)
	ret0, _ := ret[0].(*github.ActionBilling)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetActionsBillingOrg indicates an expected call of GetActionsBillingOrg.
func (mr *MockBillingServiceMockRecorder) GetActionsBillingOrg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActionsBillingOrg", reflect.TypeOf((*MockBillingService)(nil).GetActionsBillingOrg), arg0, arg1)
}

// GetPackagesBillingOrg mocks base method.
func (m *MockBillingService) GetPackagesBillingOrg(arg0 context.Context, arg1 string) (*github.PackageBilling, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPackagesBillingOrg", arg0, arg1)
	ret0, _ := ret[0].(*github.PackageBilling)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPackagesBillingOrg indicates an expected call of GetPackagesBillingOrg.
func (mr *MockBillingServiceMockRecorder) GetPackagesBillingOrg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPackagesBillingOrg", reflect.TypeOf((*MockBillingService)(nil).GetPackagesBillingOrg), arg0, arg1)
}

// GetStorageBillingOrg mocks base method.
func (m *MockBillingService) GetStorageBillingOrg(arg0 context.Context, arg1 string) (*github.StorageBilling, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageBillingOrg", arg0, arg1)
	ret0, _ := ret[0].(*github.StorageBilling)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetStorageBillingOrg indicates an expected call of GetStorageBillingOrg.
func (mr *MockBillingServiceMockRecorder) GetStorageBillingOrg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageBillingOrg", reflect.TypeOf((*MockBillingService)(nil).GetStorageBillingOrg), arg0, arg1)
}
