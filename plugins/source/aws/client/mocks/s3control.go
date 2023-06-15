// Code generated by MockGen. DO NOT EDIT.
// Source: s3control.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	s3control "github.com/aws/aws-sdk-go-v2/service/s3control"
	gomock "github.com/golang/mock/gomock"
)

// MockS3controlClient is a mock of S3controlClient interface.
type MockS3controlClient struct {
	ctrl     *gomock.Controller
	recorder *MockS3controlClientMockRecorder
}

// MockS3controlClientMockRecorder is the mock recorder for MockS3controlClient.
type MockS3controlClientMockRecorder struct {
	mock *MockS3controlClient
}

// NewMockS3controlClient creates a new mock instance.
func NewMockS3controlClient(ctrl *gomock.Controller) *MockS3controlClient {
	mock := &MockS3controlClient{ctrl: ctrl}
	mock.recorder = &MockS3controlClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3controlClient) EXPECT() *MockS3controlClientMockRecorder {
	return m.recorder
}

// DescribeJob mocks base method.
func (m *MockS3controlClient) DescribeJob(arg0 context.Context, arg1 *s3control.DescribeJobInput, arg2 ...func(*s3control.Options)) (*s3control.DescribeJobOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeJob")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeJob", varargs...)
	ret0, _ := ret[0].(*s3control.DescribeJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeJob indicates an expected call of DescribeJob.
func (mr *MockS3controlClientMockRecorder) DescribeJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeJob", reflect.TypeOf((*MockS3controlClient)(nil).DescribeJob), varargs...)
}

// DescribeMultiRegionAccessPointOperation mocks base method.
func (m *MockS3controlClient) DescribeMultiRegionAccessPointOperation(arg0 context.Context, arg1 *s3control.DescribeMultiRegionAccessPointOperationInput, arg2 ...func(*s3control.Options)) (*s3control.DescribeMultiRegionAccessPointOperationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeMultiRegionAccessPointOperation")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeMultiRegionAccessPointOperation", varargs...)
	ret0, _ := ret[0].(*s3control.DescribeMultiRegionAccessPointOperationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeMultiRegionAccessPointOperation indicates an expected call of DescribeMultiRegionAccessPointOperation.
func (mr *MockS3controlClientMockRecorder) DescribeMultiRegionAccessPointOperation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMultiRegionAccessPointOperation", reflect.TypeOf((*MockS3controlClient)(nil).DescribeMultiRegionAccessPointOperation), varargs...)
}

