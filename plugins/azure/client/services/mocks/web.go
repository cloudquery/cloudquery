// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/azure/client/services (interfaces: AppsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	web "github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
	gomock "github.com/golang/mock/gomock"
)

// MockAppsClient is a mock of AppsClient interface.
type MockAppsClient struct {
	ctrl     *gomock.Controller
	recorder *MockAppsClientMockRecorder
}

// MockAppsClientMockRecorder is the mock recorder for MockAppsClient.
type MockAppsClientMockRecorder struct {
	mock *MockAppsClient
}

// NewMockAppsClient creates a new mock instance.
func NewMockAppsClient(ctrl *gomock.Controller) *MockAppsClient {
	mock := &MockAppsClient{ctrl: ctrl}
	mock.recorder = &MockAppsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppsClient) EXPECT() *MockAppsClientMockRecorder {
	return m.recorder
}

// GetAuthSettings mocks base method.
func (m *MockAppsClient) GetAuthSettings(arg0 context.Context, arg1, arg2 string) (web.SiteAuthSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthSettings", arg0, arg1, arg2)
	ret0, _ := ret[0].(web.SiteAuthSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthSettings indicates an expected call of GetAuthSettings.
func (mr *MockAppsClientMockRecorder) GetAuthSettings(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthSettings", reflect.TypeOf((*MockAppsClient)(nil).GetAuthSettings), arg0, arg1, arg2)
}

// GetVnetConnection mocks base method.
func (m *MockAppsClient) GetVnetConnection(arg0 context.Context, arg1, arg2, arg3 string) (web.VnetInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVnetConnection", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(web.VnetInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVnetConnection indicates an expected call of GetVnetConnection.
func (mr *MockAppsClientMockRecorder) GetVnetConnection(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVnetConnection", reflect.TypeOf((*MockAppsClient)(nil).GetVnetConnection), arg0, arg1, arg2, arg3)
}

// List mocks base method.
func (m *MockAppsClient) List(arg0 context.Context) (web.AppCollectionPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(web.AppCollectionPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAppsClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAppsClient)(nil).List), arg0)
}

// ListPublishingProfileXMLWithSecrets mocks base method.
func (m *MockAppsClient) ListPublishingProfileXMLWithSecrets(arg0 context.Context, arg1, arg2 string, arg3 web.CsmPublishingProfileOptions) (web.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPublishingProfileXMLWithSecrets", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(web.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPublishingProfileXMLWithSecrets indicates an expected call of ListPublishingProfileXMLWithSecrets.
func (mr *MockAppsClientMockRecorder) ListPublishingProfileXMLWithSecrets(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPublishingProfileXMLWithSecrets", reflect.TypeOf((*MockAppsClient)(nil).ListPublishingProfileXMLWithSecrets), arg0, arg1, arg2, arg3)
}
