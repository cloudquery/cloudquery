// Auto generated code - DO NOT EDIT.

package network

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
)

func flowLogs() *schema.Table {
	return &schema.Table{
		Name:     "azure_network_flow_logs",
		Resolver: fetchNetworkFlowLogs,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "network_watcher_id",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIDResolver,
			},
			{
				Name:     "target_resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TargetResourceID"),
			},
			{
				Name:     "target_resource_guid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TargetResourceGUID"),
			},
			{
				Name:     "storage_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StorageID"),
			},
			{
				Name:     "enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Enabled"),
			},
			{
				Name:     "retention_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RetentionPolicy"),
			},
			{
				Name:     "format",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Format"),
			},
			{
				Name:     "flow_analytics_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FlowAnalyticsConfiguration"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
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
		},
	}
}

func fetchNetworkFlowLogs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.FlowLogs

	watcher := parent.Item.(network.Watcher)
	resourceDetails, err := client.ParseResourceID(*watcher.ID)
	if err != nil {
		return err
	}
	response, err := svc.List(ctx, resourceDetails.ResourceGroup, *watcher.Name)

	if err != nil {
		return err
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}

	return nil
}
