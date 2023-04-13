package computeoptimizer

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/golang/mock/gomock"
)

func buildAutoscalingGroupsRecommendations(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockComputeoptimizerClient(ctrl)
	services := client.Services{
		Computeoptimizer: m,
	}
	item := types.AutoScalingGroupRecommendation{}
	err := faker.FakeObject(&item)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetAutoScalingGroupRecommendations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&computeoptimizer.GetAutoScalingGroupRecommendationsOutput{
			AutoScalingGroupRecommendations: []types.AutoScalingGroupRecommendation{item},
		}, nil)

	return services
}

func TestAutoscalingGroupRecommendations(t *testing.T) {
	client.AwsMockTestHelper(t, AutoscalingGroupsRecommendations(), buildAutoscalingGroupsRecommendations, client.TestOptions{})
}
