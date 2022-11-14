// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/networking/v1 (interfaces: NetworkPoliciesGetter,NetworkPolicyInterface)

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

// MockNetworkPoliciesGetter is a mock of NetworkPoliciesGetter interface.
type MockNetworkPoliciesGetter struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkPoliciesGetterMockRecorder
}

// MockNetworkPoliciesGetterMockRecorder is the mock recorder for MockNetworkPoliciesGetter.
type MockNetworkPoliciesGetterMockRecorder struct {
	mock *MockNetworkPoliciesGetter
}

// NewMockNetworkPoliciesGetter creates a new mock instance.
func NewMockNetworkPoliciesGetter(ctrl *gomock.Controller) *MockNetworkPoliciesGetter {
	mock := &MockNetworkPoliciesGetter{ctrl: ctrl}
	mock.recorder = &MockNetworkPoliciesGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkPoliciesGetter) EXPECT() *MockNetworkPoliciesGetterMockRecorder {
	return m.recorder
}

// NetworkPolicies mocks base method.
func (m *MockNetworkPoliciesGetter) NetworkPolicies(arg0 string) v12.NetworkPolicyInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkPolicies", arg0)
	ret0, _ := ret[0].(v12.NetworkPolicyInterface)
	return ret0
}

// NetworkPolicies indicates an expected call of NetworkPolicies.
func (mr *MockNetworkPoliciesGetterMockRecorder) NetworkPolicies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkPolicies", reflect.TypeOf((*MockNetworkPoliciesGetter)(nil).NetworkPolicies), arg0)
}

// MockNetworkPolicyInterface is a mock of NetworkPolicyInterface interface.
type MockNetworkPolicyInterface struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkPolicyInterfaceMockRecorder
}

// MockNetworkPolicyInterfaceMockRecorder is the mock recorder for MockNetworkPolicyInterface.
type MockNetworkPolicyInterfaceMockRecorder struct {
	mock *MockNetworkPolicyInterface
}

// NewMockNetworkPolicyInterface creates a new mock instance.
func NewMockNetworkPolicyInterface(ctrl *gomock.Controller) *MockNetworkPolicyInterface {
	mock := &MockNetworkPolicyInterface{ctrl: ctrl}
	mock.recorder = &MockNetworkPolicyInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkPolicyInterface) EXPECT() *MockNetworkPolicyInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockNetworkPolicyInterface) Apply(arg0 context.Context, arg1 *v11.NetworkPolicyApplyConfiguration, arg2 v10.ApplyOptions) (*v1.NetworkPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.NetworkPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockNetworkPolicyInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockNetworkPolicyInterface)(nil).Apply), arg0, arg1, arg2)
}

// ApplyStatus mocks base method.
func (m *MockNetworkPolicyInterface) ApplyStatus(arg0 context.Context, arg1 *v11.NetworkPolicyApplyConfiguration, arg2 v10.ApplyOptions) (*v1.NetworkPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.NetworkPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyStatus indicates an expected call of ApplyStatus.
func (mr *MockNetworkPolicyInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyStatus", reflect.TypeOf((*MockNetworkPolicyInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockNetworkPolicyInterface) Create(arg0 context.Context, arg1 *v1.NetworkPolicy, arg2 v10.CreateOptions) (*v1.NetworkPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.NetworkPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockNetworkPolicyInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockNetworkPolicyInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockNetworkPolicyInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockNetworkPolicyInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNetworkPolicyInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockNetworkPolicyInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockNetworkPolicyInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockNetworkPolicyInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockNetworkPolicyInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.NetworkPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.NetworkPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockNetworkPolicyInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockNetworkPolicyInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockNetworkPolicyInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.NetworkPolicyList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.NetworkPolicyList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockNetworkPolicyInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNetworkPolicyInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockNetworkPolicyInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.NetworkPolicy, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.NetworkPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockNetworkPolicyInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockNetworkPolicyInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockNetworkPolicyInterface) Update(arg0 context.Context, arg1 *v1.NetworkPolicy, arg2 v10.UpdateOptions) (*v1.NetworkPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.NetworkPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockNetworkPolicyInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNetworkPolicyInterface)(nil).Update), arg0, arg1, arg2)
}

// UpdateStatus mocks base method.
func (m *MockNetworkPolicyInterface) UpdateStatus(arg0 context.Context, arg1 *v1.NetworkPolicy, arg2 v10.UpdateOptions) (*v1.NetworkPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.NetworkPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockNetworkPolicyInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockNetworkPolicyInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockNetworkPolicyInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockNetworkPolicyInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockNetworkPolicyInterface)(nil).Watch), arg0, arg1)
}
