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

func buildEcsServiceRecommendations(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockComputeoptimizerClient(ctrl)
	services := client.Services{
		Computeoptimizer: m,
	}
	item := types.ECSServiceRecommendation{}
	require.NoError(t, faker.FakeObject(&item))

	m.EXPECT().GetECSServiceRecommendations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&computeoptimizer.GetECSServiceRecommendationsOutput{
			EcsServiceRecommendations: []types.ECSServiceRecommendation{item},
		}, nil)

	return services
}

func TestEcsServiceRecommendations(t *testing.T) {
	client.AwsMockTestHelper(t, EcsServiceRecommendations(), buildEcsServiceRecommendations, client.TestOptions{})
}
