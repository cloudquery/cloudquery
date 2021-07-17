package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-11-01-preview/insights"
	"github.com/Azure/go-autorest/autorest"
)

type MonitorClient struct {
	LogProfiles        LogProfilesClient
	DiagnosticSettings DiagnosticSettingsClient
	ActivityLogAlerts  ActivityLogAlertsClient
}

func NewMonitorClient(subscriptionId string, auth autorest.Authorizer) MonitorClient {
	servers := insights.NewActivityLogAlertsClient(subscriptionId)
	servers.Authorizer = auth

	logProfiles := insights.NewLogProfilesClient(subscriptionId)
	logProfiles.Authorizer = auth

	diagnosticSettings := insights.NewDiagnosticSettingsClient(subscriptionId)
	diagnosticSettings.Authorizer = auth
	return MonitorClient{
		LogProfiles:        logProfiles,
		DiagnosticSettings: diagnosticSettings,
		ActivityLogAlerts:  servers,
	}
}

type ActivityLogAlertsClient interface {
	ListBySubscriptionID(ctx context.Context) (result insights.ActivityLogAlertList, err error)
}

type LogProfilesClient interface {
	List(ctx context.Context) (result insights.LogProfileCollection, err error)
}

type DiagnosticSettingsClient interface {
	List(ctx context.Context, resourceURI string) (result insights.DiagnosticSettingsResourceCollection, err error)
}
