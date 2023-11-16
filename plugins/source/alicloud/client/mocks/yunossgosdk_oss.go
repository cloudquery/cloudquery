// Code generated by MockGen. DO NOT EDIT.
// Source: yunossgosdk_oss.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	oss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	gomock "github.com/golang/mock/gomock"
)

// MockOssClient is a mock of OssClient interface.
type MockOssClient struct {
	ctrl     *gomock.Controller
	recorder *MockOssClientMockRecorder
}

// MockOssClientMockRecorder is the mock recorder for MockOssClient.
type MockOssClientMockRecorder struct {
	mock *MockOssClient
}

// NewMockOssClient creates a new mock instance.
func NewMockOssClient(ctrl *gomock.Controller) *MockOssClient {
	mock := &MockOssClient{ctrl: ctrl}
	mock.recorder = &MockOssClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOssClient) EXPECT() *MockOssClientMockRecorder {
	return m.recorder
}

// GetBucketACL mocks base method.
func (m *MockOssClient) GetBucketACL(arg0 string, arg1 ...oss.Option) (oss.GetBucketACLResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketACL", varargs...)
	ret0, _ := ret[0].(oss.GetBucketACLResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketACL indicates an expected call of GetBucketACL.
func (mr *MockOssClientMockRecorder) GetBucketACL(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketACL", reflect.TypeOf((*MockOssClient)(nil).GetBucketACL), varargs...)
}

// GetBucketAccessMonitor mocks base method.
func (m *MockOssClient) GetBucketAccessMonitor(arg0 string, arg1 ...oss.Option) (oss.GetBucketAccessMonitorResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketAccessMonitor", varargs...)
	ret0, _ := ret[0].(oss.GetBucketAccessMonitorResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketAccessMonitor indicates an expected call of GetBucketAccessMonitor.
func (mr *MockOssClientMockRecorder) GetBucketAccessMonitor(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketAccessMonitor", reflect.TypeOf((*MockOssClient)(nil).GetBucketAccessMonitor), varargs...)
}

// GetBucketAccessMonitorXml mocks base method.
func (m *MockOssClient) GetBucketAccessMonitorXml(arg0 string, arg1 ...oss.Option) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketAccessMonitorXml", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketAccessMonitorXml indicates an expected call of GetBucketAccessMonitorXml.
func (mr *MockOssClientMockRecorder) GetBucketAccessMonitorXml(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketAccessMonitorXml", reflect.TypeOf((*MockOssClient)(nil).GetBucketAccessMonitorXml), varargs...)
}

// GetBucketAsyncTask mocks base method.
func (m *MockOssClient) GetBucketAsyncTask(arg0, arg1 string, arg2 ...oss.Option) (oss.AsynFetchTaskInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketAsyncTask", varargs...)
	ret0, _ := ret[0].(oss.AsynFetchTaskInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketAsyncTask indicates an expected call of GetBucketAsyncTask.
func (mr *MockOssClientMockRecorder) GetBucketAsyncTask(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketAsyncTask", reflect.TypeOf((*MockOssClient)(nil).GetBucketAsyncTask), varargs...)
}

// GetBucketCORS mocks base method.
func (m *MockOssClient) GetBucketCORS(arg0 string, arg1 ...oss.Option) (oss.GetBucketCORSResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketCORS", varargs...)
	ret0, _ := ret[0].(oss.GetBucketCORSResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketCORS indicates an expected call of GetBucketCORS.
func (mr *MockOssClientMockRecorder) GetBucketCORS(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketCORS", reflect.TypeOf((*MockOssClient)(nil).GetBucketCORS), varargs...)
}

// GetBucketCORSXml mocks base method.
func (m *MockOssClient) GetBucketCORSXml(arg0 string, arg1 ...oss.Option) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketCORSXml", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketCORSXml indicates an expected call of GetBucketCORSXml.
func (mr *MockOssClientMockRecorder) GetBucketCORSXml(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketCORSXml", reflect.TypeOf((*MockOssClient)(nil).GetBucketCORSXml), varargs...)
}

// GetBucketCname mocks base method.
func (m *MockOssClient) GetBucketCname(arg0 string, arg1 ...oss.Option) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketCname", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketCname indicates an expected call of GetBucketCname.
func (mr *MockOssClientMockRecorder) GetBucketCname(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketCname", reflect.TypeOf((*MockOssClient)(nil).GetBucketCname), varargs...)
}

// GetBucketCnameToken mocks base method.
func (m *MockOssClient) GetBucketCnameToken(arg0, arg1 string, arg2 ...oss.Option) (oss.GetBucketCnameTokenResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketCnameToken", varargs...)
	ret0, _ := ret[0].(oss.GetBucketCnameTokenResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketCnameToken indicates an expected call of GetBucketCnameToken.
func (mr *MockOssClientMockRecorder) GetBucketCnameToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketCnameToken", reflect.TypeOf((*MockOssClient)(nil).GetBucketCnameToken), varargs...)
}

// GetBucketEncryption mocks base method.
func (m *MockOssClient) GetBucketEncryption(arg0 string, arg1 ...oss.Option) (oss.GetBucketEncryptionResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketEncryption", varargs...)
	ret0, _ := ret[0].(oss.GetBucketEncryptionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketEncryption indicates an expected call of GetBucketEncryption.
func (mr *MockOssClientMockRecorder) GetBucketEncryption(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketEncryption", reflect.TypeOf((*MockOssClient)(nil).GetBucketEncryption), varargs...)
}

// GetBucketInfo mocks base method.
func (m *MockOssClient) GetBucketInfo(arg0 string, arg1 ...oss.Option) (oss.GetBucketInfoResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketInfo", varargs...)
	ret0, _ := ret[0].(oss.GetBucketInfoResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketInfo indicates an expected call of GetBucketInfo.
func (mr *MockOssClientMockRecorder) GetBucketInfo(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketInfo", reflect.TypeOf((*MockOssClient)(nil).GetBucketInfo), varargs...)
}

// GetBucketInventory mocks base method.
func (m *MockOssClient) GetBucketInventory(arg0, arg1 string, arg2 ...oss.Option) (oss.InventoryConfiguration, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketInventory", varargs...)
	ret0, _ := ret[0].(oss.InventoryConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketInventory indicates an expected call of GetBucketInventory.
func (mr *MockOssClientMockRecorder) GetBucketInventory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketInventory", reflect.TypeOf((*MockOssClient)(nil).GetBucketInventory), varargs...)
}

// GetBucketInventoryXml mocks base method.
func (m *MockOssClient) GetBucketInventoryXml(arg0, arg1 string, arg2 ...oss.Option) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketInventoryXml", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketInventoryXml indicates an expected call of GetBucketInventoryXml.
func (mr *MockOssClientMockRecorder) GetBucketInventoryXml(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketInventoryXml", reflect.TypeOf((*MockOssClient)(nil).GetBucketInventoryXml), varargs...)
}

// GetBucketLifecycle mocks base method.
func (m *MockOssClient) GetBucketLifecycle(arg0 string, arg1 ...oss.Option) (oss.GetBucketLifecycleResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketLifecycle", varargs...)
	ret0, _ := ret[0].(oss.GetBucketLifecycleResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketLifecycle indicates an expected call of GetBucketLifecycle.
func (mr *MockOssClientMockRecorder) GetBucketLifecycle(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketLifecycle", reflect.TypeOf((*MockOssClient)(nil).GetBucketLifecycle), varargs...)
}

// GetBucketLifecycleXml mocks base method.
func (m *MockOssClient) GetBucketLifecycleXml(arg0 string, arg1 ...oss.Option) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketLifecycleXml", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketLifecycleXml indicates an expected call of GetBucketLifecycleXml.
func (mr *MockOssClientMockRecorder) GetBucketLifecycleXml(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketLifecycleXml", reflect.TypeOf((*MockOssClient)(nil).GetBucketLifecycleXml), varargs...)
}

// GetBucketLocation mocks base method.
func (m *MockOssClient) GetBucketLocation(arg0 string, arg1 ...oss.Option) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketLocation", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketLocation indicates an expected call of GetBucketLocation.
func (mr *MockOssClientMockRecorder) GetBucketLocation(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketLocation", reflect.TypeOf((*MockOssClient)(nil).GetBucketLocation), varargs...)
}

// GetBucketLogging mocks base method.
func (m *MockOssClient) GetBucketLogging(arg0 string, arg1 ...oss.Option) (oss.GetBucketLoggingResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketLogging", varargs...)
	ret0, _ := ret[0].(oss.GetBucketLoggingResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketLogging indicates an expected call of GetBucketLogging.
func (mr *MockOssClientMockRecorder) GetBucketLogging(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketLogging", reflect.TypeOf((*MockOssClient)(nil).GetBucketLogging), varargs...)
}

// GetBucketPolicy mocks base method.
func (m *MockOssClient) GetBucketPolicy(arg0 string, arg1 ...oss.Option) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketPolicy", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketPolicy indicates an expected call of GetBucketPolicy.
func (mr *MockOssClientMockRecorder) GetBucketPolicy(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketPolicy", reflect.TypeOf((*MockOssClient)(nil).GetBucketPolicy), varargs...)
}

// GetBucketQosInfo mocks base method.
func (m *MockOssClient) GetBucketQosInfo(arg0 string, arg1 ...oss.Option) (oss.BucketQoSConfiguration, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketQosInfo", varargs...)
	ret0, _ := ret[0].(oss.BucketQoSConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketQosInfo indicates an expected call of GetBucketQosInfo.
func (mr *MockOssClientMockRecorder) GetBucketQosInfo(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketQosInfo", reflect.TypeOf((*MockOssClient)(nil).GetBucketQosInfo), varargs...)
}

// GetBucketReferer mocks base method.
func (m *MockOssClient) GetBucketReferer(arg0 string, arg1 ...oss.Option) (oss.GetBucketRefererResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketReferer", varargs...)
	ret0, _ := ret[0].(oss.GetBucketRefererResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketReferer indicates an expected call of GetBucketReferer.
func (mr *MockOssClientMockRecorder) GetBucketReferer(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketReferer", reflect.TypeOf((*MockOssClient)(nil).GetBucketReferer), varargs...)
}

// GetBucketReplication mocks base method.
func (m *MockOssClient) GetBucketReplication(arg0 string, arg1 ...oss.Option) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketReplication", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketReplication indicates an expected call of GetBucketReplication.
func (mr *MockOssClientMockRecorder) GetBucketReplication(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketReplication", reflect.TypeOf((*MockOssClient)(nil).GetBucketReplication), varargs...)
}

// GetBucketReplicationLocation mocks base method.
func (m *MockOssClient) GetBucketReplicationLocation(arg0 string, arg1 ...oss.Option) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketReplicationLocation", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketReplicationLocation indicates an expected call of GetBucketReplicationLocation.
func (mr *MockOssClientMockRecorder) GetBucketReplicationLocation(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketReplicationLocation", reflect.TypeOf((*MockOssClient)(nil).GetBucketReplicationLocation), varargs...)
}

// GetBucketReplicationProgress mocks base method.
func (m *MockOssClient) GetBucketReplicationProgress(arg0, arg1 string, arg2 ...oss.Option) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketReplicationProgress", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketReplicationProgress indicates an expected call of GetBucketReplicationProgress.
func (mr *MockOssClientMockRecorder) GetBucketReplicationProgress(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketReplicationProgress", reflect.TypeOf((*MockOssClient)(nil).GetBucketReplicationProgress), varargs...)
}

// GetBucketRequestPayment mocks base method.
func (m *MockOssClient) GetBucketRequestPayment(arg0 string, arg1 ...oss.Option) (oss.RequestPaymentConfiguration, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketRequestPayment", varargs...)
	ret0, _ := ret[0].(oss.RequestPaymentConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketRequestPayment indicates an expected call of GetBucketRequestPayment.
func (mr *MockOssClientMockRecorder) GetBucketRequestPayment(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketRequestPayment", reflect.TypeOf((*MockOssClient)(nil).GetBucketRequestPayment), varargs...)
}

// GetBucketResourceGroup mocks base method.
func (m *MockOssClient) GetBucketResourceGroup(arg0 string, arg1 ...oss.Option) (oss.GetBucketResourceGroupResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketResourceGroup", varargs...)
	ret0, _ := ret[0].(oss.GetBucketResourceGroupResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketResourceGroup indicates an expected call of GetBucketResourceGroup.
func (mr *MockOssClientMockRecorder) GetBucketResourceGroup(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketResourceGroup", reflect.TypeOf((*MockOssClient)(nil).GetBucketResourceGroup), varargs...)
}

// GetBucketResourceGroupXml mocks base method.
func (m *MockOssClient) GetBucketResourceGroupXml(arg0 string, arg1 ...oss.Option) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketResourceGroupXml", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketResourceGroupXml indicates an expected call of GetBucketResourceGroupXml.
func (mr *MockOssClientMockRecorder) GetBucketResourceGroupXml(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketResourceGroupXml", reflect.TypeOf((*MockOssClient)(nil).GetBucketResourceGroupXml), varargs...)
}

// GetBucketStat mocks base method.
func (m *MockOssClient) GetBucketStat(arg0 string, arg1 ...oss.Option) (oss.GetBucketStatResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketStat", varargs...)
	ret0, _ := ret[0].(oss.GetBucketStatResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketStat indicates an expected call of GetBucketStat.
func (mr *MockOssClientMockRecorder) GetBucketStat(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketStat", reflect.TypeOf((*MockOssClient)(nil).GetBucketStat), varargs...)
}

// GetBucketStyle mocks base method.
func (m *MockOssClient) GetBucketStyle(arg0, arg1 string, arg2 ...oss.Option) (oss.GetBucketStyleResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketStyle", varargs...)
	ret0, _ := ret[0].(oss.GetBucketStyleResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketStyle indicates an expected call of GetBucketStyle.
func (mr *MockOssClientMockRecorder) GetBucketStyle(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketStyle", reflect.TypeOf((*MockOssClient)(nil).GetBucketStyle), varargs...)
}

// GetBucketStyleXml mocks base method.
func (m *MockOssClient) GetBucketStyleXml(arg0, arg1 string, arg2 ...oss.Option) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketStyleXml", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketStyleXml indicates an expected call of GetBucketStyleXml.
func (mr *MockOssClientMockRecorder) GetBucketStyleXml(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketStyleXml", reflect.TypeOf((*MockOssClient)(nil).GetBucketStyleXml), varargs...)
}

// GetBucketTagging mocks base method.
func (m *MockOssClient) GetBucketTagging(arg0 string, arg1 ...oss.Option) (oss.GetBucketTaggingResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketTagging", varargs...)
	ret0, _ := ret[0].(oss.GetBucketTaggingResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketTagging indicates an expected call of GetBucketTagging.
func (mr *MockOssClientMockRecorder) GetBucketTagging(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketTagging", reflect.TypeOf((*MockOssClient)(nil).GetBucketTagging), varargs...)
}

// GetBucketTransferAcc mocks base method.
func (m *MockOssClient) GetBucketTransferAcc(arg0 string, arg1 ...oss.Option) (oss.TransferAccConfiguration, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketTransferAcc", varargs...)
	ret0, _ := ret[0].(oss.TransferAccConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketTransferAcc indicates an expected call of GetBucketTransferAcc.
func (mr *MockOssClientMockRecorder) GetBucketTransferAcc(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketTransferAcc", reflect.TypeOf((*MockOssClient)(nil).GetBucketTransferAcc), varargs...)
}

// GetBucketVersioning mocks base method.
func (m *MockOssClient) GetBucketVersioning(arg0 string, arg1 ...oss.Option) (oss.GetBucketVersioningResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketVersioning", varargs...)
	ret0, _ := ret[0].(oss.GetBucketVersioningResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketVersioning indicates an expected call of GetBucketVersioning.
func (mr *MockOssClientMockRecorder) GetBucketVersioning(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketVersioning", reflect.TypeOf((*MockOssClient)(nil).GetBucketVersioning), varargs...)
}

// GetBucketWebsite mocks base method.
func (m *MockOssClient) GetBucketWebsite(arg0 string, arg1 ...oss.Option) (oss.GetBucketWebsiteResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketWebsite", varargs...)
	ret0, _ := ret[0].(oss.GetBucketWebsiteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketWebsite indicates an expected call of GetBucketWebsite.
func (mr *MockOssClientMockRecorder) GetBucketWebsite(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketWebsite", reflect.TypeOf((*MockOssClient)(nil).GetBucketWebsite), varargs...)
}

// GetBucketWebsiteXml mocks base method.
func (m *MockOssClient) GetBucketWebsiteXml(arg0 string, arg1 ...oss.Option) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketWebsiteXml", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketWebsiteXml indicates an expected call of GetBucketWebsiteXml.
func (mr *MockOssClientMockRecorder) GetBucketWebsiteXml(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketWebsiteXml", reflect.TypeOf((*MockOssClient)(nil).GetBucketWebsiteXml), varargs...)
}

// GetBucketWorm mocks base method.
func (m *MockOssClient) GetBucketWorm(arg0 string, arg1 ...oss.Option) (oss.WormConfiguration, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketWorm", varargs...)
	ret0, _ := ret[0].(oss.WormConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketWorm indicates an expected call of GetBucketWorm.
func (mr *MockOssClientMockRecorder) GetBucketWorm(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketWorm", reflect.TypeOf((*MockOssClient)(nil).GetBucketWorm), varargs...)
}

// GetMetaQueryStatus mocks base method.
func (m *MockOssClient) GetMetaQueryStatus(arg0 string, arg1 ...oss.Option) (oss.GetMetaQueryStatusResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMetaQueryStatus", varargs...)
	ret0, _ := ret[0].(oss.GetMetaQueryStatusResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetaQueryStatus indicates an expected call of GetMetaQueryStatus.
func (mr *MockOssClientMockRecorder) GetMetaQueryStatus(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetaQueryStatus", reflect.TypeOf((*MockOssClient)(nil).GetMetaQueryStatus), varargs...)
}

// GetUserQoSInfo mocks base method.
func (m *MockOssClient) GetUserQoSInfo(arg0 ...oss.Option) (oss.UserQoSConfiguration, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserQoSInfo", varargs...)
	ret0, _ := ret[0].(oss.UserQoSConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserQoSInfo indicates an expected call of GetUserQoSInfo.
func (mr *MockOssClientMockRecorder) GetUserQoSInfo(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserQoSInfo", reflect.TypeOf((*MockOssClient)(nil).GetUserQoSInfo), arg0...)
}

// ListBucketCname mocks base method.
func (m *MockOssClient) ListBucketCname(arg0 string, arg1 ...oss.Option) (oss.ListBucketCnameResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBucketCname", varargs...)
	ret0, _ := ret[0].(oss.ListBucketCnameResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBucketCname indicates an expected call of ListBucketCname.
func (mr *MockOssClientMockRecorder) ListBucketCname(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBucketCname", reflect.TypeOf((*MockOssClient)(nil).ListBucketCname), varargs...)
}

// ListBucketInventory mocks base method.
func (m *MockOssClient) ListBucketInventory(arg0, arg1 string, arg2 ...oss.Option) (oss.ListInventoryConfigurationsResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBucketInventory", varargs...)
	ret0, _ := ret[0].(oss.ListInventoryConfigurationsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBucketInventory indicates an expected call of ListBucketInventory.
func (mr *MockOssClientMockRecorder) ListBucketInventory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBucketInventory", reflect.TypeOf((*MockOssClient)(nil).ListBucketInventory), varargs...)
}

// ListBucketInventoryXml mocks base method.
func (m *MockOssClient) ListBucketInventoryXml(arg0, arg1 string, arg2 ...oss.Option) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBucketInventoryXml", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBucketInventoryXml indicates an expected call of ListBucketInventoryXml.
func (mr *MockOssClientMockRecorder) ListBucketInventoryXml(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBucketInventoryXml", reflect.TypeOf((*MockOssClient)(nil).ListBucketInventoryXml), varargs...)
}

// ListBucketStyle mocks base method.
func (m *MockOssClient) ListBucketStyle(arg0 string, arg1 ...oss.Option) (oss.GetBucketListStyleResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBucketStyle", varargs...)
	ret0, _ := ret[0].(oss.GetBucketListStyleResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBucketStyle indicates an expected call of ListBucketStyle.
func (mr *MockOssClientMockRecorder) ListBucketStyle(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBucketStyle", reflect.TypeOf((*MockOssClient)(nil).ListBucketStyle), varargs...)
}

// ListBucketStyleXml mocks base method.
func (m *MockOssClient) ListBucketStyleXml(arg0 string, arg1 ...oss.Option) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBucketStyleXml", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBucketStyleXml indicates an expected call of ListBucketStyleXml.
func (mr *MockOssClientMockRecorder) ListBucketStyleXml(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBucketStyleXml", reflect.TypeOf((*MockOssClient)(nil).ListBucketStyleXml), varargs...)
}

// ListBuckets mocks base method.
func (m *MockOssClient) ListBuckets(arg0 ...oss.Option) (oss.ListBucketsResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBuckets", varargs...)
	ret0, _ := ret[0].(oss.ListBucketsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBuckets indicates an expected call of ListBuckets.
func (mr *MockOssClientMockRecorder) ListBuckets(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBuckets", reflect.TypeOf((*MockOssClient)(nil).ListBuckets), arg0...)
}

// ListCloudBoxes mocks base method.
func (m *MockOssClient) ListCloudBoxes(arg0 ...oss.Option) (oss.ListCloudBoxResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCloudBoxes", varargs...)
	ret0, _ := ret[0].(oss.ListCloudBoxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCloudBoxes indicates an expected call of ListCloudBoxes.
func (mr *MockOssClientMockRecorder) ListCloudBoxes(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCloudBoxes", reflect.TypeOf((*MockOssClient)(nil).ListCloudBoxes), arg0...)
}
