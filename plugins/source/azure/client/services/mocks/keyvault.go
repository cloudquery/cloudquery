// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/azure/client/services (interfaces: KeyVault71Client,VaultsClient,ManagedHSMsClient,KeysClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	keyvault "github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault"
	keyvault0 "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"
	keyvault1 "github.com/Azure/azure-sdk-for-go/services/preview/keyvault/mgmt/2020-04-01-preview/keyvault"
	gomock "github.com/golang/mock/gomock"
)

// MockKeyVault71Client is a mock of KeyVault71Client interface.
type MockKeyVault71Client struct {
	ctrl     *gomock.Controller
	recorder *MockKeyVault71ClientMockRecorder
}

// MockKeyVault71ClientMockRecorder is the mock recorder for MockKeyVault71Client.
type MockKeyVault71ClientMockRecorder struct {
	mock *MockKeyVault71Client
}

// NewMockKeyVault71Client creates a new mock instance.
func NewMockKeyVault71Client(ctrl *gomock.Controller) *MockKeyVault71Client {
	mock := &MockKeyVault71Client{ctrl: ctrl}
	mock.recorder = &MockKeyVault71ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeyVault71Client) EXPECT() *MockKeyVault71ClientMockRecorder {
	return m.recorder
}

// GetKeys mocks base method.
func (m *MockKeyVault71Client) GetKeys(arg0 context.Context, arg1 string, arg2 *int32) (keyvault0.KeyListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeys", arg0, arg1, arg2)
	ret0, _ := ret[0].(keyvault0.KeyListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeys indicates an expected call of GetKeys.
func (mr *MockKeyVault71ClientMockRecorder) GetKeys(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeys", reflect.TypeOf((*MockKeyVault71Client)(nil).GetKeys), arg0, arg1, arg2)
}

// GetSecrets mocks base method.
func (m *MockKeyVault71Client) GetSecrets(arg0 context.Context, arg1 string, arg2 *int32) (keyvault0.SecretListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecrets", arg0, arg1, arg2)
	ret0, _ := ret[0].(keyvault0.SecretListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecrets indicates an expected call of GetSecrets.
func (mr *MockKeyVault71ClientMockRecorder) GetSecrets(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecrets", reflect.TypeOf((*MockKeyVault71Client)(nil).GetSecrets), arg0, arg1, arg2)
}

// MockVaultsClient is a mock of VaultsClient interface.
type MockVaultsClient struct {
	ctrl     *gomock.Controller
	recorder *MockVaultsClientMockRecorder
}

// MockVaultsClientMockRecorder is the mock recorder for MockVaultsClient.
type MockVaultsClientMockRecorder struct {
	mock *MockVaultsClient
}

// NewMockVaultsClient creates a new mock instance.
func NewMockVaultsClient(ctrl *gomock.Controller) *MockVaultsClient {
	mock := &MockVaultsClient{ctrl: ctrl}
	mock.recorder = &MockVaultsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVaultsClient) EXPECT() *MockVaultsClientMockRecorder {
	return m.recorder
}

// ListBySubscription mocks base method.
func (m *MockVaultsClient) ListBySubscription(arg0 context.Context, arg1 *int32) (keyvault.VaultListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubscription", arg0, arg1)
	ret0, _ := ret[0].(keyvault.VaultListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySubscription indicates an expected call of ListBySubscription.
func (mr *MockVaultsClientMockRecorder) ListBySubscription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubscription", reflect.TypeOf((*MockVaultsClient)(nil).ListBySubscription), arg0, arg1)
}

// MockManagedHSMsClient is a mock of ManagedHSMsClient interface.
type MockManagedHSMsClient struct {
	ctrl     *gomock.Controller
	recorder *MockManagedHSMsClientMockRecorder
}

// MockManagedHSMsClientMockRecorder is the mock recorder for MockManagedHSMsClient.
type MockManagedHSMsClientMockRecorder struct {
	mock *MockManagedHSMsClient
}

// NewMockManagedHSMsClient creates a new mock instance.
func NewMockManagedHSMsClient(ctrl *gomock.Controller) *MockManagedHSMsClient {
	mock := &MockManagedHSMsClient{ctrl: ctrl}
	mock.recorder = &MockManagedHSMsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManagedHSMsClient) EXPECT() *MockManagedHSMsClientMockRecorder {
	return m.recorder
}

// ListBySubscription mocks base method.
func (m *MockManagedHSMsClient) ListBySubscription(arg0 context.Context, arg1 *int32) (keyvault1.ManagedHsmListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubscription", arg0, arg1)
	ret0, _ := ret[0].(keyvault1.ManagedHsmListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySubscription indicates an expected call of ListBySubscription.
func (mr *MockManagedHSMsClientMockRecorder) ListBySubscription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubscription", reflect.TypeOf((*MockManagedHSMsClient)(nil).ListBySubscription), arg0, arg1)
}

// MockKeysClient is a mock of KeysClient interface.
type MockKeysClient struct {
	ctrl     *gomock.Controller
	recorder *MockKeysClientMockRecorder
}

// MockKeysClientMockRecorder is the mock recorder for MockKeysClient.
type MockKeysClientMockRecorder struct {
	mock *MockKeysClient
}

// NewMockKeysClient creates a new mock instance.
func NewMockKeysClient(ctrl *gomock.Controller) *MockKeysClient {
	mock := &MockKeysClient{ctrl: ctrl}
	mock.recorder = &MockKeysClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeysClient) EXPECT() *MockKeysClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockKeysClient) List(arg0 context.Context, arg1, arg2 string) (keyvault.KeyListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(keyvault.KeyListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockKeysClientMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockKeysClient)(nil).List), arg0, arg1, arg2)
}
