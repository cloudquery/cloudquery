package resources

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func NetworkWatchers() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_watchers",
		Description: "Azure network watcher",
		Resolver:    fetchNetworkWatchers,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Multiplex:   client.SubscriptionMultiplex,
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
			{
				Name:        "flow_log_storage_id",
				Description: "ID of the storage account which is used to store the flow log",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StorageID"),
			},
			{
				Name:        "flow_log_enabled",
				Description: "Flag to enable/disable flow logging",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Enabled"),
			},
			{
				Name:        "flow_log_retention_policy_days",
				Description: "Number of days to retain flow log records",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("RetentionPolicy.Days"),
			},
			{
				Name:        "flow_log_retention_policy_enabled",
				Description: "Flag to enable/disable retention",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("RetentionPolicy.Enabled"),
			},
			{
				Name:        "flow_log_format_type",
				Description: "The file type of flow log Possible values include: 'JSON'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Format.Type"),
			},
			{
				Name:        "flow_log_format_version",
				Description: "The version (revision) of the flow log",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Format.Version"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchNetworkWatchers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Network.Watchers
	result, err := svc.ListAll(ctx)
	if err != nil {
		return err
	}
	if result.Value == nil {
		return nil
	}
	for _, w := range *result.Value {
		resourceDetails, err := client.ParseResourceID(*w.ID)
		if err != nil {
			return err
		}
		result, err := svc.GetFlowLogStatus(ctx, resourceDetails.ResourceGroup, *w.Name, network.FlowLogStatusParameters{})
		if err != nil {
			return err
		}
		client, ok := svc.(network.WatchersClient)
		if !ok {
			client = network.WatchersClient{} //use a dummy network.WatchersClient with unit tests
		}
		properties, err := result.Result(client)
		if err != nil {
			return err
		}
		res <- NetworkWatcherType{
			w,
			*properties.FlowLogProperties,
		}
	}
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

type NetworkWatcherType struct {
	network.Watcher
	network.FlowLogProperties
}
