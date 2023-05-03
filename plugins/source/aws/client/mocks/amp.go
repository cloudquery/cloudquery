// Code generated by MockGen. DO NOT EDIT.
// Source: amp.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	amp "github.com/aws/aws-sdk-go-v2/service/amp"
	gomock "github.com/golang/mock/gomock"
)

// MockAmpClient is a mock of AmpClient interface.
type MockAmpClient struct {
	ctrl     *gomock.Controller
	recorder *MockAmpClientMockRecorder
}

// MockAmpClientMockRecorder is the mock recorder for MockAmpClient.
type MockAmpClientMockRecorder struct {
	mock *MockAmpClient
}

// NewMockAmpClient creates a new mock instance.
func NewMockAmpClient(ctrl *gomock.Controller) *MockAmpClient {
	mock := &MockAmpClient{ctrl: ctrl}
	mock.recorder = &MockAmpClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAmpClient) EXPECT() *MockAmpClientMockRecorder {
	return m.recorder
}

// DescribeAlertManagerDefinition mocks base method.
func (m *MockAmpClient) DescribeAlertManagerDefinition(arg0 context.Context, arg1 *amp.DescribeAlertManagerDefinitionInput, arg2 ...func(*amp.Options)) (*amp.DescribeAlertManagerDefinitionOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &amp.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAlertManagerDefinition")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAlertManagerDefinition", varargs...)
	ret0, _ := ret[0].(*amp.DescribeAlertManagerDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAlertManagerDefinition indicates an expected call of DescribeAlertManagerDefinition.
func (mr *MockAmpClientMockRecorder) DescribeAlertManagerDefinition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAlertManagerDefinition", reflect.TypeOf((*MockAmpClient)(nil).DescribeAlertManagerDefinition), varargs...)
}

// DescribeLoggingConfiguration mocks base method.
func (m *MockAmpClient) DescribeLoggingConfiguration(arg0 context.Context, arg1 *amp.DescribeLoggingConfigurationInput, arg2 ...func(*amp.Options)) (*amp.DescribeLoggingConfigurationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &amp.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeLoggingConfiguration")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLoggingConfiguration", varargs...)
	ret0, _ := ret[0].(*amp.DescribeLoggingConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLoggingConfiguration indicates an expected call of DescribeLoggingConfiguration.
func (mr *MockAmpClientMockRecorder) DescribeLoggingConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoggingConfiguration", reflect.TypeOf((*MockAmpClient)(nil).DescribeLoggingConfiguration), varargs...)
}

// DescribeRuleGroupsNamespace mocks base method.
func (m *MockAmpClient) DescribeRuleGroupsNamespace(arg0 context.Context, arg1 *amp.DescribeRuleGroupsNamespaceInput, arg2 ...func(*amp.Options)) (*amp.DescribeRuleGroupsNamespaceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &amp.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeRuleGroupsNamespace")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRuleGroupsNamespace", varargs...)
	ret0, _ := ret[0].(*amp.DescribeRuleGroupsNamespaceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRuleGroupsNamespace indicates an expected call of DescribeRuleGroupsNamespace.
func (mr *MockAmpClientMockRecorder) DescribeRuleGroupsNamespace(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRuleGroupsNamespace", reflect.TypeOf((*MockAmpClient)(nil).DescribeRuleGroupsNamespace), varargs...)
}

// DescribeWorkspace mocks base method.
func (m *MockAmpClient) DescribeWorkspace(arg0 context.Context, arg1 *amp.DescribeWorkspaceInput, arg2 ...func(*amp.Options)) (*amp.DescribeWorkspaceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &amp.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeWorkspace")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeWorkspace", varargs...)
	ret0, _ := ret[0].(*amp.DescribeWorkspaceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeWorkspace indicates an expected call of DescribeWorkspace.
func (mr *MockAmpClientMockRecorder) DescribeWorkspace(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeWorkspace", reflect.TypeOf((*MockAmpClient)(nil).DescribeWorkspace), varargs...)
}

// ListRuleGroupsNamespaces mocks base method.
func (m *MockAmpClient) ListRuleGroupsNamespaces(arg0 context.Context, arg1 *amp.ListRuleGroupsNamespacesInput, arg2 ...func(*amp.Options)) (*amp.ListRuleGroupsNamespacesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &amp.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListRuleGroupsNamespaces")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRuleGroupsNamespaces", varargs...)
	ret0, _ := ret[0].(*amp.ListRuleGroupsNamespacesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRuleGroupsNamespaces indicates an expected call of ListRuleGroupsNamespaces.
func (mr *MockAmpClientMockRecorder) ListRuleGroupsNamespaces(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRuleGroupsNamespaces", reflect.TypeOf((*MockAmpClient)(nil).ListRuleGroupsNamespaces), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockAmpClient) ListTagsForResource(arg0 context.Context, arg1 *amp.ListTagsForResourceInput, arg2 ...func(*amp.Options)) (*amp.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &amp.Options{}
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
	ret0, _ := ret[0].(*amp.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockAmpClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockAmpClient)(nil).ListTagsForResource), varargs...)
}

// ListWorkspaces mocks base method.
func (m *MockAmpClient) ListWorkspaces(arg0 context.Context, arg1 *amp.ListWorkspacesInput, arg2 ...func(*amp.Options)) (*amp.ListWorkspacesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &amp.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListWorkspaces")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListWorkspaces", varargs...)
	ret0, _ := ret[0].(*amp.ListWorkspacesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWorkspaces indicates an expected call of ListWorkspaces.
func (mr *MockAmpClientMockRecorder) ListWorkspaces(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkspaces", reflect.TypeOf((*MockAmpClient)(nil).ListWorkspaces), varargs...)
}
