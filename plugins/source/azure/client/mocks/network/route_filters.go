// Code generated by MockGen. DO NOT EDIT.
// Source: route_filters.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	network "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/network"
	gomock "github.com/golang/mock/gomock"
)

// MockRouteFiltersClient is a mock of RouteFiltersClient interface.
type MockRouteFiltersClient struct {
	ctrl     *gomock.Controller
	recorder *MockRouteFiltersClientMockRecorder
}

// MockRouteFiltersClientMockRecorder is the mock recorder for MockRouteFiltersClient.
type MockRouteFiltersClientMockRecorder struct {
	mock *MockRouteFiltersClient
}

// NewMockRouteFiltersClient creates a new mock instance.
func NewMockRouteFiltersClient(ctrl *gomock.Controller) *MockRouteFiltersClient {
	mock := &MockRouteFiltersClient{ctrl: ctrl}
	mock.recorder = &MockRouteFiltersClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteFiltersClient) EXPECT() *MockRouteFiltersClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockRouteFiltersClient) Get(arg0 context.Context, arg1, arg2 string, arg3 *armnetwork.RouteFiltersClientGetOptions) (armnetwork.RouteFiltersClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(armnetwork.RouteFiltersClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRouteFiltersClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRouteFiltersClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// NewListByResourceGroupPager mocks base method.
func (m *MockRouteFiltersClient) NewListByResourceGroupPager(arg0 string, arg1 *armnetwork.RouteFiltersClientListByResourceGroupOptions) *network.RuntimePagerArmnetworkRouteFiltersClientListByResourceGroupResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListByResourceGroupPager", arg0, arg1)
	ret0, _ := ret[0].(*network.RuntimePagerArmnetworkRouteFiltersClientListByResourceGroupResponse)
	return ret0
}

// NewListByResourceGroupPager indicates an expected call of NewListByResourceGroupPager.
func (mr *MockRouteFiltersClientMockRecorder) NewListByResourceGroupPager(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListByResourceGroupPager", reflect.TypeOf((*MockRouteFiltersClient)(nil).NewListByResourceGroupPager), arg0, arg1)
}

// NewListPager mocks base method.
func (m *MockRouteFiltersClient) NewListPager(arg0 *armnetwork.RouteFiltersClientListOptions) *network.RuntimePagerArmnetworkRouteFiltersClientListResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListPager", arg0)
	ret0, _ := ret[0].(*network.RuntimePagerArmnetworkRouteFiltersClientListResponse)
	return ret0
}

// NewListPager indicates an expected call of NewListPager.
func (mr *MockRouteFiltersClientMockRecorder) NewListPager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListPager", reflect.TypeOf((*MockRouteFiltersClient)(nil).NewListPager), arg0)
}
