package frauddetector

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildModels(t *testing.T, ctrl *gomock.Controller) client.Services {
	fdClient := mocks.NewMockFrauddetectorClient(ctrl)

	data := types.Model{}
	require.NoError(t, faker.FakeObject(&data))

	fdClient.EXPECT().GetModels(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.GetModelsOutput{Models: []types.Model{data}}, nil,
	)

	buildModelVersions(t, fdClient)

	return client.Services{
		Frauddetector: fdClient,
	}
}

func TestModels(t *testing.T) {
	client.AwsMockTestHelper(t, Models(), buildModels, client.TestOptions{})
}
