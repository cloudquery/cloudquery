// Auto generated code - DO NOT EDIT.

package network

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"
)

func PublicIPAddresses() *schema.Table {
	return &schema.Table{
		Name:      "azure_network_public_ip_addresses",
		Resolver:  fetchNetworkPublicIPAddresses,
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
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sku"),
			},
			{
				Name:     "public_ip_allocation_method",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicIPAllocationMethod"),
			},
			{
				Name:     "public_ip_address_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicIPAddressVersion"),
			},
			{
				Name:     "ip_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IPConfiguration"),
			},
			{
				Name:     "dns_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DNSSettings"),
			},
			{
				Name:     "ddos_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DdosSettings"),
			},
			{
				Name:     "ip_tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IPTags"),
			},
			{
				Name:     "ip_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IPAddress"),
			},
			{
				Name:     "public_ip_prefix",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PublicIPPrefix"),
			},
			{
				Name:     "idle_timeout_in_minutes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("IdleTimeoutInMinutes"),
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
				Name:     "service_public_ip_address",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ServicePublicIPAddress"),
			},
			{
				Name:     "nat_gateway",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NatGateway"),
			},
			{
				Name:     "migration_phase",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MigrationPhase"),
			},
			{
				Name:     "linked_public_ip_address",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LinkedPublicIPAddress"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "zones",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Zones"),
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

func fetchNetworkPublicIPAddresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.PublicIPAddresses

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
