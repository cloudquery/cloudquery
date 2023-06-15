// Code generated by MockGen. DO NOT EDIT.
// Source: account.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	account "github.com/aws/aws-sdk-go-v2/service/account"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountClient is a mock of AccountClient interface.
type MockAccountClient struct {
	ctrl     *gomock.Controller
	recorder *MockAccountClientMockRecorder
}

// MockAccountClientMockRecorder is the mock recorder for MockAccountClient.
type MockAccountClientMockRecorder struct {
	mock *MockAccountClient
}

// NewMockAccountClient creates a new mock instance.
func NewMockAccountClient(ctrl *gomock.Controller) *MockAccountClient {
	mock := &MockAccountClient{ctrl: ctrl}
	mock.recorder = &MockAccountClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountClient) EXPECT() *MockAccountClientMockRecorder {
	return m.recorder
}

// GetAlternateContact mocks base method.
func (m *MockAccountClient) GetAlternateContact(arg0 context.Context, arg1 *account.GetAlternateContactInput, arg2 ...func(*account.Options)) (*account.GetAlternateContactOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &account.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetAlternateContact")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAlternateContact", varargs...)
	ret0, _ := ret[0].(*account.GetAlternateContactOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlternateContact indicates an expected call of GetAlternateContact.
func (mr *MockAccountClientMockRecorder) GetAlternateContact(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlternateContact", reflect.TypeOf((*MockAccountClient)(nil).GetAlternateContact), varargs...)
}

// GetContactInformation mocks base method.
func (m *MockAccountClient) GetContactInformation(arg0 context.Context, arg1 *account.GetContactInformationInput, arg2 ...func(*account.Options)) (*account.GetContactInformationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &account.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetContactInformation")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContactInformation", varargs...)
	ret0, _ := ret[0].(*account.GetContactInformationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContactInformation indicates an expected call of GetContactInformation.
func (mr *MockAccountClientMockRecorder) GetContactInformation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContactInformation", reflect.TypeOf((*MockAccountClient)(nil).GetContactInformation), varargs...)
}

// GetRegionOptStatus mocks base method.
func (m *MockAccountClient) GetRegionOptStatus(arg0 context.Context, arg1 *account.GetRegionOptStatusInput, arg2 ...func(*account.Options)) (*account.GetRegionOptStatusOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &account.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetRegionOptStatus")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRegionOptStatus", varargs...)
	ret0, _ := ret[0].(*account.GetRegionOptStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegionOptStatus indicates an expected call of GetRegionOptStatus.
func (mr *MockAccountClientMockRecorder) GetRegionOptStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegionOptStatus", reflect.TypeOf((*MockAccountClient)(nil).GetRegionOptStatus), varargs...)
}

// ListRegions mocks base method.
func (m *MockAccountClient) ListRegions(arg0 context.Context, arg1 *account.ListRegionsInput, arg2 ...func(*account.Options)) (*account.ListRegionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &account.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListRegions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRegions", varargs...)
	ret0, _ := ret[0].(*account.ListRegionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRegions indicates an expected call of ListRegions.
func (mr *MockAccountClientMockRecorder) ListRegions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRegions", reflect.TypeOf((*MockAccountClient)(nil).ListRegions), varargs...)
}
