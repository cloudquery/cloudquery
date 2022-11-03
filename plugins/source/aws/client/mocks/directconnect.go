// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/aws/client (interfaces: DirectconnectClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	directconnect "github.com/aws/aws-sdk-go-v2/service/directconnect"
	gomock "github.com/golang/mock/gomock"
)

// MockDirectconnectClient is a mock of DirectconnectClient interface.
type MockDirectconnectClient struct {
	ctrl     *gomock.Controller
	recorder *MockDirectconnectClientMockRecorder
}

// MockDirectconnectClientMockRecorder is the mock recorder for MockDirectconnectClient.
type MockDirectconnectClientMockRecorder struct {
	mock *MockDirectconnectClient
}

// NewMockDirectconnectClient creates a new mock instance.
func NewMockDirectconnectClient(ctrl *gomock.Controller) *MockDirectconnectClient {
	mock := &MockDirectconnectClient{ctrl: ctrl}
	mock.recorder = &MockDirectconnectClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDirectconnectClient) EXPECT() *MockDirectconnectClientMockRecorder {
	return m.recorder
}

// DescribeConnectionLoa mocks base method.
func (m *MockDirectconnectClient) DescribeConnectionLoa(arg0 context.Context, arg1 *directconnect.DescribeConnectionLoaInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeConnectionLoaOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConnectionLoa", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeConnectionLoaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConnectionLoa indicates an expected call of DescribeConnectionLoa.
func (mr *MockDirectconnectClientMockRecorder) DescribeConnectionLoa(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConnectionLoa", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeConnectionLoa), varargs...)
}

// DescribeConnections mocks base method.
func (m *MockDirectconnectClient) DescribeConnections(arg0 context.Context, arg1 *directconnect.DescribeConnectionsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeConnectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConnections", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeConnectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConnections indicates an expected call of DescribeConnections.
func (mr *MockDirectconnectClientMockRecorder) DescribeConnections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConnections", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeConnections), varargs...)
}

// DescribeConnectionsOnInterconnect mocks base method.
func (m *MockDirectconnectClient) DescribeConnectionsOnInterconnect(arg0 context.Context, arg1 *directconnect.DescribeConnectionsOnInterconnectInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeConnectionsOnInterconnectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConnectionsOnInterconnect", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeConnectionsOnInterconnectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConnectionsOnInterconnect indicates an expected call of DescribeConnectionsOnInterconnect.
func (mr *MockDirectconnectClientMockRecorder) DescribeConnectionsOnInterconnect(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConnectionsOnInterconnect", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeConnectionsOnInterconnect), varargs...)
}

// DescribeCustomerMetadata mocks base method.
func (m *MockDirectconnectClient) DescribeCustomerMetadata(arg0 context.Context, arg1 *directconnect.DescribeCustomerMetadataInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeCustomerMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCustomerMetadata", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeCustomerMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCustomerMetadata indicates an expected call of DescribeCustomerMetadata.
func (mr *MockDirectconnectClientMockRecorder) DescribeCustomerMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCustomerMetadata", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeCustomerMetadata), varargs...)
}

// DescribeDirectConnectGatewayAssociationProposals mocks base method.
func (m *MockDirectconnectClient) DescribeDirectConnectGatewayAssociationProposals(arg0 context.Context, arg1 *directconnect.DescribeDirectConnectGatewayAssociationProposalsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDirectConnectGatewayAssociationProposals", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDirectConnectGatewayAssociationProposals indicates an expected call of DescribeDirectConnectGatewayAssociationProposals.
func (mr *MockDirectconnectClientMockRecorder) DescribeDirectConnectGatewayAssociationProposals(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDirectConnectGatewayAssociationProposals", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeDirectConnectGatewayAssociationProposals), varargs...)
}

// DescribeDirectConnectGatewayAssociations mocks base method.
func (m *MockDirectconnectClient) DescribeDirectConnectGatewayAssociations(arg0 context.Context, arg1 *directconnect.DescribeDirectConnectGatewayAssociationsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewayAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDirectConnectGatewayAssociations", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeDirectConnectGatewayAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDirectConnectGatewayAssociations indicates an expected call of DescribeDirectConnectGatewayAssociations.
func (mr *MockDirectconnectClientMockRecorder) DescribeDirectConnectGatewayAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDirectConnectGatewayAssociations", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeDirectConnectGatewayAssociations), varargs...)
}

// DescribeDirectConnectGatewayAttachments mocks base method.
func (m *MockDirectconnectClient) DescribeDirectConnectGatewayAttachments(arg0 context.Context, arg1 *directconnect.DescribeDirectConnectGatewayAttachmentsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewayAttachmentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDirectConnectGatewayAttachments", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeDirectConnectGatewayAttachmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDirectConnectGatewayAttachments indicates an expected call of DescribeDirectConnectGatewayAttachments.
func (mr *MockDirectconnectClientMockRecorder) DescribeDirectConnectGatewayAttachments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDirectConnectGatewayAttachments", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeDirectConnectGatewayAttachments), varargs...)
}

// DescribeDirectConnectGateways mocks base method.
func (m *MockDirectconnectClient) DescribeDirectConnectGateways(arg0 context.Context, arg1 *directconnect.DescribeDirectConnectGatewaysInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDirectConnectGateways", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeDirectConnectGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDirectConnectGateways indicates an expected call of DescribeDirectConnectGateways.
func (mr *MockDirectconnectClientMockRecorder) DescribeDirectConnectGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDirectConnectGateways", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeDirectConnectGateways), varargs...)
}

// DescribeHostedConnections mocks base method.
func (m *MockDirectconnectClient) DescribeHostedConnections(arg0 context.Context, arg1 *directconnect.DescribeHostedConnectionsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeHostedConnectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeHostedConnections", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeHostedConnectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeHostedConnections indicates an expected call of DescribeHostedConnections.
func (mr *MockDirectconnectClientMockRecorder) DescribeHostedConnections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeHostedConnections", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeHostedConnections), varargs...)
}

// DescribeInterconnectLoa mocks base method.
func (m *MockDirectconnectClient) DescribeInterconnectLoa(arg0 context.Context, arg1 *directconnect.DescribeInterconnectLoaInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeInterconnectLoaOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeInterconnectLoa", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeInterconnectLoaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInterconnectLoa indicates an expected call of DescribeInterconnectLoa.
func (mr *MockDirectconnectClientMockRecorder) DescribeInterconnectLoa(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInterconnectLoa", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeInterconnectLoa), varargs...)
}

// DescribeInterconnects mocks base method.
func (m *MockDirectconnectClient) DescribeInterconnects(arg0 context.Context, arg1 *directconnect.DescribeInterconnectsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeInterconnectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeInterconnects", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeInterconnectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInterconnects indicates an expected call of DescribeInterconnects.
func (mr *MockDirectconnectClientMockRecorder) DescribeInterconnects(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInterconnects", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeInterconnects), varargs...)
}

// DescribeLags mocks base method.
func (m *MockDirectconnectClient) DescribeLags(arg0 context.Context, arg1 *directconnect.DescribeLagsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeLagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLags", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeLagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLags indicates an expected call of DescribeLags.
func (mr *MockDirectconnectClientMockRecorder) DescribeLags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLags", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeLags), varargs...)
}

// DescribeLoa mocks base method.
func (m *MockDirectconnectClient) DescribeLoa(arg0 context.Context, arg1 *directconnect.DescribeLoaInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeLoaOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLoa", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeLoaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLoa indicates an expected call of DescribeLoa.
func (mr *MockDirectconnectClientMockRecorder) DescribeLoa(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoa", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeLoa), varargs...)
}

// DescribeLocations mocks base method.
func (m *MockDirectconnectClient) DescribeLocations(arg0 context.Context, arg1 *directconnect.DescribeLocationsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeLocationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLocations", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeLocationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLocations indicates an expected call of DescribeLocations.
func (mr *MockDirectconnectClientMockRecorder) DescribeLocations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLocations", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeLocations), varargs...)
}

// DescribeRouterConfiguration mocks base method.
func (m *MockDirectconnectClient) DescribeRouterConfiguration(arg0 context.Context, arg1 *directconnect.DescribeRouterConfigurationInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeRouterConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRouterConfiguration", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeRouterConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRouterConfiguration indicates an expected call of DescribeRouterConfiguration.
func (mr *MockDirectconnectClientMockRecorder) DescribeRouterConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRouterConfiguration", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeRouterConfiguration), varargs...)
}

// DescribeTags mocks base method.
func (m *MockDirectconnectClient) DescribeTags(arg0 context.Context, arg1 *directconnect.DescribeTagsInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTags", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTags indicates an expected call of DescribeTags.
func (mr *MockDirectconnectClientMockRecorder) DescribeTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTags", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeTags), varargs...)
}

// DescribeVirtualGateways mocks base method.
func (m *MockDirectconnectClient) DescribeVirtualGateways(arg0 context.Context, arg1 *directconnect.DescribeVirtualGatewaysInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeVirtualGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVirtualGateways", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeVirtualGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVirtualGateways indicates an expected call of DescribeVirtualGateways.
func (mr *MockDirectconnectClientMockRecorder) DescribeVirtualGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVirtualGateways", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeVirtualGateways), varargs...)
}

// DescribeVirtualInterfaces mocks base method.
func (m *MockDirectconnectClient) DescribeVirtualInterfaces(arg0 context.Context, arg1 *directconnect.DescribeVirtualInterfacesInput, arg2 ...func(*directconnect.Options)) (*directconnect.DescribeVirtualInterfacesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVirtualInterfaces", varargs...)
	ret0, _ := ret[0].(*directconnect.DescribeVirtualInterfacesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVirtualInterfaces indicates an expected call of DescribeVirtualInterfaces.
func (mr *MockDirectconnectClientMockRecorder) DescribeVirtualInterfaces(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVirtualInterfaces", reflect.TypeOf((*MockDirectconnectClient)(nil).DescribeVirtualInterfaces), varargs...)
}

// ListVirtualInterfaceTestHistory mocks base method.
func (m *MockDirectconnectClient) ListVirtualInterfaceTestHistory(arg0 context.Context, arg1 *directconnect.ListVirtualInterfaceTestHistoryInput, arg2 ...func(*directconnect.Options)) (*directconnect.ListVirtualInterfaceTestHistoryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListVirtualInterfaceTestHistory", varargs...)
	ret0, _ := ret[0].(*directconnect.ListVirtualInterfaceTestHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVirtualInterfaceTestHistory indicates an expected call of ListVirtualInterfaceTestHistory.
func (mr *MockDirectconnectClientMockRecorder) ListVirtualInterfaceTestHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVirtualInterfaceTestHistory", reflect.TypeOf((*MockDirectconnectClient)(nil).ListVirtualInterfaceTestHistory), varargs...)
}
