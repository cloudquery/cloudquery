// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/aws/client (interfaces: IamClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	iam "github.com/aws/aws-sdk-go-v2/service/iam"
	gomock "github.com/golang/mock/gomock"
)

// MockIamClient is a mock of IamClient interface.
type MockIamClient struct {
	ctrl     *gomock.Controller
	recorder *MockIamClientMockRecorder
}

// MockIamClientMockRecorder is the mock recorder for MockIamClient.
type MockIamClientMockRecorder struct {
	mock *MockIamClient
}

// NewMockIamClient creates a new mock instance.
func NewMockIamClient(ctrl *gomock.Controller) *MockIamClient {
	mock := &MockIamClient{ctrl: ctrl}
	mock.recorder = &MockIamClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIamClient) EXPECT() *MockIamClientMockRecorder {
	return m.recorder
}

// GenerateCredentialReport mocks base method.
func (m *MockIamClient) GenerateCredentialReport(arg0 context.Context, arg1 *iam.GenerateCredentialReportInput, arg2 ...func(*iam.Options)) (*iam.GenerateCredentialReportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GenerateCredentialReport", varargs...)
	ret0, _ := ret[0].(*iam.GenerateCredentialReportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateCredentialReport indicates an expected call of GenerateCredentialReport.
func (mr *MockIamClientMockRecorder) GenerateCredentialReport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateCredentialReport", reflect.TypeOf((*MockIamClient)(nil).GenerateCredentialReport), varargs...)
}

// GetAccessKeyLastUsed mocks base method.
func (m *MockIamClient) GetAccessKeyLastUsed(arg0 context.Context, arg1 *iam.GetAccessKeyLastUsedInput, arg2 ...func(*iam.Options)) (*iam.GetAccessKeyLastUsedOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccessKeyLastUsed", varargs...)
	ret0, _ := ret[0].(*iam.GetAccessKeyLastUsedOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessKeyLastUsed indicates an expected call of GetAccessKeyLastUsed.
func (mr *MockIamClientMockRecorder) GetAccessKeyLastUsed(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessKeyLastUsed", reflect.TypeOf((*MockIamClient)(nil).GetAccessKeyLastUsed), varargs...)
}

// GetAccountAuthorizationDetails mocks base method.
func (m *MockIamClient) GetAccountAuthorizationDetails(arg0 context.Context, arg1 *iam.GetAccountAuthorizationDetailsInput, arg2 ...func(*iam.Options)) (*iam.GetAccountAuthorizationDetailsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccountAuthorizationDetails", varargs...)
	ret0, _ := ret[0].(*iam.GetAccountAuthorizationDetailsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountAuthorizationDetails indicates an expected call of GetAccountAuthorizationDetails.
func (mr *MockIamClientMockRecorder) GetAccountAuthorizationDetails(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountAuthorizationDetails", reflect.TypeOf((*MockIamClient)(nil).GetAccountAuthorizationDetails), varargs...)
}

// GetAccountPasswordPolicy mocks base method.
func (m *MockIamClient) GetAccountPasswordPolicy(arg0 context.Context, arg1 *iam.GetAccountPasswordPolicyInput, arg2 ...func(*iam.Options)) (*iam.GetAccountPasswordPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccountPasswordPolicy", varargs...)
	ret0, _ := ret[0].(*iam.GetAccountPasswordPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountPasswordPolicy indicates an expected call of GetAccountPasswordPolicy.
func (mr *MockIamClientMockRecorder) GetAccountPasswordPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountPasswordPolicy", reflect.TypeOf((*MockIamClient)(nil).GetAccountPasswordPolicy), varargs...)
}

// GetAccountSummary mocks base method.
func (m *MockIamClient) GetAccountSummary(arg0 context.Context, arg1 *iam.GetAccountSummaryInput, arg2 ...func(*iam.Options)) (*iam.GetAccountSummaryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccountSummary", varargs...)
	ret0, _ := ret[0].(*iam.GetAccountSummaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountSummary indicates an expected call of GetAccountSummary.
func (mr *MockIamClientMockRecorder) GetAccountSummary(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountSummary", reflect.TypeOf((*MockIamClient)(nil).GetAccountSummary), varargs...)
}

// GetCredentialReport mocks base method.
func (m *MockIamClient) GetCredentialReport(arg0 context.Context, arg1 *iam.GetCredentialReportInput, arg2 ...func(*iam.Options)) (*iam.GetCredentialReportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCredentialReport", varargs...)
	ret0, _ := ret[0].(*iam.GetCredentialReportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentialReport indicates an expected call of GetCredentialReport.
func (mr *MockIamClientMockRecorder) GetCredentialReport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentialReport", reflect.TypeOf((*MockIamClient)(nil).GetCredentialReport), varargs...)
}

// GetGroupPolicy mocks base method.
func (m *MockIamClient) GetGroupPolicy(arg0 context.Context, arg1 *iam.GetGroupPolicyInput, arg2 ...func(*iam.Options)) (*iam.GetGroupPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroupPolicy", varargs...)
	ret0, _ := ret[0].(*iam.GetGroupPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupPolicy indicates an expected call of GetGroupPolicy.
func (mr *MockIamClientMockRecorder) GetGroupPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupPolicy", reflect.TypeOf((*MockIamClient)(nil).GetGroupPolicy), varargs...)
}

// GetOpenIDConnectProvider mocks base method.
func (m *MockIamClient) GetOpenIDConnectProvider(arg0 context.Context, arg1 *iam.GetOpenIDConnectProviderInput, arg2 ...func(*iam.Options)) (*iam.GetOpenIDConnectProviderOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOpenIDConnectProvider", varargs...)
	ret0, _ := ret[0].(*iam.GetOpenIDConnectProviderOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOpenIDConnectProvider indicates an expected call of GetOpenIDConnectProvider.
func (mr *MockIamClientMockRecorder) GetOpenIDConnectProvider(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpenIDConnectProvider", reflect.TypeOf((*MockIamClient)(nil).GetOpenIDConnectProvider), varargs...)
}

// GetRolePolicy mocks base method.
func (m *MockIamClient) GetRolePolicy(arg0 context.Context, arg1 *iam.GetRolePolicyInput, arg2 ...func(*iam.Options)) (*iam.GetRolePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRolePolicy", varargs...)
	ret0, _ := ret[0].(*iam.GetRolePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRolePolicy indicates an expected call of GetRolePolicy.
func (mr *MockIamClientMockRecorder) GetRolePolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRolePolicy", reflect.TypeOf((*MockIamClient)(nil).GetRolePolicy), varargs...)
}

// GetSAMLProvider mocks base method.
func (m *MockIamClient) GetSAMLProvider(arg0 context.Context, arg1 *iam.GetSAMLProviderInput, arg2 ...func(*iam.Options)) (*iam.GetSAMLProviderOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSAMLProvider", varargs...)
	ret0, _ := ret[0].(*iam.GetSAMLProviderOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSAMLProvider indicates an expected call of GetSAMLProvider.
func (mr *MockIamClientMockRecorder) GetSAMLProvider(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSAMLProvider", reflect.TypeOf((*MockIamClient)(nil).GetSAMLProvider), varargs...)
}

// GetUser mocks base method.
func (m *MockIamClient) GetUser(arg0 context.Context, arg1 *iam.GetUserInput, arg2 ...func(*iam.Options)) (*iam.GetUserOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUser", varargs...)
	ret0, _ := ret[0].(*iam.GetUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockIamClientMockRecorder) GetUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIamClient)(nil).GetUser), varargs...)
}

// GetUserPolicy mocks base method.
func (m *MockIamClient) GetUserPolicy(arg0 context.Context, arg1 *iam.GetUserPolicyInput, arg2 ...func(*iam.Options)) (*iam.GetUserPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserPolicy", varargs...)
	ret0, _ := ret[0].(*iam.GetUserPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPolicy indicates an expected call of GetUserPolicy.
func (mr *MockIamClientMockRecorder) GetUserPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPolicy", reflect.TypeOf((*MockIamClient)(nil).GetUserPolicy), varargs...)
}

// ListAccessKeys mocks base method.
func (m *MockIamClient) ListAccessKeys(arg0 context.Context, arg1 *iam.ListAccessKeysInput, arg2 ...func(*iam.Options)) (*iam.ListAccessKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccessKeys", varargs...)
	ret0, _ := ret[0].(*iam.ListAccessKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccessKeys indicates an expected call of ListAccessKeys.
func (mr *MockIamClientMockRecorder) ListAccessKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessKeys", reflect.TypeOf((*MockIamClient)(nil).ListAccessKeys), varargs...)
}

// ListAccountAliases mocks base method.
func (m *MockIamClient) ListAccountAliases(arg0 context.Context, arg1 *iam.ListAccountAliasesInput, arg2 ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccountAliases", varargs...)
	ret0, _ := ret[0].(*iam.ListAccountAliasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccountAliases indicates an expected call of ListAccountAliases.
func (mr *MockIamClientMockRecorder) ListAccountAliases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccountAliases", reflect.TypeOf((*MockIamClient)(nil).ListAccountAliases), varargs...)
}

// ListAttachedGroupPolicies mocks base method.
func (m *MockIamClient) ListAttachedGroupPolicies(arg0 context.Context, arg1 *iam.ListAttachedGroupPoliciesInput, arg2 ...func(*iam.Options)) (*iam.ListAttachedGroupPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAttachedGroupPolicies", varargs...)
	ret0, _ := ret[0].(*iam.ListAttachedGroupPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAttachedGroupPolicies indicates an expected call of ListAttachedGroupPolicies.
func (mr *MockIamClientMockRecorder) ListAttachedGroupPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttachedGroupPolicies", reflect.TypeOf((*MockIamClient)(nil).ListAttachedGroupPolicies), varargs...)
}

// ListAttachedRolePolicies mocks base method.
func (m *MockIamClient) ListAttachedRolePolicies(arg0 context.Context, arg1 *iam.ListAttachedRolePoliciesInput, arg2 ...func(*iam.Options)) (*iam.ListAttachedRolePoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAttachedRolePolicies", varargs...)
	ret0, _ := ret[0].(*iam.ListAttachedRolePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAttachedRolePolicies indicates an expected call of ListAttachedRolePolicies.
func (mr *MockIamClientMockRecorder) ListAttachedRolePolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttachedRolePolicies", reflect.TypeOf((*MockIamClient)(nil).ListAttachedRolePolicies), varargs...)
}

// ListAttachedUserPolicies mocks base method.
func (m *MockIamClient) ListAttachedUserPolicies(arg0 context.Context, arg1 *iam.ListAttachedUserPoliciesInput, arg2 ...func(*iam.Options)) (*iam.ListAttachedUserPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAttachedUserPolicies", varargs...)
	ret0, _ := ret[0].(*iam.ListAttachedUserPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAttachedUserPolicies indicates an expected call of ListAttachedUserPolicies.
func (mr *MockIamClientMockRecorder) ListAttachedUserPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttachedUserPolicies", reflect.TypeOf((*MockIamClient)(nil).ListAttachedUserPolicies), varargs...)
}

// ListGroupPolicies mocks base method.
func (m *MockIamClient) ListGroupPolicies(arg0 context.Context, arg1 *iam.ListGroupPoliciesInput, arg2 ...func(*iam.Options)) (*iam.ListGroupPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroupPolicies", varargs...)
	ret0, _ := ret[0].(*iam.ListGroupPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroupPolicies indicates an expected call of ListGroupPolicies.
func (mr *MockIamClientMockRecorder) ListGroupPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupPolicies", reflect.TypeOf((*MockIamClient)(nil).ListGroupPolicies), varargs...)
}

// ListGroups mocks base method.
func (m *MockIamClient) ListGroups(arg0 context.Context, arg1 *iam.ListGroupsInput, arg2 ...func(*iam.Options)) (*iam.ListGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroups", varargs...)
	ret0, _ := ret[0].(*iam.ListGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroups indicates an expected call of ListGroups.
func (mr *MockIamClientMockRecorder) ListGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroups", reflect.TypeOf((*MockIamClient)(nil).ListGroups), varargs...)
}

// ListGroupsForUser mocks base method.
func (m *MockIamClient) ListGroupsForUser(arg0 context.Context, arg1 *iam.ListGroupsForUserInput, arg2 ...func(*iam.Options)) (*iam.ListGroupsForUserOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroupsForUser", varargs...)
	ret0, _ := ret[0].(*iam.ListGroupsForUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroupsForUser indicates an expected call of ListGroupsForUser.
func (mr *MockIamClientMockRecorder) ListGroupsForUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupsForUser", reflect.TypeOf((*MockIamClient)(nil).ListGroupsForUser), varargs...)
}

// ListOpenIDConnectProviders mocks base method.
func (m *MockIamClient) ListOpenIDConnectProviders(arg0 context.Context, arg1 *iam.ListOpenIDConnectProvidersInput, arg2 ...func(*iam.Options)) (*iam.ListOpenIDConnectProvidersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOpenIDConnectProviders", varargs...)
	ret0, _ := ret[0].(*iam.ListOpenIDConnectProvidersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOpenIDConnectProviders indicates an expected call of ListOpenIDConnectProviders.
func (mr *MockIamClientMockRecorder) ListOpenIDConnectProviders(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOpenIDConnectProviders", reflect.TypeOf((*MockIamClient)(nil).ListOpenIDConnectProviders), varargs...)
}

// ListPolicyTags mocks base method.
func (m *MockIamClient) ListPolicyTags(arg0 context.Context, arg1 *iam.ListPolicyTagsInput, arg2 ...func(*iam.Options)) (*iam.ListPolicyTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPolicyTags", varargs...)
	ret0, _ := ret[0].(*iam.ListPolicyTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPolicyTags indicates an expected call of ListPolicyTags.
func (mr *MockIamClientMockRecorder) ListPolicyTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPolicyTags", reflect.TypeOf((*MockIamClient)(nil).ListPolicyTags), varargs...)
}

// ListRolePolicies mocks base method.
func (m *MockIamClient) ListRolePolicies(arg0 context.Context, arg1 *iam.ListRolePoliciesInput, arg2 ...func(*iam.Options)) (*iam.ListRolePoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRolePolicies", varargs...)
	ret0, _ := ret[0].(*iam.ListRolePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRolePolicies indicates an expected call of ListRolePolicies.
func (mr *MockIamClientMockRecorder) ListRolePolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRolePolicies", reflect.TypeOf((*MockIamClient)(nil).ListRolePolicies), varargs...)
}

// ListRoleTags mocks base method.
func (m *MockIamClient) ListRoleTags(arg0 context.Context, arg1 *iam.ListRoleTagsInput, arg2 ...func(*iam.Options)) (*iam.ListRoleTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRoleTags", varargs...)
	ret0, _ := ret[0].(*iam.ListRoleTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoleTags indicates an expected call of ListRoleTags.
func (mr *MockIamClientMockRecorder) ListRoleTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoleTags", reflect.TypeOf((*MockIamClient)(nil).ListRoleTags), varargs...)
}

// ListRoles mocks base method.
func (m *MockIamClient) ListRoles(arg0 context.Context, arg1 *iam.ListRolesInput, arg2 ...func(*iam.Options)) (*iam.ListRolesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRoles", varargs...)
	ret0, _ := ret[0].(*iam.ListRolesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoles indicates an expected call of ListRoles.
func (mr *MockIamClientMockRecorder) ListRoles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoles", reflect.TypeOf((*MockIamClient)(nil).ListRoles), varargs...)
}

// ListSAMLProviders mocks base method.
func (m *MockIamClient) ListSAMLProviders(arg0 context.Context, arg1 *iam.ListSAMLProvidersInput, arg2 ...func(*iam.Options)) (*iam.ListSAMLProvidersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSAMLProviders", varargs...)
	ret0, _ := ret[0].(*iam.ListSAMLProvidersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSAMLProviders indicates an expected call of ListSAMLProviders.
func (mr *MockIamClientMockRecorder) ListSAMLProviders(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSAMLProviders", reflect.TypeOf((*MockIamClient)(nil).ListSAMLProviders), varargs...)
}

// ListServerCertificates mocks base method.
func (m *MockIamClient) ListServerCertificates(arg0 context.Context, arg1 *iam.ListServerCertificatesInput, arg2 ...func(*iam.Options)) (*iam.ListServerCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListServerCertificates", varargs...)
	ret0, _ := ret[0].(*iam.ListServerCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServerCertificates indicates an expected call of ListServerCertificates.
func (mr *MockIamClientMockRecorder) ListServerCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServerCertificates", reflect.TypeOf((*MockIamClient)(nil).ListServerCertificates), varargs...)
}

// ListUserPolicies mocks base method.
func (m *MockIamClient) ListUserPolicies(arg0 context.Context, arg1 *iam.ListUserPoliciesInput, arg2 ...func(*iam.Options)) (*iam.ListUserPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUserPolicies", varargs...)
	ret0, _ := ret[0].(*iam.ListUserPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserPolicies indicates an expected call of ListUserPolicies.
func (mr *MockIamClientMockRecorder) ListUserPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserPolicies", reflect.TypeOf((*MockIamClient)(nil).ListUserPolicies), varargs...)
}

// ListUserTags mocks base method.
func (m *MockIamClient) ListUserTags(arg0 context.Context, arg1 *iam.ListUserTagsInput, arg2 ...func(*iam.Options)) (*iam.ListUserTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUserTags", varargs...)
	ret0, _ := ret[0].(*iam.ListUserTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserTags indicates an expected call of ListUserTags.
func (mr *MockIamClientMockRecorder) ListUserTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserTags", reflect.TypeOf((*MockIamClient)(nil).ListUserTags), varargs...)
}

// ListUsers mocks base method.
func (m *MockIamClient) ListUsers(arg0 context.Context, arg1 *iam.ListUsersInput, arg2 ...func(*iam.Options)) (*iam.ListUsersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUsers", varargs...)
	ret0, _ := ret[0].(*iam.ListUsersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockIamClientMockRecorder) ListUsers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockIamClient)(nil).ListUsers), varargs...)
}

// ListVirtualMFADevices mocks base method.
func (m *MockIamClient) ListVirtualMFADevices(arg0 context.Context, arg1 *iam.ListVirtualMFADevicesInput, arg2 ...func(*iam.Options)) (*iam.ListVirtualMFADevicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListVirtualMFADevices", varargs...)
	ret0, _ := ret[0].(*iam.ListVirtualMFADevicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVirtualMFADevices indicates an expected call of ListVirtualMFADevices.
func (mr *MockIamClientMockRecorder) ListVirtualMFADevices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVirtualMFADevices", reflect.TypeOf((*MockIamClient)(nil).ListVirtualMFADevices), varargs...)
}
