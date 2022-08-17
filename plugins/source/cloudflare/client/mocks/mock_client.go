// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/cloudflare/client (interfaces: Api)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	cloudflare "github.com/cloudflare/cloudflare-go"
	gomock "github.com/golang/mock/gomock"
)

// MockApi is a mock of Api interface.
type MockApi struct {
	ctrl     *gomock.Controller
	recorder *MockApiMockRecorder
}

// MockApiMockRecorder is the mock recorder for MockApi.
type MockApiMockRecorder struct {
	mock *MockApi
}

// NewMockApi creates a new mock instance.
func NewMockApi(ctrl *gomock.Controller) *MockApi {
	mock := &MockApi{ctrl: ctrl}
	mock.recorder = &MockApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApi) EXPECT() *MockApiMockRecorder {
	return m.recorder
}

// AccountMembers mocks base method.
func (m *MockApi) AccountMembers(arg0 context.Context, arg1 string, arg2 cloudflare.PaginationOptions) ([]cloudflare.AccountMember, cloudflare.ResultInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccountMembers", arg0, arg1, arg2)
	ret0, _ := ret[0].([]cloudflare.AccountMember)
	ret1, _ := ret[1].(cloudflare.ResultInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AccountMembers indicates an expected call of AccountMembers.
func (mr *MockApiMockRecorder) AccountMembers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountMembers", reflect.TypeOf((*MockApi)(nil).AccountMembers), arg0, arg1, arg2)
}

// Accounts mocks base method.
func (m *MockApi) Accounts(arg0 context.Context, arg1 cloudflare.AccountsListParams) ([]cloudflare.Account, cloudflare.ResultInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accounts", arg0, arg1)
	ret0, _ := ret[0].([]cloudflare.Account)
	ret1, _ := ret[1].(cloudflare.ResultInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Accounts indicates an expected call of Accounts.
func (mr *MockApiMockRecorder) Accounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accounts", reflect.TypeOf((*MockApi)(nil).Accounts), arg0, arg1)
}

// DNSRecords mocks base method.
func (m *MockApi) DNSRecords(arg0 context.Context, arg1 string, arg2 cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DNSRecords", arg0, arg1, arg2)
	ret0, _ := ret[0].([]cloudflare.DNSRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DNSRecords indicates an expected call of DNSRecords.
func (mr *MockApiMockRecorder) DNSRecords(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DNSRecords", reflect.TypeOf((*MockApi)(nil).DNSRecords), arg0, arg1, arg2)
}

// ListCertificatePacks mocks base method.
func (m *MockApi) ListCertificatePacks(arg0 context.Context, arg1 string) ([]cloudflare.CertificatePack, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCertificatePacks", arg0, arg1)
	ret0, _ := ret[0].([]cloudflare.CertificatePack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCertificatePacks indicates an expected call of ListCertificatePacks.
func (mr *MockApiMockRecorder) ListCertificatePacks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCertificatePacks", reflect.TypeOf((*MockApi)(nil).ListCertificatePacks), arg0, arg1)
}

// ListImages mocks base method.
func (m *MockApi) ListImages(arg0 context.Context, arg1 string, arg2 cloudflare.PaginationOptions) ([]cloudflare.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImages", arg0, arg1, arg2)
	ret0, _ := ret[0].([]cloudflare.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImages indicates an expected call of ListImages.
func (mr *MockApiMockRecorder) ListImages(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImages", reflect.TypeOf((*MockApi)(nil).ListImages), arg0, arg1, arg2)
}

// ListWAFGroups mocks base method.
func (m *MockApi) ListWAFGroups(arg0 context.Context, arg1, arg2 string) ([]cloudflare.WAFGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWAFGroups", arg0, arg1, arg2)
	ret0, _ := ret[0].([]cloudflare.WAFGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWAFGroups indicates an expected call of ListWAFGroups.
func (mr *MockApiMockRecorder) ListWAFGroups(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWAFGroups", reflect.TypeOf((*MockApi)(nil).ListWAFGroups), arg0, arg1, arg2)
}

// ListWAFOverrides mocks base method.
func (m *MockApi) ListWAFOverrides(arg0 context.Context, arg1 string) ([]cloudflare.WAFOverride, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWAFOverrides", arg0, arg1)
	ret0, _ := ret[0].([]cloudflare.WAFOverride)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWAFOverrides indicates an expected call of ListWAFOverrides.
func (mr *MockApiMockRecorder) ListWAFOverrides(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWAFOverrides", reflect.TypeOf((*MockApi)(nil).ListWAFOverrides), arg0, arg1)
}

// ListWAFPackages mocks base method.
func (m *MockApi) ListWAFPackages(arg0 context.Context, arg1 string) ([]cloudflare.WAFPackage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWAFPackages", arg0, arg1)
	ret0, _ := ret[0].([]cloudflare.WAFPackage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWAFPackages indicates an expected call of ListWAFPackages.
func (mr *MockApiMockRecorder) ListWAFPackages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWAFPackages", reflect.TypeOf((*MockApi)(nil).ListWAFPackages), arg0, arg1)
}

// ListWAFRules mocks base method.
func (m *MockApi) ListWAFRules(arg0 context.Context, arg1, arg2 string) ([]cloudflare.WAFRule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWAFRules", arg0, arg1, arg2)
	ret0, _ := ret[0].([]cloudflare.WAFRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWAFRules indicates an expected call of ListWAFRules.
func (mr *MockApiMockRecorder) ListWAFRules(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWAFRules", reflect.TypeOf((*MockApi)(nil).ListWAFRules), arg0, arg1, arg2)
}

// ListWorkerCronTriggers mocks base method.
func (m *MockApi) ListWorkerCronTriggers(arg0 context.Context, arg1, arg2 string) ([]cloudflare.WorkerCronTrigger, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWorkerCronTriggers", arg0, arg1, arg2)
	ret0, _ := ret[0].([]cloudflare.WorkerCronTrigger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWorkerCronTriggers indicates an expected call of ListWorkerCronTriggers.
func (mr *MockApiMockRecorder) ListWorkerCronTriggers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkerCronTriggers", reflect.TypeOf((*MockApi)(nil).ListWorkerCronTriggers), arg0, arg1, arg2)
}

// ListWorkerRoutes mocks base method.
func (m *MockApi) ListWorkerRoutes(arg0 context.Context, arg1 string) (cloudflare.WorkerRoutesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWorkerRoutes", arg0, arg1)
	ret0, _ := ret[0].(cloudflare.WorkerRoutesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWorkerRoutes indicates an expected call of ListWorkerRoutes.
func (mr *MockApiMockRecorder) ListWorkerRoutes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkerRoutes", reflect.TypeOf((*MockApi)(nil).ListWorkerRoutes), arg0, arg1)
}

// ListWorkerScripts mocks base method.
func (m *MockApi) ListWorkerScripts(arg0 context.Context) (cloudflare.WorkerListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWorkerScripts", arg0)
	ret0, _ := ret[0].(cloudflare.WorkerListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWorkerScripts indicates an expected call of ListWorkerScripts.
func (mr *MockApiMockRecorder) ListWorkerScripts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkerScripts", reflect.TypeOf((*MockApi)(nil).ListWorkerScripts), arg0)
}

// ListWorkersSecrets mocks base method.
func (m *MockApi) ListWorkersSecrets(arg0 context.Context, arg1 string) (cloudflare.WorkersListSecretsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWorkersSecrets", arg0, arg1)
	ret0, _ := ret[0].(cloudflare.WorkersListSecretsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWorkersSecrets indicates an expected call of ListWorkersSecrets.
func (mr *MockApiMockRecorder) ListWorkersSecrets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkersSecrets", reflect.TypeOf((*MockApi)(nil).ListWorkersSecrets), arg0, arg1)
}

// ListZonesContext mocks base method.
func (m *MockApi) ListZonesContext(arg0 context.Context, arg1 ...cloudflare.ReqOption) (cloudflare.ZonesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListZonesContext", varargs...)
	ret0, _ := ret[0].(cloudflare.ZonesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListZonesContext indicates an expected call of ListZonesContext.
func (mr *MockApiMockRecorder) ListZonesContext(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListZonesContext", reflect.TypeOf((*MockApi)(nil).ListZonesContext), varargs...)
}

// ZoneLevelAccessGroups mocks base method.
func (m *MockApi) ZoneLevelAccessGroups(arg0 context.Context, arg1 string, arg2 cloudflare.PaginationOptions) ([]cloudflare.AccessGroup, cloudflare.ResultInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZoneLevelAccessGroups", arg0, arg1, arg2)
	ret0, _ := ret[0].([]cloudflare.AccessGroup)
	ret1, _ := ret[1].(cloudflare.ResultInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ZoneLevelAccessGroups indicates an expected call of ZoneLevelAccessGroups.
func (mr *MockApiMockRecorder) ZoneLevelAccessGroups(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZoneLevelAccessGroups", reflect.TypeOf((*MockApi)(nil).ZoneLevelAccessGroups), arg0, arg1, arg2)
}
