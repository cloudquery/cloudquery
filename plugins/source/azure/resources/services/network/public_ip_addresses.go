// Code generated by codegen; DO NOT EDIT.

package network

import (
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func PublicIpAddresses() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_public_ip_addresses",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2#PublicIPAddress`,
		Resolver:    fetchPublicIpAddresses,
		Multiplex:   client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Type:        schema.TypeString,
				Resolver:    client.SubscriptionIDResolver,
				Description: `Azure subscription ID`,
			},
			{
				Name:     "extended_location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ExtendedLocation"),
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
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "dns_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.DNSSettings"),
			},
			{
				Name:     "ddos_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.DdosSettings"),
			},
			{
				Name:     "delete_option",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.DeleteOption"),
			},
			{
				Name:     "ip_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.IPAddress"),
			},
			{
				Name:     "ip_tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.IPTags"),
			},
			{
				Name:     "idle_timeout_in_minutes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Properties.IdleTimeoutInMinutes"),
			},
			{
				Name:     "linked_public_ip_address",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.LinkedPublicIPAddress"),
			},
			{
				Name:     "migration_phase",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.MigrationPhase"),
			},
			{
				Name:     "nat_gateway",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.NatGateway"),
			},
			{
				Name:     "public_ip_address_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.PublicIPAddressVersion"),
			},
			{
				Name:     "public_ip_allocation_method",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.PublicIPAllocationMethod"),
			},
			{
				Name:     "public_ip_prefix",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.PublicIPPrefix"),
			},
			{
				Name:     "service_public_ip_address",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.ServicePublicIPAddress"),
			},
			{
				Name:     "ip_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.IPConfiguration"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.ProvisioningState"),
			},
			{
				Name:     "resource_guid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.ResourceGUID"),
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SKU"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "zones",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Zones"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
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
