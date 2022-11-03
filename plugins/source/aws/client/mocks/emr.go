// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/aws/client (interfaces: EmrClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	emr "github.com/aws/aws-sdk-go-v2/service/emr"
	gomock "github.com/golang/mock/gomock"
)

// MockEmrClient is a mock of EmrClient interface.
type MockEmrClient struct {
	ctrl     *gomock.Controller
	recorder *MockEmrClientMockRecorder
}

// MockEmrClientMockRecorder is the mock recorder for MockEmrClient.
type MockEmrClientMockRecorder struct {
	mock *MockEmrClient
}

// NewMockEmrClient creates a new mock instance.
func NewMockEmrClient(ctrl *gomock.Controller) *MockEmrClient {
	mock := &MockEmrClient{ctrl: ctrl}
	mock.recorder = &MockEmrClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmrClient) EXPECT() *MockEmrClientMockRecorder {
	return m.recorder
}

// DescribeCluster mocks base method.
func (m *MockEmrClient) DescribeCluster(arg0 context.Context, arg1 *emr.DescribeClusterInput, arg2 ...func(*emr.Options)) (*emr.DescribeClusterOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCluster", varargs...)
	ret0, _ := ret[0].(*emr.DescribeClusterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCluster indicates an expected call of DescribeCluster.
func (mr *MockEmrClientMockRecorder) DescribeCluster(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCluster", reflect.TypeOf((*MockEmrClient)(nil).DescribeCluster), varargs...)
}

// DescribeJobFlows mocks base method.
func (m *MockEmrClient) DescribeJobFlows(arg0 context.Context, arg1 *emr.DescribeJobFlowsInput, arg2 ...func(*emr.Options)) (*emr.DescribeJobFlowsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeJobFlows", varargs...)
	ret0, _ := ret[0].(*emr.DescribeJobFlowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeJobFlows indicates an expected call of DescribeJobFlows.
func (mr *MockEmrClientMockRecorder) DescribeJobFlows(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeJobFlows", reflect.TypeOf((*MockEmrClient)(nil).DescribeJobFlows), varargs...)
}

// DescribeNotebookExecution mocks base method.
func (m *MockEmrClient) DescribeNotebookExecution(arg0 context.Context, arg1 *emr.DescribeNotebookExecutionInput, arg2 ...func(*emr.Options)) (*emr.DescribeNotebookExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeNotebookExecution", varargs...)
	ret0, _ := ret[0].(*emr.DescribeNotebookExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeNotebookExecution indicates an expected call of DescribeNotebookExecution.
func (mr *MockEmrClientMockRecorder) DescribeNotebookExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeNotebookExecution", reflect.TypeOf((*MockEmrClient)(nil).DescribeNotebookExecution), varargs...)
}

// DescribeReleaseLabel mocks base method.
func (m *MockEmrClient) DescribeReleaseLabel(arg0 context.Context, arg1 *emr.DescribeReleaseLabelInput, arg2 ...func(*emr.Options)) (*emr.DescribeReleaseLabelOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeReleaseLabel", varargs...)
	ret0, _ := ret[0].(*emr.DescribeReleaseLabelOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeReleaseLabel indicates an expected call of DescribeReleaseLabel.
func (mr *MockEmrClientMockRecorder) DescribeReleaseLabel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReleaseLabel", reflect.TypeOf((*MockEmrClient)(nil).DescribeReleaseLabel), varargs...)
}

// DescribeSecurityConfiguration mocks base method.
func (m *MockEmrClient) DescribeSecurityConfiguration(arg0 context.Context, arg1 *emr.DescribeSecurityConfigurationInput, arg2 ...func(*emr.Options)) (*emr.DescribeSecurityConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSecurityConfiguration", varargs...)
	ret0, _ := ret[0].(*emr.DescribeSecurityConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSecurityConfiguration indicates an expected call of DescribeSecurityConfiguration.
func (mr *MockEmrClientMockRecorder) DescribeSecurityConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSecurityConfiguration", reflect.TypeOf((*MockEmrClient)(nil).DescribeSecurityConfiguration), varargs...)
}

// DescribeStep mocks base method.
func (m *MockEmrClient) DescribeStep(arg0 context.Context, arg1 *emr.DescribeStepInput, arg2 ...func(*emr.Options)) (*emr.DescribeStepOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeStep", varargs...)
	ret0, _ := ret[0].(*emr.DescribeStepOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeStep indicates an expected call of DescribeStep.
func (mr *MockEmrClientMockRecorder) DescribeStep(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStep", reflect.TypeOf((*MockEmrClient)(nil).DescribeStep), varargs...)
}

// DescribeStudio mocks base method.
func (m *MockEmrClient) DescribeStudio(arg0 context.Context, arg1 *emr.DescribeStudioInput, arg2 ...func(*emr.Options)) (*emr.DescribeStudioOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeStudio", varargs...)
	ret0, _ := ret[0].(*emr.DescribeStudioOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeStudio indicates an expected call of DescribeStudio.
func (mr *MockEmrClientMockRecorder) DescribeStudio(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStudio", reflect.TypeOf((*MockEmrClient)(nil).DescribeStudio), varargs...)
}

// GetAutoTerminationPolicy mocks base method.
func (m *MockEmrClient) GetAutoTerminationPolicy(arg0 context.Context, arg1 *emr.GetAutoTerminationPolicyInput, arg2 ...func(*emr.Options)) (*emr.GetAutoTerminationPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAutoTerminationPolicy", varargs...)
	ret0, _ := ret[0].(*emr.GetAutoTerminationPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAutoTerminationPolicy indicates an expected call of GetAutoTerminationPolicy.
func (mr *MockEmrClientMockRecorder) GetAutoTerminationPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAutoTerminationPolicy", reflect.TypeOf((*MockEmrClient)(nil).GetAutoTerminationPolicy), varargs...)
}

// GetBlockPublicAccessConfiguration mocks base method.
func (m *MockEmrClient) GetBlockPublicAccessConfiguration(arg0 context.Context, arg1 *emr.GetBlockPublicAccessConfigurationInput, arg2 ...func(*emr.Options)) (*emr.GetBlockPublicAccessConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBlockPublicAccessConfiguration", varargs...)
	ret0, _ := ret[0].(*emr.GetBlockPublicAccessConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockPublicAccessConfiguration indicates an expected call of GetBlockPublicAccessConfiguration.
func (mr *MockEmrClientMockRecorder) GetBlockPublicAccessConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockPublicAccessConfiguration", reflect.TypeOf((*MockEmrClient)(nil).GetBlockPublicAccessConfiguration), varargs...)
}

// GetManagedScalingPolicy mocks base method.
func (m *MockEmrClient) GetManagedScalingPolicy(arg0 context.Context, arg1 *emr.GetManagedScalingPolicyInput, arg2 ...func(*emr.Options)) (*emr.GetManagedScalingPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetManagedScalingPolicy", varargs...)
	ret0, _ := ret[0].(*emr.GetManagedScalingPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetManagedScalingPolicy indicates an expected call of GetManagedScalingPolicy.
func (mr *MockEmrClientMockRecorder) GetManagedScalingPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagedScalingPolicy", reflect.TypeOf((*MockEmrClient)(nil).GetManagedScalingPolicy), varargs...)
}

// GetStudioSessionMapping mocks base method.
func (m *MockEmrClient) GetStudioSessionMapping(arg0 context.Context, arg1 *emr.GetStudioSessionMappingInput, arg2 ...func(*emr.Options)) (*emr.GetStudioSessionMappingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStudioSessionMapping", varargs...)
	ret0, _ := ret[0].(*emr.GetStudioSessionMappingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudioSessionMapping indicates an expected call of GetStudioSessionMapping.
func (mr *MockEmrClientMockRecorder) GetStudioSessionMapping(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudioSessionMapping", reflect.TypeOf((*MockEmrClient)(nil).GetStudioSessionMapping), varargs...)
}

// ListBootstrapActions mocks base method.
func (m *MockEmrClient) ListBootstrapActions(arg0 context.Context, arg1 *emr.ListBootstrapActionsInput, arg2 ...func(*emr.Options)) (*emr.ListBootstrapActionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBootstrapActions", varargs...)
	ret0, _ := ret[0].(*emr.ListBootstrapActionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBootstrapActions indicates an expected call of ListBootstrapActions.
func (mr *MockEmrClientMockRecorder) ListBootstrapActions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBootstrapActions", reflect.TypeOf((*MockEmrClient)(nil).ListBootstrapActions), varargs...)
}

// ListClusters mocks base method.
func (m *MockEmrClient) ListClusters(arg0 context.Context, arg1 *emr.ListClustersInput, arg2 ...func(*emr.Options)) (*emr.ListClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListClusters", varargs...)
	ret0, _ := ret[0].(*emr.ListClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusters indicates an expected call of ListClusters.
func (mr *MockEmrClientMockRecorder) ListClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockEmrClient)(nil).ListClusters), varargs...)
}

// ListInstanceFleets mocks base method.
func (m *MockEmrClient) ListInstanceFleets(arg0 context.Context, arg1 *emr.ListInstanceFleetsInput, arg2 ...func(*emr.Options)) (*emr.ListInstanceFleetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListInstanceFleets", varargs...)
	ret0, _ := ret[0].(*emr.ListInstanceFleetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInstanceFleets indicates an expected call of ListInstanceFleets.
func (mr *MockEmrClientMockRecorder) ListInstanceFleets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInstanceFleets", reflect.TypeOf((*MockEmrClient)(nil).ListInstanceFleets), varargs...)
}

// ListInstanceGroups mocks base method.
func (m *MockEmrClient) ListInstanceGroups(arg0 context.Context, arg1 *emr.ListInstanceGroupsInput, arg2 ...func(*emr.Options)) (*emr.ListInstanceGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListInstanceGroups", varargs...)
	ret0, _ := ret[0].(*emr.ListInstanceGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInstanceGroups indicates an expected call of ListInstanceGroups.
func (mr *MockEmrClientMockRecorder) ListInstanceGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInstanceGroups", reflect.TypeOf((*MockEmrClient)(nil).ListInstanceGroups), varargs...)
}

// ListInstances mocks base method.
func (m *MockEmrClient) ListInstances(arg0 context.Context, arg1 *emr.ListInstancesInput, arg2 ...func(*emr.Options)) (*emr.ListInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListInstances", varargs...)
	ret0, _ := ret[0].(*emr.ListInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInstances indicates an expected call of ListInstances.
func (mr *MockEmrClientMockRecorder) ListInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInstances", reflect.TypeOf((*MockEmrClient)(nil).ListInstances), varargs...)
}

// ListNotebookExecutions mocks base method.
func (m *MockEmrClient) ListNotebookExecutions(arg0 context.Context, arg1 *emr.ListNotebookExecutionsInput, arg2 ...func(*emr.Options)) (*emr.ListNotebookExecutionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListNotebookExecutions", varargs...)
	ret0, _ := ret[0].(*emr.ListNotebookExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNotebookExecutions indicates an expected call of ListNotebookExecutions.
func (mr *MockEmrClientMockRecorder) ListNotebookExecutions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNotebookExecutions", reflect.TypeOf((*MockEmrClient)(nil).ListNotebookExecutions), varargs...)
}

// ListReleaseLabels mocks base method.
func (m *MockEmrClient) ListReleaseLabels(arg0 context.Context, arg1 *emr.ListReleaseLabelsInput, arg2 ...func(*emr.Options)) (*emr.ListReleaseLabelsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListReleaseLabels", varargs...)
	ret0, _ := ret[0].(*emr.ListReleaseLabelsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReleaseLabels indicates an expected call of ListReleaseLabels.
func (mr *MockEmrClientMockRecorder) ListReleaseLabels(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReleaseLabels", reflect.TypeOf((*MockEmrClient)(nil).ListReleaseLabels), varargs...)
}

// ListSecurityConfigurations mocks base method.
func (m *MockEmrClient) ListSecurityConfigurations(arg0 context.Context, arg1 *emr.ListSecurityConfigurationsInput, arg2 ...func(*emr.Options)) (*emr.ListSecurityConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSecurityConfigurations", varargs...)
	ret0, _ := ret[0].(*emr.ListSecurityConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecurityConfigurations indicates an expected call of ListSecurityConfigurations.
func (mr *MockEmrClientMockRecorder) ListSecurityConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecurityConfigurations", reflect.TypeOf((*MockEmrClient)(nil).ListSecurityConfigurations), varargs...)
}

// ListSteps mocks base method.
func (m *MockEmrClient) ListSteps(arg0 context.Context, arg1 *emr.ListStepsInput, arg2 ...func(*emr.Options)) (*emr.ListStepsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSteps", varargs...)
	ret0, _ := ret[0].(*emr.ListStepsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSteps indicates an expected call of ListSteps.
func (mr *MockEmrClientMockRecorder) ListSteps(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSteps", reflect.TypeOf((*MockEmrClient)(nil).ListSteps), varargs...)
}

// ListStudioSessionMappings mocks base method.
func (m *MockEmrClient) ListStudioSessionMappings(arg0 context.Context, arg1 *emr.ListStudioSessionMappingsInput, arg2 ...func(*emr.Options)) (*emr.ListStudioSessionMappingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListStudioSessionMappings", varargs...)
	ret0, _ := ret[0].(*emr.ListStudioSessionMappingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStudioSessionMappings indicates an expected call of ListStudioSessionMappings.
func (mr *MockEmrClientMockRecorder) ListStudioSessionMappings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStudioSessionMappings", reflect.TypeOf((*MockEmrClient)(nil).ListStudioSessionMappings), varargs...)
}

// ListStudios mocks base method.
func (m *MockEmrClient) ListStudios(arg0 context.Context, arg1 *emr.ListStudiosInput, arg2 ...func(*emr.Options)) (*emr.ListStudiosOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListStudios", varargs...)
	ret0, _ := ret[0].(*emr.ListStudiosOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStudios indicates an expected call of ListStudios.
func (mr *MockEmrClientMockRecorder) ListStudios(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStudios", reflect.TypeOf((*MockEmrClient)(nil).ListStudios), varargs...)
}
