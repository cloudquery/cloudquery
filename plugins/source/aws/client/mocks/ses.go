// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/aws/client (interfaces: SESClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	sesv2 "github.com/aws/aws-sdk-go-v2/service/sesv2"
	gomock "github.com/golang/mock/gomock"
)

// MockSESClient is a mock of SESClient interface.
type MockSESClient struct {
	ctrl     *gomock.Controller
	recorder *MockSESClientMockRecorder
}

// MockSESClientMockRecorder is the mock recorder for MockSESClient.
type MockSESClientMockRecorder struct {
	mock *MockSESClient
}

// NewMockSESClient creates a new mock instance.
func NewMockSESClient(ctrl *gomock.Controller) *MockSESClient {
	mock := &MockSESClient{ctrl: ctrl}
	mock.recorder = &MockSESClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSESClient) EXPECT() *MockSESClientMockRecorder {
	return m.recorder
}

// GetEmailTemplate mocks base method.
func (m *MockSESClient) GetEmailTemplate(arg0 context.Context, arg1 *sesv2.GetEmailTemplateInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetEmailTemplateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEmailTemplate", varargs...)
	ret0, _ := ret[0].(*sesv2.GetEmailTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmailTemplate indicates an expected call of GetEmailTemplate.
func (mr *MockSESClientMockRecorder) GetEmailTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmailTemplate", reflect.TypeOf((*MockSESClient)(nil).GetEmailTemplate), varargs...)
}

// ListEmailTemplates mocks base method.
func (m *MockSESClient) ListEmailTemplates(arg0 context.Context, arg1 *sesv2.ListEmailTemplatesInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListEmailTemplatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEmailTemplates", varargs...)
	ret0, _ := ret[0].(*sesv2.ListEmailTemplatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEmailTemplates indicates an expected call of ListEmailTemplates.
func (mr *MockSESClientMockRecorder) ListEmailTemplates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEmailTemplates", reflect.TypeOf((*MockSESClient)(nil).ListEmailTemplates), varargs...)
}
