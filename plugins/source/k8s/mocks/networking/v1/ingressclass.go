// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/networking/v1 (interfaces: IngressClassesGetter,IngressClassInterface)

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/networking/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/networking/v1"
	v12 "k8s.io/client-go/kubernetes/typed/networking/v1"
)

// MockIngressClassesGetter is a mock of IngressClassesGetter interface.
type MockIngressClassesGetter struct {
	ctrl     *gomock.Controller
	recorder *MockIngressClassesGetterMockRecorder
}

// MockIngressClassesGetterMockRecorder is the mock recorder for MockIngressClassesGetter.
type MockIngressClassesGetterMockRecorder struct {
	mock *MockIngressClassesGetter
}

// NewMockIngressClassesGetter creates a new mock instance.
func NewMockIngressClassesGetter(ctrl *gomock.Controller) *MockIngressClassesGetter {
	mock := &MockIngressClassesGetter{ctrl: ctrl}
	mock.recorder = &MockIngressClassesGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIngressClassesGetter) EXPECT() *MockIngressClassesGetterMockRecorder {
	return m.recorder
}

// IngressClasses mocks base method.
func (m *MockIngressClassesGetter) IngressClasses() v12.IngressClassInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IngressClasses")
	ret0, _ := ret[0].(v12.IngressClassInterface)
	return ret0
}

// IngressClasses indicates an expected call of IngressClasses.
func (mr *MockIngressClassesGetterMockRecorder) IngressClasses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IngressClasses", reflect.TypeOf((*MockIngressClassesGetter)(nil).IngressClasses))
}

// MockIngressClassInterface is a mock of IngressClassInterface interface.
type MockIngressClassInterface struct {
	ctrl     *gomock.Controller
	recorder *MockIngressClassInterfaceMockRecorder
}

// MockIngressClassInterfaceMockRecorder is the mock recorder for MockIngressClassInterface.
type MockIngressClassInterfaceMockRecorder struct {
	mock *MockIngressClassInterface
}

// NewMockIngressClassInterface creates a new mock instance.
func NewMockIngressClassInterface(ctrl *gomock.Controller) *MockIngressClassInterface {
	mock := &MockIngressClassInterface{ctrl: ctrl}
	mock.recorder = &MockIngressClassInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIngressClassInterface) EXPECT() *MockIngressClassInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockIngressClassInterface) Apply(arg0 context.Context, arg1 *v11.IngressClassApplyConfiguration, arg2 v10.ApplyOptions) (*v1.IngressClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.IngressClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockIngressClassInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockIngressClassInterface)(nil).Apply), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockIngressClassInterface) Create(arg0 context.Context, arg1 *v1.IngressClass, arg2 v10.CreateOptions) (*v1.IngressClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.IngressClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIngressClassInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIngressClassInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockIngressClassInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIngressClassInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIngressClassInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockIngressClassInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockIngressClassInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockIngressClassInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockIngressClassInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.IngressClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.IngressClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIngressClassInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIngressClassInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockIngressClassInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.IngressClassList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.IngressClassList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockIngressClassInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIngressClassInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockIngressClassInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.IngressClass, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.IngressClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockIngressClassInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockIngressClassInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockIngressClassInterface) Update(arg0 context.Context, arg1 *v1.IngressClass, arg2 v10.UpdateOptions) (*v1.IngressClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.IngressClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIngressClassInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIngressClassInterface)(nil).Update), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockIngressClassInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockIngressClassInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockIngressClassInterface)(nil).Watch), arg0, arg1)
}
