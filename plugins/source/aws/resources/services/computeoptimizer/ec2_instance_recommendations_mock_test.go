package computeoptimizer

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEc2InstanceRecommendations(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockComputeoptimizerClient(ctrl)
	services := client.Services{
		Computeoptimizer: m,
	}
	item := types.InstanceRecommendation{}
	err := faker.FakeObject(&item)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetEC2InstanceRecommendations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&computeoptimizer.GetEC2InstanceRecommendationsOutput{
			InstanceRecommendations: []types.InstanceRecommendation{item},
		}, nil)

	return services
}

func TestEc2InstanceRecommendations(t *testing.T) {
	client.AwsMockTestHelper(t, Ec2InstanceRecommendations(), buildEc2InstanceRecommendations, client.TestOptions{})
}
