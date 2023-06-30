// Code generated by MockGen. DO NOT EDIT.
// Source: codecommit.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	codecommit "github.com/aws/aws-sdk-go-v2/service/codecommit"
	gomock "github.com/golang/mock/gomock"
)

// MockCodecommitClient is a mock of CodecommitClient interface.
type MockCodecommitClient struct {
	ctrl     *gomock.Controller
	recorder *MockCodecommitClientMockRecorder
}

// MockCodecommitClientMockRecorder is the mock recorder for MockCodecommitClient.
type MockCodecommitClientMockRecorder struct {
	mock *MockCodecommitClient
}

// NewMockCodecommitClient creates a new mock instance.
func NewMockCodecommitClient(ctrl *gomock.Controller) *MockCodecommitClient {
	mock := &MockCodecommitClient{ctrl: ctrl}
	mock.recorder = &MockCodecommitClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCodecommitClient) EXPECT() *MockCodecommitClientMockRecorder {
	return m.recorder
}

// BatchGetCommits mocks base method.
func (m *MockCodecommitClient) BatchGetCommits(arg0 context.Context, arg1 *codecommit.BatchGetCommitsInput, arg2 ...func(*codecommit.Options)) (*codecommit.BatchGetCommitsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to BatchGetCommits")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetCommits", varargs...)
	ret0, _ := ret[0].(*codecommit.BatchGetCommitsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetCommits indicates an expected call of BatchGetCommits.
func (mr *MockCodecommitClientMockRecorder) BatchGetCommits(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetCommits", reflect.TypeOf((*MockCodecommitClient)(nil).BatchGetCommits), varargs...)
}

// BatchGetRepositories mocks base method.
func (m *MockCodecommitClient) BatchGetRepositories(arg0 context.Context, arg1 *codecommit.BatchGetRepositoriesInput, arg2 ...func(*codecommit.Options)) (*codecommit.BatchGetRepositoriesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to BatchGetRepositories")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetRepositories", varargs...)
	ret0, _ := ret[0].(*codecommit.BatchGetRepositoriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetRepositories indicates an expected call of BatchGetRepositories.
func (mr *MockCodecommitClientMockRecorder) BatchGetRepositories(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetRepositories", reflect.TypeOf((*MockCodecommitClient)(nil).BatchGetRepositories), varargs...)
}

// DescribeMergeConflicts mocks base method.
func (m *MockCodecommitClient) DescribeMergeConflicts(arg0 context.Context, arg1 *codecommit.DescribeMergeConflictsInput, arg2 ...func(*codecommit.Options)) (*codecommit.DescribeMergeConflictsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeMergeConflicts")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeMergeConflicts", varargs...)
	ret0, _ := ret[0].(*codecommit.DescribeMergeConflictsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeMergeConflicts indicates an expected call of DescribeMergeConflicts.
func (mr *MockCodecommitClientMockRecorder) DescribeMergeConflicts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMergeConflicts", reflect.TypeOf((*MockCodecommitClient)(nil).DescribeMergeConflicts), varargs...)
}

// DescribePullRequestEvents mocks base method.
func (m *MockCodecommitClient) DescribePullRequestEvents(arg0 context.Context, arg1 *codecommit.DescribePullRequestEventsInput, arg2 ...func(*codecommit.Options)) (*codecommit.DescribePullRequestEventsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribePullRequestEvents")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribePullRequestEvents", varargs...)
	ret0, _ := ret[0].(*codecommit.DescribePullRequestEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribePullRequestEvents indicates an expected call of DescribePullRequestEvents.
func (mr *MockCodecommitClientMockRecorder) DescribePullRequestEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribePullRequestEvents", reflect.TypeOf((*MockCodecommitClient)(nil).DescribePullRequestEvents), varargs...)
}

// GetApprovalRuleTemplate mocks base method.
func (m *MockCodecommitClient) GetApprovalRuleTemplate(arg0 context.Context, arg1 *codecommit.GetApprovalRuleTemplateInput, arg2 ...func(*codecommit.Options)) (*codecommit.GetApprovalRuleTemplateOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetApprovalRuleTemplate")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetApprovalRuleTemplate", varargs...)
	ret0, _ := ret[0].(*codecommit.GetApprovalRuleTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApprovalRuleTemplate indicates an expected call of GetApprovalRuleTemplate.
func (mr *MockCodecommitClientMockRecorder) GetApprovalRuleTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApprovalRuleTemplate", reflect.TypeOf((*MockCodecommitClient)(nil).GetApprovalRuleTemplate), varargs...)
}

// GetBlob mocks base method.
func (m *MockCodecommitClient) GetBlob(arg0 context.Context, arg1 *codecommit.GetBlobInput, arg2 ...func(*codecommit.Options)) (*codecommit.GetBlobOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetBlob")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBlob", varargs...)
	ret0, _ := ret[0].(*codecommit.GetBlobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlob indicates an expected call of GetBlob.
func (mr *MockCodecommitClientMockRecorder) GetBlob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlob", reflect.TypeOf((*MockCodecommitClient)(nil).GetBlob), varargs...)
}

// GetBranch mocks base method.
func (m *MockCodecommitClient) GetBranch(arg0 context.Context, arg1 *codecommit.GetBranchInput, arg2 ...func(*codecommit.Options)) (*codecommit.GetBranchOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetBranch")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBranch", varargs...)
	ret0, _ := ret[0].(*codecommit.GetBranchOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBranch indicates an expected call of GetBranch.
func (mr *MockCodecommitClientMockRecorder) GetBranch(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBranch", reflect.TypeOf((*MockCodecommitClient)(nil).GetBranch), varargs...)
}

// GetComment mocks base method.
func (m *MockCodecommitClient) GetComment(arg0 context.Context, arg1 *codecommit.GetCommentInput, arg2 ...func(*codecommit.Options)) (*codecommit.GetCommentOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetComment")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetComment", varargs...)
	ret0, _ := ret[0].(*codecommit.GetCommentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComment indicates an expected call of GetComment.
func (mr *MockCodecommitClientMockRecorder) GetComment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComment", reflect.TypeOf((*MockCodecommitClient)(nil).GetComment), varargs...)
}

// GetCommentReactions mocks base method.
func (m *MockCodecommitClient) GetCommentReactions(arg0 context.Context, arg1 *codecommit.GetCommentReactionsInput, arg2 ...func(*codecommit.Options)) (*codecommit.GetCommentReactionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetCommentReactions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCommentReactions", varargs...)
	ret0, _ := ret[0].(*codecommit.GetCommentReactionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentReactions indicates an expected call of GetCommentReactions.
func (mr *MockCodecommitClientMockRecorder) GetCommentReactions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentReactions", reflect.TypeOf((*MockCodecommitClient)(nil).GetCommentReactions), varargs...)
}

// GetCommentsForComparedCommit mocks base method.
func (m *MockCodecommitClient) GetCommentsForComparedCommit(arg0 context.Context, arg1 *codecommit.GetCommentsForComparedCommitInput, arg2 ...func(*codecommit.Options)) (*codecommit.GetCommentsForComparedCommitOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetCommentsForComparedCommit")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCommentsForComparedCommit", varargs...)
	ret0, _ := ret[0].(*codecommit.GetCommentsForComparedCommitOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentsForComparedCommit indicates an expected call of GetCommentsForComparedCommit.
func (mr *MockCodecommitClientMockRecorder) GetCommentsForComparedCommit(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentsForComparedCommit", reflect.TypeOf((*MockCodecommitClient)(nil).GetCommentsForComparedCommit), varargs...)
}

// GetCommentsForPullRequest mocks base method.
func (m *MockCodecommitClient) GetCommentsForPullRequest(arg0 context.Context, arg1 *codecommit.GetCommentsForPullRequestInput, arg2 ...func(*codecommit.Options)) (*codecommit.GetCommentsForPullRequestOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetCommentsForPullRequest")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCommentsForPullRequest", varargs...)
	ret0, _ := ret[0].(*codecommit.GetCommentsForPullRequestOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentsForPullRequest indicates an expected call of GetCommentsForPullRequest.
func (mr *MockCodecommitClientMockRecorder) GetCommentsForPullRequest(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentsForPullRequest", reflect.TypeOf((*MockCodecommitClient)(nil).GetCommentsForPullRequest), varargs...)
}

// GetCommit mocks base method.
func (m *MockCodecommitClient) GetCommit(arg0 context.Context, arg1 *codecommit.GetCommitInput, arg2 ...func(*codecommit.Options)) (*codecommit.GetCommitOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetCommit")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCommit", varargs...)
	ret0, _ := ret[0].(*codecommit.GetCommitOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommit indicates an expected call of GetCommit.
func (mr *MockCodecommitClientMockRecorder) GetCommit(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommit", reflect.TypeOf((*MockCodecommitClient)(nil).GetCommit), varargs...)
}

// GetDifferences mocks base method.
func (m *MockCodecommitClient) GetDifferences(arg0 context.Context, arg1 *codecommit.GetDifferencesInput, arg2 ...func(*codecommit.Options)) (*codecommit.GetDifferencesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetDifferences")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDifferences", varargs...)
	ret0, _ := ret[0].(*codecommit.GetDifferencesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDifferences indicates an expected call of GetDifferences.
func (mr *MockCodecommitClientMockRecorder) GetDifferences(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDifferences", reflect.TypeOf((*MockCodecommitClient)(nil).GetDifferences), varargs...)
}

// GetFile mocks base method.
func (m *MockCodecommitClient) GetFile(arg0 context.Context, arg1 *codecommit.GetFileInput, arg2 ...func(*codecommit.Options)) (*codecommit.GetFileOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetFile")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFile", varargs...)
	ret0, _ := ret[0].(*codecommit.GetFileOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile.
func (mr *MockCodecommitClientMockRecorder) GetFile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*MockCodecommitClient)(nil).GetFile), varargs...)
}

// GetFolder mocks base method.
func (m *MockCodecommitClient) GetFolder(arg0 context.Context, arg1 *codecommit.GetFolderInput, arg2 ...func(*codecommit.Options)) (*codecommit.GetFolderOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetFolder")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFolder", varargs...)
	ret0, _ := ret[0].(*codecommit.GetFolderOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFolder indicates an expected call of GetFolder.
func (mr *MockCodecommitClientMockRecorder) GetFolder(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFolder", reflect.TypeOf((*MockCodecommitClient)(nil).GetFolder), varargs...)
}

// GetMergeCommit mocks base method.
func (m *MockCodecommitClient) GetMergeCommit(arg0 context.Context, arg1 *codecommit.GetMergeCommitInput, arg2 ...func(*codecommit.Options)) (*codecommit.GetMergeCommitOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetMergeCommit")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMergeCommit", varargs...)
	ret0, _ := ret[0].(*codecommit.GetMergeCommitOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMergeCommit indicates an expected call of GetMergeCommit.
func (mr *MockCodecommitClientMockRecorder) GetMergeCommit(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeCommit", reflect.TypeOf((*MockCodecommitClient)(nil).GetMergeCommit), varargs...)
}

// GetMergeConflicts mocks base method.
func (m *MockCodecommitClient) GetMergeConflicts(arg0 context.Context, arg1 *codecommit.GetMergeConflictsInput, arg2 ...func(*codecommit.Options)) (*codecommit.GetMergeConflictsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetMergeConflicts")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMergeConflicts", varargs...)
	ret0, _ := ret[0].(*codecommit.GetMergeConflictsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMergeConflicts indicates an expected call of GetMergeConflicts.
func (mr *MockCodecommitClientMockRecorder) GetMergeConflicts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeConflicts", reflect.TypeOf((*MockCodecommitClient)(nil).GetMergeConflicts), varargs...)
}

// GetMergeOptions mocks base method.
func (m *MockCodecommitClient) GetMergeOptions(arg0 context.Context, arg1 *codecommit.GetMergeOptionsInput, arg2 ...func(*codecommit.Options)) (*codecommit.GetMergeOptionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetMergeOptions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMergeOptions", varargs...)
	ret0, _ := ret[0].(*codecommit.GetMergeOptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMergeOptions indicates an expected call of GetMergeOptions.
func (mr *MockCodecommitClientMockRecorder) GetMergeOptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeOptions", reflect.TypeOf((*MockCodecommitClient)(nil).GetMergeOptions), varargs...)
}

// GetPullRequest mocks base method.
func (m *MockCodecommitClient) GetPullRequest(arg0 context.Context, arg1 *codecommit.GetPullRequestInput, arg2 ...func(*codecommit.Options)) (*codecommit.GetPullRequestOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetPullRequest")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPullRequest", varargs...)
	ret0, _ := ret[0].(*codecommit.GetPullRequestOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPullRequest indicates an expected call of GetPullRequest.
func (mr *MockCodecommitClientMockRecorder) GetPullRequest(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPullRequest", reflect.TypeOf((*MockCodecommitClient)(nil).GetPullRequest), varargs...)
}

// GetPullRequestApprovalStates mocks base method.
func (m *MockCodecommitClient) GetPullRequestApprovalStates(arg0 context.Context, arg1 *codecommit.GetPullRequestApprovalStatesInput, arg2 ...func(*codecommit.Options)) (*codecommit.GetPullRequestApprovalStatesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetPullRequestApprovalStates")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPullRequestApprovalStates", varargs...)
	ret0, _ := ret[0].(*codecommit.GetPullRequestApprovalStatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPullRequestApprovalStates indicates an expected call of GetPullRequestApprovalStates.
func (mr *MockCodecommitClientMockRecorder) GetPullRequestApprovalStates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPullRequestApprovalStates", reflect.TypeOf((*MockCodecommitClient)(nil).GetPullRequestApprovalStates), varargs...)
}

// GetPullRequestOverrideState mocks base method.
func (m *MockCodecommitClient) GetPullRequestOverrideState(arg0 context.Context, arg1 *codecommit.GetPullRequestOverrideStateInput, arg2 ...func(*codecommit.Options)) (*codecommit.GetPullRequestOverrideStateOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetPullRequestOverrideState")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPullRequestOverrideState", varargs...)
	ret0, _ := ret[0].(*codecommit.GetPullRequestOverrideStateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPullRequestOverrideState indicates an expected call of GetPullRequestOverrideState.
func (mr *MockCodecommitClientMockRecorder) GetPullRequestOverrideState(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPullRequestOverrideState", reflect.TypeOf((*MockCodecommitClient)(nil).GetPullRequestOverrideState), varargs...)
}

// GetRepository mocks base method.
func (m *MockCodecommitClient) GetRepository(arg0 context.Context, arg1 *codecommit.GetRepositoryInput, arg2 ...func(*codecommit.Options)) (*codecommit.GetRepositoryOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetRepository")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRepository", varargs...)
	ret0, _ := ret[0].(*codecommit.GetRepositoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepository indicates an expected call of GetRepository.
func (mr *MockCodecommitClientMockRecorder) GetRepository(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepository", reflect.TypeOf((*MockCodecommitClient)(nil).GetRepository), varargs...)
}

// GetRepositoryTriggers mocks base method.
func (m *MockCodecommitClient) GetRepositoryTriggers(arg0 context.Context, arg1 *codecommit.GetRepositoryTriggersInput, arg2 ...func(*codecommit.Options)) (*codecommit.GetRepositoryTriggersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetRepositoryTriggers")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRepositoryTriggers", varargs...)
	ret0, _ := ret[0].(*codecommit.GetRepositoryTriggersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoryTriggers indicates an expected call of GetRepositoryTriggers.
func (mr *MockCodecommitClientMockRecorder) GetRepositoryTriggers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositoryTriggers", reflect.TypeOf((*MockCodecommitClient)(nil).GetRepositoryTriggers), varargs...)
}

// ListApprovalRuleTemplates mocks base method.
func (m *MockCodecommitClient) ListApprovalRuleTemplates(arg0 context.Context, arg1 *codecommit.ListApprovalRuleTemplatesInput, arg2 ...func(*codecommit.Options)) (*codecommit.ListApprovalRuleTemplatesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListApprovalRuleTemplates")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListApprovalRuleTemplates", varargs...)
	ret0, _ := ret[0].(*codecommit.ListApprovalRuleTemplatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApprovalRuleTemplates indicates an expected call of ListApprovalRuleTemplates.
func (mr *MockCodecommitClientMockRecorder) ListApprovalRuleTemplates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApprovalRuleTemplates", reflect.TypeOf((*MockCodecommitClient)(nil).ListApprovalRuleTemplates), varargs...)
}

// ListAssociatedApprovalRuleTemplatesForRepository mocks base method.
func (m *MockCodecommitClient) ListAssociatedApprovalRuleTemplatesForRepository(arg0 context.Context, arg1 *codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryInput, arg2 ...func(*codecommit.Options)) (*codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAssociatedApprovalRuleTemplatesForRepository")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAssociatedApprovalRuleTemplatesForRepository", varargs...)
	ret0, _ := ret[0].(*codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAssociatedApprovalRuleTemplatesForRepository indicates an expected call of ListAssociatedApprovalRuleTemplatesForRepository.
func (mr *MockCodecommitClientMockRecorder) ListAssociatedApprovalRuleTemplatesForRepository(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAssociatedApprovalRuleTemplatesForRepository", reflect.TypeOf((*MockCodecommitClient)(nil).ListAssociatedApprovalRuleTemplatesForRepository), varargs...)
}

// ListBranches mocks base method.
func (m *MockCodecommitClient) ListBranches(arg0 context.Context, arg1 *codecommit.ListBranchesInput, arg2 ...func(*codecommit.Options)) (*codecommit.ListBranchesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListBranches")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBranches", varargs...)
	ret0, _ := ret[0].(*codecommit.ListBranchesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBranches indicates an expected call of ListBranches.
func (mr *MockCodecommitClientMockRecorder) ListBranches(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBranches", reflect.TypeOf((*MockCodecommitClient)(nil).ListBranches), varargs...)
}

// ListPullRequests mocks base method.
func (m *MockCodecommitClient) ListPullRequests(arg0 context.Context, arg1 *codecommit.ListPullRequestsInput, arg2 ...func(*codecommit.Options)) (*codecommit.ListPullRequestsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListPullRequests")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPullRequests", varargs...)
	ret0, _ := ret[0].(*codecommit.ListPullRequestsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPullRequests indicates an expected call of ListPullRequests.
func (mr *MockCodecommitClientMockRecorder) ListPullRequests(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPullRequests", reflect.TypeOf((*MockCodecommitClient)(nil).ListPullRequests), varargs...)
}

// ListRepositories mocks base method.
func (m *MockCodecommitClient) ListRepositories(arg0 context.Context, arg1 *codecommit.ListRepositoriesInput, arg2 ...func(*codecommit.Options)) (*codecommit.ListRepositoriesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListRepositories")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRepositories", varargs...)
	ret0, _ := ret[0].(*codecommit.ListRepositoriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRepositories indicates an expected call of ListRepositories.
func (mr *MockCodecommitClientMockRecorder) ListRepositories(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRepositories", reflect.TypeOf((*MockCodecommitClient)(nil).ListRepositories), varargs...)
}

// ListRepositoriesForApprovalRuleTemplate mocks base method.
func (m *MockCodecommitClient) ListRepositoriesForApprovalRuleTemplate(arg0 context.Context, arg1 *codecommit.ListRepositoriesForApprovalRuleTemplateInput, arg2 ...func(*codecommit.Options)) (*codecommit.ListRepositoriesForApprovalRuleTemplateOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListRepositoriesForApprovalRuleTemplate")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRepositoriesForApprovalRuleTemplate", varargs...)
	ret0, _ := ret[0].(*codecommit.ListRepositoriesForApprovalRuleTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRepositoriesForApprovalRuleTemplate indicates an expected call of ListRepositoriesForApprovalRuleTemplate.
func (mr *MockCodecommitClientMockRecorder) ListRepositoriesForApprovalRuleTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRepositoriesForApprovalRuleTemplate", reflect.TypeOf((*MockCodecommitClient)(nil).ListRepositoriesForApprovalRuleTemplate), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockCodecommitClient) ListTagsForResource(arg0 context.Context, arg1 *codecommit.ListTagsForResourceInput, arg2 ...func(*codecommit.Options)) (*codecommit.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &codecommit.Options{}
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
	ret0, _ := ret[0].(*codecommit.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockCodecommitClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockCodecommitClient)(nil).ListTagsForResource), varargs...)
}
