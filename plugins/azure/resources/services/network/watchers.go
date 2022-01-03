package network

import (
	"context"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func NetworkWatchers() *schema.Table {
	return &schema.Table{
		Name:         "azure_network_watchers",
		Description:  "Azure network watcher",
		Resolver:     fetchNetworkWatchers,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated",
				Type:        schema.TypeString,
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the network watcher resource Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("WatcherPropertiesFormat.ProvisioningState"),
			},
			{
				Name:        "id",
				Description: "Resource ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource name",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "Resource location",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchNetworkWatchers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.Watchers
	result, err := svc.ListAll(ctx)
	if err != nil {
		return err
	}
	if result.Value == nil {
		return nil
	}
	res <- *result.Value
	return nil
}
