// Code generated by MockGen. DO NOT EDIT.
// Source: rum_api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	reflect "reflect"

	datadog "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	datadogV2 "github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	gomock "github.com/golang/mock/gomock"
)

// MockRumAPIClient is a mock of RumAPIClient interface.
type MockRumAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockRumAPIClientMockRecorder
}

// MockRumAPIClientMockRecorder is the mock recorder for MockRumAPIClient.
type MockRumAPIClientMockRecorder struct {
	mock *MockRumAPIClient
}

// NewMockRumAPIClient creates a new mock instance.
func NewMockRumAPIClient(ctrl *gomock.Controller) *MockRumAPIClient {
	mock := &MockRumAPIClient{ctrl: ctrl}
	mock.recorder = &MockRumAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRumAPIClient) EXPECT() *MockRumAPIClientMockRecorder {
	return m.recorder
}

// ListRUMEvents mocks base method.
func (m *MockRumAPIClient) ListRUMEvents(arg0 context.Context, arg1 ...datadogV2.ListRUMEventsOptionalParameters) (datadogV2.RUMEventsResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRUMEvents", varargs...)
	ret0, _ := ret[0].(datadogV2.RUMEventsResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListRUMEvents indicates an expected call of ListRUMEvents.
func (mr *MockRumAPIClientMockRecorder) ListRUMEvents(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRUMEvents", reflect.TypeOf((*MockRumAPIClient)(nil).ListRUMEvents), varargs...)
}

// ListRUMEventsWithPagination mocks base method.
func (m *MockRumAPIClient) ListRUMEventsWithPagination(arg0 context.Context, arg1 ...datadogV2.ListRUMEventsOptionalParameters) (<-chan datadog.PaginationResult[datadogV2.RUMEvent], func()) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRUMEventsWithPagination", varargs...)
	ret0, _ := ret[0].(<-chan datadog.PaginationResult[datadogV2.RUMEvent])
	ret1, _ := ret[1].(func())
	return ret0, ret1
}

// ListRUMEventsWithPagination indicates an expected call of ListRUMEventsWithPagination.
func (mr *MockRumAPIClientMockRecorder) ListRUMEventsWithPagination(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRUMEventsWithPagination", reflect.TypeOf((*MockRumAPIClient)(nil).ListRUMEventsWithPagination), varargs...)
}
