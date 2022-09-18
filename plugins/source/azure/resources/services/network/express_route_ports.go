// Auto generated code - DO NOT EDIT.

package network

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"
)

func ExpressRoutePorts() *schema.Table {
	return &schema.Table{
		Name:      "azure_network_express_route_ports",
		Resolver:  fetchNetworkExpressRoutePorts,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "peering_location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PeeringLocation"),
			},
			{
				Name:     "bandwidth_in_gbps",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("BandwidthInGbps"),
			},
			{
				Name:     "provisioned_bandwidth_in_gbps",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("ProvisionedBandwidthInGbps"),
			},
			{
				Name:     "mtu",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Mtu"),
			},
			{
				Name:     "encapsulation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Encapsulation"),
			},
			{
				Name:     "ether_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EtherType"),
			},
			{
				Name:     "allocation_date",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AllocationDate"),
			},
			{
				Name:     "links",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Links"),
			},
			{
				Name:     "circuits",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Circuits"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "resource_guid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceGUID"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "identity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Identity"),
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

func fetchNetworkExpressRoutePorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.ExpressRoutePorts

	response, err := svc.List(ctx)

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
