// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/gandi/client (interfaces: SimpleHostingClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	simplehosting "github.com/go-gandi/go-gandi/simplehosting"
	gomock "github.com/golang/mock/gomock"
)

// MockSimpleHostingClient is a mock of SimpleHostingClient interface.
type MockSimpleHostingClient struct {
	ctrl     *gomock.Controller
	recorder *MockSimpleHostingClientMockRecorder
}

// MockSimpleHostingClientMockRecorder is the mock recorder for MockSimpleHostingClient.
type MockSimpleHostingClientMockRecorder struct {
	mock *MockSimpleHostingClient
}

// NewMockSimpleHostingClient creates a new mock instance.
func NewMockSimpleHostingClient(ctrl *gomock.Controller) *MockSimpleHostingClient {
	mock := &MockSimpleHostingClient{ctrl: ctrl}
	mock.recorder = &MockSimpleHostingClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSimpleHostingClient) EXPECT() *MockSimpleHostingClientMockRecorder {
	return m.recorder
}

// ListInstances mocks base method.
func (m *MockSimpleHostingClient) ListInstances() ([]simplehosting.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInstances")
	ret0, _ := ret[0].([]simplehosting.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInstances indicates an expected call of ListInstances.
func (mr *MockSimpleHostingClientMockRecorder) ListInstances() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInstances", reflect.TypeOf((*MockSimpleHostingClient)(nil).ListInstances))
}

// ListVhosts mocks base method.
func (m *MockSimpleHostingClient) ListVhosts(arg0 string) ([]simplehosting.Vhost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVhosts", arg0)
	ret0, _ := ret[0].([]simplehosting.Vhost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVhosts indicates an expected call of ListVhosts.
func (mr *MockSimpleHostingClientMockRecorder) ListVhosts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVhosts", reflect.TypeOf((*MockSimpleHostingClient)(nil).ListVhosts), arg0)
}
