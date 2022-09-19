// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/azure/client/services (interfaces: SearchServicesClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	search "github.com/Azure/azure-sdk-for-go/services/search/mgmt/2020-08-01/search"
	uuid "github.com/gofrs/uuid"
	gomock "github.com/golang/mock/gomock"
)

// MockSearchServicesClient is a mock of SearchServicesClient interface.
type MockSearchServicesClient struct {
	ctrl     *gomock.Controller
	recorder *MockSearchServicesClientMockRecorder
}

// MockSearchServicesClientMockRecorder is the mock recorder for MockSearchServicesClient.
type MockSearchServicesClientMockRecorder struct {
	mock *MockSearchServicesClient
}

// NewMockSearchServicesClient creates a new mock instance.
func NewMockSearchServicesClient(ctrl *gomock.Controller) *MockSearchServicesClient {
	mock := &MockSearchServicesClient{ctrl: ctrl}
	mock.recorder = &MockSearchServicesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchServicesClient) EXPECT() *MockSearchServicesClientMockRecorder {
	return m.recorder
}

// ListBySubscription mocks base method.
func (m *MockSearchServicesClient) ListBySubscription(arg0 context.Context, arg1 *uuid.UUID) (search.ServiceListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubscription", arg0, arg1)
	ret0, _ := ret[0].(search.ServiceListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySubscription indicates an expected call of ListBySubscription.
func (mr *MockSearchServicesClientMockRecorder) ListBySubscription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubscription", reflect.TypeOf((*MockSearchServicesClient)(nil).ListBySubscription), arg0, arg1)
}
