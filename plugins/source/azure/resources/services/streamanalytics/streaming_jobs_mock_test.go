// Auto generated code - DO NOT EDIT.

package streamanalytics

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/go-faker/faker/v4"
	fakerOptions "github.com/go-faker/faker/v4/pkg/options"
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
	fieldsToIgnore := []string{"Response", "Properties", "Datasource", "Serialization"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := streamanalytics.NewStreamingJobListResultPage(streamanalytics.StreamingJobListResult{Value: &[]streamanalytics.StreamingJob{data}}, func(ctx context.Context, result streamanalytics.StreamingJobListResult) (streamanalytics.StreamingJobListResult, error) {
		return streamanalytics.StreamingJobListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "").Return(result, nil)
	return s
}