// GetAccessPoint mocks base method.
func (m *MockS3controlClient) GetAccessPoint(arg0 context.Context, arg1 *s3control.GetAccessPointInput, arg2 ...func(*s3control.Options)) (*s3control.GetAccessPointOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetAccessPoint")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccessPoint", varargs...)
	ret0, _ := ret[0].(*s3control.GetAccessPointOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessPoint indicates an expected call of GetAccessPoint.
func (mr *MockS3controlClientMockRecorder) GetAccessPoint(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessPoint", reflect.TypeOf((*MockS3controlClient)(nil).GetAccessPoint), varargs...)
}

// GetAccessPointConfigurationForObjectLambda mocks base method.
func (m *MockS3controlClient) GetAccessPointConfigurationForObjectLambda(arg0 context.Context, arg1 *s3control.GetAccessPointConfigurationForObjectLambdaInput, arg2 ...func(*s3control.Options)) (*s3control.GetAccessPointConfigurationForObjectLambdaOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetAccessPointConfigurationForObjectLambda")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccessPointConfigurationForObjectLambda", varargs...)
	ret0, _ := ret[0].(*s3control.GetAccessPointConfigurationForObjectLambdaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessPointConfigurationForObjectLambda indicates an expected call of GetAccessPointConfigurationForObjectLambda.
func (mr *MockS3controlClientMockRecorder) GetAccessPointConfigurationForObjectLambda(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessPointConfigurationForObjectLambda", reflect.TypeOf((*MockS3controlClient)(nil).GetAccessPointConfigurationForObjectLambda), varargs...)
}

// GetAccessPointForObjectLambda mocks base method.
func (m *MockS3controlClient) GetAccessPointForObjectLambda(arg0 context.Context, arg1 *s3control.GetAccessPointForObjectLambdaInput, arg2 ...func(*s3control.Options)) (*s3control.GetAccessPointForObjectLambdaOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetAccessPointForObjectLambda")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccessPointForObjectLambda", varargs...)
	ret0, _ := ret[0].(*s3control.GetAccessPointForObjectLambdaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessPointForObjectLambda indicates an expected call of GetAccessPointForObjectLambda.
func (mr *MockS3controlClientMockRecorder) GetAccessPointForObjectLambda(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessPointForObjectLambda", reflect.TypeOf((*MockS3controlClient)(nil).GetAccessPointForObjectLambda), varargs...)
}

// GetAccessPointPolicy mocks base method.
func (m *MockS3controlClient) GetAccessPointPolicy(arg0 context.Context, arg1 *s3control.GetAccessPointPolicyInput, arg2 ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetAccessPointPolicy")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccessPointPolicy", varargs...)
	ret0, _ := ret[0].(*s3control.GetAccessPointPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessPointPolicy indicates an expected call of GetAccessPointPolicy.
func (mr *MockS3controlClientMockRecorder) GetAccessPointPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessPointPolicy", reflect.TypeOf((*MockS3controlClient)(nil).GetAccessPointPolicy), varargs...)
}

// GetAccessPointPolicyForObjectLambda mocks base method.
func (m *MockS3controlClient) GetAccessPointPolicyForObjectLambda(arg0 context.Context, arg1 *s3control.GetAccessPointPolicyForObjectLambdaInput, arg2 ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyForObjectLambdaOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetAccessPointPolicyForObjectLambda")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccessPointPolicyForObjectLambda", varargs...)
	ret0, _ := ret[0].(*s3control.GetAccessPointPolicyForObjectLambdaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessPointPolicyForObjectLambda indicates an expected call of GetAccessPointPolicyForObjectLambda.
func (mr *MockS3controlClientMockRecorder) GetAccessPointPolicyForObjectLambda(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessPointPolicyForObjectLambda", reflect.TypeOf((*MockS3controlClient)(nil).GetAccessPointPolicyForObjectLambda), varargs...)
}

// GetAccessPointPolicyStatus mocks base method.
func (m *MockS3controlClient) GetAccessPointPolicyStatus(arg0 context.Context, arg1 *s3control.GetAccessPointPolicyStatusInput, arg2 ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyStatusOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetAccessPointPolicyStatus")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccessPointPolicyStatus", varargs...)
	ret0, _ := ret[0].(*s3control.GetAccessPointPolicyStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessPointPolicyStatus indicates an expected call of GetAccessPointPolicyStatus.
func (mr *MockS3controlClientMockRecorder) GetAccessPointPolicyStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessPointPolicyStatus", reflect.TypeOf((*MockS3controlClient)(nil).GetAccessPointPolicyStatus), varargs...)
}

// GetAccessPointPolicyStatusForObjectLambda mocks base method.
func (m *MockS3controlClient) GetAccessPointPolicyStatusForObjectLambda(arg0 context.Context, arg1 *s3control.GetAccessPointPolicyStatusForObjectLambdaInput, arg2 ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyStatusForObjectLambdaOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetAccessPointPolicyStatusForObjectLambda")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccessPointPolicyStatusForObjectLambda", varargs...)
	ret0, _ := ret[0].(*s3control.GetAccessPointPolicyStatusForObjectLambdaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessPointPolicyStatusForObjectLambda indicates an expected call of GetAccessPointPolicyStatusForObjectLambda.
func (mr *MockS3controlClientMockRecorder) GetAccessPointPolicyStatusForObjectLambda(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessPointPolicyStatusForObjectLambda", reflect.TypeOf((*MockS3controlClient)(nil).GetAccessPointPolicyStatusForObjectLambda), varargs...)
}

// GetBucket mocks base method.
func (m *MockS3controlClient) GetBucket(arg0 context.Context, arg1 *s3control.GetBucketInput, arg2 ...func(*s3control.Options)) (*s3control.GetBucketOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetBucket")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucket", varargs...)
	ret0, _ := ret[0].(*s3control.GetBucketOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucket indicates an expected call of GetBucket.
func (mr *MockS3controlClientMockRecorder) GetBucket(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucket", reflect.TypeOf((*MockS3controlClient)(nil).GetBucket), varargs...)
}

// GetBucketLifecycleConfiguration mocks base method.
func (m *MockS3controlClient) GetBucketLifecycleConfiguration(arg0 context.Context, arg1 *s3control.GetBucketLifecycleConfigurationInput, arg2 ...func(*s3control.Options)) (*s3control.GetBucketLifecycleConfigurationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetBucketLifecycleConfiguration")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketLifecycleConfiguration", varargs...)
	ret0, _ := ret[0].(*s3control.GetBucketLifecycleConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketLifecycleConfiguration indicates an expected call of GetBucketLifecycleConfiguration.
func (mr *MockS3controlClientMockRecorder) GetBucketLifecycleConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketLifecycleConfiguration", reflect.TypeOf((*MockS3controlClient)(nil).GetBucketLifecycleConfiguration), varargs...)
}

// GetBucketPolicy mocks base method.
func (m *MockS3controlClient) GetBucketPolicy(arg0 context.Context, arg1 *s3control.GetBucketPolicyInput, arg2 ...func(*s3control.Options)) (*s3control.GetBucketPolicyOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetBucketPolicy")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketPolicy", varargs...)
	ret0, _ := ret[0].(*s3control.GetBucketPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketPolicy indicates an expected call of GetBucketPolicy.
func (mr *MockS3controlClientMockRecorder) GetBucketPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketPolicy", reflect.TypeOf((*MockS3controlClient)(nil).GetBucketPolicy), varargs...)
}

// GetBucketReplication mocks base method.
func (m *MockS3controlClient) GetBucketReplication(arg0 context.Context, arg1 *s3control.GetBucketReplicationInput, arg2 ...func(*s3control.Options)) (*s3control.GetBucketReplicationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetBucketReplication")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketReplication", varargs...)
	ret0, _ := ret[0].(*s3control.GetBucketReplicationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketReplication indicates an expected call of GetBucketReplication.
func (mr *MockS3controlClientMockRecorder) GetBucketReplication(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketReplication", reflect.TypeOf((*MockS3controlClient)(nil).GetBucketReplication), varargs...)
}

// GetBucketTagging mocks base method.
func (m *MockS3controlClient) GetBucketTagging(arg0 context.Context, arg1 *s3control.GetBucketTaggingInput, arg2 ...func(*s3control.Options)) (*s3control.GetBucketTaggingOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetBucketTagging")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketTagging", varargs...)
	ret0, _ := ret[0].(*s3control.GetBucketTaggingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketTagging indicates an expected call of GetBucketTagging.
func (mr *MockS3controlClientMockRecorder) GetBucketTagging(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketTagging", reflect.TypeOf((*MockS3controlClient)(nil).GetBucketTagging), varargs...)
}

// GetBucketVersioning mocks base method.
func (m *MockS3controlClient) GetBucketVersioning(arg0 context.Context, arg1 *s3control.GetBucketVersioningInput, arg2 ...func(*s3control.Options)) (*s3control.GetBucketVersioningOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetBucketVersioning")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketVersioning", varargs...)
	ret0, _ := ret[0].(*s3control.GetBucketVersioningOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketVersioning indicates an expected call of GetBucketVersioning.
func (mr *MockS3controlClientMockRecorder) GetBucketVersioning(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketVersioning", reflect.TypeOf((*MockS3controlClient)(nil).GetBucketVersioning), varargs...)
}

// GetJobTagging mocks base method.
func (m *MockS3controlClient) GetJobTagging(arg0 context.Context, arg1 *s3control.GetJobTaggingInput, arg2 ...func(*s3control.Options)) (*s3control.GetJobTaggingOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetJobTagging")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetJobTagging", varargs...)
	ret0, _ := ret[0].(*s3control.GetJobTaggingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobTagging indicates an expected call of GetJobTagging.
func (mr *MockS3controlClientMockRecorder) GetJobTagging(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobTagging", reflect.TypeOf((*MockS3controlClient)(nil).GetJobTagging), varargs...)
}

// GetMultiRegionAccessPoint mocks base method.
func (m *MockS3controlClient) GetMultiRegionAccessPoint(arg0 context.Context, arg1 *s3control.GetMultiRegionAccessPointInput, arg2 ...func(*s3control.Options)) (*s3control.GetMultiRegionAccessPointOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetMultiRegionAccessPoint")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMultiRegionAccessPoint", varargs...)
	ret0, _ := ret[0].(*s3control.GetMultiRegionAccessPointOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMultiRegionAccessPoint indicates an expected call of GetMultiRegionAccessPoint.
func (mr *MockS3controlClientMockRecorder) GetMultiRegionAccessPoint(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMultiRegionAccessPoint", reflect.TypeOf((*MockS3controlClient)(nil).GetMultiRegionAccessPoint), varargs...)
}

// GetMultiRegionAccessPointPolicy mocks base method.
func (m *MockS3controlClient) GetMultiRegionAccessPointPolicy(arg0 context.Context, arg1 *s3control.GetMultiRegionAccessPointPolicyInput, arg2 ...func(*s3control.Options)) (*s3control.GetMultiRegionAccessPointPolicyOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetMultiRegionAccessPointPolicy")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMultiRegionAccessPointPolicy", varargs...)
	ret0, _ := ret[0].(*s3control.GetMultiRegionAccessPointPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMultiRegionAccessPointPolicy indicates an expected call of GetMultiRegionAccessPointPolicy.
func (mr *MockS3controlClientMockRecorder) GetMultiRegionAccessPointPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMultiRegionAccessPointPolicy", reflect.TypeOf((*MockS3controlClient)(nil).GetMultiRegionAccessPointPolicy), varargs...)
}

// GetMultiRegionAccessPointPolicyStatus mocks base method.
func (m *MockS3controlClient) GetMultiRegionAccessPointPolicyStatus(arg0 context.Context, arg1 *s3control.GetMultiRegionAccessPointPolicyStatusInput, arg2 ...func(*s3control.Options)) (*s3control.GetMultiRegionAccessPointPolicyStatusOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetMultiRegionAccessPointPolicyStatus")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMultiRegionAccessPointPolicyStatus", varargs...)
	ret0, _ := ret[0].(*s3control.GetMultiRegionAccessPointPolicyStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMultiRegionAccessPointPolicyStatus indicates an expected call of GetMultiRegionAccessPointPolicyStatus.
func (mr *MockS3controlClientMockRecorder) GetMultiRegionAccessPointPolicyStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMultiRegionAccessPointPolicyStatus", reflect.TypeOf((*MockS3controlClient)(nil).GetMultiRegionAccessPointPolicyStatus), varargs...)
}

// GetMultiRegionAccessPointRoutes mocks base method.
func (m *MockS3controlClient) GetMultiRegionAccessPointRoutes(arg0 context.Context, arg1 *s3control.GetMultiRegionAccessPointRoutesInput, arg2 ...func(*s3control.Options)) (*s3control.GetMultiRegionAccessPointRoutesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetMultiRegionAccessPointRoutes")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMultiRegionAccessPointRoutes", varargs...)
	ret0, _ := ret[0].(*s3control.GetMultiRegionAccessPointRoutesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMultiRegionAccessPointRoutes indicates an expected call of GetMultiRegionAccessPointRoutes.
func (mr *MockS3controlClientMockRecorder) GetMultiRegionAccessPointRoutes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMultiRegionAccessPointRoutes", reflect.TypeOf((*MockS3controlClient)(nil).GetMultiRegionAccessPointRoutes), varargs...)
}

// GetPublicAccessBlock mocks base method.
func (m *MockS3controlClient) GetPublicAccessBlock(arg0 context.Context, arg1 *s3control.GetPublicAccessBlockInput, arg2 ...func(*s3control.Options)) (*s3control.GetPublicAccessBlockOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetPublicAccessBlock")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPublicAccessBlock", varargs...)
	ret0, _ := ret[0].(*s3control.GetPublicAccessBlockOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicAccessBlock indicates an expected call of GetPublicAccessBlock.
func (mr *MockS3controlClientMockRecorder) GetPublicAccessBlock(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicAccessBlock", reflect.TypeOf((*MockS3controlClient)(nil).GetPublicAccessBlock), varargs...)
}

// GetStorageLensConfiguration mocks base method.
func (m *MockS3controlClient) GetStorageLensConfiguration(arg0 context.Context, arg1 *s3control.GetStorageLensConfigurationInput, arg2 ...func(*s3control.Options)) (*s3control.GetStorageLensConfigurationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetStorageLensConfiguration")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStorageLensConfiguration", varargs...)
	ret0, _ := ret[0].(*s3control.GetStorageLensConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorageLensConfiguration indicates an expected call of GetStorageLensConfiguration.
func (mr *MockS3controlClientMockRecorder) GetStorageLensConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageLensConfiguration", reflect.TypeOf((*MockS3controlClient)(nil).GetStorageLensConfiguration), varargs...)
}

// GetStorageLensConfigurationTagging mocks base method.
func (m *MockS3controlClient) GetStorageLensConfigurationTagging(arg0 context.Context, arg1 *s3control.GetStorageLensConfigurationTaggingInput, arg2 ...func(*s3control.Options)) (*s3control.GetStorageLensConfigurationTaggingOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetStorageLensConfigurationTagging")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStorageLensConfigurationTagging", varargs...)
	ret0, _ := ret[0].(*s3control.GetStorageLensConfigurationTaggingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorageLensConfigurationTagging indicates an expected call of GetStorageLensConfigurationTagging.
func (mr *MockS3controlClientMockRecorder) GetStorageLensConfigurationTagging(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageLensConfigurationTagging", reflect.TypeOf((*MockS3controlClient)(nil).GetStorageLensConfigurationTagging), varargs...)
}

// ListAccessPoints mocks base method.
func (m *MockS3controlClient) ListAccessPoints(arg0 context.Context, arg1 *s3control.ListAccessPointsInput, arg2 ...func(*s3control.Options)) (*s3control.ListAccessPointsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAccessPoints")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccessPoints", varargs...)
	ret0, _ := ret[0].(*s3control.ListAccessPointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccessPoints indicates an expected call of ListAccessPoints.
func (mr *MockS3controlClientMockRecorder) ListAccessPoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessPoints", reflect.TypeOf((*MockS3controlClient)(nil).ListAccessPoints), varargs...)
}

// ListAccessPointsForObjectLambda mocks base method.
func (m *MockS3controlClient) ListAccessPointsForObjectLambda(arg0 context.Context, arg1 *s3control.ListAccessPointsForObjectLambdaInput, arg2 ...func(*s3control.Options)) (*s3control.ListAccessPointsForObjectLambdaOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAccessPointsForObjectLambda")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccessPointsForObjectLambda", varargs...)
	ret0, _ := ret[0].(*s3control.ListAccessPointsForObjectLambdaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccessPointsForObjectLambda indicates an expected call of ListAccessPointsForObjectLambda.
func (mr *MockS3controlClientMockRecorder) ListAccessPointsForObjectLambda(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessPointsForObjectLambda", reflect.TypeOf((*MockS3controlClient)(nil).ListAccessPointsForObjectLambda), varargs...)
}

// ListJobs mocks base method.
func (m *MockS3controlClient) ListJobs(arg0 context.Context, arg1 *s3control.ListJobsInput, arg2 ...func(*s3control.Options)) (*s3control.ListJobsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListJobs")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListJobs", varargs...)
	ret0, _ := ret[0].(*s3control.ListJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListJobs indicates an expected call of ListJobs.
func (mr *MockS3controlClientMockRecorder) ListJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListJobs", reflect.TypeOf((*MockS3controlClient)(nil).ListJobs), varargs...)
}

// ListMultiRegionAccessPoints mocks base method.
func (m *MockS3controlClient) ListMultiRegionAccessPoints(arg0 context.Context, arg1 *s3control.ListMultiRegionAccessPointsInput, arg2 ...func(*s3control.Options)) (*s3control.ListMultiRegionAccessPointsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListMultiRegionAccessPoints")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMultiRegionAccessPoints", varargs...)
	ret0, _ := ret[0].(*s3control.ListMultiRegionAccessPointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMultiRegionAccessPoints indicates an expected call of ListMultiRegionAccessPoints.
func (mr *MockS3controlClientMockRecorder) ListMultiRegionAccessPoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMultiRegionAccessPoints", reflect.TypeOf((*MockS3controlClient)(nil).ListMultiRegionAccessPoints), varargs...)
}

// ListRegionalBuckets mocks base method.
func (m *MockS3controlClient) ListRegionalBuckets(arg0 context.Context, arg1 *s3control.ListRegionalBucketsInput, arg2 ...func(*s3control.Options)) (*s3control.ListRegionalBucketsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListRegionalBuckets")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRegionalBuckets", varargs...)
	ret0, _ := ret[0].(*s3control.ListRegionalBucketsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRegionalBuckets indicates an expected call of ListRegionalBuckets.
func (mr *MockS3controlClientMockRecorder) ListRegionalBuckets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRegionalBuckets", reflect.TypeOf((*MockS3controlClient)(nil).ListRegionalBuckets), varargs...)
}

// ListStorageLensConfigurations mocks base method.
func (m *MockS3controlClient) ListStorageLensConfigurations(arg0 context.Context, arg1 *s3control.ListStorageLensConfigurationsInput, arg2 ...func(*s3control.Options)) (*s3control.ListStorageLensConfigurationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &s3control.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListStorageLensConfigurations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListStorageLensConfigurations", varargs...)
	ret0, _ := ret[0].(*s3control.ListStorageLensConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStorageLensConfigurations indicates an expected call of ListStorageLensConfigurations.
func (mr *MockS3controlClientMockRecorder) ListStorageLensConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStorageLensConfigurations", reflect.TypeOf((*MockS3controlClient)(nil).ListStorageLensConfigurations), varargs...)
}
