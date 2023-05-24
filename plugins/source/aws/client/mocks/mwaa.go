// Code generated by MockGen. DO NOT EDIT.
// Source: mwaa.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	
	reflect "reflect"

	mwaa "github.com/aws/aws-sdk-go-v2/service/mwaa"
	gomock "github.com/golang/mock/gomock"
)

// MockMwaaClient is a mock of MwaaClient interface.
type MockMwaaClient struct {
	ctrl     *gomock.Controller
	recorder *MockMwaaClientMockRecorder
}

// MockMwaaClientMockRecorder is the mock recorder for MockMwaaClient.
type MockMwaaClientMockRecorder struct {
	mock *MockMwaaClient
}

// NewMockMwaaClient creates a new mock instance.
func NewMockMwaaClient(ctrl *gomock.Controller) *MockMwaaClient {
	mock := &MockMwaaClient{ctrl: ctrl}
	mock.recorder = &MockMwaaClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMwaaClient) EXPECT() *MockMwaaClientMockRecorder {
	return m.recorder
}

// GetEnvironment mocks base method.
func (m *MockMwaaClient) GetEnvironment(arg0 context.Context, arg1 *mwaa.GetEnvironmentInput, arg2 ...func(*mwaa.Options)) (*mwaa.GetEnvironmentOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &mwaa.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetEnvironment")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEnvironment", varargs...)
	ret0, _ := ret[0].(*mwaa.GetEnvironmentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnvironment indicates an expected call of GetEnvironment.
func (mr *MockMwaaClientMockRecorder) GetEnvironment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvironment", reflect.TypeOf((*MockMwaaClient)(nil).GetEnvironment), varargs...)
}

// ListEnvironments mocks base method.
func (m *MockMwaaClient) ListEnvironments(arg0 context.Context, arg1 *mwaa.ListEnvironmentsInput, arg2 ...func(*mwaa.Options)) (*mwaa.ListEnvironmentsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &mwaa.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListEnvironments")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEnvironments", varargs...)
	ret0, _ := ret[0].(*mwaa.ListEnvironmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEnvironments indicates an expected call of ListEnvironments.
func (mr *MockMwaaClientMockRecorder) ListEnvironments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnvironments", reflect.TypeOf((*MockMwaaClient)(nil).ListEnvironments), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockMwaaClient) ListTagsForResource(arg0 context.Context, arg1 *mwaa.ListTagsForResourceInput, arg2 ...func(*mwaa.Options)) (*mwaa.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &mwaa.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListTagsForResource")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*mwaa.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockMwaaClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockMwaaClient)(nil).ListTagsForResource), varargs...)
}
