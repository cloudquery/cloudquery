// Auto generated code - DO NOT EDIT.

package search

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:        "azure_search_services",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/search/mgmt/2020-08-01/search#Service`,
		Resolver:    fetchSearchServices,
		Multiplex:   client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "replica_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ReplicaCount"),
			},
			{
				Name:     "partition_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PartitionCount"),
			},
			{
				Name:     "hosting_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HostingMode"),
			},
			{
				Name:     "public_network_access",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicNetworkAccess"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "status_details",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StatusDetails"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "network_rule_set",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkRuleSet"),
			},
			{
				Name:     "private_endpoint_connections",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrivateEndpointConnections"),
			},
			{
				Name:     "shared_private_link_resources",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SharedPrivateLinkResources"),
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sku"),
			},
			{
				Name:     "identity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Identity"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
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
		},
	}
}

func fetchSearchServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Search.Services

	response, err := svc.ListBySubscription(ctx, nil)

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
