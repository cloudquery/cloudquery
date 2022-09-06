//go:generate mockgen -destination=./mocks/monitor.go -package=mocks . MonitorActivityLogAlertsClient,MonitorLogProfilesClient,MonitorDiagnosticSettingsClient,MonitorActivityLogsClient,MonitorResourcesClient
package services

import (
	"context"

	o "github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-11-01-preview/insights"
	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"
	"github.com/Azure/go-autorest/autorest"
)

type MonitorClient struct {
	ActivityLogAlerts  MonitorActivityLogAlertsClient
	LogProfiles        MonitorLogProfilesClient
	ActivityLogs       MonitorActivityLogsClient
	DiagnosticSettings MonitorDiagnosticSettingsClient
	Resources          MonitorResourcesClient
}

type MonitorActivityLogAlertsClient interface {
	ListBySubscriptionID(ctx context.Context) (result o.ActivityLogAlertList, err error)
}
type MonitorActivityLogsClient interface {
	List(ctx context.Context, filter string, selectParameter string) (result insights.EventDataCollectionPage, err error)
}
type MonitorLogProfilesClient interface {
	List(ctx context.Context) (result insights.LogProfileCollection, err error)
}

type MonitorDiagnosticSettingsClient interface {
	List(ctx context.Context, resourceURI string) (result insights.DiagnosticSettingsResourceCollection, err error)
}

type MonitorResourcesClient interface {
	List(ctx context.Context, filter string, expand string, top *int32) (result resources.ListResultPage, err error)
}

func NewMonitorClient(subscriptionId string, auth autorest.Authorizer) MonitorClient {
	logAlerts := o.NewActivityLogAlertsClient(subscriptionId)
	logAlerts.Authorizer = auth
	logProfiles := insights.NewLogProfilesClient(subscriptionId)
	logProfiles.Authorizer = auth
	activityLogs := insights.NewActivityLogsClient(subscriptionId)
	activityLogs.Authorizer = auth
	diagnosticSettings := insights.NewDiagnosticSettingsClient(subscriptionId)
	diagnosticSettings.Authorizer = auth
	resourcesClient := resources.NewClient(subscriptionId)
	resourcesClient.Authorizer = auth
	return MonitorClient{
		ActivityLogAlerts:  logAlerts,
		LogProfiles:        logProfiles,
		ActivityLogs:       activityLogs,
		DiagnosticSettings: diagnosticSettings,
		Resources:          resourcesClient,
	}
}
