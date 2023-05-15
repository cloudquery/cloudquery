// Code generated by MockGen. DO NOT EDIT.
// Source: dynamodbstreams.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	
	reflect "reflect"

	dynamodbstreams "github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
	gomock "github.com/golang/mock/gomock"
)

// MockDynamodbstreamsClient is a mock of DynamodbstreamsClient interface.
type MockDynamodbstreamsClient struct {
	ctrl     *gomock.Controller
	recorder *MockDynamodbstreamsClientMockRecorder
}

// MockDynamodbstreamsClientMockRecorder is the mock recorder for MockDynamodbstreamsClient.
type MockDynamodbstreamsClientMockRecorder struct {
	mock *MockDynamodbstreamsClient
}

// NewMockDynamodbstreamsClient creates a new mock instance.
func NewMockDynamodbstreamsClient(ctrl *gomock.Controller) *MockDynamodbstreamsClient {
	mock := &MockDynamodbstreamsClient{ctrl: ctrl}
	mock.recorder = &MockDynamodbstreamsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDynamodbstreamsClient) EXPECT() *MockDynamodbstreamsClientMockRecorder {
	return m.recorder
}

// DescribeStream mocks base method.
func (m *MockDynamodbstreamsClient) DescribeStream(arg0 context.Context, arg1 *dynamodbstreams.DescribeStreamInput, arg2 ...func(*dynamodbstreams.Options)) (*dynamodbstreams.DescribeStreamOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &dynamodbstreams.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeStream")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeStream", varargs...)
	ret0, _ := ret[0].(*dynamodbstreams.DescribeStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeStream indicates an expected call of DescribeStream.
func (mr *MockDynamodbstreamsClientMockRecorder) DescribeStream(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStream", reflect.TypeOf((*MockDynamodbstreamsClient)(nil).DescribeStream), varargs...)
}

// GetRecords mocks base method.
func (m *MockDynamodbstreamsClient) GetRecords(arg0 context.Context, arg1 *dynamodbstreams.GetRecordsInput, arg2 ...func(*dynamodbstreams.Options)) (*dynamodbstreams.GetRecordsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &dynamodbstreams.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetRecords")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRecords", varargs...)
	ret0, _ := ret[0].(*dynamodbstreams.GetRecordsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecords indicates an expected call of GetRecords.
func (mr *MockDynamodbstreamsClientMockRecorder) GetRecords(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecords", reflect.TypeOf((*MockDynamodbstreamsClient)(nil).GetRecords), varargs...)
}

// GetShardIterator mocks base method.
func (m *MockDynamodbstreamsClient) GetShardIterator(arg0 context.Context, arg1 *dynamodbstreams.GetShardIteratorInput, arg2 ...func(*dynamodbstreams.Options)) (*dynamodbstreams.GetShardIteratorOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &dynamodbstreams.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetShardIterator")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetShardIterator", varargs...)
	ret0, _ := ret[0].(*dynamodbstreams.GetShardIteratorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShardIterator indicates an expected call of GetShardIterator.
func (mr *MockDynamodbstreamsClientMockRecorder) GetShardIterator(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShardIterator", reflect.TypeOf((*MockDynamodbstreamsClient)(nil).GetShardIterator), varargs...)
}

// ListStreams mocks base method.
func (m *MockDynamodbstreamsClient) ListStreams(arg0 context.Context, arg1 *dynamodbstreams.ListStreamsInput, arg2 ...func(*dynamodbstreams.Options)) (*dynamodbstreams.ListStreamsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &dynamodbstreams.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListStreams")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListStreams", varargs...)
	ret0, _ := ret[0].(*dynamodbstreams.ListStreamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStreams indicates an expected call of ListStreams.
func (mr *MockDynamodbstreamsClientMockRecorder) ListStreams(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStreams", reflect.TypeOf((*MockDynamodbstreamsClient)(nil).ListStreams), varargs...)
}
