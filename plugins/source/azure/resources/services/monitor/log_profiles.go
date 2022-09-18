// Auto generated code - DO NOT EDIT.

package monitor

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func LogProfiles() *schema.Table {
	return &schema.Table{
		Name:      "azure_monitor_log_profiles",
		Resolver:  fetchMonitorLogProfiles,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "storage_account_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StorageAccountID"),
			},
			{
				Name:     "service_bus_rule_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceBusRuleID"),
			},
			{
				Name:     "locations",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Locations"),
			},
			{
				Name:     "categories",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Categories"),
			},
			{
				Name:     "retention_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RetentionPolicy"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
		},
	}
}

func fetchMonitorLogProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Monitor.LogProfiles

	response, err := svc.List(ctx)
	if err != nil {
		return err
	}
	if response.Value == nil {
		return nil
	}
	res <- *response.Value

	return nil
}
