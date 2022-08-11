package autoscaling

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildAutoscalingGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAutoscalingClient(ctrl)

	groups := autoscaling.DescribeAutoScalingGroupsOutput{}
	err := faker.FakeData(&groups)
	if err != nil {
		t.Fatal(err)
	}
	groups.NextToken = nil
	m.EXPECT().DescribeAutoScalingGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&groups, nil)

	configurations := autoscaling.DescribeNotificationConfigurationsOutput{}
	err = faker.FakeData(&configurations)
	if err != nil {
		t.Fatal(err)
	}
	configurations.NextToken = nil
	configurations.NotificationConfigurations[0].AutoScalingGroupName = groups.AutoScalingGroups[0].AutoScalingGroupName
	m.EXPECT().DescribeNotificationConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).Return(&configurations, nil)

	loadBalancers := autoscaling.DescribeLoadBalancersOutput{}
	err = faker.FakeData(&loadBalancers)
	if err != nil {
		t.Fatal(err)
	}
	loadBalancers.NextToken = nil
	m.EXPECT().DescribeLoadBalancers(gomock.Any(), gomock.Any(), gomock.Any()).Return(&loadBalancers, nil)

	loadBalancerTargetGroups := autoscaling.DescribeLoadBalancerTargetGroupsOutput{}
	err = faker.FakeData(&loadBalancerTargetGroups)
	if err != nil {
		t.Fatal(err)
	}
	loadBalancerTargetGroups.NextToken = nil
	m.EXPECT().DescribeLoadBalancerTargetGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&loadBalancerTargetGroups, nil)

	policies := autoscaling.DescribePoliciesOutput{}
	err = faker.FakeData(&policies)
	if err != nil {
		t.Fatal(err)
	}
	policies.NextToken = nil
	m.EXPECT().DescribePolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(&policies, nil)

	lifecycleHooks := autoscaling.DescribeLifecycleHooksOutput{}
	err = faker.FakeData(&lifecycleHooks)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeLifecycleHooks(gomock.Any(), gomock.Any(), gomock.Any()).Return(&lifecycleHooks, nil)

	return client.Services{
		Autoscaling: m,
	}
}

func TestAutoscalingGroups(t *testing.T) {
	client.AwsMockTestHelper(t, AutoscalingGroups(), buildAutoscalingGroups, client.TestOptions{})
}
