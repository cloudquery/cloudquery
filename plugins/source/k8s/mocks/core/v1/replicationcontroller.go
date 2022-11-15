// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/core/v1 (interfaces: ReplicationControllersGetter,ReplicationControllerInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/autoscaling/v1"
	v10 "k8s.io/api/core/v1"
	v11 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v12 "k8s.io/client-go/applyconfigurations/core/v1"
	v13 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// MockReplicationControllersGetter is a mock of ReplicationControllersGetter interface.
type MockReplicationControllersGetter struct {
	ctrl     *gomock.Controller
	recorder *MockReplicationControllersGetterMockRecorder
}

// MockReplicationControllersGetterMockRecorder is the mock recorder for MockReplicationControllersGetter.
type MockReplicationControllersGetterMockRecorder struct {
	mock *MockReplicationControllersGetter
}

// NewMockReplicationControllersGetter creates a new mock instance.
func NewMockReplicationControllersGetter(ctrl *gomock.Controller) *MockReplicationControllersGetter {
	mock := &MockReplicationControllersGetter{ctrl: ctrl}
	mock.recorder = &MockReplicationControllersGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReplicationControllersGetter) EXPECT() *MockReplicationControllersGetterMockRecorder {
	return m.recorder
}

// ReplicationControllers mocks base method.
func (m *MockReplicationControllersGetter) ReplicationControllers(arg0 string) v13.ReplicationControllerInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplicationControllers", arg0)
	ret0, _ := ret[0].(v13.ReplicationControllerInterface)
	return ret0
}

// ReplicationControllers indicates an expected call of ReplicationControllers.
func (mr *MockReplicationControllersGetterMockRecorder) ReplicationControllers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicationControllers", reflect.TypeOf((*MockReplicationControllersGetter)(nil).ReplicationControllers), arg0)
}

// MockReplicationControllerInterface is a mock of ReplicationControllerInterface interface.
type MockReplicationControllerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockReplicationControllerInterfaceMockRecorder
}

// MockReplicationControllerInterfaceMockRecorder is the mock recorder for MockReplicationControllerInterface.
type MockReplicationControllerInterfaceMockRecorder struct {
	mock *MockReplicationControllerInterface
}

// NewMockReplicationControllerInterface creates a new mock instance.
func NewMockReplicationControllerInterface(ctrl *gomock.Controller) *MockReplicationControllerInterface {
	mock := &MockReplicationControllerInterface{ctrl: ctrl}
	mock.recorder = &MockReplicationControllerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReplicationControllerInterface) EXPECT() *MockReplicationControllerInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockReplicationControllerInterface) Apply(arg0 context.Context, arg1 *v12.ReplicationControllerApplyConfiguration, arg2 v11.ApplyOptions) (*v10.ReplicationController, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v10.ReplicationController)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockReplicationControllerInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockReplicationControllerInterface)(nil).Apply), arg0, arg1, arg2)
}

// ApplyStatus mocks base method.
func (m *MockReplicationControllerInterface) ApplyStatus(arg0 context.Context, arg1 *v12.ReplicationControllerApplyConfiguration, arg2 v11.ApplyOptions) (*v10.ReplicationController, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v10.ReplicationController)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyStatus indicates an expected call of ApplyStatus.
func (mr *MockReplicationControllerInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyStatus", reflect.TypeOf((*MockReplicationControllerInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockReplicationControllerInterface) Create(arg0 context.Context, arg1 *v10.ReplicationController, arg2 v11.CreateOptions) (*v10.ReplicationController, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v10.ReplicationController)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockReplicationControllerInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockReplicationControllerInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockReplicationControllerInterface) Delete(arg0 context.Context, arg1 string, arg2 v11.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockReplicationControllerInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockReplicationControllerInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockReplicationControllerInterface) DeleteCollection(arg0 context.Context, arg1 v11.DeleteOptions, arg2 v11.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockReplicationControllerInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockReplicationControllerInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockReplicationControllerInterface) Get(arg0 context.Context, arg1 string, arg2 v11.GetOptions) (*v10.ReplicationController, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v10.ReplicationController)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockReplicationControllerInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockReplicationControllerInterface)(nil).Get), arg0, arg1, arg2)
}

// GetScale mocks base method.
func (m *MockReplicationControllerInterface) GetScale(arg0 context.Context, arg1 string, arg2 v11.GetOptions) (*v1.Scale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScale", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Scale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScale indicates an expected call of GetScale.
func (mr *MockReplicationControllerInterfaceMockRecorder) GetScale(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScale", reflect.TypeOf((*MockReplicationControllerInterface)(nil).GetScale), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockReplicationControllerInterface) List(arg0 context.Context, arg1 v11.ListOptions) (*v10.ReplicationControllerList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v10.ReplicationControllerList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockReplicationControllerInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockReplicationControllerInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockReplicationControllerInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v11.PatchOptions, arg5 ...string) (*v10.ReplicationController, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v10.ReplicationController)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockReplicationControllerInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockReplicationControllerInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockReplicationControllerInterface) Update(arg0 context.Context, arg1 *v10.ReplicationController, arg2 v11.UpdateOptions) (*v10.ReplicationController, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v10.ReplicationController)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockReplicationControllerInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockReplicationControllerInterface)(nil).Update), arg0, arg1, arg2)
}

// UpdateScale mocks base method.
func (m *MockReplicationControllerInterface) UpdateScale(arg0 context.Context, arg1 string, arg2 *v1.Scale, arg3 v11.UpdateOptions) (*v1.Scale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateScale", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1.Scale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateScale indicates an expected call of UpdateScale.
func (mr *MockReplicationControllerInterfaceMockRecorder) UpdateScale(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScale", reflect.TypeOf((*MockReplicationControllerInterface)(nil).UpdateScale), arg0, arg1, arg2, arg3)
}

// UpdateStatus mocks base method.
func (m *MockReplicationControllerInterface) UpdateStatus(arg0 context.Context, arg1 *v10.ReplicationController, arg2 v11.UpdateOptions) (*v10.ReplicationController, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v10.ReplicationController)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockReplicationControllerInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockReplicationControllerInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockReplicationControllerInterface) Watch(arg0 context.Context, arg1 v11.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockReplicationControllerInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockReplicationControllerInterface)(nil).Watch), arg0, arg1)
}
