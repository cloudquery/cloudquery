// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/networking/v1beta1 (interfaces: NetworkingV1beta1Interface)

// Package v1beta1 is a generated GoMock package.
package v1beta1

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	rest "k8s.io/client-go/rest"
)

// MockNetworkingV1beta1Interface is a mock of NetworkingV1beta1Interface interface.
type MockNetworkingV1beta1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkingV1beta1InterfaceMockRecorder
}

// MockNetworkingV1beta1InterfaceMockRecorder is the mock recorder for MockNetworkingV1beta1Interface.
type MockNetworkingV1beta1InterfaceMockRecorder struct {
	mock *MockNetworkingV1beta1Interface
}

// NewMockNetworkingV1beta1Interface creates a new mock instance.
func NewMockNetworkingV1beta1Interface(ctrl *gomock.Controller) *MockNetworkingV1beta1Interface {
	mock := &MockNetworkingV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockNetworkingV1beta1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkingV1beta1Interface) EXPECT() *MockNetworkingV1beta1InterfaceMockRecorder {
	return m.recorder
}

// IngressClasses mocks base method.
func (m *MockNetworkingV1beta1Interface) IngressClasses() v1beta1.IngressClassInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IngressClasses")
	ret0, _ := ret[0].(v1beta1.IngressClassInterface)
	return ret0
}

// IngressClasses indicates an expected call of IngressClasses.
func (mr *MockNetworkingV1beta1InterfaceMockRecorder) IngressClasses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IngressClasses", reflect.TypeOf((*MockNetworkingV1beta1Interface)(nil).IngressClasses))
}

// Ingresses mocks base method.
func (m *MockNetworkingV1beta1Interface) Ingresses(arg0 string) v1beta1.IngressInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ingresses", arg0)
	ret0, _ := ret[0].(v1beta1.IngressInterface)
	return ret0
}

// Ingresses indicates an expected call of Ingresses.
func (mr *MockNetworkingV1beta1InterfaceMockRecorder) Ingresses(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ingresses", reflect.TypeOf((*MockNetworkingV1beta1Interface)(nil).Ingresses), arg0)
}

// RESTClient mocks base method.
func (m *MockNetworkingV1beta1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient.
func (mr *MockNetworkingV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockNetworkingV1beta1Interface)(nil).RESTClient))
}
