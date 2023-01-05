// Code generated by MockGen. DO NOT EDIT.
// Source: incidents_api.go

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

// MockIncidentsAPIClient is a mock of IncidentsAPIClient interface.
type MockIncidentsAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockIncidentsAPIClientMockRecorder
}

// MockIncidentsAPIClientMockRecorder is the mock recorder for MockIncidentsAPIClient.
type MockIncidentsAPIClientMockRecorder struct {
	mock *MockIncidentsAPIClient
}

// NewMockIncidentsAPIClient creates a new mock instance.
func NewMockIncidentsAPIClient(ctrl *gomock.Controller) *MockIncidentsAPIClient {
	mock := &MockIncidentsAPIClient{ctrl: ctrl}
	mock.recorder = &MockIncidentsAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIncidentsAPIClient) EXPECT() *MockIncidentsAPIClientMockRecorder {
	return m.recorder
}

// ListIncidentAttachments mocks base method.
func (m *MockIncidentsAPIClient) ListIncidentAttachments(arg0 context.Context, arg1 string, arg2 ...datadogV2.ListIncidentAttachmentsOptionalParameters) (datadogV2.IncidentAttachmentsResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIncidentAttachments", varargs...)
	ret0, _ := ret[0].(datadogV2.IncidentAttachmentsResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListIncidentAttachments indicates an expected call of ListIncidentAttachments.
func (mr *MockIncidentsAPIClientMockRecorder) ListIncidentAttachments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIncidentAttachments", reflect.TypeOf((*MockIncidentsAPIClient)(nil).ListIncidentAttachments), varargs...)
}

// ListIncidents mocks base method.
func (m *MockIncidentsAPIClient) ListIncidents(arg0 context.Context, arg1 ...datadogV2.ListIncidentsOptionalParameters) (datadogV2.IncidentsResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIncidents", varargs...)
	ret0, _ := ret[0].(datadogV2.IncidentsResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListIncidents indicates an expected call of ListIncidents.
func (mr *MockIncidentsAPIClientMockRecorder) ListIncidents(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIncidents", reflect.TypeOf((*MockIncidentsAPIClient)(nil).ListIncidents), varargs...)
}

// ListIncidentsWithPagination mocks base method.
func (m *MockIncidentsAPIClient) ListIncidentsWithPagination(arg0 context.Context, arg1 ...datadogV2.ListIncidentsOptionalParameters) (<-chan datadog.PaginationResult[datadogV2.IncidentResponseData], func()) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIncidentsWithPagination", varargs...)
	ret0, _ := ret[0].(<-chan datadog.PaginationResult[datadogV2.IncidentResponseData])
	ret1, _ := ret[1].(func())
	return ret0, ret1
}

// ListIncidentsWithPagination indicates an expected call of ListIncidentsWithPagination.
func (mr *MockIncidentsAPIClientMockRecorder) ListIncidentsWithPagination(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIncidentsWithPagination", reflect.TypeOf((*MockIncidentsAPIClient)(nil).ListIncidentsWithPagination), varargs...)
}
