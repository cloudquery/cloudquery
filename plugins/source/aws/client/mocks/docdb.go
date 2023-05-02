// Code generated by MockGen. DO NOT EDIT.
// Source: docdb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	docdb "github.com/aws/aws-sdk-go-v2/service/docdb"
	gomock "github.com/golang/mock/gomock"
)

// MockDocdbClient is a mock of DocdbClient interface.
type MockDocdbClient struct {
	ctrl     *gomock.Controller
	recorder *MockDocdbClientMockRecorder
}

// MockDocdbClientMockRecorder is the mock recorder for MockDocdbClient.
type MockDocdbClientMockRecorder struct {
	mock *MockDocdbClient
}

// NewMockDocdbClient creates a new mock instance.
func NewMockDocdbClient(ctrl *gomock.Controller) *MockDocdbClient {
	mock := &MockDocdbClient{ctrl: ctrl}
	mock.recorder = &MockDocdbClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDocdbClient) EXPECT() *MockDocdbClientMockRecorder {
	return m.recorder
}

// DescribeCertificates mocks base method.
func (m *MockDocdbClient) DescribeCertificates(arg0 context.Context, arg1 *docdb.DescribeCertificatesInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeCertificatesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &docdb.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeCertificates")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCertificates", varargs...)
	ret0, _ := ret[0].(*docdb.DescribeCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCertificates indicates an expected call of DescribeCertificates.
func (mr *MockDocdbClientMockRecorder) DescribeCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCertificates", reflect.TypeOf((*MockDocdbClient)(nil).DescribeCertificates), varargs...)
}

// DescribeDBClusterParameterGroups mocks base method.
func (m *MockDocdbClient) DescribeDBClusterParameterGroups(arg0 context.Context, arg1 *docdb.DescribeDBClusterParameterGroupsInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeDBClusterParameterGroupsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &docdb.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeDBClusterParameterGroups")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBClusterParameterGroups", varargs...)
	ret0, _ := ret[0].(*docdb.DescribeDBClusterParameterGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDBClusterParameterGroups indicates an expected call of DescribeDBClusterParameterGroups.
func (mr *MockDocdbClientMockRecorder) DescribeDBClusterParameterGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBClusterParameterGroups", reflect.TypeOf((*MockDocdbClient)(nil).DescribeDBClusterParameterGroups), varargs...)
}

// DescribeDBClusterParameters mocks base method.
func (m *MockDocdbClient) DescribeDBClusterParameters(arg0 context.Context, arg1 *docdb.DescribeDBClusterParametersInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeDBClusterParametersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &docdb.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeDBClusterParameters")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBClusterParameters", varargs...)
	ret0, _ := ret[0].(*docdb.DescribeDBClusterParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDBClusterParameters indicates an expected call of DescribeDBClusterParameters.
func (mr *MockDocdbClientMockRecorder) DescribeDBClusterParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBClusterParameters", reflect.TypeOf((*MockDocdbClient)(nil).DescribeDBClusterParameters), varargs...)
}

// DescribeDBClusterSnapshotAttributes mocks base method.
func (m *MockDocdbClient) DescribeDBClusterSnapshotAttributes(arg0 context.Context, arg1 *docdb.DescribeDBClusterSnapshotAttributesInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeDBClusterSnapshotAttributesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &docdb.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeDBClusterSnapshotAttributes")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBClusterSnapshotAttributes", varargs...)
	ret0, _ := ret[0].(*docdb.DescribeDBClusterSnapshotAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDBClusterSnapshotAttributes indicates an expected call of DescribeDBClusterSnapshotAttributes.
func (mr *MockDocdbClientMockRecorder) DescribeDBClusterSnapshotAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBClusterSnapshotAttributes", reflect.TypeOf((*MockDocdbClient)(nil).DescribeDBClusterSnapshotAttributes), varargs...)
}

