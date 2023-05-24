// Code generated by MockGen. DO NOT EDIT.
// Source: elasticbeanstalk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	
	reflect "reflect"

	elasticbeanstalk "github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	gomock "github.com/golang/mock/gomock"
)

// MockElasticbeanstalkClient is a mock of ElasticbeanstalkClient interface.
type MockElasticbeanstalkClient struct {
	ctrl     *gomock.Controller
	recorder *MockElasticbeanstalkClientMockRecorder
}

// MockElasticbeanstalkClientMockRecorder is the mock recorder for MockElasticbeanstalkClient.
type MockElasticbeanstalkClientMockRecorder struct {
	mock *MockElasticbeanstalkClient
}

// NewMockElasticbeanstalkClient creates a new mock instance.
func NewMockElasticbeanstalkClient(ctrl *gomock.Controller) *MockElasticbeanstalkClient {
	mock := &MockElasticbeanstalkClient{ctrl: ctrl}
	mock.recorder = &MockElasticbeanstalkClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockElasticbeanstalkClient) EXPECT() *MockElasticbeanstalkClientMockRecorder {
	return m.recorder
}

// DescribeAccountAttributes mocks base method.
func (m *MockElasticbeanstalkClient) DescribeAccountAttributes(arg0 context.Context, arg1 *elasticbeanstalk.DescribeAccountAttributesInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeAccountAttributesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticbeanstalk.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAccountAttributes")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAccountAttributes", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeAccountAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAccountAttributes indicates an expected call of DescribeAccountAttributes.
func (mr *MockElasticbeanstalkClientMockRecorder) DescribeAccountAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAccountAttributes", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeAccountAttributes), varargs...)
}

// DescribeApplicationVersions mocks base method.
func (m *MockElasticbeanstalkClient) DescribeApplicationVersions(arg0 context.Context, arg1 *elasticbeanstalk.DescribeApplicationVersionsInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeApplicationVersionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticbeanstalk.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeApplicationVersions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeApplicationVersions", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeApplicationVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeApplicationVersions indicates an expected call of DescribeApplicationVersions.
func (mr *MockElasticbeanstalkClientMockRecorder) DescribeApplicationVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeApplicationVersions", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeApplicationVersions), varargs...)
}

// DescribeApplications mocks base method.
func (m *MockElasticbeanstalkClient) DescribeApplications(arg0 context.Context, arg1 *elasticbeanstalk.DescribeApplicationsInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeApplicationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticbeanstalk.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeApplications")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeApplications", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeApplicationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeApplications indicates an expected call of DescribeApplications.
func (mr *MockElasticbeanstalkClientMockRecorder) DescribeApplications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeApplications", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeApplications), varargs...)
}

// DescribeConfigurationOptions mocks base method.
func (m *MockElasticbeanstalkClient) DescribeConfigurationOptions(arg0 context.Context, arg1 *elasticbeanstalk.DescribeConfigurationOptionsInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeConfigurationOptionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticbeanstalk.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeConfigurationOptions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConfigurationOptions", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeConfigurationOptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConfigurationOptions indicates an expected call of DescribeConfigurationOptions.
func (mr *MockElasticbeanstalkClientMockRecorder) DescribeConfigurationOptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConfigurationOptions", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeConfigurationOptions), varargs...)
}

// DescribeConfigurationSettings mocks base method.
func (m *MockElasticbeanstalkClient) DescribeConfigurationSettings(arg0 context.Context, arg1 *elasticbeanstalk.DescribeConfigurationSettingsInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeConfigurationSettingsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticbeanstalk.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeConfigurationSettings")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConfigurationSettings", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeConfigurationSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConfigurationSettings indicates an expected call of DescribeConfigurationSettings.
func (mr *MockElasticbeanstalkClientMockRecorder) DescribeConfigurationSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConfigurationSettings", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeConfigurationSettings), varargs...)
}

