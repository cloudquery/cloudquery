// Code generated by MockGen. DO NOT EDIT.
// Source: virtual_network_gateway_connections.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	network "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/network"
	gomock "github.com/golang/mock/gomock"
)

// MockVirtualNetworkGatewayConnectionsClient is a mock of VirtualNetworkGatewayConnectionsClient interface.
type MockVirtualNetworkGatewayConnectionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualNetworkGatewayConnectionsClientMockRecorder
}

// MockVirtualNetworkGatewayConnectionsClientMockRecorder is the mock recorder for MockVirtualNetworkGatewayConnectionsClient.
type MockVirtualNetworkGatewayConnectionsClientMockRecorder struct {
	mock *MockVirtualNetworkGatewayConnectionsClient
}

// NewMockVirtualNetworkGatewayConnectionsClient creates a new mock instance.
func NewMockVirtualNetworkGatewayConnectionsClient(ctrl *gomock.Controller) *MockVirtualNetworkGatewayConnectionsClient {
	mock := &MockVirtualNetworkGatewayConnectionsClient{ctrl: ctrl}
	mock.recorder = &MockVirtualNetworkGatewayConnectionsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualNetworkGatewayConnectionsClient) EXPECT() *MockVirtualNetworkGatewayConnectionsClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockVirtualNetworkGatewayConnectionsClient) Get(arg0 context.Context, arg1, arg2 string, arg3 *armnetwork.VirtualNetworkGatewayConnectionsClientGetOptions) (armnetwork.VirtualNetworkGatewayConnectionsClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(armnetwork.VirtualNetworkGatewayConnectionsClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockVirtualNetworkGatewayConnectionsClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVirtualNetworkGatewayConnectionsClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// GetSharedKey mocks base method.
func (m *MockVirtualNetworkGatewayConnectionsClient) GetSharedKey(arg0 context.Context, arg1, arg2 string, arg3 *armnetwork.VirtualNetworkGatewayConnectionsClientGetSharedKeyOptions) (armnetwork.VirtualNetworkGatewayConnectionsClientGetSharedKeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSharedKey", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(armnetwork.VirtualNetworkGatewayConnectionsClientGetSharedKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSharedKey indicates an expected call of GetSharedKey.
func (mr *MockVirtualNetworkGatewayConnectionsClientMockRecorder) GetSharedKey(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSharedKey", reflect.TypeOf((*MockVirtualNetworkGatewayConnectionsClient)(nil).GetSharedKey), arg0, arg1, arg2, arg3)
}

// NewListPager mocks base method.
func (m *MockVirtualNetworkGatewayConnectionsClient) NewListPager(arg0 string, arg1 *armnetwork.VirtualNetworkGatewayConnectionsClientListOptions) *network.RuntimePagerArmnetworkVirtualNetworkGatewayConnectionsClientListResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListPager", arg0, arg1)
	ret0, _ := ret[0].(*network.RuntimePagerArmnetworkVirtualNetworkGatewayConnectionsClientListResponse)
	return ret0
}

// NewListPager indicates an expected call of NewListPager.
func (mr *MockVirtualNetworkGatewayConnectionsClientMockRecorder) NewListPager(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListPager", reflect.TypeOf((*MockVirtualNetworkGatewayConnectionsClient)(nil).NewListPager), arg0, arg1)
}