// DescribeDBClusterSnapshots mocks base method.
func (m *MockDocdbClient) DescribeDBClusterSnapshots(arg0 context.Context, arg1 *docdb.DescribeDBClusterSnapshotsInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeDBClusterSnapshotsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &docdb.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeDBClusterSnapshots")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBClusterSnapshots", varargs...)
	ret0, _ := ret[0].(*docdb.DescribeDBClusterSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDBClusterSnapshots indicates an expected call of DescribeDBClusterSnapshots.
func (mr *MockDocdbClientMockRecorder) DescribeDBClusterSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBClusterSnapshots", reflect.TypeOf((*MockDocdbClient)(nil).DescribeDBClusterSnapshots), varargs...)
}

// DescribeDBClusters mocks base method.
func (m *MockDocdbClient) DescribeDBClusters(arg0 context.Context, arg1 *docdb.DescribeDBClustersInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeDBClustersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &docdb.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeDBClusters")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBClusters", varargs...)
	ret0, _ := ret[0].(*docdb.DescribeDBClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDBClusters indicates an expected call of DescribeDBClusters.
func (mr *MockDocdbClientMockRecorder) DescribeDBClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBClusters", reflect.TypeOf((*MockDocdbClient)(nil).DescribeDBClusters), varargs...)
}

