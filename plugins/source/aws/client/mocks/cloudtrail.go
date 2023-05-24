// Code generated by MockGen. DO NOT EDIT.
// Source: cloudtrail.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	
	reflect "reflect"

	cloudtrail "github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	gomock "github.com/golang/mock/gomock"
)

// MockCloudtrailClient is a mock of CloudtrailClient interface.
type MockCloudtrailClient struct {
	ctrl     *gomock.Controller
	recorder *MockCloudtrailClientMockRecorder
}

// MockCloudtrailClientMockRecorder is the mock recorder for MockCloudtrailClient.
type MockCloudtrailClientMockRecorder struct {
	mock *MockCloudtrailClient
}

// NewMockCloudtrailClient creates a new mock instance.
func NewMockCloudtrailClient(ctrl *gomock.Controller) *MockCloudtrailClient {
	mock := &MockCloudtrailClient{ctrl: ctrl}
	mock.recorder = &MockCloudtrailClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudtrailClient) EXPECT() *MockCloudtrailClientMockRecorder {
	return m.recorder
}

// DescribeQuery mocks base method.
func (m *MockCloudtrailClient) DescribeQuery(arg0 context.Context, arg1 *cloudtrail.DescribeQueryInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.DescribeQueryOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudtrail.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeQuery")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeQuery", varargs...)
	ret0, _ := ret[0].(*cloudtrail.DescribeQueryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeQuery indicates an expected call of DescribeQuery.
func (mr *MockCloudtrailClientMockRecorder) DescribeQuery(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeQuery", reflect.TypeOf((*MockCloudtrailClient)(nil).DescribeQuery), varargs...)
}

// DescribeTrails mocks base method.
func (m *MockCloudtrailClient) DescribeTrails(arg0 context.Context, arg1 *cloudtrail.DescribeTrailsInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.DescribeTrailsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudtrail.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeTrails")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTrails", varargs...)
	ret0, _ := ret[0].(*cloudtrail.DescribeTrailsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTrails indicates an expected call of DescribeTrails.
func (mr *MockCloudtrailClientMockRecorder) DescribeTrails(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTrails", reflect.TypeOf((*MockCloudtrailClient)(nil).DescribeTrails), varargs...)
}

// GetChannel mocks base method.
func (m *MockCloudtrailClient) GetChannel(arg0 context.Context, arg1 *cloudtrail.GetChannelInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.GetChannelOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudtrail.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetChannel")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetChannel", varargs...)
	ret0, _ := ret[0].(*cloudtrail.GetChannelOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChannel indicates an expected call of GetChannel.
func (mr *MockCloudtrailClientMockRecorder) GetChannel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChannel", reflect.TypeOf((*MockCloudtrailClient)(nil).GetChannel), varargs...)
}

// GetEventDataStore mocks base method.
func (m *MockCloudtrailClient) GetEventDataStore(arg0 context.Context, arg1 *cloudtrail.GetEventDataStoreInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.GetEventDataStoreOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudtrail.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetEventDataStore")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEventDataStore", varargs...)
	ret0, _ := ret[0].(*cloudtrail.GetEventDataStoreOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventDataStore indicates an expected call of GetEventDataStore.
func (mr *MockCloudtrailClientMockRecorder) GetEventDataStore(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventDataStore", reflect.TypeOf((*MockCloudtrailClient)(nil).GetEventDataStore), varargs...)
}

// GetEventSelectors mocks base method.
func (m *MockCloudtrailClient) GetEventSelectors(arg0 context.Context, arg1 *cloudtrail.GetEventSelectorsInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.GetEventSelectorsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudtrail.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetEventSelectors")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEventSelectors", varargs...)
	ret0, _ := ret[0].(*cloudtrail.GetEventSelectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventSelectors indicates an expected call of GetEventSelectors.
func (mr *MockCloudtrailClientMockRecorder) GetEventSelectors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventSelectors", reflect.TypeOf((*MockCloudtrailClient)(nil).GetEventSelectors), varargs...)
}

// GetImport mocks base method.
func (m *MockCloudtrailClient) GetImport(arg0 context.Context, arg1 *cloudtrail.GetImportInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.GetImportOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudtrail.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetImport")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetImport", varargs...)
	ret0, _ := ret[0].(*cloudtrail.GetImportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImport indicates an expected call of GetImport.
func (mr *MockCloudtrailClientMockRecorder) GetImport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImport", reflect.TypeOf((*MockCloudtrailClient)(nil).GetImport), varargs...)
}

// GetInsightSelectors mocks base method.
func (m *MockCloudtrailClient) GetInsightSelectors(arg0 context.Context, arg1 *cloudtrail.GetInsightSelectorsInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.GetInsightSelectorsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudtrail.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetInsightSelectors")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInsightSelectors", varargs...)
	ret0, _ := ret[0].(*cloudtrail.GetInsightSelectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInsightSelectors indicates an expected call of GetInsightSelectors.
func (mr *MockCloudtrailClientMockRecorder) GetInsightSelectors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInsightSelectors", reflect.TypeOf((*MockCloudtrailClient)(nil).GetInsightSelectors), varargs...)
}

// GetQueryResults mocks base method.
func (m *MockCloudtrailClient) GetQueryResults(arg0 context.Context, arg1 *cloudtrail.GetQueryResultsInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.GetQueryResultsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudtrail.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetQueryResults")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetQueryResults", varargs...)
	ret0, _ := ret[0].(*cloudtrail.GetQueryResultsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueryResults indicates an expected call of GetQueryResults.
func (mr *MockCloudtrailClientMockRecorder) GetQueryResults(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryResults", reflect.TypeOf((*MockCloudtrailClient)(nil).GetQueryResults), varargs...)
}

// GetTrail mocks base method.
func (m *MockCloudtrailClient) GetTrail(arg0 context.Context, arg1 *cloudtrail.GetTrailInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.GetTrailOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudtrail.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetTrail")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTrail", varargs...)
	ret0, _ := ret[0].(*cloudtrail.GetTrailOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrail indicates an expected call of GetTrail.
func (mr *MockCloudtrailClientMockRecorder) GetTrail(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrail", reflect.TypeOf((*MockCloudtrailClient)(nil).GetTrail), varargs...)
}

// GetTrailStatus mocks base method.
func (m *MockCloudtrailClient) GetTrailStatus(arg0 context.Context, arg1 *cloudtrail.GetTrailStatusInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.GetTrailStatusOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudtrail.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetTrailStatus")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTrailStatus", varargs...)
	ret0, _ := ret[0].(*cloudtrail.GetTrailStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrailStatus indicates an expected call of GetTrailStatus.
func (mr *MockCloudtrailClientMockRecorder) GetTrailStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrailStatus", reflect.TypeOf((*MockCloudtrailClient)(nil).GetTrailStatus), varargs...)
}

// ListChannels mocks base method.
func (m *MockCloudtrailClient) ListChannels(arg0 context.Context, arg1 *cloudtrail.ListChannelsInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.ListChannelsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudtrail.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListChannels")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListChannels", varargs...)
	ret0, _ := ret[0].(*cloudtrail.ListChannelsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListChannels indicates an expected call of ListChannels.
func (mr *MockCloudtrailClientMockRecorder) ListChannels(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChannels", reflect.TypeOf((*MockCloudtrailClient)(nil).ListChannels), varargs...)
}

// ListEventDataStores mocks base method.
func (m *MockCloudtrailClient) ListEventDataStores(arg0 context.Context, arg1 *cloudtrail.ListEventDataStoresInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.ListEventDataStoresOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudtrail.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListEventDataStores")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEventDataStores", varargs...)
	ret0, _ := ret[0].(*cloudtrail.ListEventDataStoresOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEventDataStores indicates an expected call of ListEventDataStores.
func (mr *MockCloudtrailClientMockRecorder) ListEventDataStores(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEventDataStores", reflect.TypeOf((*MockCloudtrailClient)(nil).ListEventDataStores), varargs...)
}

// ListImportFailures mocks base method.
func (m *MockCloudtrailClient) ListImportFailures(arg0 context.Context, arg1 *cloudtrail.ListImportFailuresInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.ListImportFailuresOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudtrail.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListImportFailures")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListImportFailures", varargs...)
	ret0, _ := ret[0].(*cloudtrail.ListImportFailuresOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImportFailures indicates an expected call of ListImportFailures.
func (mr *MockCloudtrailClientMockRecorder) ListImportFailures(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImportFailures", reflect.TypeOf((*MockCloudtrailClient)(nil).ListImportFailures), varargs...)
}

// ListImports mocks base method.
func (m *MockCloudtrailClient) ListImports(arg0 context.Context, arg1 *cloudtrail.ListImportsInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.ListImportsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudtrail.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListImports")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListImports", varargs...)
	ret0, _ := ret[0].(*cloudtrail.ListImportsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImports indicates an expected call of ListImports.
func (mr *MockCloudtrailClientMockRecorder) ListImports(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImports", reflect.TypeOf((*MockCloudtrailClient)(nil).ListImports), varargs...)
}

// ListPublicKeys mocks base method.
func (m *MockCloudtrailClient) ListPublicKeys(arg0 context.Context, arg1 *cloudtrail.ListPublicKeysInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.ListPublicKeysOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudtrail.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListPublicKeys")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPublicKeys", varargs...)
	ret0, _ := ret[0].(*cloudtrail.ListPublicKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPublicKeys indicates an expected call of ListPublicKeys.
func (mr *MockCloudtrailClientMockRecorder) ListPublicKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPublicKeys", reflect.TypeOf((*MockCloudtrailClient)(nil).ListPublicKeys), varargs...)
}

// ListQueries mocks base method.
func (m *MockCloudtrailClient) ListQueries(arg0 context.Context, arg1 *cloudtrail.ListQueriesInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.ListQueriesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudtrail.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListQueries")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListQueries", varargs...)
	ret0, _ := ret[0].(*cloudtrail.ListQueriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListQueries indicates an expected call of ListQueries.
func (mr *MockCloudtrailClientMockRecorder) ListQueries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListQueries", reflect.TypeOf((*MockCloudtrailClient)(nil).ListQueries), varargs...)
}

// ListTags mocks base method.
func (m *MockCloudtrailClient) ListTags(arg0 context.Context, arg1 *cloudtrail.ListTagsInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.ListTagsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudtrail.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListTags")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTags", varargs...)
	ret0, _ := ret[0].(*cloudtrail.ListTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTags indicates an expected call of ListTags.
func (mr *MockCloudtrailClientMockRecorder) ListTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockCloudtrailClient)(nil).ListTags), varargs...)
}

// ListTrails mocks base method.
func (m *MockCloudtrailClient) ListTrails(arg0 context.Context, arg1 *cloudtrail.ListTrailsInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.ListTrailsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudtrail.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListTrails")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTrails", varargs...)
	ret0, _ := ret[0].(*cloudtrail.ListTrailsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTrails indicates an expected call of ListTrails.
func (mr *MockCloudtrailClientMockRecorder) ListTrails(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTrails", reflect.TypeOf((*MockCloudtrailClient)(nil).ListTrails), varargs...)
}

// LookupEvents mocks base method.
func (m *MockCloudtrailClient) LookupEvents(arg0 context.Context, arg1 *cloudtrail.LookupEventsInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.LookupEventsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudtrail.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to LookupEvents")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LookupEvents", varargs...)
	ret0, _ := ret[0].(*cloudtrail.LookupEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookupEvents indicates an expected call of LookupEvents.
func (mr *MockCloudtrailClientMockRecorder) LookupEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupEvents", reflect.TypeOf((*MockCloudtrailClient)(nil).LookupEvents), varargs...)
}
