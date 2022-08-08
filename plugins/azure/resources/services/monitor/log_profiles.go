package monitor

import (
	"context"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func MonitorLogProfiles() *schema.Table {
	return &schema.Table{
		Name:         "azure_monitor_log_profiles",
		Description:  "LogProfileResource the log profile resource",
		Resolver:     fetchMonitorLogProfiles,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "storage_account_id",
				Description: "the resource id of the storage account to which you would like to send the Activity Log",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LogProfileProperties.StorageAccountID"),
			},
			{
				Name:          "service_bus_rule_id",
				Description:   "The service bus rule ID of the service bus namespace in which you would like to have Event Hubs created for streaming the Activity Log The rule ID is of the format: '{service bus resource ID}/authorizationrules/{key name}'",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("LogProfileProperties.ServiceBusRuleID"),
				IgnoreInTests: true,
			},
			{
				Name:        "locations",
				Description: "List of regions for which Activity Log events should be stored or streamed It is a comma separated list of valid ARM locations including the 'global' location",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("LogProfileProperties.Locations"),
			},
			{
				Name:        "categories",
				Description: "the categories of the logs These categories are created as is convenient to the user Some values are: 'Write', 'Delete', and/or 'Action'",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("LogProfileProperties.Categories"),
			},
			{
				Name:        "retention_policy_enabled",
				Description: "a value indicating whether the retention policy is enabled",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("LogProfileProperties.RetentionPolicy.Enabled"),
			},
			{
				Name:        "retention_policy_days",
				Description: "the number of days for the retention in days A value of 0 will retain the events indefinitely",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("LogProfileProperties.RetentionPolicy.Days"),
			},
			{
				Name:        "id",
				Description: "Azure resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Azure resource name",
				Type:        schema.TypeString,
			},
			{
				Name:          "type",
				Description:   "Azure resource type",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "location",
				Description:   "Resource location",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "tags",
				Description:   "Resource tags",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchMonitorLogProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Monitor.LogProfiles
	result, err := svc.List(ctx)
	if err != nil {
		return diag.WrapError(err)
	}
	if result.Value == nil {
		return nil
	}
	res <- *result.Value
	return nil
}
