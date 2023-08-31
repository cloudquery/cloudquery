package autoscalingplans

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildPlans(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAutoscalingplansClient(ctrl)
	services := client.Services{
		Autoscalingplans: m,
	}
	p := types.ScalingPlan{}
	require.NoError(t, faker.FakeObject(&p))

	m.EXPECT().DescribeScalingPlans(gomock.Any(), gomock.Any(), gomock.Any()).Return(&autoscalingplans.DescribeScalingPlansOutput{
		ScalingPlans: []types.ScalingPlan{p},
	}, nil)

	pr := types.ScalingPlanResource{}
	require.NoError(t, faker.FakeObject(&pr))

	m.EXPECT().DescribeScalingPlanResources(gomock.Any(), &autoscalingplans.DescribeScalingPlanResourcesInput{
		ScalingPlanName: p.ScalingPlanName,
	}, gomock.Any()).Return(&autoscalingplans.DescribeScalingPlanResourcesOutput{
		ScalingPlanResources: []types.ScalingPlanResource{pr},
	}, nil)

	return services
}

func TestPlans(t *testing.T) {
	client.AwsMockTestHelper(t, Plans(), buildPlans, client.TestOptions{})
}
