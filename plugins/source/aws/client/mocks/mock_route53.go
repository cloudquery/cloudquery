// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/aws/client (interfaces: Route53Client)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	route53 "github.com/aws/aws-sdk-go-v2/service/route53"
	gomock "github.com/golang/mock/gomock"
)

// MockRoute53Client is a mock of Route53Client interface.
type MockRoute53Client struct {
	ctrl     *gomock.Controller
	recorder *MockRoute53ClientMockRecorder
}

// MockRoute53ClientMockRecorder is the mock recorder for MockRoute53Client.
type MockRoute53ClientMockRecorder struct {
	mock *MockRoute53Client
}

// NewMockRoute53Client creates a new mock instance.
func NewMockRoute53Client(ctrl *gomock.Controller) *MockRoute53Client {
	mock := &MockRoute53Client{ctrl: ctrl}
	mock.recorder = &MockRoute53ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoute53Client) EXPECT() *MockRoute53ClientMockRecorder {
	return m.recorder
}

// GetHostedZone mocks base method.
func (m *MockRoute53Client) GetHostedZone(arg0 context.Context, arg1 *route53.GetHostedZoneInput, arg2 ...func(*route53.Options)) (*route53.GetHostedZoneOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetHostedZone", varargs...)
	ret0, _ := ret[0].(*route53.GetHostedZoneOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostedZone indicates an expected call of GetHostedZone.
func (mr *MockRoute53ClientMockRecorder) GetHostedZone(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostedZone", reflect.TypeOf((*MockRoute53Client)(nil).GetHostedZone), varargs...)
}

// GetTrafficPolicy mocks base method.
func (m *MockRoute53Client) GetTrafficPolicy(arg0 context.Context, arg1 *route53.GetTrafficPolicyInput, arg2 ...func(*route53.Options)) (*route53.GetTrafficPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTrafficPolicy", varargs...)
	ret0, _ := ret[0].(*route53.GetTrafficPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrafficPolicy indicates an expected call of GetTrafficPolicy.
func (mr *MockRoute53ClientMockRecorder) GetTrafficPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrafficPolicy", reflect.TypeOf((*MockRoute53Client)(nil).GetTrafficPolicy), varargs...)
}

// ListHealthChecks mocks base method.
func (m *MockRoute53Client) ListHealthChecks(arg0 context.Context, arg1 *route53.ListHealthChecksInput, arg2 ...func(*route53.Options)) (*route53.ListHealthChecksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListHealthChecks", varargs...)
	ret0, _ := ret[0].(*route53.ListHealthChecksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHealthChecks indicates an expected call of ListHealthChecks.
func (mr *MockRoute53ClientMockRecorder) ListHealthChecks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHealthChecks", reflect.TypeOf((*MockRoute53Client)(nil).ListHealthChecks), varargs...)
}

// ListHostedZones mocks base method.
func (m *MockRoute53Client) ListHostedZones(arg0 context.Context, arg1 *route53.ListHostedZonesInput, arg2 ...func(*route53.Options)) (*route53.ListHostedZonesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListHostedZones", varargs...)
	ret0, _ := ret[0].(*route53.ListHostedZonesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHostedZones indicates an expected call of ListHostedZones.
func (mr *MockRoute53ClientMockRecorder) ListHostedZones(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHostedZones", reflect.TypeOf((*MockRoute53Client)(nil).ListHostedZones), varargs...)
}

// ListQueryLoggingConfigs mocks base method.
func (m *MockRoute53Client) ListQueryLoggingConfigs(arg0 context.Context, arg1 *route53.ListQueryLoggingConfigsInput, arg2 ...func(*route53.Options)) (*route53.ListQueryLoggingConfigsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListQueryLoggingConfigs", varargs...)
	ret0, _ := ret[0].(*route53.ListQueryLoggingConfigsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListQueryLoggingConfigs indicates an expected call of ListQueryLoggingConfigs.
func (mr *MockRoute53ClientMockRecorder) ListQueryLoggingConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListQueryLoggingConfigs", reflect.TypeOf((*MockRoute53Client)(nil).ListQueryLoggingConfigs), varargs...)
}

// ListResourceRecordSets mocks base method.
func (m *MockRoute53Client) ListResourceRecordSets(arg0 context.Context, arg1 *route53.ListResourceRecordSetsInput, arg2 ...func(*route53.Options)) (*route53.ListResourceRecordSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListResourceRecordSets", varargs...)
	ret0, _ := ret[0].(*route53.ListResourceRecordSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResourceRecordSets indicates an expected call of ListResourceRecordSets.
func (mr *MockRoute53ClientMockRecorder) ListResourceRecordSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceRecordSets", reflect.TypeOf((*MockRoute53Client)(nil).ListResourceRecordSets), varargs...)
}

// ListReusableDelegationSets mocks base method.
func (m *MockRoute53Client) ListReusableDelegationSets(arg0 context.Context, arg1 *route53.ListReusableDelegationSetsInput, arg2 ...func(*route53.Options)) (*route53.ListReusableDelegationSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListReusableDelegationSets", varargs...)
	ret0, _ := ret[0].(*route53.ListReusableDelegationSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReusableDelegationSets indicates an expected call of ListReusableDelegationSets.
func (mr *MockRoute53ClientMockRecorder) ListReusableDelegationSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReusableDelegationSets", reflect.TypeOf((*MockRoute53Client)(nil).ListReusableDelegationSets), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockRoute53Client) ListTagsForResource(arg0 context.Context, arg1 *route53.ListTagsForResourceInput, arg2 ...func(*route53.Options)) (*route53.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*route53.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockRoute53ClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockRoute53Client)(nil).ListTagsForResource), varargs...)
}

// ListTagsForResources mocks base method.
func (m *MockRoute53Client) ListTagsForResources(arg0 context.Context, arg1 *route53.ListTagsForResourcesInput, arg2 ...func(*route53.Options)) (*route53.ListTagsForResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResources", varargs...)
	ret0, _ := ret[0].(*route53.ListTagsForResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResources indicates an expected call of ListTagsForResources.
func (mr *MockRoute53ClientMockRecorder) ListTagsForResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResources", reflect.TypeOf((*MockRoute53Client)(nil).ListTagsForResources), varargs...)
}

// ListTrafficPolicies mocks base method.
func (m *MockRoute53Client) ListTrafficPolicies(arg0 context.Context, arg1 *route53.ListTrafficPoliciesInput, arg2 ...func(*route53.Options)) (*route53.ListTrafficPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTrafficPolicies", varargs...)
	ret0, _ := ret[0].(*route53.ListTrafficPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTrafficPolicies indicates an expected call of ListTrafficPolicies.
func (mr *MockRoute53ClientMockRecorder) ListTrafficPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTrafficPolicies", reflect.TypeOf((*MockRoute53Client)(nil).ListTrafficPolicies), varargs...)
}

// ListTrafficPolicyInstancesByHostedZone mocks base method.
func (m *MockRoute53Client) ListTrafficPolicyInstancesByHostedZone(arg0 context.Context, arg1 *route53.ListTrafficPolicyInstancesByHostedZoneInput, arg2 ...func(*route53.Options)) (*route53.ListTrafficPolicyInstancesByHostedZoneOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTrafficPolicyInstancesByHostedZone", varargs...)
	ret0, _ := ret[0].(*route53.ListTrafficPolicyInstancesByHostedZoneOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTrafficPolicyInstancesByHostedZone indicates an expected call of ListTrafficPolicyInstancesByHostedZone.
func (mr *MockRoute53ClientMockRecorder) ListTrafficPolicyInstancesByHostedZone(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTrafficPolicyInstancesByHostedZone", reflect.TypeOf((*MockRoute53Client)(nil).ListTrafficPolicyInstancesByHostedZone), varargs...)
}

// ListTrafficPolicyVersions mocks base method.
func (m *MockRoute53Client) ListTrafficPolicyVersions(arg0 context.Context, arg1 *route53.ListTrafficPolicyVersionsInput, arg2 ...func(*route53.Options)) (*route53.ListTrafficPolicyVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTrafficPolicyVersions", varargs...)
	ret0, _ := ret[0].(*route53.ListTrafficPolicyVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTrafficPolicyVersions indicates an expected call of ListTrafficPolicyVersions.
func (mr *MockRoute53ClientMockRecorder) ListTrafficPolicyVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTrafficPolicyVersions", reflect.TypeOf((*MockRoute53Client)(nil).ListTrafficPolicyVersions), varargs...)
}

// ListVPCAssociationAuthorizations mocks base method.
func (m *MockRoute53Client) ListVPCAssociationAuthorizations(arg0 context.Context, arg1 *route53.ListVPCAssociationAuthorizationsInput, arg2 ...func(*route53.Options)) (*route53.ListVPCAssociationAuthorizationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListVPCAssociationAuthorizations", varargs...)
	ret0, _ := ret[0].(*route53.ListVPCAssociationAuthorizationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVPCAssociationAuthorizations indicates an expected call of ListVPCAssociationAuthorizations.
func (mr *MockRoute53ClientMockRecorder) ListVPCAssociationAuthorizations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVPCAssociationAuthorizations", reflect.TypeOf((*MockRoute53Client)(nil).ListVPCAssociationAuthorizations), varargs...)
}
