// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/aws/client (interfaces: EcsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	ecs "github.com/aws/aws-sdk-go-v2/service/ecs"
	gomock "github.com/golang/mock/gomock"
)

// MockEcsClient is a mock of EcsClient interface.
type MockEcsClient struct {
	ctrl     *gomock.Controller
	recorder *MockEcsClientMockRecorder
}

// MockEcsClientMockRecorder is the mock recorder for MockEcsClient.
type MockEcsClientMockRecorder struct {
	mock *MockEcsClient
}

// NewMockEcsClient creates a new mock instance.
func NewMockEcsClient(ctrl *gomock.Controller) *MockEcsClient {
	mock := &MockEcsClient{ctrl: ctrl}
	mock.recorder = &MockEcsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEcsClient) EXPECT() *MockEcsClientMockRecorder {
	return m.recorder
}

// DescribeClusters mocks base method.
func (m *MockEcsClient) DescribeClusters(arg0 context.Context, arg1 *ecs.DescribeClustersInput, arg2 ...func(*ecs.Options)) (*ecs.DescribeClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeClusters", varargs...)
	ret0, _ := ret[0].(*ecs.DescribeClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeClusters indicates an expected call of DescribeClusters.
func (mr *MockEcsClientMockRecorder) DescribeClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeClusters", reflect.TypeOf((*MockEcsClient)(nil).DescribeClusters), varargs...)
}

// DescribeContainerInstances mocks base method.
func (m *MockEcsClient) DescribeContainerInstances(arg0 context.Context, arg1 *ecs.DescribeContainerInstancesInput, arg2 ...func(*ecs.Options)) (*ecs.DescribeContainerInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeContainerInstances", varargs...)
	ret0, _ := ret[0].(*ecs.DescribeContainerInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeContainerInstances indicates an expected call of DescribeContainerInstances.
func (mr *MockEcsClientMockRecorder) DescribeContainerInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeContainerInstances", reflect.TypeOf((*MockEcsClient)(nil).DescribeContainerInstances), varargs...)
}

// DescribeServices mocks base method.
func (m *MockEcsClient) DescribeServices(arg0 context.Context, arg1 *ecs.DescribeServicesInput, arg2 ...func(*ecs.Options)) (*ecs.DescribeServicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeServices", varargs...)
	ret0, _ := ret[0].(*ecs.DescribeServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeServices indicates an expected call of DescribeServices.
func (mr *MockEcsClientMockRecorder) DescribeServices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeServices", reflect.TypeOf((*MockEcsClient)(nil).DescribeServices), varargs...)
}

// DescribeTaskDefinition mocks base method.
func (m *MockEcsClient) DescribeTaskDefinition(arg0 context.Context, arg1 *ecs.DescribeTaskDefinitionInput, arg2 ...func(*ecs.Options)) (*ecs.DescribeTaskDefinitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTaskDefinition", varargs...)
	ret0, _ := ret[0].(*ecs.DescribeTaskDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTaskDefinition indicates an expected call of DescribeTaskDefinition.
func (mr *MockEcsClientMockRecorder) DescribeTaskDefinition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTaskDefinition", reflect.TypeOf((*MockEcsClient)(nil).DescribeTaskDefinition), varargs...)
}

// DescribeTasks mocks base method.
func (m *MockEcsClient) DescribeTasks(arg0 context.Context, arg1 *ecs.DescribeTasksInput, arg2 ...func(*ecs.Options)) (*ecs.DescribeTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTasks", varargs...)
	ret0, _ := ret[0].(*ecs.DescribeTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTasks indicates an expected call of DescribeTasks.
func (mr *MockEcsClientMockRecorder) DescribeTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTasks", reflect.TypeOf((*MockEcsClient)(nil).DescribeTasks), varargs...)
}

// ListClusters mocks base method.
func (m *MockEcsClient) ListClusters(arg0 context.Context, arg1 *ecs.ListClustersInput, arg2 ...func(*ecs.Options)) (*ecs.ListClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListClusters", varargs...)
	ret0, _ := ret[0].(*ecs.ListClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusters indicates an expected call of ListClusters.
func (mr *MockEcsClientMockRecorder) ListClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockEcsClient)(nil).ListClusters), varargs...)
}

// ListContainerInstances mocks base method.
func (m *MockEcsClient) ListContainerInstances(arg0 context.Context, arg1 *ecs.ListContainerInstancesInput, arg2 ...func(*ecs.Options)) (*ecs.ListContainerInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListContainerInstances", varargs...)
	ret0, _ := ret[0].(*ecs.ListContainerInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContainerInstances indicates an expected call of ListContainerInstances.
func (mr *MockEcsClientMockRecorder) ListContainerInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContainerInstances", reflect.TypeOf((*MockEcsClient)(nil).ListContainerInstances), varargs...)
}

// ListServices mocks base method.
func (m *MockEcsClient) ListServices(arg0 context.Context, arg1 *ecs.ListServicesInput, arg2 ...func(*ecs.Options)) (*ecs.ListServicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListServices", varargs...)
	ret0, _ := ret[0].(*ecs.ListServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices.
func (mr *MockEcsClientMockRecorder) ListServices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockEcsClient)(nil).ListServices), varargs...)
}

// ListTaskDefinitions mocks base method.
func (m *MockEcsClient) ListTaskDefinitions(arg0 context.Context, arg1 *ecs.ListTaskDefinitionsInput, arg2 ...func(*ecs.Options)) (*ecs.ListTaskDefinitionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTaskDefinitions", varargs...)
	ret0, _ := ret[0].(*ecs.ListTaskDefinitionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTaskDefinitions indicates an expected call of ListTaskDefinitions.
func (mr *MockEcsClientMockRecorder) ListTaskDefinitions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTaskDefinitions", reflect.TypeOf((*MockEcsClient)(nil).ListTaskDefinitions), varargs...)
}

// ListTasks mocks base method.
func (m *MockEcsClient) ListTasks(arg0 context.Context, arg1 *ecs.ListTasksInput, arg2 ...func(*ecs.Options)) (*ecs.ListTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTasks", varargs...)
	ret0, _ := ret[0].(*ecs.ListTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTasks indicates an expected call of ListTasks.
func (mr *MockEcsClientMockRecorder) ListTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTasks", reflect.TypeOf((*MockEcsClient)(nil).ListTasks), varargs...)
}
