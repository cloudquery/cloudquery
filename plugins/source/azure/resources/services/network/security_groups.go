// Auto generated code - DO NOT EDIT.

package network

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SecurityGroups() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_security_groups",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network#SecurityGroup`,
		Resolver:    fetchNetworkSecurityGroups,
		Multiplex:   client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "security_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SecurityRules"),
			},
			{
				Name:     "default_security_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DefaultSecurityRules"),
			},
			{
				Name:     "network_interfaces",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkInterfaces"),
			},
			{
				Name:     "subnets",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Subnets"),
			},
			{
				Name:     "flow_logs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FlowLogs"),
			},
			{
				Name:     "resource_guid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceGUID"),
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

func fetchNetworkSecurityGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.SecurityGroups

	response, err := svc.ListAll(ctx)

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