// DescribeEnvironmentHealth mocks base method.
func (m *MockElasticbeanstalkClient) DescribeEnvironmentHealth(arg0 context.Context, arg1 *elasticbeanstalk.DescribeEnvironmentHealthInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeEnvironmentHealthOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticbeanstalk.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeEnvironmentHealth")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEnvironmentHealth", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeEnvironmentHealthOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEnvironmentHealth indicates an expected call of DescribeEnvironmentHealth.
func (mr *MockElasticbeanstalkClientMockRecorder) DescribeEnvironmentHealth(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEnvironmentHealth", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeEnvironmentHealth), varargs...)
}

// DescribeEnvironmentManagedActionHistory mocks base method.
func (m *MockElasticbeanstalkClient) DescribeEnvironmentManagedActionHistory(arg0 context.Context, arg1 *elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticbeanstalk.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeEnvironmentManagedActionHistory")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEnvironmentManagedActionHistory", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEnvironmentManagedActionHistory indicates an expected call of DescribeEnvironmentManagedActionHistory.
func (mr *MockElasticbeanstalkClientMockRecorder) DescribeEnvironmentManagedActionHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEnvironmentManagedActionHistory", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeEnvironmentManagedActionHistory), varargs...)
}

// DescribeEnvironmentManagedActions mocks base method.
func (m *MockElasticbeanstalkClient) DescribeEnvironmentManagedActions(arg0 context.Context, arg1 *elasticbeanstalk.DescribeEnvironmentManagedActionsInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeEnvironmentManagedActionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticbeanstalk.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeEnvironmentManagedActions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEnvironmentManagedActions", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeEnvironmentManagedActionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEnvironmentManagedActions indicates an expected call of DescribeEnvironmentManagedActions.
func (mr *MockElasticbeanstalkClientMockRecorder) DescribeEnvironmentManagedActions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEnvironmentManagedActions", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeEnvironmentManagedActions), varargs...)
}

// DescribeEnvironmentResources mocks base method.
func (m *MockElasticbeanstalkClient) DescribeEnvironmentResources(arg0 context.Context, arg1 *elasticbeanstalk.DescribeEnvironmentResourcesInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeEnvironmentResourcesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticbeanstalk.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeEnvironmentResources")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEnvironmentResources", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeEnvironmentResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEnvironmentResources indicates an expected call of DescribeEnvironmentResources.
func (mr *MockElasticbeanstalkClientMockRecorder) DescribeEnvironmentResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEnvironmentResources", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeEnvironmentResources), varargs...)
}

// DescribeEnvironments mocks base method.
func (m *MockElasticbeanstalkClient) DescribeEnvironments(arg0 context.Context, arg1 *elasticbeanstalk.DescribeEnvironmentsInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeEnvironmentsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticbeanstalk.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeEnvironments")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEnvironments", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeEnvironmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEnvironments indicates an expected call of DescribeEnvironments.
func (mr *MockElasticbeanstalkClientMockRecorder) DescribeEnvironments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEnvironments", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeEnvironments), varargs...)
}

// DescribeEvents mocks base method.
func (m *MockElasticbeanstalkClient) DescribeEvents(arg0 context.Context, arg1 *elasticbeanstalk.DescribeEventsInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeEventsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticbeanstalk.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeEvents")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEvents", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEvents indicates an expected call of DescribeEvents.
func (mr *MockElasticbeanstalkClientMockRecorder) DescribeEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEvents", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeEvents), varargs...)
}

// DescribeInstancesHealth mocks base method.
func (m *MockElasticbeanstalkClient) DescribeInstancesHealth(arg0 context.Context, arg1 *elasticbeanstalk.DescribeInstancesHealthInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeInstancesHealthOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticbeanstalk.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeInstancesHealth")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeInstancesHealth", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeInstancesHealthOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstancesHealth indicates an expected call of DescribeInstancesHealth.
func (mr *MockElasticbeanstalkClientMockRecorder) DescribeInstancesHealth(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstancesHealth", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeInstancesHealth), varargs...)
}

// DescribePlatformVersion mocks base method.
func (m *MockElasticbeanstalkClient) DescribePlatformVersion(arg0 context.Context, arg1 *elasticbeanstalk.DescribePlatformVersionInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribePlatformVersionOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticbeanstalk.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribePlatformVersion")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribePlatformVersion", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribePlatformVersionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribePlatformVersion indicates an expected call of DescribePlatformVersion.
func (mr *MockElasticbeanstalkClientMockRecorder) DescribePlatformVersion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribePlatformVersion", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribePlatformVersion), varargs...)
}

// ListAvailableSolutionStacks mocks base method.
func (m *MockElasticbeanstalkClient) ListAvailableSolutionStacks(arg0 context.Context, arg1 *elasticbeanstalk.ListAvailableSolutionStacksInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.ListAvailableSolutionStacksOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticbeanstalk.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAvailableSolutionStacks")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAvailableSolutionStacks", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.ListAvailableSolutionStacksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAvailableSolutionStacks indicates an expected call of ListAvailableSolutionStacks.
func (mr *MockElasticbeanstalkClientMockRecorder) ListAvailableSolutionStacks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAvailableSolutionStacks", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).ListAvailableSolutionStacks), varargs...)
}

// ListPlatformBranches mocks base method.
func (m *MockElasticbeanstalkClient) ListPlatformBranches(arg0 context.Context, arg1 *elasticbeanstalk.ListPlatformBranchesInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.ListPlatformBranchesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticbeanstalk.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListPlatformBranches")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPlatformBranches", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.ListPlatformBranchesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPlatformBranches indicates an expected call of ListPlatformBranches.
func (mr *MockElasticbeanstalkClientMockRecorder) ListPlatformBranches(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPlatformBranches", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).ListPlatformBranches), varargs...)
}

// ListPlatformVersions mocks base method.
func (m *MockElasticbeanstalkClient) ListPlatformVersions(arg0 context.Context, arg1 *elasticbeanstalk.ListPlatformVersionsInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.ListPlatformVersionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticbeanstalk.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListPlatformVersions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPlatformVersions", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.ListPlatformVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPlatformVersions indicates an expected call of ListPlatformVersions.
func (mr *MockElasticbeanstalkClientMockRecorder) ListPlatformVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPlatformVersions", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).ListPlatformVersions), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockElasticbeanstalkClient) ListTagsForResource(arg0 context.Context, arg1 *elasticbeanstalk.ListTagsForResourceInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticbeanstalk.Options{}
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
	ret0, _ := ret[0].(*elasticbeanstalk.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockElasticbeanstalkClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).ListTagsForResource), varargs...)
}