// DescribeDBEngineVersions mocks base method.
func (m *MockDocdbClient) DescribeDBEngineVersions(arg0 context.Context, arg1 *docdb.DescribeDBEngineVersionsInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeDBEngineVersionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &docdb.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeDBEngineVersions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBEngineVersions", varargs...)
	ret0, _ := ret[0].(*docdb.DescribeDBEngineVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDBEngineVersions indicates an expected call of DescribeDBEngineVersions.
func (mr *MockDocdbClientMockRecorder) DescribeDBEngineVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBEngineVersions", reflect.TypeOf((*MockDocdbClient)(nil).DescribeDBEngineVersions), varargs...)
}

// DescribeDBInstances mocks base method.
func (m *MockDocdbClient) DescribeDBInstances(arg0 context.Context, arg1 *docdb.DescribeDBInstancesInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeDBInstancesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &docdb.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeDBInstances")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBInstances", varargs...)
	ret0, _ := ret[0].(*docdb.DescribeDBInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDBInstances indicates an expected call of DescribeDBInstances.
func (mr *MockDocdbClientMockRecorder) DescribeDBInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBInstances", reflect.TypeOf((*MockDocdbClient)(nil).DescribeDBInstances), varargs...)
}

// DescribeDBSubnetGroups mocks base method.
func (m *MockDocdbClient) DescribeDBSubnetGroups(arg0 context.Context, arg1 *docdb.DescribeDBSubnetGroupsInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeDBSubnetGroupsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &docdb.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeDBSubnetGroups")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBSubnetGroups", varargs...)
	ret0, _ := ret[0].(*docdb.DescribeDBSubnetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDBSubnetGroups indicates an expected call of DescribeDBSubnetGroups.
func (mr *MockDocdbClientMockRecorder) DescribeDBSubnetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBSubnetGroups", reflect.TypeOf((*MockDocdbClient)(nil).DescribeDBSubnetGroups), varargs...)
}

// DescribeEngineDefaultClusterParameters mocks base method.
func (m *MockDocdbClient) DescribeEngineDefaultClusterParameters(arg0 context.Context, arg1 *docdb.DescribeEngineDefaultClusterParametersInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeEngineDefaultClusterParametersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &docdb.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeEngineDefaultClusterParameters")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEngineDefaultClusterParameters", varargs...)
	ret0, _ := ret[0].(*docdb.DescribeEngineDefaultClusterParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEngineDefaultClusterParameters indicates an expected call of DescribeEngineDefaultClusterParameters.
func (mr *MockDocdbClientMockRecorder) DescribeEngineDefaultClusterParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEngineDefaultClusterParameters", reflect.TypeOf((*MockDocdbClient)(nil).DescribeEngineDefaultClusterParameters), varargs...)
}

// DescribeEventCategories mocks base method.
func (m *MockDocdbClient) DescribeEventCategories(arg0 context.Context, arg1 *docdb.DescribeEventCategoriesInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeEventCategoriesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &docdb.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeEventCategories")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEventCategories", varargs...)
	ret0, _ := ret[0].(*docdb.DescribeEventCategoriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEventCategories indicates an expected call of DescribeEventCategories.
func (mr *MockDocdbClientMockRecorder) DescribeEventCategories(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEventCategories", reflect.TypeOf((*MockDocdbClient)(nil).DescribeEventCategories), varargs...)
}

// DescribeEventSubscriptions mocks base method.
func (m *MockDocdbClient) DescribeEventSubscriptions(arg0 context.Context, arg1 *docdb.DescribeEventSubscriptionsInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeEventSubscriptionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &docdb.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeEventSubscriptions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEventSubscriptions", varargs...)
	ret0, _ := ret[0].(*docdb.DescribeEventSubscriptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEventSubscriptions indicates an expected call of DescribeEventSubscriptions.
func (mr *MockDocdbClientMockRecorder) DescribeEventSubscriptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEventSubscriptions", reflect.TypeOf((*MockDocdbClient)(nil).DescribeEventSubscriptions), varargs...)
}

// DescribeEvents mocks base method.
func (m *MockDocdbClient) DescribeEvents(arg0 context.Context, arg1 *docdb.DescribeEventsInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeEventsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &docdb.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeEvents")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEvents", varargs...)
	ret0, _ := ret[0].(*docdb.DescribeEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEvents indicates an expected call of DescribeEvents.
func (mr *MockDocdbClientMockRecorder) DescribeEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEvents", reflect.TypeOf((*MockDocdbClient)(nil).DescribeEvents), varargs...)
}

// DescribeGlobalClusters mocks base method.
func (m *MockDocdbClient) DescribeGlobalClusters(arg0 context.Context, arg1 *docdb.DescribeGlobalClustersInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeGlobalClustersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &docdb.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeGlobalClusters")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeGlobalClusters", varargs...)
	ret0, _ := ret[0].(*docdb.DescribeGlobalClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeGlobalClusters indicates an expected call of DescribeGlobalClusters.
func (mr *MockDocdbClientMockRecorder) DescribeGlobalClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeGlobalClusters", reflect.TypeOf((*MockDocdbClient)(nil).DescribeGlobalClusters), varargs...)
}

// DescribeOrderableDBInstanceOptions mocks base method.
func (m *MockDocdbClient) DescribeOrderableDBInstanceOptions(arg0 context.Context, arg1 *docdb.DescribeOrderableDBInstanceOptionsInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeOrderableDBInstanceOptionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &docdb.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeOrderableDBInstanceOptions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeOrderableDBInstanceOptions", varargs...)
	ret0, _ := ret[0].(*docdb.DescribeOrderableDBInstanceOptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeOrderableDBInstanceOptions indicates an expected call of DescribeOrderableDBInstanceOptions.
func (mr *MockDocdbClientMockRecorder) DescribeOrderableDBInstanceOptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeOrderableDBInstanceOptions", reflect.TypeOf((*MockDocdbClient)(nil).DescribeOrderableDBInstanceOptions), varargs...)
}

// DescribePendingMaintenanceActions mocks base method.
func (m *MockDocdbClient) DescribePendingMaintenanceActions(arg0 context.Context, arg1 *docdb.DescribePendingMaintenanceActionsInput, arg2 ...func(*docdb.Options)) (*docdb.DescribePendingMaintenanceActionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &docdb.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribePendingMaintenanceActions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribePendingMaintenanceActions", varargs...)
	ret0, _ := ret[0].(*docdb.DescribePendingMaintenanceActionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribePendingMaintenanceActions indicates an expected call of DescribePendingMaintenanceActions.
func (mr *MockDocdbClientMockRecorder) DescribePendingMaintenanceActions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribePendingMaintenanceActions", reflect.TypeOf((*MockDocdbClient)(nil).DescribePendingMaintenanceActions), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockDocdbClient) ListTagsForResource(arg0 context.Context, arg1 *docdb.ListTagsForResourceInput, arg2 ...func(*docdb.Options)) (*docdb.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &docdb.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListTagsForResource")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*docdb.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockDocdbClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockDocdbClient)(nil).ListTagsForResource), varargs...)
}
