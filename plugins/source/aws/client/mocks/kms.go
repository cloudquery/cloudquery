// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/aws/client (interfaces: KmsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	kms "github.com/aws/aws-sdk-go-v2/service/kms"
	gomock "github.com/golang/mock/gomock"
)

// MockKmsClient is a mock of KmsClient interface.
type MockKmsClient struct {
	ctrl     *gomock.Controller
	recorder *MockKmsClientMockRecorder
}

// MockKmsClientMockRecorder is the mock recorder for MockKmsClient.
type MockKmsClientMockRecorder struct {
	mock *MockKmsClient
}

// NewMockKmsClient creates a new mock instance.
func NewMockKmsClient(ctrl *gomock.Controller) *MockKmsClient {
	mock := &MockKmsClient{ctrl: ctrl}
	mock.recorder = &MockKmsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKmsClient) EXPECT() *MockKmsClientMockRecorder {
	return m.recorder
}

// DescribeCustomKeyStores mocks base method.
func (m *MockKmsClient) DescribeCustomKeyStores(arg0 context.Context, arg1 *kms.DescribeCustomKeyStoresInput, arg2 ...func(*kms.Options)) (*kms.DescribeCustomKeyStoresOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCustomKeyStores", varargs...)
	ret0, _ := ret[0].(*kms.DescribeCustomKeyStoresOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCustomKeyStores indicates an expected call of DescribeCustomKeyStores.
func (mr *MockKmsClientMockRecorder) DescribeCustomKeyStores(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCustomKeyStores", reflect.TypeOf((*MockKmsClient)(nil).DescribeCustomKeyStores), varargs...)
}

// DescribeKey mocks base method.
func (m *MockKmsClient) DescribeKey(arg0 context.Context, arg1 *kms.DescribeKeyInput, arg2 ...func(*kms.Options)) (*kms.DescribeKeyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeKey", varargs...)
	ret0, _ := ret[0].(*kms.DescribeKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeKey indicates an expected call of DescribeKey.
func (mr *MockKmsClientMockRecorder) DescribeKey(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeKey", reflect.TypeOf((*MockKmsClient)(nil).DescribeKey), varargs...)
}

// GetKeyPolicy mocks base method.
func (m *MockKmsClient) GetKeyPolicy(arg0 context.Context, arg1 *kms.GetKeyPolicyInput, arg2 ...func(*kms.Options)) (*kms.GetKeyPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetKeyPolicy", varargs...)
	ret0, _ := ret[0].(*kms.GetKeyPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeyPolicy indicates an expected call of GetKeyPolicy.
func (mr *MockKmsClientMockRecorder) GetKeyPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeyPolicy", reflect.TypeOf((*MockKmsClient)(nil).GetKeyPolicy), varargs...)
}

// GetKeyRotationStatus mocks base method.
func (m *MockKmsClient) GetKeyRotationStatus(arg0 context.Context, arg1 *kms.GetKeyRotationStatusInput, arg2 ...func(*kms.Options)) (*kms.GetKeyRotationStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetKeyRotationStatus", varargs...)
	ret0, _ := ret[0].(*kms.GetKeyRotationStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeyRotationStatus indicates an expected call of GetKeyRotationStatus.
func (mr *MockKmsClientMockRecorder) GetKeyRotationStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeyRotationStatus", reflect.TypeOf((*MockKmsClient)(nil).GetKeyRotationStatus), varargs...)
}

// GetParametersForImport mocks base method.
func (m *MockKmsClient) GetParametersForImport(arg0 context.Context, arg1 *kms.GetParametersForImportInput, arg2 ...func(*kms.Options)) (*kms.GetParametersForImportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetParametersForImport", varargs...)
	ret0, _ := ret[0].(*kms.GetParametersForImportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParametersForImport indicates an expected call of GetParametersForImport.
func (mr *MockKmsClientMockRecorder) GetParametersForImport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParametersForImport", reflect.TypeOf((*MockKmsClient)(nil).GetParametersForImport), varargs...)
}

// GetPublicKey mocks base method.
func (m *MockKmsClient) GetPublicKey(arg0 context.Context, arg1 *kms.GetPublicKeyInput, arg2 ...func(*kms.Options)) (*kms.GetPublicKeyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPublicKey", varargs...)
	ret0, _ := ret[0].(*kms.GetPublicKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicKey indicates an expected call of GetPublicKey.
func (mr *MockKmsClientMockRecorder) GetPublicKey(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicKey", reflect.TypeOf((*MockKmsClient)(nil).GetPublicKey), varargs...)
}

// ListAliases mocks base method.
func (m *MockKmsClient) ListAliases(arg0 context.Context, arg1 *kms.ListAliasesInput, arg2 ...func(*kms.Options)) (*kms.ListAliasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAliases", varargs...)
	ret0, _ := ret[0].(*kms.ListAliasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAliases indicates an expected call of ListAliases.
func (mr *MockKmsClientMockRecorder) ListAliases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAliases", reflect.TypeOf((*MockKmsClient)(nil).ListAliases), varargs...)
}

// ListGrants mocks base method.
func (m *MockKmsClient) ListGrants(arg0 context.Context, arg1 *kms.ListGrantsInput, arg2 ...func(*kms.Options)) (*kms.ListGrantsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGrants", varargs...)
	ret0, _ := ret[0].(*kms.ListGrantsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGrants indicates an expected call of ListGrants.
func (mr *MockKmsClientMockRecorder) ListGrants(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGrants", reflect.TypeOf((*MockKmsClient)(nil).ListGrants), varargs...)
}

// ListKeyPolicies mocks base method.
func (m *MockKmsClient) ListKeyPolicies(arg0 context.Context, arg1 *kms.ListKeyPoliciesInput, arg2 ...func(*kms.Options)) (*kms.ListKeyPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListKeyPolicies", varargs...)
	ret0, _ := ret[0].(*kms.ListKeyPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKeyPolicies indicates an expected call of ListKeyPolicies.
func (mr *MockKmsClientMockRecorder) ListKeyPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeyPolicies", reflect.TypeOf((*MockKmsClient)(nil).ListKeyPolicies), varargs...)
}

// ListKeys mocks base method.
func (m *MockKmsClient) ListKeys(arg0 context.Context, arg1 *kms.ListKeysInput, arg2 ...func(*kms.Options)) (*kms.ListKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListKeys", varargs...)
	ret0, _ := ret[0].(*kms.ListKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKeys indicates an expected call of ListKeys.
func (mr *MockKmsClientMockRecorder) ListKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockKmsClient)(nil).ListKeys), varargs...)
}

// ListResourceTags mocks base method.
func (m *MockKmsClient) ListResourceTags(arg0 context.Context, arg1 *kms.ListResourceTagsInput, arg2 ...func(*kms.Options)) (*kms.ListResourceTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListResourceTags", varargs...)
	ret0, _ := ret[0].(*kms.ListResourceTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResourceTags indicates an expected call of ListResourceTags.
func (mr *MockKmsClientMockRecorder) ListResourceTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceTags", reflect.TypeOf((*MockKmsClient)(nil).ListResourceTags), varargs...)
}

// ListRetirableGrants mocks base method.
func (m *MockKmsClient) ListRetirableGrants(arg0 context.Context, arg1 *kms.ListRetirableGrantsInput, arg2 ...func(*kms.Options)) (*kms.ListRetirableGrantsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRetirableGrants", varargs...)
	ret0, _ := ret[0].(*kms.ListRetirableGrantsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRetirableGrants indicates an expected call of ListRetirableGrants.
func (mr *MockKmsClientMockRecorder) ListRetirableGrants(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRetirableGrants", reflect.TypeOf((*MockKmsClient)(nil).ListRetirableGrants), varargs...)
}
