// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/flowcontrol/v1 (interfaces: FlowcontrolV1Interface)

// Package v1 is a generated GoMock package.
package v1

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1"
	rest "k8s.io/client-go/rest"
)

// MockFlowcontrolV1Interface is a mock of FlowcontrolV1Interface interface.
type MockFlowcontrolV1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockFlowcontrolV1InterfaceMockRecorder
}

// MockFlowcontrolV1InterfaceMockRecorder is the mock recorder for MockFlowcontrolV1Interface.
type MockFlowcontrolV1InterfaceMockRecorder struct {
	mock *MockFlowcontrolV1Interface
}

// NewMockFlowcontrolV1Interface creates a new mock instance.
func NewMockFlowcontrolV1Interface(ctrl *gomock.Controller) *MockFlowcontrolV1Interface {
	mock := &MockFlowcontrolV1Interface{ctrl: ctrl}
	mock.recorder = &MockFlowcontrolV1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFlowcontrolV1Interface) EXPECT() *MockFlowcontrolV1InterfaceMockRecorder {
	return m.recorder
}

// FlowSchemas mocks base method.
func (m *MockFlowcontrolV1Interface) FlowSchemas() v1.FlowSchemaInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlowSchemas")
	ret0, _ := ret[0].(v1.FlowSchemaInterface)
	return ret0
}

// FlowSchemas indicates an expected call of FlowSchemas.
func (mr *MockFlowcontrolV1InterfaceMockRecorder) FlowSchemas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlowSchemas", reflect.TypeOf((*MockFlowcontrolV1Interface)(nil).FlowSchemas))
}

// PriorityLevelConfigurations mocks base method.
func (m *MockFlowcontrolV1Interface) PriorityLevelConfigurations() v1.PriorityLevelConfigurationInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PriorityLevelConfigurations")
	ret0, _ := ret[0].(v1.PriorityLevelConfigurationInterface)
	return ret0
}

// PriorityLevelConfigurations indicates an expected call of PriorityLevelConfigurations.
func (mr *MockFlowcontrolV1InterfaceMockRecorder) PriorityLevelConfigurations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PriorityLevelConfigurations", reflect.TypeOf((*MockFlowcontrolV1Interface)(nil).PriorityLevelConfigurations))
}

// RESTClient mocks base method.
func (m *MockFlowcontrolV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient.
func (mr *MockFlowcontrolV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockFlowcontrolV1Interface)(nil).RESTClient))
}
