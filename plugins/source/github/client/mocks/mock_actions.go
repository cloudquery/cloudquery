// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/github/client (interfaces: ActionsService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v59/github"
)

// MockActionsService is a mock of ActionsService interface.
type MockActionsService struct {
	ctrl     *gomock.Controller
	recorder *MockActionsServiceMockRecorder
}

// MockActionsServiceMockRecorder is the mock recorder for MockActionsService.
type MockActionsServiceMockRecorder struct {
	mock *MockActionsService
}

// NewMockActionsService creates a new mock instance.
func NewMockActionsService(ctrl *gomock.Controller) *MockActionsService {
	mock := &MockActionsService{ctrl: ctrl}
	mock.recorder = &MockActionsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActionsService) EXPECT() *MockActionsServiceMockRecorder {
	return m.recorder
}

// GetWorkflowRunUsageByID mocks base method.
func (m *MockActionsService) GetWorkflowRunUsageByID(arg0 context.Context, arg1, arg2 string, arg3 int64) (*github.WorkflowRunUsage, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkflowRunUsageByID", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*github.WorkflowRunUsage)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetWorkflowRunUsageByID indicates an expected call of GetWorkflowRunUsageByID.
func (mr *MockActionsServiceMockRecorder) GetWorkflowRunUsageByID(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowRunUsageByID", reflect.TypeOf((*MockActionsService)(nil).GetWorkflowRunUsageByID), arg0, arg1, arg2, arg3)
}

// ListRepositoryWorkflowRuns mocks base method.
func (m *MockActionsService) ListRepositoryWorkflowRuns(arg0 context.Context, arg1, arg2 string, arg3 *github.ListWorkflowRunsOptions) (*github.WorkflowRuns, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRepositoryWorkflowRuns", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*github.WorkflowRuns)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListRepositoryWorkflowRuns indicates an expected call of ListRepositoryWorkflowRuns.
func (mr *MockActionsServiceMockRecorder) ListRepositoryWorkflowRuns(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRepositoryWorkflowRuns", reflect.TypeOf((*MockActionsService)(nil).ListRepositoryWorkflowRuns), arg0, arg1, arg2, arg3)
}

// ListWorkflowJobs mocks base method.
func (m *MockActionsService) ListWorkflowJobs(arg0 context.Context, arg1, arg2 string, arg3 int64, arg4 *github.ListWorkflowJobsOptions) (*github.Jobs, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWorkflowJobs", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*github.Jobs)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListWorkflowJobs indicates an expected call of ListWorkflowJobs.
func (mr *MockActionsServiceMockRecorder) ListWorkflowJobs(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkflowJobs", reflect.TypeOf((*MockActionsService)(nil).ListWorkflowJobs), arg0, arg1, arg2, arg3, arg4)
}

// ListWorkflows mocks base method.
func (m *MockActionsService) ListWorkflows(arg0 context.Context, arg1, arg2 string, arg3 *github.ListOptions) (*github.Workflows, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWorkflows", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*github.Workflows)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListWorkflows indicates an expected call of ListWorkflows.
func (mr *MockActionsServiceMockRecorder) ListWorkflows(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkflows", reflect.TypeOf((*MockActionsService)(nil).ListWorkflows), arg0, arg1, arg2, arg3)
}
