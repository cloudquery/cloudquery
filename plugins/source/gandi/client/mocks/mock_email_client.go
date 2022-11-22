// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/gandi/client (interfaces: EmailClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	email "github.com/go-gandi/go-gandi/email"
	gomock "github.com/golang/mock/gomock"
)

// MockEmailClient is a mock of EmailClient interface.
type MockEmailClient struct {
	ctrl     *gomock.Controller
	recorder *MockEmailClientMockRecorder
}

// MockEmailClientMockRecorder is the mock recorder for MockEmailClient.
type MockEmailClientMockRecorder struct {
	mock *MockEmailClient
}

// NewMockEmailClient creates a new mock instance.
func NewMockEmailClient(ctrl *gomock.Controller) *MockEmailClient {
	mock := &MockEmailClient{ctrl: ctrl}
	mock.recorder = &MockEmailClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailClient) EXPECT() *MockEmailClientMockRecorder {
	return m.recorder
}

// GetForwards mocks base method.
func (m *MockEmailClient) GetForwards(arg0 string) ([]email.GetForwardRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetForwards", arg0)
	ret0, _ := ret[0].([]email.GetForwardRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetForwards indicates an expected call of GetForwards.
func (mr *MockEmailClientMockRecorder) GetForwards(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetForwards", reflect.TypeOf((*MockEmailClient)(nil).GetForwards), arg0)
}

// GetMailbox mocks base method.
func (m *MockEmailClient) GetMailbox(arg0, arg1 string) (email.MailboxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMailbox", arg0, arg1)
	ret0, _ := ret[0].(email.MailboxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMailbox indicates an expected call of GetMailbox.
func (mr *MockEmailClientMockRecorder) GetMailbox(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMailbox", reflect.TypeOf((*MockEmailClient)(nil).GetMailbox), arg0, arg1)
}

// ListMailboxes mocks base method.
func (m *MockEmailClient) ListMailboxes(arg0 string) ([]email.ListMailboxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMailboxes", arg0)
	ret0, _ := ret[0].([]email.ListMailboxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMailboxes indicates an expected call of ListMailboxes.
func (mr *MockEmailClientMockRecorder) ListMailboxes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMailboxes", reflect.TypeOf((*MockEmailClient)(nil).ListMailboxes), arg0)
}
