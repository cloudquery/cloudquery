// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/aws/client (interfaces: LambdaClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	lambda "github.com/aws/aws-sdk-go-v2/service/lambda"
	gomock "github.com/golang/mock/gomock"
)

// MockLambdaClient is a mock of LambdaClient interface.
type MockLambdaClient struct {
	ctrl     *gomock.Controller
	recorder *MockLambdaClientMockRecorder
}

// MockLambdaClientMockRecorder is the mock recorder for MockLambdaClient.
type MockLambdaClientMockRecorder struct {
	mock *MockLambdaClient
}

// NewMockLambdaClient creates a new mock instance.
func NewMockLambdaClient(ctrl *gomock.Controller) *MockLambdaClient {
	mock := &MockLambdaClient{ctrl: ctrl}
	mock.recorder = &MockLambdaClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLambdaClient) EXPECT() *MockLambdaClientMockRecorder {
	return m.recorder
}

// GetCodeSigningConfig mocks base method.
func (m *MockLambdaClient) GetCodeSigningConfig(arg0 context.Context, arg1 *lambda.GetCodeSigningConfigInput, arg2 ...func(*lambda.Options)) (*lambda.GetCodeSigningConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCodeSigningConfig", varargs...)
	ret0, _ := ret[0].(*lambda.GetCodeSigningConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCodeSigningConfig indicates an expected call of GetCodeSigningConfig.
func (mr *MockLambdaClientMockRecorder) GetCodeSigningConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCodeSigningConfig", reflect.TypeOf((*MockLambdaClient)(nil).GetCodeSigningConfig), varargs...)
}

// GetFunction mocks base method.
func (m *MockLambdaClient) GetFunction(arg0 context.Context, arg1 *lambda.GetFunctionInput, arg2 ...func(*lambda.Options)) (*lambda.GetFunctionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFunction", varargs...)
	ret0, _ := ret[0].(*lambda.GetFunctionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFunction indicates an expected call of GetFunction.
func (mr *MockLambdaClientMockRecorder) GetFunction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFunction", reflect.TypeOf((*MockLambdaClient)(nil).GetFunction), varargs...)
}

// GetFunctionCodeSigningConfig mocks base method.
func (m *MockLambdaClient) GetFunctionCodeSigningConfig(arg0 context.Context, arg1 *lambda.GetFunctionCodeSigningConfigInput, arg2 ...func(*lambda.Options)) (*lambda.GetFunctionCodeSigningConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFunctionCodeSigningConfig", varargs...)
	ret0, _ := ret[0].(*lambda.GetFunctionCodeSigningConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFunctionCodeSigningConfig indicates an expected call of GetFunctionCodeSigningConfig.
func (mr *MockLambdaClientMockRecorder) GetFunctionCodeSigningConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFunctionCodeSigningConfig", reflect.TypeOf((*MockLambdaClient)(nil).GetFunctionCodeSigningConfig), varargs...)
}

// GetFunctionUrlConfig mocks base method.
func (m *MockLambdaClient) GetFunctionUrlConfig(arg0 context.Context, arg1 *lambda.GetFunctionUrlConfigInput, arg2 ...func(*lambda.Options)) (*lambda.GetFunctionUrlConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFunctionUrlConfig", varargs...)
	ret0, _ := ret[0].(*lambda.GetFunctionUrlConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFunctionUrlConfig indicates an expected call of GetFunctionUrlConfig.
func (mr *MockLambdaClientMockRecorder) GetFunctionUrlConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFunctionUrlConfig", reflect.TypeOf((*MockLambdaClient)(nil).GetFunctionUrlConfig), varargs...)
}

// GetLayerVersionPolicy mocks base method.
func (m *MockLambdaClient) GetLayerVersionPolicy(arg0 context.Context, arg1 *lambda.GetLayerVersionPolicyInput, arg2 ...func(*lambda.Options)) (*lambda.GetLayerVersionPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLayerVersionPolicy", varargs...)
	ret0, _ := ret[0].(*lambda.GetLayerVersionPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLayerVersionPolicy indicates an expected call of GetLayerVersionPolicy.
func (mr *MockLambdaClientMockRecorder) GetLayerVersionPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLayerVersionPolicy", reflect.TypeOf((*MockLambdaClient)(nil).GetLayerVersionPolicy), varargs...)
}

// GetPolicy mocks base method.
func (m *MockLambdaClient) GetPolicy(arg0 context.Context, arg1 *lambda.GetPolicyInput, arg2 ...func(*lambda.Options)) (*lambda.GetPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPolicy", varargs...)
	ret0, _ := ret[0].(*lambda.GetPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPolicy indicates an expected call of GetPolicy.
func (mr *MockLambdaClientMockRecorder) GetPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicy", reflect.TypeOf((*MockLambdaClient)(nil).GetPolicy), varargs...)
}

// ListAliases mocks base method.
func (m *MockLambdaClient) ListAliases(arg0 context.Context, arg1 *lambda.ListAliasesInput, arg2 ...func(*lambda.Options)) (*lambda.ListAliasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAliases", varargs...)
	ret0, _ := ret[0].(*lambda.ListAliasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAliases indicates an expected call of ListAliases.
func (mr *MockLambdaClientMockRecorder) ListAliases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAliases", reflect.TypeOf((*MockLambdaClient)(nil).ListAliases), varargs...)
}

// ListEventSourceMappings mocks base method.
func (m *MockLambdaClient) ListEventSourceMappings(arg0 context.Context, arg1 *lambda.ListEventSourceMappingsInput, arg2 ...func(*lambda.Options)) (*lambda.ListEventSourceMappingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEventSourceMappings", varargs...)
	ret0, _ := ret[0].(*lambda.ListEventSourceMappingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEventSourceMappings indicates an expected call of ListEventSourceMappings.
func (mr *MockLambdaClientMockRecorder) ListEventSourceMappings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEventSourceMappings", reflect.TypeOf((*MockLambdaClient)(nil).ListEventSourceMappings), varargs...)
}

// ListFunctionEventInvokeConfigs mocks base method.
func (m *MockLambdaClient) ListFunctionEventInvokeConfigs(arg0 context.Context, arg1 *lambda.ListFunctionEventInvokeConfigsInput, arg2 ...func(*lambda.Options)) (*lambda.ListFunctionEventInvokeConfigsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFunctionEventInvokeConfigs", varargs...)
	ret0, _ := ret[0].(*lambda.ListFunctionEventInvokeConfigsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFunctionEventInvokeConfigs indicates an expected call of ListFunctionEventInvokeConfigs.
func (mr *MockLambdaClientMockRecorder) ListFunctionEventInvokeConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFunctionEventInvokeConfigs", reflect.TypeOf((*MockLambdaClient)(nil).ListFunctionEventInvokeConfigs), varargs...)
}

// ListFunctions mocks base method.
func (m *MockLambdaClient) ListFunctions(arg0 context.Context, arg1 *lambda.ListFunctionsInput, arg2 ...func(*lambda.Options)) (*lambda.ListFunctionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFunctions", varargs...)
	ret0, _ := ret[0].(*lambda.ListFunctionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFunctions indicates an expected call of ListFunctions.
func (mr *MockLambdaClientMockRecorder) ListFunctions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFunctions", reflect.TypeOf((*MockLambdaClient)(nil).ListFunctions), varargs...)
}

// ListLayerVersions mocks base method.
func (m *MockLambdaClient) ListLayerVersions(arg0 context.Context, arg1 *lambda.ListLayerVersionsInput, arg2 ...func(*lambda.Options)) (*lambda.ListLayerVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListLayerVersions", varargs...)
	ret0, _ := ret[0].(*lambda.ListLayerVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLayerVersions indicates an expected call of ListLayerVersions.
func (mr *MockLambdaClientMockRecorder) ListLayerVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLayerVersions", reflect.TypeOf((*MockLambdaClient)(nil).ListLayerVersions), varargs...)
}

// ListLayers mocks base method.
func (m *MockLambdaClient) ListLayers(arg0 context.Context, arg1 *lambda.ListLayersInput, arg2 ...func(*lambda.Options)) (*lambda.ListLayersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListLayers", varargs...)
	ret0, _ := ret[0].(*lambda.ListLayersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLayers indicates an expected call of ListLayers.
func (mr *MockLambdaClientMockRecorder) ListLayers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLayers", reflect.TypeOf((*MockLambdaClient)(nil).ListLayers), varargs...)
}

// ListProvisionedConcurrencyConfigs mocks base method.
func (m *MockLambdaClient) ListProvisionedConcurrencyConfigs(arg0 context.Context, arg1 *lambda.ListProvisionedConcurrencyConfigsInput, arg2 ...func(*lambda.Options)) (*lambda.ListProvisionedConcurrencyConfigsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProvisionedConcurrencyConfigs", varargs...)
	ret0, _ := ret[0].(*lambda.ListProvisionedConcurrencyConfigsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProvisionedConcurrencyConfigs indicates an expected call of ListProvisionedConcurrencyConfigs.
func (mr *MockLambdaClientMockRecorder) ListProvisionedConcurrencyConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProvisionedConcurrencyConfigs", reflect.TypeOf((*MockLambdaClient)(nil).ListProvisionedConcurrencyConfigs), varargs...)
}

// ListVersionsByFunction mocks base method.
func (m *MockLambdaClient) ListVersionsByFunction(arg0 context.Context, arg1 *lambda.ListVersionsByFunctionInput, arg2 ...func(*lambda.Options)) (*lambda.ListVersionsByFunctionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListVersionsByFunction", varargs...)
	ret0, _ := ret[0].(*lambda.ListVersionsByFunctionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVersionsByFunction indicates an expected call of ListVersionsByFunction.
func (mr *MockLambdaClientMockRecorder) ListVersionsByFunction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVersionsByFunction", reflect.TypeOf((*MockLambdaClient)(nil).ListVersionsByFunction), varargs...)
}
