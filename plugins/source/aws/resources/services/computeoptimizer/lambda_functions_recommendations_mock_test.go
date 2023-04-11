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

func buildLambdaFunctionRecommendations(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockComputeoptimizerClient(ctrl)
	services := client.Services{
		Computeoptimizer: m,
	}
	item := types.LambdaFunctionRecommendation{}
	err := faker.FakeObject(&item)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetLambdaFunctionRecommendations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&computeoptimizer.GetLambdaFunctionRecommendationsOutput{
			LambdaFunctionRecommendations: []types.LambdaFunctionRecommendation{item},
		}, nil)

	return services
}

func TestLambdaFunctionRecommendations(t *testing.T) {
	client.AwsMockTestHelper(t, LambdaFunctionsRecommendations(), buildLambdaFunctionRecommendations, client.TestOptions{})
}
