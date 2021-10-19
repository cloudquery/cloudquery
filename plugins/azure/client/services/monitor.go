package services

import (
	"context"

	o "github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-11-01-preview/insights"
	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights"
	"github.com/Azure/go-autorest/autorest"
)

type MonitorClient struct {
	ActivityLogAlerts  ActivityLogAlertsClient
	LogProfiles        LogProfilesClient
	ActivityLogs       ActivityLogClient
	DiagnosticSettings DiagnosticSettingsClient
}

func NewMonitorClient(subscriptionId string, auth autorest.Authorizer) MonitorClient {
	servers := o.NewActivityLogAlertsClient(subscriptionId)
	servers.Authorizer = auth
	logProfiles := insights.NewLogProfilesClient(subscriptionId)
	logProfiles.Authorizer = auth
	activityLogs := insights.NewActivityLogsClient(subscriptionId)
	activityLogs.Authorizer = auth
	diagnosticSettings := insights.NewDiagnosticSettingsClient(subscriptionId)
	diagnosticSettings.Authorizer = auth
	return MonitorClient{
		ActivityLogAlerts:  servers,
		LogProfiles:        logProfiles,
		ActivityLogs:       activityLogs,
		DiagnosticSettings: diagnosticSettings,
	}
}

type ActivityLogAlertsClient interface {
	ListBySubscriptionID(ctx context.Context) (result o.ActivityLogAlertList, err error)
}
type ActivityLogClient interface {
	List(ctx context.Context, filter string, selectParameter string) (result insights.EventDataCollectionPage, err error)
}
type LogProfilesClient interface {
	List(ctx context.Context) (result insights.LogProfileCollection, err error)
}

type DiagnosticSettingsClient interface {
	List(ctx context.Context, resourceURI string) (result insights.DiagnosticSettingsResourceCollection, err error)
}
