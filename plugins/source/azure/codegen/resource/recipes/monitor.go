package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/monitor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/resources"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func Monitor() []*resource.Resource {
	return []*resource.Resource{
		{
			SubService: "activity_log_alerts",
			Struct:     new(armmonitor.ActivityLogAlertResource),
			Resolver: &resource.FuncParams{
				Func: monitor.ActivityLogAlertsClient.NewListBySubscriptionIDPager,
			},
		},
		{
			SubService: "activity_logs",
			Struct:     new(armmonitor.EventData),
			Resolver: &resource.FuncParams{
				Func: monitor.ActivityLogsClient.NewListPager,
				Params: []string{`func()string {
		const fetchWindow = 24 * time.Hour
		now := time.Now().UTC()
		past := now.Add(-fetchWindow)
		return fmt.Sprintf("eventTimestamp ge '%s' and eventTimestamp le '%s'",
			past.Format(time.RFC3339Nano),
			now.Format(time.RFC3339Nano),
		)
	}()`},
			},
		},
		{
			SubService: "log_profiles",
			Struct:     new(armmonitor.LogProfileResource),
			Resolver: &resource.FuncParams{
				Func: monitor.LogProfilesClient.NewListPager,
			},
		},
		{
			Service:    "monitor",
			SubService: "resources",
			Struct:     new(armresources.GenericResourceExpanded),
			Resolver: &resource.FuncParams{
				Func:        resources.Client.NewListPager,
				ExtraValues: []string{"&armresources.GenericResourceExpanded{ID: to.Ptr(c.ScopeSubscription())}"},
			},
			Children: []*resource.Resource{
				{
					SubService: "diagnostic_settings",
					Struct:     new(armmonitor.DiagnosticSettingsResource),
					Resolver: &resource.FuncParams{
						Func:   monitor.DiagnosticSettingsClient.NewListPager,
						Params: []string{"id.String()"},
					},
				},
			},
		},
	}
}
