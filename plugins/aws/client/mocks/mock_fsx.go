// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/aws/client (interfaces: FsxClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	fsx "github.com/aws/aws-sdk-go-v2/service/fsx"
	gomock "github.com/golang/mock/gomock"
)

// MockFsxClient is a mock of FsxClient interface.
type MockFsxClient struct {
	ctrl     *gomock.Controller
	recorder *MockFsxClientMockRecorder
}

// MockFsxClientMockRecorder is the mock recorder for MockFsxClient.
type MockFsxClientMockRecorder struct {
	mock *MockFsxClient
}

// NewMockFsxClient creates a new mock instance.
func NewMockFsxClient(ctrl *gomock.Controller) *MockFsxClient {
	mock := &MockFsxClient{ctrl: ctrl}
	mock.recorder = &MockFsxClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFsxClient) EXPECT() *MockFsxClientMockRecorder {
	return m.recorder
}

// DescribeBackups mocks base method.
func (m *MockFsxClient) DescribeBackups(arg0 context.Context, arg1 *fsx.DescribeBackupsInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeBackupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeBackups", varargs...)
	ret0, _ := ret[0].(*fsx.DescribeBackupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeBackups indicates an expected call of DescribeBackups.
func (mr *MockFsxClientMockRecorder) DescribeBackups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeBackups", reflect.TypeOf((*MockFsxClient)(nil).DescribeBackups), varargs...)
}
