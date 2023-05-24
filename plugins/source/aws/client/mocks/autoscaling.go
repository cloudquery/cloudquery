// Code generated by MockGen. DO NOT EDIT.
// Source: autoscaling.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	
	reflect "reflect"

	autoscaling "github.com/aws/aws-sdk-go-v2/service/autoscaling"
	gomock "github.com/golang/mock/gomock"
)

// MockAutoscalingClient is a mock of AutoscalingClient interface.
type MockAutoscalingClient struct {
	ctrl     *gomock.Controller
	recorder *MockAutoscalingClientMockRecorder
}

// MockAutoscalingClientMockRecorder is the mock recorder for MockAutoscalingClient.
type MockAutoscalingClientMockRecorder struct {
	mock *MockAutoscalingClient
}

// NewMockAutoscalingClient creates a new mock instance.
func NewMockAutoscalingClient(ctrl *gomock.Controller) *MockAutoscalingClient {
	mock := &MockAutoscalingClient{ctrl: ctrl}
	mock.recorder = &MockAutoscalingClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAutoscalingClient) EXPECT() *MockAutoscalingClientMockRecorder {
	return m.recorder
}

// DescribeAccountLimits mocks base method.
func (m *MockAutoscalingClient) DescribeAccountLimits(arg0 context.Context, arg1 *autoscaling.DescribeAccountLimitsInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeAccountLimitsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAccountLimits")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAccountLimits", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeAccountLimitsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAccountLimits indicates an expected call of DescribeAccountLimits.
func (mr *MockAutoscalingClientMockRecorder) DescribeAccountLimits(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAccountLimits", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeAccountLimits), varargs...)
}

// DescribeAdjustmentTypes mocks base method.
func (m *MockAutoscalingClient) DescribeAdjustmentTypes(arg0 context.Context, arg1 *autoscaling.DescribeAdjustmentTypesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeAdjustmentTypesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAdjustmentTypes")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAdjustmentTypes", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeAdjustmentTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAdjustmentTypes indicates an expected call of DescribeAdjustmentTypes.
func (mr *MockAutoscalingClientMockRecorder) DescribeAdjustmentTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAdjustmentTypes", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeAdjustmentTypes), varargs...)
}

// DescribeAutoScalingGroups mocks base method.
func (m *MockAutoscalingClient) DescribeAutoScalingGroups(arg0 context.Context, arg1 *autoscaling.DescribeAutoScalingGroupsInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeAutoScalingGroupsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAutoScalingGroups")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAutoScalingGroups", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeAutoScalingGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAutoScalingGroups indicates an expected call of DescribeAutoScalingGroups.
func (mr *MockAutoscalingClientMockRecorder) DescribeAutoScalingGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAutoScalingGroups", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeAutoScalingGroups), varargs...)
}

// DescribeAutoScalingInstances mocks base method.
func (m *MockAutoscalingClient) DescribeAutoScalingInstances(arg0 context.Context, arg1 *autoscaling.DescribeAutoScalingInstancesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeAutoScalingInstancesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAutoScalingInstances")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAutoScalingInstances", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeAutoScalingInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAutoScalingInstances indicates an expected call of DescribeAutoScalingInstances.
func (mr *MockAutoscalingClientMockRecorder) DescribeAutoScalingInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAutoScalingInstances", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeAutoScalingInstances), varargs...)
}

// DescribeAutoScalingNotificationTypes mocks base method.
func (m *MockAutoscalingClient) DescribeAutoScalingNotificationTypes(arg0 context.Context, arg1 *autoscaling.DescribeAutoScalingNotificationTypesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeAutoScalingNotificationTypesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAutoScalingNotificationTypes")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAutoScalingNotificationTypes", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeAutoScalingNotificationTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAutoScalingNotificationTypes indicates an expected call of DescribeAutoScalingNotificationTypes.
func (mr *MockAutoscalingClientMockRecorder) DescribeAutoScalingNotificationTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAutoScalingNotificationTypes", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeAutoScalingNotificationTypes), varargs...)
}

// DescribeInstanceRefreshes mocks base method.
func (m *MockAutoscalingClient) DescribeInstanceRefreshes(arg0 context.Context, arg1 *autoscaling.DescribeInstanceRefreshesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeInstanceRefreshesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeInstanceRefreshes")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeInstanceRefreshes", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeInstanceRefreshesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstanceRefreshes indicates an expected call of DescribeInstanceRefreshes.
func (mr *MockAutoscalingClientMockRecorder) DescribeInstanceRefreshes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstanceRefreshes", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeInstanceRefreshes), varargs...)
}

// DescribeLaunchConfigurations mocks base method.
func (m *MockAutoscalingClient) DescribeLaunchConfigurations(arg0 context.Context, arg1 *autoscaling.DescribeLaunchConfigurationsInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeLaunchConfigurationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeLaunchConfigurations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLaunchConfigurations", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeLaunchConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLaunchConfigurations indicates an expected call of DescribeLaunchConfigurations.
func (mr *MockAutoscalingClientMockRecorder) DescribeLaunchConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLaunchConfigurations", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeLaunchConfigurations), varargs...)
}

// DescribeLifecycleHookTypes mocks base method.
func (m *MockAutoscalingClient) DescribeLifecycleHookTypes(arg0 context.Context, arg1 *autoscaling.DescribeLifecycleHookTypesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeLifecycleHookTypesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeLifecycleHookTypes")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLifecycleHookTypes", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeLifecycleHookTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLifecycleHookTypes indicates an expected call of DescribeLifecycleHookTypes.
func (mr *MockAutoscalingClientMockRecorder) DescribeLifecycleHookTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLifecycleHookTypes", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeLifecycleHookTypes), varargs...)
}

// DescribeLifecycleHooks mocks base method.
func (m *MockAutoscalingClient) DescribeLifecycleHooks(arg0 context.Context, arg1 *autoscaling.DescribeLifecycleHooksInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeLifecycleHooksOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeLifecycleHooks")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLifecycleHooks", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeLifecycleHooksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLifecycleHooks indicates an expected call of DescribeLifecycleHooks.
func (mr *MockAutoscalingClientMockRecorder) DescribeLifecycleHooks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLifecycleHooks", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeLifecycleHooks), varargs...)
}

// DescribeLoadBalancerTargetGroups mocks base method.
func (m *MockAutoscalingClient) DescribeLoadBalancerTargetGroups(arg0 context.Context, arg1 *autoscaling.DescribeLoadBalancerTargetGroupsInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeLoadBalancerTargetGroupsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeLoadBalancerTargetGroups")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLoadBalancerTargetGroups", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeLoadBalancerTargetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLoadBalancerTargetGroups indicates an expected call of DescribeLoadBalancerTargetGroups.
func (mr *MockAutoscalingClientMockRecorder) DescribeLoadBalancerTargetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoadBalancerTargetGroups", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeLoadBalancerTargetGroups), varargs...)
}

// DescribeLoadBalancers mocks base method.
func (m *MockAutoscalingClient) DescribeLoadBalancers(arg0 context.Context, arg1 *autoscaling.DescribeLoadBalancersInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeLoadBalancersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeLoadBalancers")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLoadBalancers", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeLoadBalancersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLoadBalancers indicates an expected call of DescribeLoadBalancers.
func (mr *MockAutoscalingClientMockRecorder) DescribeLoadBalancers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoadBalancers", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeLoadBalancers), varargs...)
}

// DescribeMetricCollectionTypes mocks base method.
func (m *MockAutoscalingClient) DescribeMetricCollectionTypes(arg0 context.Context, arg1 *autoscaling.DescribeMetricCollectionTypesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeMetricCollectionTypesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeMetricCollectionTypes")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeMetricCollectionTypes", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeMetricCollectionTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeMetricCollectionTypes indicates an expected call of DescribeMetricCollectionTypes.
func (mr *MockAutoscalingClientMockRecorder) DescribeMetricCollectionTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMetricCollectionTypes", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeMetricCollectionTypes), varargs...)
}

// DescribeNotificationConfigurations mocks base method.
func (m *MockAutoscalingClient) DescribeNotificationConfigurations(arg0 context.Context, arg1 *autoscaling.DescribeNotificationConfigurationsInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeNotificationConfigurationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeNotificationConfigurations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeNotificationConfigurations", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeNotificationConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeNotificationConfigurations indicates an expected call of DescribeNotificationConfigurations.
func (mr *MockAutoscalingClientMockRecorder) DescribeNotificationConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeNotificationConfigurations", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeNotificationConfigurations), varargs...)
}

// DescribePolicies mocks base method.
func (m *MockAutoscalingClient) DescribePolicies(arg0 context.Context, arg1 *autoscaling.DescribePoliciesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribePoliciesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribePolicies")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribePolicies", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribePolicies indicates an expected call of DescribePolicies.
func (mr *MockAutoscalingClientMockRecorder) DescribePolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribePolicies", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribePolicies), varargs...)
}

// DescribeScalingActivities mocks base method.
func (m *MockAutoscalingClient) DescribeScalingActivities(arg0 context.Context, arg1 *autoscaling.DescribeScalingActivitiesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeScalingActivitiesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeScalingActivities")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeScalingActivities", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeScalingActivitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScalingActivities indicates an expected call of DescribeScalingActivities.
func (mr *MockAutoscalingClientMockRecorder) DescribeScalingActivities(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingActivities", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeScalingActivities), varargs...)
}

// DescribeScalingProcessTypes mocks base method.
func (m *MockAutoscalingClient) DescribeScalingProcessTypes(arg0 context.Context, arg1 *autoscaling.DescribeScalingProcessTypesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeScalingProcessTypesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeScalingProcessTypes")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeScalingProcessTypes", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeScalingProcessTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScalingProcessTypes indicates an expected call of DescribeScalingProcessTypes.
func (mr *MockAutoscalingClientMockRecorder) DescribeScalingProcessTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingProcessTypes", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeScalingProcessTypes), varargs...)
}

// DescribeScheduledActions mocks base method.
func (m *MockAutoscalingClient) DescribeScheduledActions(arg0 context.Context, arg1 *autoscaling.DescribeScheduledActionsInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeScheduledActionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeScheduledActions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeScheduledActions", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeScheduledActionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScheduledActions indicates an expected call of DescribeScheduledActions.
func (mr *MockAutoscalingClientMockRecorder) DescribeScheduledActions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScheduledActions", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeScheduledActions), varargs...)
}

// DescribeTags mocks base method.
func (m *MockAutoscalingClient) DescribeTags(arg0 context.Context, arg1 *autoscaling.DescribeTagsInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeTagsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeTags")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTags", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTags indicates an expected call of DescribeTags.
func (mr *MockAutoscalingClientMockRecorder) DescribeTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTags", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeTags), varargs...)
}

// DescribeTerminationPolicyTypes mocks base method.
func (m *MockAutoscalingClient) DescribeTerminationPolicyTypes(arg0 context.Context, arg1 *autoscaling.DescribeTerminationPolicyTypesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeTerminationPolicyTypesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeTerminationPolicyTypes")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTerminationPolicyTypes", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeTerminationPolicyTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTerminationPolicyTypes indicates an expected call of DescribeTerminationPolicyTypes.
func (mr *MockAutoscalingClientMockRecorder) DescribeTerminationPolicyTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTerminationPolicyTypes", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeTerminationPolicyTypes), varargs...)
}

// DescribeTrafficSources mocks base method.
func (m *MockAutoscalingClient) DescribeTrafficSources(arg0 context.Context, arg1 *autoscaling.DescribeTrafficSourcesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeTrafficSourcesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeTrafficSources")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTrafficSources", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeTrafficSourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTrafficSources indicates an expected call of DescribeTrafficSources.
func (mr *MockAutoscalingClientMockRecorder) DescribeTrafficSources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTrafficSources", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeTrafficSources), varargs...)
}

// DescribeWarmPool mocks base method.
func (m *MockAutoscalingClient) DescribeWarmPool(arg0 context.Context, arg1 *autoscaling.DescribeWarmPoolInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeWarmPoolOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeWarmPool")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeWarmPool", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeWarmPoolOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeWarmPool indicates an expected call of DescribeWarmPool.
func (mr *MockAutoscalingClientMockRecorder) DescribeWarmPool(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeWarmPool", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeWarmPool), varargs...)
}

// GetPredictiveScalingForecast mocks base method.
func (m *MockAutoscalingClient) GetPredictiveScalingForecast(arg0 context.Context, arg1 *autoscaling.GetPredictiveScalingForecastInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.GetPredictiveScalingForecastOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &autoscaling.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetPredictiveScalingForecast")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPredictiveScalingForecast", varargs...)
	ret0, _ := ret[0].(*autoscaling.GetPredictiveScalingForecastOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPredictiveScalingForecast indicates an expected call of GetPredictiveScalingForecast.
func (mr *MockAutoscalingClientMockRecorder) GetPredictiveScalingForecast(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPredictiveScalingForecast", reflect.TypeOf((*MockAutoscalingClient)(nil).GetPredictiveScalingForecast), varargs...)
}
