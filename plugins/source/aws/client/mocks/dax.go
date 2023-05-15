// Code generated by MockGen. DO NOT EDIT.
// Source: dax.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	
	reflect "reflect"

	dax "github.com/aws/aws-sdk-go-v2/service/dax"
	gomock "github.com/golang/mock/gomock"
)

// MockDaxClient is a mock of DaxClient interface.
type MockDaxClient struct {
	ctrl     *gomock.Controller
	recorder *MockDaxClientMockRecorder
}

// MockDaxClientMockRecorder is the mock recorder for MockDaxClient.
type MockDaxClientMockRecorder struct {
	mock *MockDaxClient
}

// NewMockDaxClient creates a new mock instance.
func NewMockDaxClient(ctrl *gomock.Controller) *MockDaxClient {
	mock := &MockDaxClient{ctrl: ctrl}
	mock.recorder = &MockDaxClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDaxClient) EXPECT() *MockDaxClientMockRecorder {
	return m.recorder
}

// DescribeClusters mocks base method.
func (m *MockDaxClient) DescribeClusters(arg0 context.Context, arg1 *dax.DescribeClustersInput, arg2 ...func(*dax.Options)) (*dax.DescribeClustersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &dax.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeClusters")
	}

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
func (mr *MockDaxClientMockRecorder) DescribeClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeClusters", reflect.TypeOf((*MockDaxClient)(nil).DescribeClusters), varargs...)
}

// DescribeDefaultParameters mocks base method.
func (m *MockDaxClient) DescribeDefaultParameters(arg0 context.Context, arg1 *dax.DescribeDefaultParametersInput, arg2 ...func(*dax.Options)) (*dax.DescribeDefaultParametersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &dax.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeDefaultParameters")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDefaultParameters", varargs...)
	ret0, _ := ret[0].(*dax.DescribeDefaultParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDefaultParameters indicates an expected call of DescribeDefaultParameters.
func (mr *MockDaxClientMockRecorder) DescribeDefaultParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDefaultParameters", reflect.TypeOf((*MockDaxClient)(nil).DescribeDefaultParameters), varargs...)
}

// DescribeEvents mocks base method.
func (m *MockDaxClient) DescribeEvents(arg0 context.Context, arg1 *dax.DescribeEventsInput, arg2 ...func(*dax.Options)) (*dax.DescribeEventsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &dax.Options{}
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
	ret0, _ := ret[0].(*dax.DescribeEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEvents indicates an expected call of DescribeEvents.
func (mr *MockDaxClientMockRecorder) DescribeEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEvents", reflect.TypeOf((*MockDaxClient)(nil).DescribeEvents), varargs...)
}

// DescribeParameterGroups mocks base method.
func (m *MockDaxClient) DescribeParameterGroups(arg0 context.Context, arg1 *dax.DescribeParameterGroupsInput, arg2 ...func(*dax.Options)) (*dax.DescribeParameterGroupsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &dax.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeParameterGroups")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeParameterGroups", varargs...)
	ret0, _ := ret[0].(*dax.DescribeParameterGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeParameterGroups indicates an expected call of DescribeParameterGroups.
func (mr *MockDaxClientMockRecorder) DescribeParameterGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeParameterGroups", reflect.TypeOf((*MockDaxClient)(nil).DescribeParameterGroups), varargs...)
}

// DescribeParameters mocks base method.
func (m *MockDaxClient) DescribeParameters(arg0 context.Context, arg1 *dax.DescribeParametersInput, arg2 ...func(*dax.Options)) (*dax.DescribeParametersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &dax.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeParameters")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeParameters", varargs...)
	ret0, _ := ret[0].(*dax.DescribeParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeParameters indicates an expected call of DescribeParameters.
func (mr *MockDaxClientMockRecorder) DescribeParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeParameters", reflect.TypeOf((*MockDaxClient)(nil).DescribeParameters), varargs...)
}

// DescribeSubnetGroups mocks base method.
func (m *MockDaxClient) DescribeSubnetGroups(arg0 context.Context, arg1 *dax.DescribeSubnetGroupsInput, arg2 ...func(*dax.Options)) (*dax.DescribeSubnetGroupsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &dax.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeSubnetGroups")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSubnetGroups", varargs...)
	ret0, _ := ret[0].(*dax.DescribeSubnetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSubnetGroups indicates an expected call of DescribeSubnetGroups.
func (mr *MockDaxClientMockRecorder) DescribeSubnetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSubnetGroups", reflect.TypeOf((*MockDaxClient)(nil).DescribeSubnetGroups), varargs...)
}

// ListTags mocks base method.
func (m *MockDaxClient) ListTags(arg0 context.Context, arg1 *dax.ListTagsInput, arg2 ...func(*dax.Options)) (*dax.ListTagsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &dax.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListTags")
	}

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
func (mr *MockDaxClientMockRecorder) ListTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockDaxClient)(nil).ListTags), varargs...)
}
