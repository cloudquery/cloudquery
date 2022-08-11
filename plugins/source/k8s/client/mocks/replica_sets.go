// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/k8s/client (interfaces: ReplicaSetsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MockReplicaSetsClient is a mock of ReplicaSetsClient interface.
type MockReplicaSetsClient struct {
	ctrl     *gomock.Controller
	recorder *MockReplicaSetsClientMockRecorder
}

// MockReplicaSetsClientMockRecorder is the mock recorder for MockReplicaSetsClient.
type MockReplicaSetsClientMockRecorder struct {
	mock *MockReplicaSetsClient
}

// NewMockReplicaSetsClient creates a new mock instance.
func NewMockReplicaSetsClient(ctrl *gomock.Controller) *MockReplicaSetsClient {
	mock := &MockReplicaSetsClient{ctrl: ctrl}
	mock.recorder = &MockReplicaSetsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReplicaSetsClient) EXPECT() *MockReplicaSetsClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockReplicaSetsClient) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.ReplicaSetList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.ReplicaSetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockReplicaSetsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockReplicaSetsClient)(nil).List), arg0, arg1)
}
