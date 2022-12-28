package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"

func init() {
	tables := []Table{
		{
			Service:        "armmonitor",
			Name:           "tenant_activity_log_alerts",
			Struct:         &armmonitor.ActivityLogAlertResource{},
			ResponseStruct: &armmonitor.ActivityLogAlertsClientListBySubscriptionIDResponse{},
			Client:         &armmonitor.ActivityLogAlertsClient{},
			ListFunc:       (&armmonitor.ActivityLogAlertsClient{}).NewListBySubscriptionIDPager,
			NewFunc:        armmonitor.NewActivityLogAlertsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/activityLogAlerts",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_insights)`,
		},
	}
	Tables = append(Tables, tables...)
}
