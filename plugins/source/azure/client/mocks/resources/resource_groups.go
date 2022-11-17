// Code generated by MockGen. DO NOT EDIT.
// Source: resource_groups.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	armresources "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	resources "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/resources"
	gomock "github.com/golang/mock/gomock"
)

// MockResourceGroupsClient is a mock of ResourceGroupsClient interface.
type MockResourceGroupsClient struct {
	ctrl     *gomock.Controller
	recorder *MockResourceGroupsClientMockRecorder
}

// MockResourceGroupsClientMockRecorder is the mock recorder for MockResourceGroupsClient.
type MockResourceGroupsClientMockRecorder struct {
	mock *MockResourceGroupsClient
}

// NewMockResourceGroupsClient creates a new mock instance.
func NewMockResourceGroupsClient(ctrl *gomock.Controller) *MockResourceGroupsClient {
	mock := &MockResourceGroupsClient{ctrl: ctrl}
	mock.recorder = &MockResourceGroupsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceGroupsClient) EXPECT() *MockResourceGroupsClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockResourceGroupsClient) Get(arg0 context.Context, arg1 string, arg2 *armresources.ResourceGroupsClientGetOptions) (armresources.ResourceGroupsClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(armresources.ResourceGroupsClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockResourceGroupsClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockResourceGroupsClient)(nil).Get), arg0, arg1, arg2)
}

// NewListPager mocks base method.
func (m *MockResourceGroupsClient) NewListPager(arg0 *armresources.ResourceGroupsClientListOptions) *resources.RuntimePagerArmresourcesResourceGroupsClientListResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListPager", arg0)
	ret0, _ := ret[0].(*resources.RuntimePagerArmresourcesResourceGroupsClientListResponse)
	return ret0
}

// NewListPager indicates an expected call of NewListPager.
func (mr *MockResourceGroupsClientMockRecorder) NewListPager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListPager", reflect.TypeOf((*MockResourceGroupsClient)(nil).NewListPager), arg0)
}
