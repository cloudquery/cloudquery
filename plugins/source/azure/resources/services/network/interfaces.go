// Auto generated code - DO NOT EDIT.

package network

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Interfaces() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_interfaces",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network#Interface`,
		Resolver:    fetchNetworkInterfaces,
		Multiplex:   client.SubscriptionMultiplex,
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
				Name:     "virtual_machine",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VirtualMachine"),
			},
			{
				Name:     "network_security_group",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkSecurityGroup"),
			},
			{
				Name:     "private_endpoint",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrivateEndpoint"),
			},
			{
				Name:     "ip_configurations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IPConfigurations"),
			},
			{
				Name:     "tap_configurations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TapConfigurations"),
			},
			{
				Name:     "dns_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DNSSettings"),
			},
			{
				Name:     "mac_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MacAddress"),
			},
			{
				Name:     "primary",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Primary"),
			},
			{
				Name:     "enable_accelerated_networking",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableAcceleratedNetworking"),
			},
			{
				Name:     "enable_ip_forwarding",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableIPForwarding"),
			},
			{
				Name:     "hosted_workloads",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("HostedWorkloads"),
			},
			{
				Name:     "dscp_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DscpConfiguration"),
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
				Name:     "nic_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NicType"),
			},
			{
				Name:     "private_link_service",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrivateLinkService"),
			},
			{
				Name:     "migration_phase",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MigrationPhase"),
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

func fetchNetworkInterfaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.Interfaces

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
