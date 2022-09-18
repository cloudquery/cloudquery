// Auto generated code - DO NOT EDIT.

package streamanalytics

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/streamanalytics/mgmt/2020-03-01/streamanalytics"
)

func TestStreamAnalyticsStreamingJobs(t *testing.T) {
	client.MockTestHelper(t, StreamingJobs(), createStreamingJobsMock)
}

func createStreamingJobsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockStreamAnalyticsStreamingJobsClient(ctrl)
	s := services.Services{
		StreamAnalytics: services.StreamAnalyticsClient{
			StreamingJobs: mockClient,
		},
	}

	data := streamanalytics.StreamingJob{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	result := streamanalytics.NewStreamingJobListResultPage(streamanalytics.StreamingJobListResult{Value: &[]streamanalytics.StreamingJob{data}}, func(ctx context.Context, result streamanalytics.StreamingJobListResult) (streamanalytics.StreamingJobListResult, error) {
		return streamanalytics.StreamingJobListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "").Return(result, nil)
	return s
}
