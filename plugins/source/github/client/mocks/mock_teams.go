// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/github/client (interfaces: TeamsService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v45/github"
)

// MockTeamsService is a mock of TeamsService interface.
type MockTeamsService struct {
	ctrl     *gomock.Controller
	recorder *MockTeamsServiceMockRecorder
}

// MockTeamsServiceMockRecorder is the mock recorder for MockTeamsService.
type MockTeamsServiceMockRecorder struct {
	mock *MockTeamsService
}

// NewMockTeamsService creates a new mock instance.
func NewMockTeamsService(ctrl *gomock.Controller) *MockTeamsService {
	mock := &MockTeamsService{ctrl: ctrl}
	mock.recorder = &MockTeamsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamsService) EXPECT() *MockTeamsServiceMockRecorder {
	return m.recorder
}

// ListExternalGroups mocks base method.
func (m *MockTeamsService) ListExternalGroups(arg0 context.Context, arg1 string, arg2 *github.ListExternalGroupsOptions) (*github.ExternalGroupList, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListExternalGroups", arg0, arg1, arg2)
	ret0, _ := ret[0].(*github.ExternalGroupList)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListExternalGroups indicates an expected call of ListExternalGroups.
func (mr *MockTeamsServiceMockRecorder) ListExternalGroups(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListExternalGroups", reflect.TypeOf((*MockTeamsService)(nil).ListExternalGroups), arg0, arg1, arg2)
}

// ListTeamMembersByID mocks base method.
func (m *MockTeamsService) ListTeamMembersByID(arg0 context.Context, arg1, arg2 int64, arg3 *github.TeamListTeamMembersOptions) ([]*github.User, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTeamMembersByID", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*github.User)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListTeamMembersByID indicates an expected call of ListTeamMembersByID.
func (mr *MockTeamsServiceMockRecorder) ListTeamMembersByID(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTeamMembersByID", reflect.TypeOf((*MockTeamsService)(nil).ListTeamMembersByID), arg0, arg1, arg2, arg3)
}

// ListTeamReposByID mocks base method.
func (m *MockTeamsService) ListTeamReposByID(arg0 context.Context, arg1, arg2 int64, arg3 *github.ListOptions) ([]*github.Repository, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTeamReposByID", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*github.Repository)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListTeamReposByID indicates an expected call of ListTeamReposByID.
func (mr *MockTeamsServiceMockRecorder) ListTeamReposByID(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTeamReposByID", reflect.TypeOf((*MockTeamsService)(nil).ListTeamReposByID), arg0, arg1, arg2, arg3)
}

// ListTeams mocks base method.
func (m *MockTeamsService) ListTeams(arg0 context.Context, arg1 string, arg2 *github.ListOptions) ([]*github.Team, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTeams", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*github.Team)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListTeams indicates an expected call of ListTeams.
func (mr *MockTeamsServiceMockRecorder) ListTeams(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTeams", reflect.TypeOf((*MockTeamsService)(nil).ListTeams), arg0, arg1, arg2)
}
