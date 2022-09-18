// Auto generated code - DO NOT EDIT.

package network

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"
)

func VirtualNetworks() *schema.Table {
	return &schema.Table{
		Name:      "azure_network_virtual_networks",
		Resolver:  fetchNetworkVirtualNetworks,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "extended_location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ExtendedLocation"),
			},
			{
				Name:     "address_space",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AddressSpace"),
			},
			{
				Name:     "dhcp_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DhcpOptions"),
			},
			{
				Name:     "subnets",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Subnets"),
			},
			{
				Name:     "virtual_network_peerings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VirtualNetworkPeerings"),
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
				Name:     "enable_ddos_protection",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableDdosProtection"),
			},
			{
				Name:     "enable_vm_protection",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableVMProtection"),
			},
			{
				Name:     "ddos_protection_plan",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DdosProtectionPlan"),
			},
			{
				Name:     "bgp_communities",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BgpCommunities"),
			},
			{
				Name:     "ip_allocations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IPAllocations"),
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

		Relations: []*schema.Table{
			virtualNetworkGateways(),
		},
	}
}

func fetchNetworkVirtualNetworks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.VirtualNetworks

	response, err := svc.ListAll(ctx)

	if err != nil {
		return errors.WithStack(err)
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
