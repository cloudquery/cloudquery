package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-11-01-preview/insights"
	"github.com/Azure/go-autorest/autorest"
)

type MonitorClient struct {
	LogProfiles LogProfilesClient
}

func NewMonitorClient(subscriptionID string, auth autorest.Authorizer) MonitorClient {
	logProfiles := insights.NewLogProfilesClient(subscriptionID)
	logProfiles.Authorizer = auth
	return MonitorClient{
		LogProfiles: logProfiles,
	}
}

type LogProfilesClient interface {
	List(ctx context.Context) (result insights.LogProfileCollection, err error)
}
