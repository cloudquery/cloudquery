package autoscaling

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildAutoscalingGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAutoscalingClient(ctrl)

	groups := autoscaling.DescribeAutoScalingGroupsOutput{}
	err := faker.FakeObject(&groups)
	if err != nil {
		t.Fatal(err)
	}
	groups.NextToken = nil
	m.EXPECT().DescribeAutoScalingGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&groups, nil)

	configurations := autoscaling.DescribeNotificationConfigurationsOutput{}
	err = faker.FakeObject(&configurations)
	if err != nil {
		t.Fatal(err)
	}
	configurations.NextToken = nil
	configurations.NotificationConfigurations[0].AutoScalingGroupName = groups.AutoScalingGroups[0].AutoScalingGroupName
	m.EXPECT().DescribeNotificationConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).Return(&configurations, nil)

	loadBalancers := autoscaling.DescribeLoadBalancersOutput{}
	err = faker.FakeObject(&loadBalancers)
	if err != nil {
		t.Fatal(err)
	}
	loadBalancers.NextToken = nil
	m.EXPECT().DescribeLoadBalancers(gomock.Any(), gomock.Any(), gomock.Any()).Return(&loadBalancers, nil)

	loadBalancerTargetGroups := autoscaling.DescribeLoadBalancerTargetGroupsOutput{}
	err = faker.FakeObject(&loadBalancerTargetGroups)
	if err != nil {
		t.Fatal(err)
	}
	loadBalancerTargetGroups.NextToken = nil
	m.EXPECT().DescribeLoadBalancerTargetGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&loadBalancerTargetGroups, nil)

	policies := autoscaling.DescribePoliciesOutput{}
	err = faker.FakeObject(&policies)
	if err != nil {
		t.Fatal(err)
	}
	policies.NextToken = nil
	m.EXPECT().DescribePolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(&policies, nil)

	lifecycleHooks := autoscaling.DescribeLifecycleHooksOutput{}
	err = faker.FakeObject(&lifecycleHooks)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeLifecycleHooks(gomock.Any(), gomock.Any(), gomock.Any()).Return(&lifecycleHooks, nil)

	return client.Services{
		Autoscaling: m,
	}
}

func TestAutoscalingGroups(t *testing.T) {
	client.AwsMockTestHelper(t, Groups(), buildAutoscalingGroups, client.TestOptions{})
}
