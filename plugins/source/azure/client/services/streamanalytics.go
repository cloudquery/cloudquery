//go:generate mockgen -destination=./mocks/streamanalytics.go -package=mocks . StreamAnalyticsJobsClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/streamanalytics/mgmt/2020-03-01/streamanalytics"
	"github.com/Azure/go-autorest/autorest"
)

type StreamAnalyticsClient struct {
	StreamingJobs StreamAnalyticsJobsClient
}

type StreamAnalyticsJobsClient interface {
	List(ctx context.Context, expand string) (result streamanalytics.StreamingJobListResultPage, err error)
}

func NewStreamAnalyticsClient(subscriptionID string, auth autorest.Authorizer) StreamAnalyticsClient {
	jobs := streamanalytics.NewStreamingJobsClient(subscriptionID)
	jobs.Authorizer = auth
	return StreamAnalyticsClient{
		StreamingJobs: jobs,
	}
}
