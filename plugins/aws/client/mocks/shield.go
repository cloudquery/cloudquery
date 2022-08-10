// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/aws/client (interfaces: ShieldClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	shield "github.com/aws/aws-sdk-go-v2/service/shield"
	gomock "github.com/golang/mock/gomock"
)

// MockShieldClient is a mock of ShieldClient interface.
type MockShieldClient struct {
	ctrl     *gomock.Controller
	recorder *MockShieldClientMockRecorder
}

// MockShieldClientMockRecorder is the mock recorder for MockShieldClient.
type MockShieldClientMockRecorder struct {
	mock *MockShieldClient
}

// NewMockShieldClient creates a new mock instance.
func NewMockShieldClient(ctrl *gomock.Controller) *MockShieldClient {
	mock := &MockShieldClient{ctrl: ctrl}
	mock.recorder = &MockShieldClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShieldClient) EXPECT() *MockShieldClientMockRecorder {
	return m.recorder
}

// DescribeAttack mocks base method.
func (m *MockShieldClient) DescribeAttack(arg0 context.Context, arg1 *shield.DescribeAttackInput, arg2 ...func(*shield.Options)) (*shield.DescribeAttackOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAttack", varargs...)
	ret0, _ := ret[0].(*shield.DescribeAttackOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAttack indicates an expected call of DescribeAttack.
func (mr *MockShieldClientMockRecorder) DescribeAttack(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAttack", reflect.TypeOf((*MockShieldClient)(nil).DescribeAttack), varargs...)
}

// DescribeSubscription mocks base method.
func (m *MockShieldClient) DescribeSubscription(arg0 context.Context, arg1 *shield.DescribeSubscriptionInput, arg2 ...func(*shield.Options)) (*shield.DescribeSubscriptionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSubscription", varargs...)
	ret0, _ := ret[0].(*shield.DescribeSubscriptionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSubscription indicates an expected call of DescribeSubscription.
func (mr *MockShieldClientMockRecorder) DescribeSubscription(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSubscription", reflect.TypeOf((*MockShieldClient)(nil).DescribeSubscription), varargs...)
}

// ListAttacks mocks base method.
func (m *MockShieldClient) ListAttacks(arg0 context.Context, arg1 *shield.ListAttacksInput, arg2 ...func(*shield.Options)) (*shield.ListAttacksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAttacks", varargs...)
	ret0, _ := ret[0].(*shield.ListAttacksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAttacks indicates an expected call of ListAttacks.
func (mr *MockShieldClientMockRecorder) ListAttacks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttacks", reflect.TypeOf((*MockShieldClient)(nil).ListAttacks), varargs...)
}

// ListProtectionGroups mocks base method.
func (m *MockShieldClient) ListProtectionGroups(arg0 context.Context, arg1 *shield.ListProtectionGroupsInput, arg2 ...func(*shield.Options)) (*shield.ListProtectionGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProtectionGroups", varargs...)
	ret0, _ := ret[0].(*shield.ListProtectionGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProtectionGroups indicates an expected call of ListProtectionGroups.
func (mr *MockShieldClientMockRecorder) ListProtectionGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProtectionGroups", reflect.TypeOf((*MockShieldClient)(nil).ListProtectionGroups), varargs...)
}

// ListProtections mocks base method.
func (m *MockShieldClient) ListProtections(arg0 context.Context, arg1 *shield.ListProtectionsInput, arg2 ...func(*shield.Options)) (*shield.ListProtectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProtections", varargs...)
	ret0, _ := ret[0].(*shield.ListProtectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProtections indicates an expected call of ListProtections.
func (mr *MockShieldClientMockRecorder) ListProtections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProtections", reflect.TypeOf((*MockShieldClient)(nil).ListProtections), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockShieldClient) ListTagsForResource(arg0 context.Context, arg1 *shield.ListTagsForResourceInput, arg2 ...func(*shield.Options)) (*shield.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*shield.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockShieldClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockShieldClient)(nil).ListTagsForResource), varargs...)
}
