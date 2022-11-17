// Code generated by MockGen. DO NOT EDIT.
// Source: security_groups.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	network "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/network"
	gomock "github.com/golang/mock/gomock"
)

// MockSecurityGroupsClient is a mock of SecurityGroupsClient interface.
type MockSecurityGroupsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecurityGroupsClientMockRecorder
}

// MockSecurityGroupsClientMockRecorder is the mock recorder for MockSecurityGroupsClient.
type MockSecurityGroupsClientMockRecorder struct {
	mock *MockSecurityGroupsClient
}

// NewMockSecurityGroupsClient creates a new mock instance.
func NewMockSecurityGroupsClient(ctrl *gomock.Controller) *MockSecurityGroupsClient {
	mock := &MockSecurityGroupsClient{ctrl: ctrl}
	mock.recorder = &MockSecurityGroupsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecurityGroupsClient) EXPECT() *MockSecurityGroupsClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockSecurityGroupsClient) Get(arg0 context.Context, arg1, arg2 string, arg3 *armnetwork.SecurityGroupsClientGetOptions) (armnetwork.SecurityGroupsClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(armnetwork.SecurityGroupsClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSecurityGroupsClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSecurityGroupsClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// NewListAllPager mocks base method.
func (m *MockSecurityGroupsClient) NewListAllPager(arg0 *armnetwork.SecurityGroupsClientListAllOptions) *network.RuntimePagerArmnetworkSecurityGroupsClientListAllResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListAllPager", arg0)
	ret0, _ := ret[0].(*network.RuntimePagerArmnetworkSecurityGroupsClientListAllResponse)
	return ret0
}

// NewListAllPager indicates an expected call of NewListAllPager.
func (mr *MockSecurityGroupsClientMockRecorder) NewListAllPager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListAllPager", reflect.TypeOf((*MockSecurityGroupsClient)(nil).NewListAllPager), arg0)
}

// NewListPager mocks base method.
func (m *MockSecurityGroupsClient) NewListPager(arg0 string, arg1 *armnetwork.SecurityGroupsClientListOptions) *network.RuntimePagerArmnetworkSecurityGroupsClientListResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListPager", arg0, arg1)
	ret0, _ := ret[0].(*network.RuntimePagerArmnetworkSecurityGroupsClientListResponse)
	return ret0
}

// NewListPager indicates an expected call of NewListPager.
func (mr *MockSecurityGroupsClientMockRecorder) NewListPager(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListPager", reflect.TypeOf((*MockSecurityGroupsClient)(nil).NewListPager), arg0, arg1)
}
