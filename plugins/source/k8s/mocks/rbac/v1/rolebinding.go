// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/rbac/v1 (interfaces: RoleBindingsGetter,RoleBindingInterface)

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/rbac/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/rbac/v1"
	v12 "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

// MockRoleBindingsGetter is a mock of RoleBindingsGetter interface.
type MockRoleBindingsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockRoleBindingsGetterMockRecorder
}

// MockRoleBindingsGetterMockRecorder is the mock recorder for MockRoleBindingsGetter.
type MockRoleBindingsGetterMockRecorder struct {
	mock *MockRoleBindingsGetter
}

// NewMockRoleBindingsGetter creates a new mock instance.
func NewMockRoleBindingsGetter(ctrl *gomock.Controller) *MockRoleBindingsGetter {
	mock := &MockRoleBindingsGetter{ctrl: ctrl}
	mock.recorder = &MockRoleBindingsGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleBindingsGetter) EXPECT() *MockRoleBindingsGetterMockRecorder {
	return m.recorder
}

// RoleBindings mocks base method.
func (m *MockRoleBindingsGetter) RoleBindings(arg0 string) v12.RoleBindingInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoleBindings", arg0)
	ret0, _ := ret[0].(v12.RoleBindingInterface)
	return ret0
}

// RoleBindings indicates an expected call of RoleBindings.
func (mr *MockRoleBindingsGetterMockRecorder) RoleBindings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoleBindings", reflect.TypeOf((*MockRoleBindingsGetter)(nil).RoleBindings), arg0)
}

// MockRoleBindingInterface is a mock of RoleBindingInterface interface.
type MockRoleBindingInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRoleBindingInterfaceMockRecorder
}

// MockRoleBindingInterfaceMockRecorder is the mock recorder for MockRoleBindingInterface.
type MockRoleBindingInterfaceMockRecorder struct {
	mock *MockRoleBindingInterface
}

// NewMockRoleBindingInterface creates a new mock instance.
func NewMockRoleBindingInterface(ctrl *gomock.Controller) *MockRoleBindingInterface {
	mock := &MockRoleBindingInterface{ctrl: ctrl}
	mock.recorder = &MockRoleBindingInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleBindingInterface) EXPECT() *MockRoleBindingInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockRoleBindingInterface) Apply(arg0 context.Context, arg1 *v11.RoleBindingApplyConfiguration, arg2 v10.ApplyOptions) (*v1.RoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.RoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockRoleBindingInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockRoleBindingInterface)(nil).Apply), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockRoleBindingInterface) Create(arg0 context.Context, arg1 *v1.RoleBinding, arg2 v10.CreateOptions) (*v1.RoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.RoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRoleBindingInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRoleBindingInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockRoleBindingInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRoleBindingInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRoleBindingInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockRoleBindingInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockRoleBindingInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockRoleBindingInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockRoleBindingInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.RoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.RoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRoleBindingInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRoleBindingInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockRoleBindingInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.RoleBindingList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.RoleBindingList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRoleBindingInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRoleBindingInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockRoleBindingInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.RoleBinding, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.RoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockRoleBindingInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockRoleBindingInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockRoleBindingInterface) Update(arg0 context.Context, arg1 *v1.RoleBinding, arg2 v10.UpdateOptions) (*v1.RoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.RoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRoleBindingInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRoleBindingInterface)(nil).Update), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockRoleBindingInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockRoleBindingInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockRoleBindingInterface)(nil).Watch), arg0, arg1)
}
