// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/aws/client (interfaces: DAXClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	dax "github.com/aws/aws-sdk-go-v2/service/dax"
	gomock "github.com/golang/mock/gomock"
)

// MockDAXClient is a mock of DAXClient interface.
type MockDAXClient struct {
	ctrl     *gomock.Controller
	recorder *MockDAXClientMockRecorder
}

// MockDAXClientMockRecorder is the mock recorder for MockDAXClient.
type MockDAXClientMockRecorder struct {
	mock *MockDAXClient
}

// NewMockDAXClient creates a new mock instance.
func NewMockDAXClient(ctrl *gomock.Controller) *MockDAXClient {
	mock := &MockDAXClient{ctrl: ctrl}
	mock.recorder = &MockDAXClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDAXClient) EXPECT() *MockDAXClientMockRecorder {
	return m.recorder
}

// DescribeClusters mocks base method.
func (m *MockDAXClient) DescribeClusters(arg0 context.Context, arg1 *dax.DescribeClustersInput, arg2 ...func(*dax.Options)) (*dax.DescribeClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeClusters", varargs...)
	ret0, _ := ret[0].(*dax.DescribeClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeClusters indicates an expected call of DescribeClusters.
func (mr *MockDAXClientMockRecorder) DescribeClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeClusters", reflect.TypeOf((*MockDAXClient)(nil).DescribeClusters), varargs...)
}

// ListTags mocks base method.
func (m *MockDAXClient) ListTags(arg0 context.Context, arg1 *dax.ListTagsInput, arg2 ...func(*dax.Options)) (*dax.ListTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTags", varargs...)
	ret0, _ := ret[0].(*dax.ListTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTags indicates an expected call of ListTags.
func (mr *MockDAXClientMockRecorder) ListTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockDAXClient)(nil).ListTags), varargs...)
}
