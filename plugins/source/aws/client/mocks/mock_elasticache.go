// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/aws/client (interfaces: ElastiCache)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	elasticache "github.com/aws/aws-sdk-go-v2/service/elasticache"
	gomock "github.com/golang/mock/gomock"
)

// MockElastiCache is a mock of ElastiCache interface.
type MockElastiCache struct {
	ctrl     *gomock.Controller
	recorder *MockElastiCacheMockRecorder
}

// MockElastiCacheMockRecorder is the mock recorder for MockElastiCache.
type MockElastiCacheMockRecorder struct {
	mock *MockElastiCache
}

// NewMockElastiCache creates a new mock instance.
func NewMockElastiCache(ctrl *gomock.Controller) *MockElastiCache {
	mock := &MockElastiCache{ctrl: ctrl}
	mock.recorder = &MockElastiCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockElastiCache) EXPECT() *MockElastiCacheMockRecorder {
	return m.recorder
}

// DescribeCacheClusters mocks base method.
func (m *MockElastiCache) DescribeCacheClusters(arg0 context.Context, arg1 *elasticache.DescribeCacheClustersInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCacheClusters", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCacheClusters indicates an expected call of DescribeCacheClusters.
func (mr *MockElastiCacheMockRecorder) DescribeCacheClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCacheClusters", reflect.TypeOf((*MockElastiCache)(nil).DescribeCacheClusters), varargs...)
}

// DescribeCacheEngineVersions mocks base method.
func (m *MockElastiCache) DescribeCacheEngineVersions(arg0 context.Context, arg1 *elasticache.DescribeCacheEngineVersionsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheEngineVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCacheEngineVersions", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheEngineVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCacheEngineVersions indicates an expected call of DescribeCacheEngineVersions.
func (mr *MockElastiCacheMockRecorder) DescribeCacheEngineVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCacheEngineVersions", reflect.TypeOf((*MockElastiCache)(nil).DescribeCacheEngineVersions), varargs...)
}

// DescribeCacheParameterGroups mocks base method.
func (m *MockElastiCache) DescribeCacheParameterGroups(arg0 context.Context, arg1 *elasticache.DescribeCacheParameterGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheParameterGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCacheParameterGroups", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheParameterGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCacheParameterGroups indicates an expected call of DescribeCacheParameterGroups.
func (mr *MockElastiCacheMockRecorder) DescribeCacheParameterGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCacheParameterGroups", reflect.TypeOf((*MockElastiCache)(nil).DescribeCacheParameterGroups), varargs...)
}

// DescribeCacheParameters mocks base method.
func (m *MockElastiCache) DescribeCacheParameters(arg0 context.Context, arg1 *elasticache.DescribeCacheParametersInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCacheParameters", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCacheParameters indicates an expected call of DescribeCacheParameters.
func (mr *MockElastiCacheMockRecorder) DescribeCacheParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCacheParameters", reflect.TypeOf((*MockElastiCache)(nil).DescribeCacheParameters), varargs...)
}

// DescribeCacheSubnetGroups mocks base method.
func (m *MockElastiCache) DescribeCacheSubnetGroups(arg0 context.Context, arg1 *elasticache.DescribeCacheSubnetGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheSubnetGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCacheSubnetGroups", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheSubnetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCacheSubnetGroups indicates an expected call of DescribeCacheSubnetGroups.
func (mr *MockElastiCacheMockRecorder) DescribeCacheSubnetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCacheSubnetGroups", reflect.TypeOf((*MockElastiCache)(nil).DescribeCacheSubnetGroups), varargs...)
}

// DescribeGlobalReplicationGroups mocks base method.
func (m *MockElastiCache) DescribeGlobalReplicationGroups(arg0 context.Context, arg1 *elasticache.DescribeGlobalReplicationGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeGlobalReplicationGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeGlobalReplicationGroups", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeGlobalReplicationGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeGlobalReplicationGroups indicates an expected call of DescribeGlobalReplicationGroups.
func (mr *MockElastiCacheMockRecorder) DescribeGlobalReplicationGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeGlobalReplicationGroups", reflect.TypeOf((*MockElastiCache)(nil).DescribeGlobalReplicationGroups), varargs...)
}

