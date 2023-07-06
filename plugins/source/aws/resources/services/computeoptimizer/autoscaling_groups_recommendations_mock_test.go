package computeoptimizer

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildAutoscalingGroupsRecommendations(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockComputeoptimizerClient(ctrl)
	services := client.Services{
		Computeoptimizer: m,
	}
	item := types.AutoScalingGroupRecommendation{}
	require.NoError(t, faker.FakeObject(&item))

	m.EXPECT().GetAutoScalingGroupRecommendations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&computeoptimizer.GetAutoScalingGroupRecommendationsOutput{
			AutoScalingGroupRecommendations: []types.AutoScalingGroupRecommendation{item},
		}, nil)

	return services
}

func TestAutoscalingGroupRecommendations(t *testing.T) {
	client.AwsMockTestHelper(t, AutoscalingGroupsRecommendations(), buildAutoscalingGroupsRecommendations, client.TestOptions{})
}
