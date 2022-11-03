// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/aws/client (interfaces: SecretsmanagerClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	secretsmanager "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	gomock "github.com/golang/mock/gomock"
)

// MockSecretsmanagerClient is a mock of SecretsmanagerClient interface.
type MockSecretsmanagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsmanagerClientMockRecorder
}

// MockSecretsmanagerClientMockRecorder is the mock recorder for MockSecretsmanagerClient.
type MockSecretsmanagerClientMockRecorder struct {
	mock *MockSecretsmanagerClient
}

// NewMockSecretsmanagerClient creates a new mock instance.
func NewMockSecretsmanagerClient(ctrl *gomock.Controller) *MockSecretsmanagerClient {
	mock := &MockSecretsmanagerClient{ctrl: ctrl}
	mock.recorder = &MockSecretsmanagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretsmanagerClient) EXPECT() *MockSecretsmanagerClientMockRecorder {
	return m.recorder
}

// DescribeSecret mocks base method.
func (m *MockSecretsmanagerClient) DescribeSecret(arg0 context.Context, arg1 *secretsmanager.DescribeSecretInput, arg2 ...func(*secretsmanager.Options)) (*secretsmanager.DescribeSecretOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSecret", varargs...)
	ret0, _ := ret[0].(*secretsmanager.DescribeSecretOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSecret indicates an expected call of DescribeSecret.
func (mr *MockSecretsmanagerClientMockRecorder) DescribeSecret(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSecret", reflect.TypeOf((*MockSecretsmanagerClient)(nil).DescribeSecret), varargs...)
}

// GetRandomPassword mocks base method.
func (m *MockSecretsmanagerClient) GetRandomPassword(arg0 context.Context, arg1 *secretsmanager.GetRandomPasswordInput, arg2 ...func(*secretsmanager.Options)) (*secretsmanager.GetRandomPasswordOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRandomPassword", varargs...)
	ret0, _ := ret[0].(*secretsmanager.GetRandomPasswordOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRandomPassword indicates an expected call of GetRandomPassword.
func (mr *MockSecretsmanagerClientMockRecorder) GetRandomPassword(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandomPassword", reflect.TypeOf((*MockSecretsmanagerClient)(nil).GetRandomPassword), varargs...)
}

// GetResourcePolicy mocks base method.
func (m *MockSecretsmanagerClient) GetResourcePolicy(arg0 context.Context, arg1 *secretsmanager.GetResourcePolicyInput, arg2 ...func(*secretsmanager.Options)) (*secretsmanager.GetResourcePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetResourcePolicy", varargs...)
	ret0, _ := ret[0].(*secretsmanager.GetResourcePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourcePolicy indicates an expected call of GetResourcePolicy.
func (mr *MockSecretsmanagerClientMockRecorder) GetResourcePolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourcePolicy", reflect.TypeOf((*MockSecretsmanagerClient)(nil).GetResourcePolicy), varargs...)
}

// GetSecretValue mocks base method.
func (m *MockSecretsmanagerClient) GetSecretValue(arg0 context.Context, arg1 *secretsmanager.GetSecretValueInput, arg2 ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSecretValue", varargs...)
	ret0, _ := ret[0].(*secretsmanager.GetSecretValueOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretValue indicates an expected call of GetSecretValue.
func (mr *MockSecretsmanagerClientMockRecorder) GetSecretValue(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretValue", reflect.TypeOf((*MockSecretsmanagerClient)(nil).GetSecretValue), varargs...)
}

// ListSecretVersionIds mocks base method.
func (m *MockSecretsmanagerClient) ListSecretVersionIds(arg0 context.Context, arg1 *secretsmanager.ListSecretVersionIdsInput, arg2 ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretVersionIdsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSecretVersionIds", varargs...)
	ret0, _ := ret[0].(*secretsmanager.ListSecretVersionIdsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecretVersionIds indicates an expected call of ListSecretVersionIds.
func (mr *MockSecretsmanagerClientMockRecorder) ListSecretVersionIds(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecretVersionIds", reflect.TypeOf((*MockSecretsmanagerClient)(nil).ListSecretVersionIds), varargs...)
}

// ListSecrets mocks base method.
func (m *MockSecretsmanagerClient) ListSecrets(arg0 context.Context, arg1 *secretsmanager.ListSecretsInput, arg2 ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSecrets", varargs...)
	ret0, _ := ret[0].(*secretsmanager.ListSecretsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets.
func (mr *MockSecretsmanagerClientMockRecorder) ListSecrets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockSecretsmanagerClient)(nil).ListSecrets), varargs...)
}
