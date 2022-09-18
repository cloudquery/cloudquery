// Auto generated code - DO NOT EDIT.

package network

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"
)

func ExpressRouteCircuits() *schema.Table {
	return &schema.Table{
		Name:      "azure_network_express_route_circuits",
		Resolver:  fetchNetworkExpressRouteCircuits,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sku"),
			},
			{
				Name:     "allow_classic_operations",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AllowClassicOperations"),
			},
			{
				Name:     "circuit_provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CircuitProvisioningState"),
			},
			{
				Name:     "service_provider_provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceProviderProvisioningState"),
			},
			{
				Name:     "authorizations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Authorizations"),
			},
			{
				Name:     "peerings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Peerings"),
			},
			{
				Name:     "service_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceKey"),
			},
			{
				Name:     "service_provider_notes",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceProviderNotes"),
			},
			{
				Name:     "service_provider_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ServiceProviderProperties"),
			},
			{
				Name:     "express_route_port",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ExpressRoutePort"),
			},
			{
				Name:     "bandwidth_in_gbps",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("BandwidthInGbps"),
			},
			{
				Name:     "stag",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Stag"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "gateway_manager_etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GatewayManagerEtag"),
			},
			{
				Name:     "global_reach_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("GlobalReachEnabled"),
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

func fetchNetworkExpressRouteCircuits(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.ExpressRouteCircuits

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
