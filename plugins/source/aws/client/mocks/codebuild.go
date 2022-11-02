// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/aws/client (interfaces: CodebuildClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	codebuild "github.com/aws/aws-sdk-go-v2/service/codebuild"
	gomock "github.com/golang/mock/gomock"
)

// MockCodebuildClient is a mock of CodebuildClient interface.
type MockCodebuildClient struct {
	ctrl     *gomock.Controller
	recorder *MockCodebuildClientMockRecorder
}

// MockCodebuildClientMockRecorder is the mock recorder for MockCodebuildClient.
type MockCodebuildClientMockRecorder struct {
	mock *MockCodebuildClient
}

// NewMockCodebuildClient creates a new mock instance.
func NewMockCodebuildClient(ctrl *gomock.Controller) *MockCodebuildClient {
	mock := &MockCodebuildClient{ctrl: ctrl}
	mock.recorder = &MockCodebuildClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCodebuildClient) EXPECT() *MockCodebuildClientMockRecorder {
	return m.recorder
}

// DescribeCodeCoverages mocks base method.
func (m *MockCodebuildClient) DescribeCodeCoverages(arg0 context.Context, arg1 *codebuild.DescribeCodeCoveragesInput, arg2 ...func(*codebuild.Options)) (*codebuild.DescribeCodeCoveragesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCodeCoverages", varargs...)
	ret0, _ := ret[0].(*codebuild.DescribeCodeCoveragesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCodeCoverages indicates an expected call of DescribeCodeCoverages.
func (mr *MockCodebuildClientMockRecorder) DescribeCodeCoverages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCodeCoverages", reflect.TypeOf((*MockCodebuildClient)(nil).DescribeCodeCoverages), varargs...)
}

// DescribeTestCases mocks base method.
func (m *MockCodebuildClient) DescribeTestCases(arg0 context.Context, arg1 *codebuild.DescribeTestCasesInput, arg2 ...func(*codebuild.Options)) (*codebuild.DescribeTestCasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTestCases", varargs...)
	ret0, _ := ret[0].(*codebuild.DescribeTestCasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTestCases indicates an expected call of DescribeTestCases.
func (mr *MockCodebuildClientMockRecorder) DescribeTestCases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTestCases", reflect.TypeOf((*MockCodebuildClient)(nil).DescribeTestCases), varargs...)
}

// GetReportGroupTrend mocks base method.
func (m *MockCodebuildClient) GetReportGroupTrend(arg0 context.Context, arg1 *codebuild.GetReportGroupTrendInput, arg2 ...func(*codebuild.Options)) (*codebuild.GetReportGroupTrendOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetReportGroupTrend", varargs...)
	ret0, _ := ret[0].(*codebuild.GetReportGroupTrendOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReportGroupTrend indicates an expected call of GetReportGroupTrend.
func (mr *MockCodebuildClientMockRecorder) GetReportGroupTrend(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReportGroupTrend", reflect.TypeOf((*MockCodebuildClient)(nil).GetReportGroupTrend), varargs...)
}

// GetResourcePolicy mocks base method.
func (m *MockCodebuildClient) GetResourcePolicy(arg0 context.Context, arg1 *codebuild.GetResourcePolicyInput, arg2 ...func(*codebuild.Options)) (*codebuild.GetResourcePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetResourcePolicy", varargs...)
	ret0, _ := ret[0].(*codebuild.GetResourcePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourcePolicy indicates an expected call of GetResourcePolicy.
func (mr *MockCodebuildClientMockRecorder) GetResourcePolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourcePolicy", reflect.TypeOf((*MockCodebuildClient)(nil).GetResourcePolicy), varargs...)
}

// ListBuildBatches mocks base method.
func (m *MockCodebuildClient) ListBuildBatches(arg0 context.Context, arg1 *codebuild.ListBuildBatchesInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListBuildBatchesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBuildBatches", varargs...)
	ret0, _ := ret[0].(*codebuild.ListBuildBatchesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBuildBatches indicates an expected call of ListBuildBatches.
func (mr *MockCodebuildClientMockRecorder) ListBuildBatches(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBuildBatches", reflect.TypeOf((*MockCodebuildClient)(nil).ListBuildBatches), varargs...)
}

// ListBuildBatchesForProject mocks base method.
func (m *MockCodebuildClient) ListBuildBatchesForProject(arg0 context.Context, arg1 *codebuild.ListBuildBatchesForProjectInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListBuildBatchesForProjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBuildBatchesForProject", varargs...)
	ret0, _ := ret[0].(*codebuild.ListBuildBatchesForProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBuildBatchesForProject indicates an expected call of ListBuildBatchesForProject.
func (mr *MockCodebuildClientMockRecorder) ListBuildBatchesForProject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBuildBatchesForProject", reflect.TypeOf((*MockCodebuildClient)(nil).ListBuildBatchesForProject), varargs...)
}

// ListBuilds mocks base method.
func (m *MockCodebuildClient) ListBuilds(arg0 context.Context, arg1 *codebuild.ListBuildsInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListBuildsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBuilds", varargs...)
	ret0, _ := ret[0].(*codebuild.ListBuildsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBuilds indicates an expected call of ListBuilds.
func (mr *MockCodebuildClientMockRecorder) ListBuilds(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBuilds", reflect.TypeOf((*MockCodebuildClient)(nil).ListBuilds), varargs...)
}

// ListBuildsForProject mocks base method.
func (m *MockCodebuildClient) ListBuildsForProject(arg0 context.Context, arg1 *codebuild.ListBuildsForProjectInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListBuildsForProjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBuildsForProject", varargs...)
	ret0, _ := ret[0].(*codebuild.ListBuildsForProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBuildsForProject indicates an expected call of ListBuildsForProject.
func (mr *MockCodebuildClientMockRecorder) ListBuildsForProject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBuildsForProject", reflect.TypeOf((*MockCodebuildClient)(nil).ListBuildsForProject), varargs...)
}

// ListCuratedEnvironmentImages mocks base method.
func (m *MockCodebuildClient) ListCuratedEnvironmentImages(arg0 context.Context, arg1 *codebuild.ListCuratedEnvironmentImagesInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListCuratedEnvironmentImagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCuratedEnvironmentImages", varargs...)
	ret0, _ := ret[0].(*codebuild.ListCuratedEnvironmentImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCuratedEnvironmentImages indicates an expected call of ListCuratedEnvironmentImages.
func (mr *MockCodebuildClientMockRecorder) ListCuratedEnvironmentImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCuratedEnvironmentImages", reflect.TypeOf((*MockCodebuildClient)(nil).ListCuratedEnvironmentImages), varargs...)
}

// ListProjects mocks base method.
func (m *MockCodebuildClient) ListProjects(arg0 context.Context, arg1 *codebuild.ListProjectsInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListProjectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjects", varargs...)
	ret0, _ := ret[0].(*codebuild.ListProjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjects indicates an expected call of ListProjects.
func (mr *MockCodebuildClientMockRecorder) ListProjects(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockCodebuildClient)(nil).ListProjects), varargs...)
}

// ListReportGroups mocks base method.
func (m *MockCodebuildClient) ListReportGroups(arg0 context.Context, arg1 *codebuild.ListReportGroupsInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListReportGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListReportGroups", varargs...)
	ret0, _ := ret[0].(*codebuild.ListReportGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReportGroups indicates an expected call of ListReportGroups.
func (mr *MockCodebuildClientMockRecorder) ListReportGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReportGroups", reflect.TypeOf((*MockCodebuildClient)(nil).ListReportGroups), varargs...)
}

// ListReports mocks base method.
func (m *MockCodebuildClient) ListReports(arg0 context.Context, arg1 *codebuild.ListReportsInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListReportsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListReports", varargs...)
	ret0, _ := ret[0].(*codebuild.ListReportsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReports indicates an expected call of ListReports.
func (mr *MockCodebuildClientMockRecorder) ListReports(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReports", reflect.TypeOf((*MockCodebuildClient)(nil).ListReports), varargs...)
}

// ListReportsForReportGroup mocks base method.
func (m *MockCodebuildClient) ListReportsForReportGroup(arg0 context.Context, arg1 *codebuild.ListReportsForReportGroupInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListReportsForReportGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListReportsForReportGroup", varargs...)
	ret0, _ := ret[0].(*codebuild.ListReportsForReportGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReportsForReportGroup indicates an expected call of ListReportsForReportGroup.
func (mr *MockCodebuildClientMockRecorder) ListReportsForReportGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReportsForReportGroup", reflect.TypeOf((*MockCodebuildClient)(nil).ListReportsForReportGroup), varargs...)
}

// ListSharedProjects mocks base method.
func (m *MockCodebuildClient) ListSharedProjects(arg0 context.Context, arg1 *codebuild.ListSharedProjectsInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListSharedProjectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSharedProjects", varargs...)
	ret0, _ := ret[0].(*codebuild.ListSharedProjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSharedProjects indicates an expected call of ListSharedProjects.
func (mr *MockCodebuildClientMockRecorder) ListSharedProjects(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSharedProjects", reflect.TypeOf((*MockCodebuildClient)(nil).ListSharedProjects), varargs...)
}

// ListSharedReportGroups mocks base method.
func (m *MockCodebuildClient) ListSharedReportGroups(arg0 context.Context, arg1 *codebuild.ListSharedReportGroupsInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListSharedReportGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSharedReportGroups", varargs...)
	ret0, _ := ret[0].(*codebuild.ListSharedReportGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSharedReportGroups indicates an expected call of ListSharedReportGroups.
func (mr *MockCodebuildClientMockRecorder) ListSharedReportGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSharedReportGroups", reflect.TypeOf((*MockCodebuildClient)(nil).ListSharedReportGroups), varargs...)
}

// ListSourceCredentials mocks base method.
func (m *MockCodebuildClient) ListSourceCredentials(arg0 context.Context, arg1 *codebuild.ListSourceCredentialsInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListSourceCredentialsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSourceCredentials", varargs...)
	ret0, _ := ret[0].(*codebuild.ListSourceCredentialsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSourceCredentials indicates an expected call of ListSourceCredentials.
func (mr *MockCodebuildClientMockRecorder) ListSourceCredentials(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSourceCredentials", reflect.TypeOf((*MockCodebuildClient)(nil).ListSourceCredentials), varargs...)
}
