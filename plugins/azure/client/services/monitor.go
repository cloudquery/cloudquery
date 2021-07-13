package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-11-01-preview/insights"
	"github.com/Azure/go-autorest/autorest"
)

type MonitorClient struct {
	ActivityLogAlerts ActivityLogAlertsClient
	LogProfiles       LogProfilesClient
}

func NewMonitorClient(subscriptionId string, auth autorest.Authorizer) MonitorClient {
	servers := insights.NewActivityLogAlertsClient(subscriptionId)
	servers.Authorizer = auth
	logProfiles := insights.NewLogProfilesClient(subscriptionId)
	logProfiles.Authorizer = auth
	return MonitorClient{
		ActivityLogAlerts: servers,
		LogProfiles:       logProfiles,
	}
}

type ActivityLogAlertsClient interface {
	ListBySubscriptionID(ctx context.Context) (result insights.ActivityLogAlertList, err error)
}

type LogProfilesClient interface {
	List(ctx context.Context) (result insights.LogProfileCollection, err error)
}
