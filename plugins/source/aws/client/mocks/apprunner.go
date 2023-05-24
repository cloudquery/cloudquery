// Code generated by MockGen. DO NOT EDIT.
// Source: apprunner.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	
	reflect "reflect"

	apprunner "github.com/aws/aws-sdk-go-v2/service/apprunner"
	gomock "github.com/golang/mock/gomock"
)

// MockApprunnerClient is a mock of ApprunnerClient interface.
type MockApprunnerClient struct {
	ctrl     *gomock.Controller
	recorder *MockApprunnerClientMockRecorder
}

// MockApprunnerClientMockRecorder is the mock recorder for MockApprunnerClient.
type MockApprunnerClientMockRecorder struct {
	mock *MockApprunnerClient
}

// NewMockApprunnerClient creates a new mock instance.
func NewMockApprunnerClient(ctrl *gomock.Controller) *MockApprunnerClient {
	mock := &MockApprunnerClient{ctrl: ctrl}
	mock.recorder = &MockApprunnerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApprunnerClient) EXPECT() *MockApprunnerClientMockRecorder {
	return m.recorder
}

// DescribeAutoScalingConfiguration mocks base method.
func (m *MockApprunnerClient) DescribeAutoScalingConfiguration(arg0 context.Context, arg1 *apprunner.DescribeAutoScalingConfigurationInput, arg2 ...func(*apprunner.Options)) (*apprunner.DescribeAutoScalingConfigurationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &apprunner.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAutoScalingConfiguration")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAutoScalingConfiguration", varargs...)
	ret0, _ := ret[0].(*apprunner.DescribeAutoScalingConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAutoScalingConfiguration indicates an expected call of DescribeAutoScalingConfiguration.
func (mr *MockApprunnerClientMockRecorder) DescribeAutoScalingConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAutoScalingConfiguration", reflect.TypeOf((*MockApprunnerClient)(nil).DescribeAutoScalingConfiguration), varargs...)
}

