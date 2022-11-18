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
