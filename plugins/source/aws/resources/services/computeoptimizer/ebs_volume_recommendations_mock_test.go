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

func buildEbsVolumeRecommendations(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockComputeoptimizerClient(ctrl)
	services := client.Services{
		Computeoptimizer: m,
	}
	item := types.VolumeRecommendation{}
	err := faker.FakeObject(&item)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetEBSVolumeRecommendations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&computeoptimizer.GetEBSVolumeRecommendationsOutput{
			VolumeRecommendations: []types.VolumeRecommendation{item},
		}, nil)

	return services
}

func TestEbsVolumeRecommendations(t *testing.T) {
	client.AwsMockTestHelper(t, EbsVolumeRecommendations(), buildEbsVolumeRecommendations, client.TestOptions{})
}
