package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildApplicationAutoscalingPoliciesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApplicationAutoscalingClient(ctrl)
	services := client.Services{
		ApplicationAutoscaling: m,
	}
	c := types.ScalingPolicy{}
	if err := faker.FakeData(&c); err != nil {
		t.Fatal(err)
	}
	output := &applicationautoscaling.DescribeScalingPoliciesOutput{
		ScalingPolicies: []types.ScalingPolicy{c},
	}
	m.EXPECT().DescribeScalingPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		output,
		nil,
	)

	return services
}

func TestApplicationAutoscalingPolicies(t *testing.T) {
	client.AllNamespaces = []string{"test-namespace"} // Just one

	awsTestHelper(t, ApplicationautoscalingPolicies(), buildApplicationAutoscalingPoliciesMock, TestOptions{})
}
