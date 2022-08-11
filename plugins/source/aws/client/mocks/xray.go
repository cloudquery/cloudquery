// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/aws/client (interfaces: XrayClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	xray "github.com/aws/aws-sdk-go-v2/service/xray"
	gomock "github.com/golang/mock/gomock"
)

// MockXrayClient is a mock of XrayClient interface.
type MockXrayClient struct {
	ctrl     *gomock.Controller
	recorder *MockXrayClientMockRecorder
}

// MockXrayClientMockRecorder is the mock recorder for MockXrayClient.
type MockXrayClientMockRecorder struct {
	mock *MockXrayClient
}

// NewMockXrayClient creates a new mock instance.
func NewMockXrayClient(ctrl *gomock.Controller) *MockXrayClient {
	mock := &MockXrayClient{ctrl: ctrl}
	mock.recorder = &MockXrayClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockXrayClient) EXPECT() *MockXrayClientMockRecorder {
	return m.recorder
}

// GetEncryptionConfig mocks base method.
func (m *MockXrayClient) GetEncryptionConfig(arg0 context.Context, arg1 *xray.GetEncryptionConfigInput, arg2 ...func(*xray.Options)) (*xray.GetEncryptionConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEncryptionConfig", varargs...)
	ret0, _ := ret[0].(*xray.GetEncryptionConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEncryptionConfig indicates an expected call of GetEncryptionConfig.
func (mr *MockXrayClientMockRecorder) GetEncryptionConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEncryptionConfig", reflect.TypeOf((*MockXrayClient)(nil).GetEncryptionConfig), varargs...)
}

// GetGroups mocks base method.
func (m *MockXrayClient) GetGroups(arg0 context.Context, arg1 *xray.GetGroupsInput, arg2 ...func(*xray.Options)) (*xray.GetGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroups", varargs...)
	ret0, _ := ret[0].(*xray.GetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroups indicates an expected call of GetGroups.
func (mr *MockXrayClientMockRecorder) GetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroups", reflect.TypeOf((*MockXrayClient)(nil).GetGroups), varargs...)
}

// GetSamplingRules mocks base method.
func (m *MockXrayClient) GetSamplingRules(arg0 context.Context, arg1 *xray.GetSamplingRulesInput, arg2 ...func(*xray.Options)) (*xray.GetSamplingRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSamplingRules", varargs...)
	ret0, _ := ret[0].(*xray.GetSamplingRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSamplingRules indicates an expected call of GetSamplingRules.
func (mr *MockXrayClientMockRecorder) GetSamplingRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSamplingRules", reflect.TypeOf((*MockXrayClient)(nil).GetSamplingRules), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockXrayClient) ListTagsForResource(arg0 context.Context, arg1 *xray.ListTagsForResourceInput, arg2 ...func(*xray.Options)) (*xray.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*xray.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockXrayClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockXrayClient)(nil).ListTagsForResource), varargs...)
}
