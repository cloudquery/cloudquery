// Code generated by MockGen. DO NOT EDIT.
// Source: amplify.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	amplify "github.com/aws/aws-sdk-go-v2/service/amplify"
	gomock "github.com/golang/mock/gomock"
)

// MockAmplifyClient is a mock of AmplifyClient interface.
type MockAmplifyClient struct {
	ctrl     *gomock.Controller
	recorder *MockAmplifyClientMockRecorder
}

// MockAmplifyClientMockRecorder is the mock recorder for MockAmplifyClient.
type MockAmplifyClientMockRecorder struct {
	mock *MockAmplifyClient
}

// NewMockAmplifyClient creates a new mock instance.
func NewMockAmplifyClient(ctrl *gomock.Controller) *MockAmplifyClient {
	mock := &MockAmplifyClient{ctrl: ctrl}
	mock.recorder = &MockAmplifyClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAmplifyClient) EXPECT() *MockAmplifyClientMockRecorder {
	return m.recorder
}

// GetApp mocks base method.
func (m *MockAmplifyClient) GetApp(arg0 context.Context, arg1 *amplify.GetAppInput, arg2 ...func(*amplify.Options)) (*amplify.GetAppOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetApp", varargs...)
	ret0, _ := ret[0].(*amplify.GetAppOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApp indicates an expected call of GetApp.
func (mr *MockAmplifyClientMockRecorder) GetApp(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApp", reflect.TypeOf((*MockAmplifyClient)(nil).GetApp), varargs...)
}

// GetArtifactUrl mocks base method.
func (m *MockAmplifyClient) GetArtifactUrl(arg0 context.Context, arg1 *amplify.GetArtifactUrlInput, arg2 ...func(*amplify.Options)) (*amplify.GetArtifactUrlOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetArtifactUrl", varargs...)
	ret0, _ := ret[0].(*amplify.GetArtifactUrlOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArtifactUrl indicates an expected call of GetArtifactUrl.
func (mr *MockAmplifyClientMockRecorder) GetArtifactUrl(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArtifactUrl", reflect.TypeOf((*MockAmplifyClient)(nil).GetArtifactUrl), varargs...)
}

// GetBackendEnvironment mocks base method.
func (m *MockAmplifyClient) GetBackendEnvironment(arg0 context.Context, arg1 *amplify.GetBackendEnvironmentInput, arg2 ...func(*amplify.Options)) (*amplify.GetBackendEnvironmentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBackendEnvironment", varargs...)
	ret0, _ := ret[0].(*amplify.GetBackendEnvironmentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackendEnvironment indicates an expected call of GetBackendEnvironment.
func (mr *MockAmplifyClientMockRecorder) GetBackendEnvironment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackendEnvironment", reflect.TypeOf((*MockAmplifyClient)(nil).GetBackendEnvironment), varargs...)
}

// GetBranch mocks base method.
func (m *MockAmplifyClient) GetBranch(arg0 context.Context, arg1 *amplify.GetBranchInput, arg2 ...func(*amplify.Options)) (*amplify.GetBranchOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBranch", varargs...)
	ret0, _ := ret[0].(*amplify.GetBranchOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBranch indicates an expected call of GetBranch.
func (mr *MockAmplifyClientMockRecorder) GetBranch(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBranch", reflect.TypeOf((*MockAmplifyClient)(nil).GetBranch), varargs...)
}

// GetDomainAssociation mocks base method.
func (m *MockAmplifyClient) GetDomainAssociation(arg0 context.Context, arg1 *amplify.GetDomainAssociationInput, arg2 ...func(*amplify.Options)) (*amplify.GetDomainAssociationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDomainAssociation", varargs...)
	ret0, _ := ret[0].(*amplify.GetDomainAssociationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDomainAssociation indicates an expected call of GetDomainAssociation.
func (mr *MockAmplifyClientMockRecorder) GetDomainAssociation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainAssociation", reflect.TypeOf((*MockAmplifyClient)(nil).GetDomainAssociation), varargs...)
}

// GetJob mocks base method.
func (m *MockAmplifyClient) GetJob(arg0 context.Context, arg1 *amplify.GetJobInput, arg2 ...func(*amplify.Options)) (*amplify.GetJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetJob", varargs...)
	ret0, _ := ret[0].(*amplify.GetJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJob indicates an expected call of GetJob.
func (mr *MockAmplifyClientMockRecorder) GetJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJob", reflect.TypeOf((*MockAmplifyClient)(nil).GetJob), varargs...)
}

// GetWebhook mocks base method.
func (m *MockAmplifyClient) GetWebhook(arg0 context.Context, arg1 *amplify.GetWebhookInput, arg2 ...func(*amplify.Options)) (*amplify.GetWebhookOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWebhook", varargs...)
	ret0, _ := ret[0].(*amplify.GetWebhookOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWebhook indicates an expected call of GetWebhook.
func (mr *MockAmplifyClientMockRecorder) GetWebhook(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWebhook", reflect.TypeOf((*MockAmplifyClient)(nil).GetWebhook), varargs...)
}

// ListApps mocks base method.
func (m *MockAmplifyClient) ListApps(arg0 context.Context, arg1 *amplify.ListAppsInput, arg2 ...func(*amplify.Options)) (*amplify.ListAppsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListApps", varargs...)
	ret0, _ := ret[0].(*amplify.ListAppsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApps indicates an expected call of ListApps.
func (mr *MockAmplifyClientMockRecorder) ListApps(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApps", reflect.TypeOf((*MockAmplifyClient)(nil).ListApps), varargs...)
}

// ListArtifacts mocks base method.
func (m *MockAmplifyClient) ListArtifacts(arg0 context.Context, arg1 *amplify.ListArtifactsInput, arg2 ...func(*amplify.Options)) (*amplify.ListArtifactsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListArtifacts", varargs...)
	ret0, _ := ret[0].(*amplify.ListArtifactsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListArtifacts indicates an expected call of ListArtifacts.
func (mr *MockAmplifyClientMockRecorder) ListArtifacts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListArtifacts", reflect.TypeOf((*MockAmplifyClient)(nil).ListArtifacts), varargs...)
}

// ListBackendEnvironments mocks base method.
func (m *MockAmplifyClient) ListBackendEnvironments(arg0 context.Context, arg1 *amplify.ListBackendEnvironmentsInput, arg2 ...func(*amplify.Options)) (*amplify.ListBackendEnvironmentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBackendEnvironments", varargs...)
	ret0, _ := ret[0].(*amplify.ListBackendEnvironmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBackendEnvironments indicates an expected call of ListBackendEnvironments.
func (mr *MockAmplifyClientMockRecorder) ListBackendEnvironments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBackendEnvironments", reflect.TypeOf((*MockAmplifyClient)(nil).ListBackendEnvironments), varargs...)
}

// ListBranches mocks base method.
func (m *MockAmplifyClient) ListBranches(arg0 context.Context, arg1 *amplify.ListBranchesInput, arg2 ...func(*amplify.Options)) (*amplify.ListBranchesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBranches", varargs...)
	ret0, _ := ret[0].(*amplify.ListBranchesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBranches indicates an expected call of ListBranches.
func (mr *MockAmplifyClientMockRecorder) ListBranches(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBranches", reflect.TypeOf((*MockAmplifyClient)(nil).ListBranches), varargs...)
}

// ListDomainAssociations mocks base method.
func (m *MockAmplifyClient) ListDomainAssociations(arg0 context.Context, arg1 *amplify.ListDomainAssociationsInput, arg2 ...func(*amplify.Options)) (*amplify.ListDomainAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDomainAssociations", varargs...)
	ret0, _ := ret[0].(*amplify.ListDomainAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDomainAssociations indicates an expected call of ListDomainAssociations.
func (mr *MockAmplifyClientMockRecorder) ListDomainAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDomainAssociations", reflect.TypeOf((*MockAmplifyClient)(nil).ListDomainAssociations), varargs...)
}

// ListJobs mocks base method.
func (m *MockAmplifyClient) ListJobs(arg0 context.Context, arg1 *amplify.ListJobsInput, arg2 ...func(*amplify.Options)) (*amplify.ListJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListJobs", varargs...)
	ret0, _ := ret[0].(*amplify.ListJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListJobs indicates an expected call of ListJobs.
func (mr *MockAmplifyClientMockRecorder) ListJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListJobs", reflect.TypeOf((*MockAmplifyClient)(nil).ListJobs), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockAmplifyClient) ListTagsForResource(arg0 context.Context, arg1 *amplify.ListTagsForResourceInput, arg2 ...func(*amplify.Options)) (*amplify.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*amplify.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockAmplifyClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockAmplifyClient)(nil).ListTagsForResource), varargs...)
}

// ListWebhooks mocks base method.
func (m *MockAmplifyClient) ListWebhooks(arg0 context.Context, arg1 *amplify.ListWebhooksInput, arg2 ...func(*amplify.Options)) (*amplify.ListWebhooksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListWebhooks", varargs...)
	ret0, _ := ret[0].(*amplify.ListWebhooksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWebhooks indicates an expected call of ListWebhooks.
func (mr *MockAmplifyClientMockRecorder) ListWebhooks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWebhooks", reflect.TypeOf((*MockAmplifyClient)(nil).ListWebhooks), varargs...)
}
