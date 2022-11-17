// Code generated by MockGen. DO NOT EDIT.
// Source: express_route_circuits.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	network "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/network"
	gomock "github.com/golang/mock/gomock"
)

// MockExpressRouteCircuitsClient is a mock of ExpressRouteCircuitsClient interface.
type MockExpressRouteCircuitsClient struct {
	ctrl     *gomock.Controller
	recorder *MockExpressRouteCircuitsClientMockRecorder
}

// MockExpressRouteCircuitsClientMockRecorder is the mock recorder for MockExpressRouteCircuitsClient.
type MockExpressRouteCircuitsClientMockRecorder struct {
	mock *MockExpressRouteCircuitsClient
}

// NewMockExpressRouteCircuitsClient creates a new mock instance.
func NewMockExpressRouteCircuitsClient(ctrl *gomock.Controller) *MockExpressRouteCircuitsClient {
	mock := &MockExpressRouteCircuitsClient{ctrl: ctrl}
	mock.recorder = &MockExpressRouteCircuitsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExpressRouteCircuitsClient) EXPECT() *MockExpressRouteCircuitsClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockExpressRouteCircuitsClient) Get(arg0 context.Context, arg1, arg2 string, arg3 *armnetwork.ExpressRouteCircuitsClientGetOptions) (armnetwork.ExpressRouteCircuitsClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(armnetwork.ExpressRouteCircuitsClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockExpressRouteCircuitsClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockExpressRouteCircuitsClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// GetPeeringStats mocks base method.
func (m *MockExpressRouteCircuitsClient) GetPeeringStats(arg0 context.Context, arg1, arg2, arg3 string, arg4 *armnetwork.ExpressRouteCircuitsClientGetPeeringStatsOptions) (armnetwork.ExpressRouteCircuitsClientGetPeeringStatsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeeringStats", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(armnetwork.ExpressRouteCircuitsClientGetPeeringStatsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeeringStats indicates an expected call of GetPeeringStats.
func (mr *MockExpressRouteCircuitsClientMockRecorder) GetPeeringStats(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeeringStats", reflect.TypeOf((*MockExpressRouteCircuitsClient)(nil).GetPeeringStats), arg0, arg1, arg2, arg3, arg4)
}

// GetStats mocks base method.
func (m *MockExpressRouteCircuitsClient) GetStats(arg0 context.Context, arg1, arg2 string, arg3 *armnetwork.ExpressRouteCircuitsClientGetStatsOptions) (armnetwork.ExpressRouteCircuitsClientGetStatsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStats", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(armnetwork.ExpressRouteCircuitsClientGetStatsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStats indicates an expected call of GetStats.
func (mr *MockExpressRouteCircuitsClientMockRecorder) GetStats(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStats", reflect.TypeOf((*MockExpressRouteCircuitsClient)(nil).GetStats), arg0, arg1, arg2, arg3)
}

// NewListAllPager mocks base method.
func (m *MockExpressRouteCircuitsClient) NewListAllPager(arg0 *armnetwork.ExpressRouteCircuitsClientListAllOptions) *network.RuntimePagerArmnetworkExpressRouteCircuitsClientListAllResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListAllPager", arg0)
	ret0, _ := ret[0].(*network.RuntimePagerArmnetworkExpressRouteCircuitsClientListAllResponse)
	return ret0
}

// NewListAllPager indicates an expected call of NewListAllPager.
func (mr *MockExpressRouteCircuitsClientMockRecorder) NewListAllPager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListAllPager", reflect.TypeOf((*MockExpressRouteCircuitsClient)(nil).NewListAllPager), arg0)
}

// NewListPager mocks base method.
func (m *MockExpressRouteCircuitsClient) NewListPager(arg0 string, arg1 *armnetwork.ExpressRouteCircuitsClientListOptions) *network.RuntimePagerArmnetworkExpressRouteCircuitsClientListResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListPager", arg0, arg1)
	ret0, _ := ret[0].(*network.RuntimePagerArmnetworkExpressRouteCircuitsClientListResponse)
	return ret0
}

// NewListPager indicates an expected call of NewListPager.
func (mr *MockExpressRouteCircuitsClientMockRecorder) NewListPager(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListPager", reflect.TypeOf((*MockExpressRouteCircuitsClient)(nil).NewListPager), arg0, arg1)
}
