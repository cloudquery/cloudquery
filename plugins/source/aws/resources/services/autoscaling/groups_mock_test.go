package autoscaling

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildAutoscalingGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAutoscalingClient(ctrl)

	groups := autoscaling.DescribeAutoScalingGroupsOutput{}
	require.NoError(t, faker.FakeObject(&groups))
	groups.NextToken = nil
	m.EXPECT().DescribeAutoScalingGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&groups, nil)

	configurations := autoscaling.DescribeNotificationConfigurationsOutput{}
	require.NoError(t, faker.FakeObject(&configurations))
	configurations.NextToken = nil
	configurations.NotificationConfigurations[0].AutoScalingGroupName = groups.AutoScalingGroups[0].AutoScalingGroupName
	m.EXPECT().DescribeNotificationConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).Return(&configurations, nil)

	loadBalancers := autoscaling.DescribeLoadBalancersOutput{}
	require.NoError(t, faker.FakeObject(&loadBalancers))
	loadBalancers.NextToken = nil
	m.EXPECT().DescribeLoadBalancers(gomock.Any(), gomock.Any(), gomock.Any()).Return(&loadBalancers, nil)

	loadBalancerTargetGroups := autoscaling.DescribeLoadBalancerTargetGroupsOutput{}
	require.NoError(t, faker.FakeObject(&loadBalancerTargetGroups))
	loadBalancerTargetGroups.NextToken = nil
	m.EXPECT().DescribeLoadBalancerTargetGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&loadBalancerTargetGroups, nil)

	policies := autoscaling.DescribePoliciesOutput{}
	require.NoError(t, faker.FakeObject(&policies))
	policies.NextToken = nil
	m.EXPECT().DescribePolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(&policies, nil)

	lifecycleHooks := autoscaling.DescribeLifecycleHooksOutput{}
	require.NoError(t, faker.FakeObject(&lifecycleHooks))
	m.EXPECT().DescribeLifecycleHooks(gomock.Any(), gomock.Any(), gomock.Any()).Return(&lifecycleHooks, nil)

	return client.Services{
		Autoscaling: m,
	}
}

func TestAutoscalingGroups(t *testing.T) {
	client.AwsMockTestHelper(t, Groups(), buildAutoscalingGroups, client.TestOptions{})
}
