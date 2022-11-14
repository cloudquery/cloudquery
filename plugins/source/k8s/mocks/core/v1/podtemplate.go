// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/core/v1 (interfaces: PodTemplatesGetter,PodTemplateInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/core/v1"
	v12 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// MockPodTemplatesGetter is a mock of PodTemplatesGetter interface.
type MockPodTemplatesGetter struct {
	ctrl     *gomock.Controller
	recorder *MockPodTemplatesGetterMockRecorder
}

// MockPodTemplatesGetterMockRecorder is the mock recorder for MockPodTemplatesGetter.
type MockPodTemplatesGetterMockRecorder struct {
	mock *MockPodTemplatesGetter
}

// NewMockPodTemplatesGetter creates a new mock instance.
func NewMockPodTemplatesGetter(ctrl *gomock.Controller) *MockPodTemplatesGetter {
	mock := &MockPodTemplatesGetter{ctrl: ctrl}
	mock.recorder = &MockPodTemplatesGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodTemplatesGetter) EXPECT() *MockPodTemplatesGetterMockRecorder {
	return m.recorder
}

// PodTemplates mocks base method.
func (m *MockPodTemplatesGetter) PodTemplates(arg0 string) v12.PodTemplateInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PodTemplates", arg0)
	ret0, _ := ret[0].(v12.PodTemplateInterface)
	return ret0
}

// PodTemplates indicates an expected call of PodTemplates.
func (mr *MockPodTemplatesGetterMockRecorder) PodTemplates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PodTemplates", reflect.TypeOf((*MockPodTemplatesGetter)(nil).PodTemplates), arg0)
}

// MockPodTemplateInterface is a mock of PodTemplateInterface interface.
type MockPodTemplateInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPodTemplateInterfaceMockRecorder
}

// MockPodTemplateInterfaceMockRecorder is the mock recorder for MockPodTemplateInterface.
type MockPodTemplateInterfaceMockRecorder struct {
	mock *MockPodTemplateInterface
}

// NewMockPodTemplateInterface creates a new mock instance.
func NewMockPodTemplateInterface(ctrl *gomock.Controller) *MockPodTemplateInterface {
	mock := &MockPodTemplateInterface{ctrl: ctrl}
	mock.recorder = &MockPodTemplateInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodTemplateInterface) EXPECT() *MockPodTemplateInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockPodTemplateInterface) Apply(arg0 context.Context, arg1 *v11.PodTemplateApplyConfiguration, arg2 v10.ApplyOptions) (*v1.PodTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PodTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockPodTemplateInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockPodTemplateInterface)(nil).Apply), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockPodTemplateInterface) Create(arg0 context.Context, arg1 *v1.PodTemplate, arg2 v10.CreateOptions) (*v1.PodTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PodTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPodTemplateInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPodTemplateInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockPodTemplateInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPodTemplateInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPodTemplateInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockPodTemplateInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockPodTemplateInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockPodTemplateInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockPodTemplateInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.PodTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PodTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPodTemplateInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPodTemplateInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockPodTemplateInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.PodTemplateList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.PodTemplateList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPodTemplateInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPodTemplateInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockPodTemplateInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.PodTemplate, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.PodTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockPodTemplateInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockPodTemplateInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockPodTemplateInterface) Update(arg0 context.Context, arg1 *v1.PodTemplate, arg2 v10.UpdateOptions) (*v1.PodTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PodTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPodTemplateInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPodTemplateInterface)(nil).Update), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockPodTemplateInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockPodTemplateInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockPodTemplateInterface)(nil).Watch), arg0, arg1)
}
