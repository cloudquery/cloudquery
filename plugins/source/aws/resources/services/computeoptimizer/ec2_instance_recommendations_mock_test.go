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

func buildEc2InstanceRecommendations(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockComputeoptimizerClient(ctrl)
	services := client.Services{
		Computeoptimizer: m,
	}
	item := types.InstanceRecommendation{}
	require.NoError(t, faker.FakeObject(&item))

	m.EXPECT().GetEC2InstanceRecommendations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&computeoptimizer.GetEC2InstanceRecommendationsOutput{
			InstanceRecommendations: []types.InstanceRecommendation{item},
		}, nil)

	return services
}

func TestEc2InstanceRecommendations(t *testing.T) {
	client.AwsMockTestHelper(t, Ec2InstanceRecommendations(), buildEc2InstanceRecommendations, client.TestOptions{})
}
