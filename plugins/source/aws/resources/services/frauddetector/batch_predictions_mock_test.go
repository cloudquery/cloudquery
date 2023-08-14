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

func buildBatchPredictions(t *testing.T, ctrl *gomock.Controller) client.Services {
	fdClient := mocks.NewMockFrauddetectorClient(ctrl)

	data := types.BatchPrediction{}
	require.NoError(t, faker.FakeObject(&data))

	fdClient.EXPECT().GetBatchPredictionJobs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.GetBatchPredictionJobsOutput{BatchPredictions: []types.BatchPrediction{data}}, nil,
	)

	return client.Services{
		Frauddetector: fdClient,
	}
}

func TestBatchPredictions(t *testing.T) {
	client.AwsMockTestHelper(t, BatchPredictions(), buildBatchPredictions, client.TestOptions{})
}
