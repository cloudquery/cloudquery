package streamanalytics

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/streamanalytics/mgmt/2020-03-01/streamanalytics"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildStreamAnalyticsJobsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockJobsClient(ctrl)
	var p streamanalytics.StreamingJobProperties
	if err := faker.FakeDataSkipFields(&p, []string{"Inputs", "Outputs", "Functions", "JobType", "OutputStartMode", "EventsOutOfOrderPolicy", "OutputErrorPolicy", "CompatibilityLevel", "ContentStoragePolicy"}); err != nil {
		t.Fatal(err)
	}
	p.JobType = streamanalytics.JobTypeCloud
	p.OutputStartMode = streamanalytics.OutputStartModeCustomTime
	p.EventsOutOfOrderPolicy = streamanalytics.EventsOutOfOrderPolicyAdjust
	p.OutputErrorPolicy = streamanalytics.OutputErrorPolicyDrop
	p.CompatibilityLevel = streamanalytics.CompatibilityLevelOneFullStopTwo
	p.ContentStoragePolicy = streamanalytics.ContentStoragePolicyJobStorageAccount
	var j streamanalytics.StreamingJob
	if err := faker.FakeDataSkipFields(&j, []string{"StreamingJobProperties"}); err != nil {
		t.Fatal(err)
	}
	j.StreamingJobProperties = &p
	m.EXPECT().List(gomock.Any(), "").Return(
		streamanalytics.NewStreamingJobListResultPage(
			streamanalytics.StreamingJobListResult{Value: &[]streamanalytics.StreamingJob{j}},
			func(c context.Context, sjlr streamanalytics.StreamingJobListResult) (streamanalytics.StreamingJobListResult, error) {
				return streamanalytics.StreamingJobListResult{}, nil
			},
		),
		nil,
	)
	return services.Services{
		StreamAnalytics: services.StreamAnalyticsClient{
			Jobs: m,
		},
	}
}

func TestStreamAnalyticsJobs(t *testing.T) {
	client.AzureMockTestHelper(t, StreamanalyticsJobs(), buildStreamAnalyticsJobsMock, client.TestOptions{})
}