// DescribeReplicationGroups mocks base method.
func (m *MockElastiCache) DescribeReplicationGroups(arg0 context.Context, arg1 *elasticache.DescribeReplicationGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeReplicationGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeReplicationGroups", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeReplicationGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeReplicationGroups indicates an expected call of DescribeReplicationGroups.
func (mr *MockElastiCacheMockRecorder) DescribeReplicationGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReplicationGroups", reflect.TypeOf((*MockElastiCache)(nil).DescribeReplicationGroups), varargs...)
}

// DescribeReservedCacheNodes mocks base method.
func (m *MockElastiCache) DescribeReservedCacheNodes(arg0 context.Context, arg1 *elasticache.DescribeReservedCacheNodesInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeReservedCacheNodesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeReservedCacheNodes", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeReservedCacheNodesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeReservedCacheNodes indicates an expected call of DescribeReservedCacheNodes.
func (mr *MockElastiCacheMockRecorder) DescribeReservedCacheNodes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReservedCacheNodes", reflect.TypeOf((*MockElastiCache)(nil).DescribeReservedCacheNodes), varargs...)
}

// DescribeReservedCacheNodesOfferings mocks base method.
func (m *MockElastiCache) DescribeReservedCacheNodesOfferings(arg0 context.Context, arg1 *elasticache.DescribeReservedCacheNodesOfferingsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeReservedCacheNodesOfferingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeReservedCacheNodesOfferings", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeReservedCacheNodesOfferingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeReservedCacheNodesOfferings indicates an expected call of DescribeReservedCacheNodesOfferings.
func (mr *MockElastiCacheMockRecorder) DescribeReservedCacheNodesOfferings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReservedCacheNodesOfferings", reflect.TypeOf((*MockElastiCache)(nil).DescribeReservedCacheNodesOfferings), varargs...)
}

// DescribeServiceUpdates mocks base method.
func (m *MockElastiCache) DescribeServiceUpdates(arg0 context.Context, arg1 *elasticache.DescribeServiceUpdatesInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeServiceUpdatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeServiceUpdates", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeServiceUpdatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeServiceUpdates indicates an expected call of DescribeServiceUpdates.
func (mr *MockElastiCacheMockRecorder) DescribeServiceUpdates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeServiceUpdates", reflect.TypeOf((*MockElastiCache)(nil).DescribeServiceUpdates), varargs...)
}

// DescribeSnapshots mocks base method.
func (m *MockElastiCache) DescribeSnapshots(arg0 context.Context, arg1 *elasticache.DescribeSnapshotsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSnapshots", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSnapshots indicates an expected call of DescribeSnapshots.
func (mr *MockElastiCacheMockRecorder) DescribeSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSnapshots", reflect.TypeOf((*MockElastiCache)(nil).DescribeSnapshots), varargs...)
}

// DescribeUserGroups mocks base method.
func (m *MockElastiCache) DescribeUserGroups(arg0 context.Context, arg1 *elasticache.DescribeUserGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeUserGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeUserGroups", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeUserGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeUserGroups indicates an expected call of DescribeUserGroups.
func (mr *MockElastiCacheMockRecorder) DescribeUserGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUserGroups", reflect.TypeOf((*MockElastiCache)(nil).DescribeUserGroups), varargs...)
}

// DescribeUsers mocks base method.
func (m *MockElastiCache) DescribeUsers(arg0 context.Context, arg1 *elasticache.DescribeUsersInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeUsersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeUsers", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeUsersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeUsers indicates an expected call of DescribeUsers.
func (mr *MockElastiCacheMockRecorder) DescribeUsers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUsers", reflect.TypeOf((*MockElastiCache)(nil).DescribeUsers), varargs...)
}