// DescribeCustomDomains mocks base method.
func (m *MockApprunnerClient) DescribeCustomDomains(arg0 context.Context, arg1 *apprunner.DescribeCustomDomainsInput, arg2 ...func(*apprunner.Options)) (*apprunner.DescribeCustomDomainsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &apprunner.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeCustomDomains")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCustomDomains", varargs...)
	ret0, _ := ret[0].(*apprunner.DescribeCustomDomainsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCustomDomains indicates an expected call of DescribeCustomDomains.
func (mr *MockApprunnerClientMockRecorder) DescribeCustomDomains(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCustomDomains", reflect.TypeOf((*MockApprunnerClient)(nil).DescribeCustomDomains), varargs...)
}

// DescribeObservabilityConfiguration mocks base method.
func (m *MockApprunnerClient) DescribeObservabilityConfiguration(arg0 context.Context, arg1 *apprunner.DescribeObservabilityConfigurationInput, arg2 ...func(*apprunner.Options)) (*apprunner.DescribeObservabilityConfigurationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &apprunner.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeObservabilityConfiguration")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeObservabilityConfiguration", varargs...)
	ret0, _ := ret[0].(*apprunner.DescribeObservabilityConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeObservabilityConfiguration indicates an expected call of DescribeObservabilityConfiguration.
func (mr *MockApprunnerClientMockRecorder) DescribeObservabilityConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeObservabilityConfiguration", reflect.TypeOf((*MockApprunnerClient)(nil).DescribeObservabilityConfiguration), varargs...)
}

// DescribeService mocks base method.
func (m *MockApprunnerClient) DescribeService(arg0 context.Context, arg1 *apprunner.DescribeServiceInput, arg2 ...func(*apprunner.Options)) (*apprunner.DescribeServiceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &apprunner.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeService")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeService", varargs...)
	ret0, _ := ret[0].(*apprunner.DescribeServiceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeService indicates an expected call of DescribeService.
func (mr *MockApprunnerClientMockRecorder) DescribeService(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeService", reflect.TypeOf((*MockApprunnerClient)(nil).DescribeService), varargs...)
}

// DescribeVpcConnector mocks base method.
func (m *MockApprunnerClient) DescribeVpcConnector(arg0 context.Context, arg1 *apprunner.DescribeVpcConnectorInput, arg2 ...func(*apprunner.Options)) (*apprunner.DescribeVpcConnectorOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &apprunner.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeVpcConnector")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVpcConnector", varargs...)
	ret0, _ := ret[0].(*apprunner.DescribeVpcConnectorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVpcConnector indicates an expected call of DescribeVpcConnector.
func (mr *MockApprunnerClientMockRecorder) DescribeVpcConnector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpcConnector", reflect.TypeOf((*MockApprunnerClient)(nil).DescribeVpcConnector), varargs...)
}

// DescribeVpcIngressConnection mocks base method.
func (m *MockApprunnerClient) DescribeVpcIngressConnection(arg0 context.Context, arg1 *apprunner.DescribeVpcIngressConnectionInput, arg2 ...func(*apprunner.Options)) (*apprunner.DescribeVpcIngressConnectionOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &apprunner.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeVpcIngressConnection")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVpcIngressConnection", varargs...)
	ret0, _ := ret[0].(*apprunner.DescribeVpcIngressConnectionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVpcIngressConnection indicates an expected call of DescribeVpcIngressConnection.
func (mr *MockApprunnerClientMockRecorder) DescribeVpcIngressConnection(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpcIngressConnection", reflect.TypeOf((*MockApprunnerClient)(nil).DescribeVpcIngressConnection), varargs...)
}

// ListAutoScalingConfigurations mocks base method.
func (m *MockApprunnerClient) ListAutoScalingConfigurations(arg0 context.Context, arg1 *apprunner.ListAutoScalingConfigurationsInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListAutoScalingConfigurationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &apprunner.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAutoScalingConfigurations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAutoScalingConfigurations", varargs...)
	ret0, _ := ret[0].(*apprunner.ListAutoScalingConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAutoScalingConfigurations indicates an expected call of ListAutoScalingConfigurations.
func (mr *MockApprunnerClientMockRecorder) ListAutoScalingConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAutoScalingConfigurations", reflect.TypeOf((*MockApprunnerClient)(nil).ListAutoScalingConfigurations), varargs...)
}

// ListConnections mocks base method.
func (m *MockApprunnerClient) ListConnections(arg0 context.Context, arg1 *apprunner.ListConnectionsInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListConnectionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &apprunner.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListConnections")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListConnections", varargs...)
	ret0, _ := ret[0].(*apprunner.ListConnectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConnections indicates an expected call of ListConnections.
func (mr *MockApprunnerClientMockRecorder) ListConnections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConnections", reflect.TypeOf((*MockApprunnerClient)(nil).ListConnections), varargs...)
}

// ListObservabilityConfigurations mocks base method.
func (m *MockApprunnerClient) ListObservabilityConfigurations(arg0 context.Context, arg1 *apprunner.ListObservabilityConfigurationsInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListObservabilityConfigurationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &apprunner.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListObservabilityConfigurations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListObservabilityConfigurations", varargs...)
	ret0, _ := ret[0].(*apprunner.ListObservabilityConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListObservabilityConfigurations indicates an expected call of ListObservabilityConfigurations.
func (mr *MockApprunnerClientMockRecorder) ListObservabilityConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObservabilityConfigurations", reflect.TypeOf((*MockApprunnerClient)(nil).ListObservabilityConfigurations), varargs...)
}

// ListOperations mocks base method.
func (m *MockApprunnerClient) ListOperations(arg0 context.Context, arg1 *apprunner.ListOperationsInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListOperationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &apprunner.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListOperations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOperations", varargs...)
	ret0, _ := ret[0].(*apprunner.ListOperationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOperations indicates an expected call of ListOperations.
func (mr *MockApprunnerClientMockRecorder) ListOperations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOperations", reflect.TypeOf((*MockApprunnerClient)(nil).ListOperations), varargs...)
}

// ListServices mocks base method.
func (m *MockApprunnerClient) ListServices(arg0 context.Context, arg1 *apprunner.ListServicesInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListServicesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &apprunner.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListServices")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListServices", varargs...)
	ret0, _ := ret[0].(*apprunner.ListServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices.
func (mr *MockApprunnerClientMockRecorder) ListServices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockApprunnerClient)(nil).ListServices), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockApprunnerClient) ListTagsForResource(arg0 context.Context, arg1 *apprunner.ListTagsForResourceInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &apprunner.Options{}
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
	ret0, _ := ret[0].(*apprunner.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockApprunnerClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockApprunnerClient)(nil).ListTagsForResource), varargs...)
}

// ListVpcConnectors mocks base method.
func (m *MockApprunnerClient) ListVpcConnectors(arg0 context.Context, arg1 *apprunner.ListVpcConnectorsInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListVpcConnectorsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &apprunner.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListVpcConnectors")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListVpcConnectors", varargs...)
	ret0, _ := ret[0].(*apprunner.ListVpcConnectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVpcConnectors indicates an expected call of ListVpcConnectors.
func (mr *MockApprunnerClientMockRecorder) ListVpcConnectors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVpcConnectors", reflect.TypeOf((*MockApprunnerClient)(nil).ListVpcConnectors), varargs...)
}

// ListVpcIngressConnections mocks base method.
func (m *MockApprunnerClient) ListVpcIngressConnections(arg0 context.Context, arg1 *apprunner.ListVpcIngressConnectionsInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListVpcIngressConnectionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &apprunner.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListVpcIngressConnections")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListVpcIngressConnections", varargs...)
	ret0, _ := ret[0].(*apprunner.ListVpcIngressConnectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVpcIngressConnections indicates an expected call of ListVpcIngressConnections.
func (mr *MockApprunnerClientMockRecorder) ListVpcIngressConnections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVpcIngressConnections", reflect.TypeOf((*MockApprunnerClient)(nil).ListVpcIngressConnections), varargs...)
}
