// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/rbac/v1 (interfaces: RbacV1Interface)

// Package v1 is a generated GoMock package.
package v1

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	rest "k8s.io/client-go/rest"
)

// MockRbacV1Interface is a mock of RbacV1Interface interface.
type MockRbacV1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockRbacV1InterfaceMockRecorder
}

// MockRbacV1InterfaceMockRecorder is the mock recorder for MockRbacV1Interface.
type MockRbacV1InterfaceMockRecorder struct {
	mock *MockRbacV1Interface
}

// NewMockRbacV1Interface creates a new mock instance.
func NewMockRbacV1Interface(ctrl *gomock.Controller) *MockRbacV1Interface {
	mock := &MockRbacV1Interface{ctrl: ctrl}
	mock.recorder = &MockRbacV1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRbacV1Interface) EXPECT() *MockRbacV1InterfaceMockRecorder {
	return m.recorder
}

// ClusterRoleBindings mocks base method.
func (m *MockRbacV1Interface) ClusterRoleBindings() v1.ClusterRoleBindingInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterRoleBindings")
	ret0, _ := ret[0].(v1.ClusterRoleBindingInterface)
	return ret0
}

// ClusterRoleBindings indicates an expected call of ClusterRoleBindings.
func (mr *MockRbacV1InterfaceMockRecorder) ClusterRoleBindings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterRoleBindings", reflect.TypeOf((*MockRbacV1Interface)(nil).ClusterRoleBindings))
}

// ClusterRoles mocks base method.
func (m *MockRbacV1Interface) ClusterRoles() v1.ClusterRoleInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterRoles")
	ret0, _ := ret[0].(v1.ClusterRoleInterface)
	return ret0
}

// ClusterRoles indicates an expected call of ClusterRoles.
func (mr *MockRbacV1InterfaceMockRecorder) ClusterRoles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterRoles", reflect.TypeOf((*MockRbacV1Interface)(nil).ClusterRoles))
}

// RESTClient mocks base method.
func (m *MockRbacV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient.
func (mr *MockRbacV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockRbacV1Interface)(nil).RESTClient))
}

// RoleBindings mocks base method.
func (m *MockRbacV1Interface) RoleBindings(arg0 string) v1.RoleBindingInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoleBindings", arg0)
	ret0, _ := ret[0].(v1.RoleBindingInterface)
	return ret0
}

// RoleBindings indicates an expected call of RoleBindings.
func (mr *MockRbacV1InterfaceMockRecorder) RoleBindings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoleBindings", reflect.TypeOf((*MockRbacV1Interface)(nil).RoleBindings), arg0)
}

// Roles mocks base method.
func (m *MockRbacV1Interface) Roles(arg0 string) v1.RoleInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Roles", arg0)
	ret0, _ := ret[0].(v1.RoleInterface)
	return ret0
}

// Roles indicates an expected call of Roles.
func (mr *MockRbacV1InterfaceMockRecorder) Roles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Roles", reflect.TypeOf((*MockRbacV1Interface)(nil).Roles), arg0)
}
