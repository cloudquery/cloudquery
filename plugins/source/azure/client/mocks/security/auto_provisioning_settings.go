// Code generated by MockGen. DO NOT EDIT.
// Source: auto_provisioning_settings.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	armsecurity "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	security "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/security"
	gomock "github.com/golang/mock/gomock"
)

// MockAutoProvisioningSettingsClient is a mock of AutoProvisioningSettingsClient interface.
type MockAutoProvisioningSettingsClient struct {
	ctrl     *gomock.Controller
	recorder *MockAutoProvisioningSettingsClientMockRecorder
}

// MockAutoProvisioningSettingsClientMockRecorder is the mock recorder for MockAutoProvisioningSettingsClient.
type MockAutoProvisioningSettingsClientMockRecorder struct {
	mock *MockAutoProvisioningSettingsClient
}

// NewMockAutoProvisioningSettingsClient creates a new mock instance.
func NewMockAutoProvisioningSettingsClient(ctrl *gomock.Controller) *MockAutoProvisioningSettingsClient {
	mock := &MockAutoProvisioningSettingsClient{ctrl: ctrl}
	mock.recorder = &MockAutoProvisioningSettingsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAutoProvisioningSettingsClient) EXPECT() *MockAutoProvisioningSettingsClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockAutoProvisioningSettingsClient) Get(arg0 context.Context, arg1 string, arg2 *armsecurity.AutoProvisioningSettingsClientGetOptions) (armsecurity.AutoProvisioningSettingsClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(armsecurity.AutoProvisioningSettingsClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAutoProvisioningSettingsClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAutoProvisioningSettingsClient)(nil).Get), arg0, arg1, arg2)
}

// NewListPager mocks base method.
func (m *MockAutoProvisioningSettingsClient) NewListPager(arg0 *armsecurity.AutoProvisioningSettingsClientListOptions) *security.RuntimePagerArmsecurityAutoProvisioningSettingsClientListResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListPager", arg0)
	ret0, _ := ret[0].(*security.RuntimePagerArmsecurityAutoProvisioningSettingsClientListResponse)
	return ret0
}

// NewListPager indicates an expected call of NewListPager.
func (mr *MockAutoProvisioningSettingsClientMockRecorder) NewListPager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListPager", reflect.TypeOf((*MockAutoProvisioningSettingsClient)(nil).NewListPager), arg0)
}
