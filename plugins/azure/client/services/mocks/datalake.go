// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/azure/client/services (interfaces: DataLakeStorageAccountsClient,DataLakeAnalyticsAccountsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	account "github.com/Azure/azure-sdk-for-go/services/datalake/analytics/mgmt/2016-11-01/account"
	account0 "github.com/Azure/azure-sdk-for-go/services/datalake/store/mgmt/2016-11-01/account"
	gomock "github.com/golang/mock/gomock"
)

// MockDataLakeStorageAccountsClient is a mock of DataLakeStorageAccountsClient interface.
type MockDataLakeStorageAccountsClient struct {
	ctrl     *gomock.Controller
	recorder *MockDataLakeStorageAccountsClientMockRecorder
}

// MockDataLakeStorageAccountsClientMockRecorder is the mock recorder for MockDataLakeStorageAccountsClient.
type MockDataLakeStorageAccountsClientMockRecorder struct {
	mock *MockDataLakeStorageAccountsClient
}

// NewMockDataLakeStorageAccountsClient creates a new mock instance.
func NewMockDataLakeStorageAccountsClient(ctrl *gomock.Controller) *MockDataLakeStorageAccountsClient {
	mock := &MockDataLakeStorageAccountsClient{ctrl: ctrl}
	mock.recorder = &MockDataLakeStorageAccountsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataLakeStorageAccountsClient) EXPECT() *MockDataLakeStorageAccountsClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockDataLakeStorageAccountsClient) Get(arg0 context.Context, arg1, arg2 string) (account0.DataLakeStoreAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(account0.DataLakeStoreAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDataLakeStorageAccountsClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDataLakeStorageAccountsClient)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockDataLakeStorageAccountsClient) List(arg0 context.Context, arg1 string, arg2, arg3 *int32, arg4, arg5 string, arg6 *bool) (account0.DataLakeStoreAccountListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(account0.DataLakeStoreAccountListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockDataLakeStorageAccountsClientMockRecorder) List(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDataLakeStorageAccountsClient)(nil).List), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// MockDataLakeAnalyticsAccountsClient is a mock of DataLakeAnalyticsAccountsClient interface.
type MockDataLakeAnalyticsAccountsClient struct {
	ctrl     *gomock.Controller
	recorder *MockDataLakeAnalyticsAccountsClientMockRecorder
}

// MockDataLakeAnalyticsAccountsClientMockRecorder is the mock recorder for MockDataLakeAnalyticsAccountsClient.
type MockDataLakeAnalyticsAccountsClientMockRecorder struct {
	mock *MockDataLakeAnalyticsAccountsClient
}

// NewMockDataLakeAnalyticsAccountsClient creates a new mock instance.
func NewMockDataLakeAnalyticsAccountsClient(ctrl *gomock.Controller) *MockDataLakeAnalyticsAccountsClient {
	mock := &MockDataLakeAnalyticsAccountsClient{ctrl: ctrl}
	mock.recorder = &MockDataLakeAnalyticsAccountsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataLakeAnalyticsAccountsClient) EXPECT() *MockDataLakeAnalyticsAccountsClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockDataLakeAnalyticsAccountsClient) Get(arg0 context.Context, arg1, arg2 string) (account.DataLakeAnalyticsAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(account.DataLakeAnalyticsAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDataLakeAnalyticsAccountsClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDataLakeAnalyticsAccountsClient)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockDataLakeAnalyticsAccountsClient) List(arg0 context.Context, arg1 string, arg2, arg3 *int32, arg4, arg5 string, arg6 *bool) (account.DataLakeAnalyticsAccountListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(account.DataLakeAnalyticsAccountListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockDataLakeAnalyticsAccountsClientMockRecorder) List(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDataLakeAnalyticsAccountsClient)(nil).List), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}
