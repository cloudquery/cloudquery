// Code generated by MockGen. DO NOT EDIT.
// Source: account.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	armbatch "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
	batch "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/batch"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountClient is a mock of AccountClient interface.
type MockAccountClient struct {
	ctrl     *gomock.Controller
	recorder *MockAccountClientMockRecorder
}

// MockAccountClientMockRecorder is the mock recorder for MockAccountClient.
type MockAccountClientMockRecorder struct {
	mock *MockAccountClient
}

// NewMockAccountClient creates a new mock instance.
func NewMockAccountClient(ctrl *gomock.Controller) *MockAccountClient {
	mock := &MockAccountClient{ctrl: ctrl}
	mock.recorder = &MockAccountClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountClient) EXPECT() *MockAccountClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockAccountClient) Get(arg0 context.Context, arg1, arg2 string, arg3 *armbatch.AccountClientGetOptions) (armbatch.AccountClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(armbatch.AccountClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAccountClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAccountClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// GetDetector mocks base method.
func (m *MockAccountClient) GetDetector(arg0 context.Context, arg1, arg2, arg3 string, arg4 *armbatch.AccountClientGetDetectorOptions) (armbatch.AccountClientGetDetectorResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetector", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(armbatch.AccountClientGetDetectorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetector indicates an expected call of GetDetector.
func (mr *MockAccountClientMockRecorder) GetDetector(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetector", reflect.TypeOf((*MockAccountClient)(nil).GetDetector), arg0, arg1, arg2, arg3, arg4)
}

// GetKeys mocks base method.
func (m *MockAccountClient) GetKeys(arg0 context.Context, arg1, arg2 string, arg3 *armbatch.AccountClientGetKeysOptions) (armbatch.AccountClientGetKeysResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeys", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(armbatch.AccountClientGetKeysResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeys indicates an expected call of GetKeys.
func (mr *MockAccountClientMockRecorder) GetKeys(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeys", reflect.TypeOf((*MockAccountClient)(nil).GetKeys), arg0, arg1, arg2, arg3)
}

// NewListByResourceGroupPager mocks base method.
func (m *MockAccountClient) NewListByResourceGroupPager(arg0 string, arg1 *armbatch.AccountClientListByResourceGroupOptions) *batch.RuntimePagerArmbatchAccountClientListByResourceGroupResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListByResourceGroupPager", arg0, arg1)
	ret0, _ := ret[0].(*batch.RuntimePagerArmbatchAccountClientListByResourceGroupResponse)
	return ret0
}

// NewListByResourceGroupPager indicates an expected call of NewListByResourceGroupPager.
func (mr *MockAccountClientMockRecorder) NewListByResourceGroupPager(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListByResourceGroupPager", reflect.TypeOf((*MockAccountClient)(nil).NewListByResourceGroupPager), arg0, arg1)
}

// NewListDetectorsPager mocks base method.
func (m *MockAccountClient) NewListDetectorsPager(arg0, arg1 string, arg2 *armbatch.AccountClientListDetectorsOptions) *batch.RuntimePagerArmbatchAccountClientListDetectorsResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListDetectorsPager", arg0, arg1, arg2)
	ret0, _ := ret[0].(*batch.RuntimePagerArmbatchAccountClientListDetectorsResponse)
	return ret0
}

// NewListDetectorsPager indicates an expected call of NewListDetectorsPager.
func (mr *MockAccountClientMockRecorder) NewListDetectorsPager(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListDetectorsPager", reflect.TypeOf((*MockAccountClient)(nil).NewListDetectorsPager), arg0, arg1, arg2)
}

// NewListOutboundNetworkDependenciesEndpointsPager mocks base method.
func (m *MockAccountClient) NewListOutboundNetworkDependenciesEndpointsPager(arg0, arg1 string, arg2 *armbatch.AccountClientListOutboundNetworkDependenciesEndpointsOptions) *batch.RuntimePagerArmbatchAccountClientListOutboundNetworkDependenciesEndpointsResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListOutboundNetworkDependenciesEndpointsPager", arg0, arg1, arg2)
	ret0, _ := ret[0].(*batch.RuntimePagerArmbatchAccountClientListOutboundNetworkDependenciesEndpointsResponse)
	return ret0
}

// NewListOutboundNetworkDependenciesEndpointsPager indicates an expected call of NewListOutboundNetworkDependenciesEndpointsPager.
func (mr *MockAccountClientMockRecorder) NewListOutboundNetworkDependenciesEndpointsPager(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListOutboundNetworkDependenciesEndpointsPager", reflect.TypeOf((*MockAccountClient)(nil).NewListOutboundNetworkDependenciesEndpointsPager), arg0, arg1, arg2)
}

// NewListPager mocks base method.
func (m *MockAccountClient) NewListPager(arg0 *armbatch.AccountClientListOptions) *batch.RuntimePagerArmbatchAccountClientListResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListPager", arg0)
	ret0, _ := ret[0].(*batch.RuntimePagerArmbatchAccountClientListResponse)
	return ret0
}

// NewListPager indicates an expected call of NewListPager.
func (mr *MockAccountClientMockRecorder) NewListPager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListPager", reflect.TypeOf((*MockAccountClient)(nil).NewListPager), arg0)
}
