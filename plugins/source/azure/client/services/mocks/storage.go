// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/azure/client/services (interfaces: StorageAccountsClient,StorageBlobServicePropertiesClient,StorageBlobServicesClient,StorageContainerClient,StorageQueueServicePropertiesClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	storage "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
	gomock "github.com/golang/mock/gomock"
	accounts "github.com/tombuildsstuff/giovanni/storage/2020-08-04/blob/accounts"
	queues "github.com/tombuildsstuff/giovanni/storage/2020-08-04/queue/queues"
)

// MockStorageAccountsClient is a mock of StorageAccountsClient interface.
type MockStorageAccountsClient struct {
	ctrl     *gomock.Controller
	recorder *MockStorageAccountsClientMockRecorder
}

// MockStorageAccountsClientMockRecorder is the mock recorder for MockStorageAccountsClient.
type MockStorageAccountsClientMockRecorder struct {
	mock *MockStorageAccountsClient
}

// NewMockStorageAccountsClient creates a new mock instance.
func NewMockStorageAccountsClient(ctrl *gomock.Controller) *MockStorageAccountsClient {
	mock := &MockStorageAccountsClient{ctrl: ctrl}
	mock.recorder = &MockStorageAccountsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageAccountsClient) EXPECT() *MockStorageAccountsClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockStorageAccountsClient) List(arg0 context.Context) (storage.AccountListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(storage.AccountListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockStorageAccountsClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStorageAccountsClient)(nil).List), arg0)
}

// ListKeys mocks base method.
func (m *MockStorageAccountsClient) ListKeys(arg0 context.Context, arg1, arg2 string, arg3 storage.ListKeyExpand) (storage.AccountListKeysResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKeys", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(storage.AccountListKeysResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKeys indicates an expected call of ListKeys.
func (mr *MockStorageAccountsClientMockRecorder) ListKeys(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockStorageAccountsClient)(nil).ListKeys), arg0, arg1, arg2, arg3)
}

// MockStorageBlobServicePropertiesClient is a mock of StorageBlobServicePropertiesClient interface.
type MockStorageBlobServicePropertiesClient struct {
	ctrl     *gomock.Controller
	recorder *MockStorageBlobServicePropertiesClientMockRecorder
}

// MockStorageBlobServicePropertiesClientMockRecorder is the mock recorder for MockStorageBlobServicePropertiesClient.
type MockStorageBlobServicePropertiesClientMockRecorder struct {
	mock *MockStorageBlobServicePropertiesClient
}

// NewMockStorageBlobServicePropertiesClient creates a new mock instance.
func NewMockStorageBlobServicePropertiesClient(ctrl *gomock.Controller) *MockStorageBlobServicePropertiesClient {
	mock := &MockStorageBlobServicePropertiesClient{ctrl: ctrl}
	mock.recorder = &MockStorageBlobServicePropertiesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageBlobServicePropertiesClient) EXPECT() *MockStorageBlobServicePropertiesClientMockRecorder {
	return m.recorder
}

// GetServiceProperties mocks base method.
func (m *MockStorageBlobServicePropertiesClient) GetServiceProperties(arg0 context.Context, arg1 string) (accounts.GetServicePropertiesResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceProperties", arg0, arg1)
	ret0, _ := ret[0].(accounts.GetServicePropertiesResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceProperties indicates an expected call of GetServiceProperties.
func (mr *MockStorageBlobServicePropertiesClientMockRecorder) GetServiceProperties(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceProperties", reflect.TypeOf((*MockStorageBlobServicePropertiesClient)(nil).GetServiceProperties), arg0, arg1)
}

// MockStorageBlobServicesClient is a mock of StorageBlobServicesClient interface.
type MockStorageBlobServicesClient struct {
	ctrl     *gomock.Controller
	recorder *MockStorageBlobServicesClientMockRecorder
}

// MockStorageBlobServicesClientMockRecorder is the mock recorder for MockStorageBlobServicesClient.
type MockStorageBlobServicesClientMockRecorder struct {
	mock *MockStorageBlobServicesClient
}

// NewMockStorageBlobServicesClient creates a new mock instance.
func NewMockStorageBlobServicesClient(ctrl *gomock.Controller) *MockStorageBlobServicesClient {
	mock := &MockStorageBlobServicesClient{ctrl: ctrl}
	mock.recorder = &MockStorageBlobServicesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageBlobServicesClient) EXPECT() *MockStorageBlobServicesClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockStorageBlobServicesClient) List(arg0 context.Context, arg1, arg2 string) (storage.BlobServiceItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(storage.BlobServiceItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockStorageBlobServicesClientMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStorageBlobServicesClient)(nil).List), arg0, arg1, arg2)
}

// MockStorageContainerClient is a mock of StorageContainerClient interface.
type MockStorageContainerClient struct {
	ctrl     *gomock.Controller
	recorder *MockStorageContainerClientMockRecorder
}

// MockStorageContainerClientMockRecorder is the mock recorder for MockStorageContainerClient.
type MockStorageContainerClientMockRecorder struct {
	mock *MockStorageContainerClient
}

// NewMockStorageContainerClient creates a new mock instance.
func NewMockStorageContainerClient(ctrl *gomock.Controller) *MockStorageContainerClient {
	mock := &MockStorageContainerClient{ctrl: ctrl}
	mock.recorder = &MockStorageContainerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageContainerClient) EXPECT() *MockStorageContainerClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockStorageContainerClient) List(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 storage.ListContainersInclude) (storage.ListContainerItemsPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(storage.ListContainerItemsPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockStorageContainerClientMockRecorder) List(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStorageContainerClient)(nil).List), arg0, arg1, arg2, arg3, arg4, arg5)
}

// MockStorageQueueServicePropertiesClient is a mock of StorageQueueServicePropertiesClient interface.
type MockStorageQueueServicePropertiesClient struct {
	ctrl     *gomock.Controller
	recorder *MockStorageQueueServicePropertiesClientMockRecorder
}

// MockStorageQueueServicePropertiesClientMockRecorder is the mock recorder for MockStorageQueueServicePropertiesClient.
type MockStorageQueueServicePropertiesClientMockRecorder struct {
	mock *MockStorageQueueServicePropertiesClient
}

// NewMockStorageQueueServicePropertiesClient creates a new mock instance.
func NewMockStorageQueueServicePropertiesClient(ctrl *gomock.Controller) *MockStorageQueueServicePropertiesClient {
	mock := &MockStorageQueueServicePropertiesClient{ctrl: ctrl}
	mock.recorder = &MockStorageQueueServicePropertiesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageQueueServicePropertiesClient) EXPECT() *MockStorageQueueServicePropertiesClientMockRecorder {
	return m.recorder
}

// GetServiceProperties mocks base method.
func (m *MockStorageQueueServicePropertiesClient) GetServiceProperties(arg0 context.Context, arg1 string) (queues.StorageServicePropertiesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceProperties", arg0, arg1)
	ret0, _ := ret[0].(queues.StorageServicePropertiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceProperties indicates an expected call of GetServiceProperties.
func (mr *MockStorageQueueServicePropertiesClientMockRecorder) GetServiceProperties(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceProperties", reflect.TypeOf((*MockStorageQueueServicePropertiesClient)(nil).GetServiceProperties), arg0, arg1)
}
