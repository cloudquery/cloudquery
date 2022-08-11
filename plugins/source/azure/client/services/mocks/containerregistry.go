// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/azure/client/services (interfaces: ContainerRegistriesClient,ContainerReplicationsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	containerregistry "github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"
	gomock "github.com/golang/mock/gomock"
)

// MockContainerRegistriesClient is a mock of ContainerRegistriesClient interface.
type MockContainerRegistriesClient struct {
	ctrl     *gomock.Controller
	recorder *MockContainerRegistriesClientMockRecorder
}

// MockContainerRegistriesClientMockRecorder is the mock recorder for MockContainerRegistriesClient.
type MockContainerRegistriesClientMockRecorder struct {
	mock *MockContainerRegistriesClient
}

// NewMockContainerRegistriesClient creates a new mock instance.
func NewMockContainerRegistriesClient(ctrl *gomock.Controller) *MockContainerRegistriesClient {
	mock := &MockContainerRegistriesClient{ctrl: ctrl}
	mock.recorder = &MockContainerRegistriesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContainerRegistriesClient) EXPECT() *MockContainerRegistriesClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockContainerRegistriesClient) List(arg0 context.Context) (containerregistry.RegistryListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(containerregistry.RegistryListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockContainerRegistriesClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockContainerRegistriesClient)(nil).List), arg0)
}

// MockContainerReplicationsClient is a mock of ContainerReplicationsClient interface.
type MockContainerReplicationsClient struct {
	ctrl     *gomock.Controller
	recorder *MockContainerReplicationsClientMockRecorder
}

// MockContainerReplicationsClientMockRecorder is the mock recorder for MockContainerReplicationsClient.
type MockContainerReplicationsClientMockRecorder struct {
	mock *MockContainerReplicationsClient
}

// NewMockContainerReplicationsClient creates a new mock instance.
func NewMockContainerReplicationsClient(ctrl *gomock.Controller) *MockContainerReplicationsClient {
	mock := &MockContainerReplicationsClient{ctrl: ctrl}
	mock.recorder = &MockContainerReplicationsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContainerReplicationsClient) EXPECT() *MockContainerReplicationsClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockContainerReplicationsClient) List(arg0 context.Context, arg1, arg2 string) (containerregistry.ReplicationListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(containerregistry.ReplicationListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockContainerReplicationsClientMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockContainerReplicationsClient)(nil).List), arg0, arg1, arg2)
}
