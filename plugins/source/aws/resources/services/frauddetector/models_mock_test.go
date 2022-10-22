package frauddetector

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildModels(t *testing.T, ctrl *gomock.Controller) client.Services {
	fdClient := mocks.NewMockFraudDetectorClient(ctrl)

	data := types.Model{}
	err := faker.FakeObject(&data)
	if err != nil {
		t.Fatal(err)
	}

	fdClient.EXPECT().GetModels(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.GetModelsOutput{Models: []types.Model{data}}, nil,
	)

	buildModelVersions(t, fdClient)

	return client.Services{
		FraudDetector: fdClient,
	}
}

func TestModels(t *testing.T) {
	client.AwsMockTestHelper(t, Models(), buildModels, client.TestOptions{})
}
