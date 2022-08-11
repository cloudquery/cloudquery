// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/aws/client (interfaces: ConfigServiceClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	configservice "github.com/aws/aws-sdk-go-v2/service/configservice"
	gomock "github.com/golang/mock/gomock"
)

// MockConfigServiceClient is a mock of ConfigServiceClient interface.
type MockConfigServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockConfigServiceClientMockRecorder
}

// MockConfigServiceClientMockRecorder is the mock recorder for MockConfigServiceClient.
type MockConfigServiceClientMockRecorder struct {
	mock *MockConfigServiceClient
}

// NewMockConfigServiceClient creates a new mock instance.
func NewMockConfigServiceClient(ctrl *gomock.Controller) *MockConfigServiceClient {
	mock := &MockConfigServiceClient{ctrl: ctrl}
	mock.recorder = &MockConfigServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigServiceClient) EXPECT() *MockConfigServiceClientMockRecorder {
	return m.recorder
}

// DescribeConfigurationRecorderStatus mocks base method.
func (m *MockConfigServiceClient) DescribeConfigurationRecorderStatus(arg0 context.Context, arg1 *configservice.DescribeConfigurationRecorderStatusInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeConfigurationRecorderStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConfigurationRecorderStatus", varargs...)
	ret0, _ := ret[0].(*configservice.DescribeConfigurationRecorderStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConfigurationRecorderStatus indicates an expected call of DescribeConfigurationRecorderStatus.
func (mr *MockConfigServiceClientMockRecorder) DescribeConfigurationRecorderStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConfigurationRecorderStatus", reflect.TypeOf((*MockConfigServiceClient)(nil).DescribeConfigurationRecorderStatus), varargs...)
}

// DescribeConfigurationRecorders mocks base method.
func (m *MockConfigServiceClient) DescribeConfigurationRecorders(arg0 context.Context, arg1 *configservice.DescribeConfigurationRecordersInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeConfigurationRecordersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConfigurationRecorders", varargs...)
	ret0, _ := ret[0].(*configservice.DescribeConfigurationRecordersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConfigurationRecorders indicates an expected call of DescribeConfigurationRecorders.
func (mr *MockConfigServiceClientMockRecorder) DescribeConfigurationRecorders(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConfigurationRecorders", reflect.TypeOf((*MockConfigServiceClient)(nil).DescribeConfigurationRecorders), varargs...)
}

// DescribeConformancePackCompliance mocks base method.
func (m *MockConfigServiceClient) DescribeConformancePackCompliance(arg0 context.Context, arg1 *configservice.DescribeConformancePackComplianceInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeConformancePackComplianceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConformancePackCompliance", varargs...)
	ret0, _ := ret[0].(*configservice.DescribeConformancePackComplianceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConformancePackCompliance indicates an expected call of DescribeConformancePackCompliance.
func (mr *MockConfigServiceClientMockRecorder) DescribeConformancePackCompliance(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConformancePackCompliance", reflect.TypeOf((*MockConfigServiceClient)(nil).DescribeConformancePackCompliance), varargs...)
}

// DescribeConformancePacks mocks base method.
func (m *MockConfigServiceClient) DescribeConformancePacks(arg0 context.Context, arg1 *configservice.DescribeConformancePacksInput, arg2 ...func(*configservice.Options)) (*configservice.DescribeConformancePacksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConformancePacks", varargs...)
	ret0, _ := ret[0].(*configservice.DescribeConformancePacksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConformancePacks indicates an expected call of DescribeConformancePacks.
func (mr *MockConfigServiceClientMockRecorder) DescribeConformancePacks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConformancePacks", reflect.TypeOf((*MockConfigServiceClient)(nil).DescribeConformancePacks), varargs...)
}

// GetConformancePackComplianceDetails mocks base method.
func (m *MockConfigServiceClient) GetConformancePackComplianceDetails(arg0 context.Context, arg1 *configservice.GetConformancePackComplianceDetailsInput, arg2 ...func(*configservice.Options)) (*configservice.GetConformancePackComplianceDetailsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConformancePackComplianceDetails", varargs...)
	ret0, _ := ret[0].(*configservice.GetConformancePackComplianceDetailsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConformancePackComplianceDetails indicates an expected call of GetConformancePackComplianceDetails.
func (mr *MockConfigServiceClientMockRecorder) GetConformancePackComplianceDetails(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConformancePackComplianceDetails", reflect.TypeOf((*MockConfigServiceClient)(nil).GetConformancePackComplianceDetails), varargs...)
}
