package network

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func NetworkInterfaces() *schema.Table {
	return &schema.Table{
		Name:         "azure_network_interfaces",
		Description:  "Azure Network Interface",
		Resolver:     fetchNetworkInterfaces,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated.",
				Type:        schema.TypeString,
			},
			{
				Name:          "extended_location_name",
				Description:   "The name of the extended location",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ExtendedLocation.Name"),
				IgnoreInTests: true,
			},
			{
				Name:          "extended_location_type",
				Description:   "The type of the extended location",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ExtendedLocation.Type"),
				IgnoreInTests: true,
			},
			{
				Name:        "id",
				Description: "Resource ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "location",
				Description: "Resource location.",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "Resource name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "dns_settings_applied_dns_servers",
				Description: "The servers that are part of the same availability set.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.DNSSettings.AppliedDNSServers"),
			},
			{
				Name:        "dns_settings_dns_servers",
				Description: "List of DNS servers IP addresses.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.DNSSettings.DNSServers"),
			},
			{
				Name:          "dns_settings_internal_dns_name_label",
				Description:   "The internal dns name label.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("InterfacePropertiesFormat.DNSSettings.InternalDNSNameLabel"),
				IgnoreInTests: true,
			},
			{
				Name:        "dns_settings_internal_domain_name_suffix",
				Description: "The internal domain name suffix.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.DNSSettings.InternalDomainNameSuffix"),
			},
			{
				Name:          "dns_settings_internal_fqdn",
				Description:   "Fully qualified DNS name supporting internal communications between VMs in the same virtual network.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("InterfacePropertiesFormat.DNSSettings.InternalFqdn"),
				IgnoreInTests: true,
			},

			{
				Name:          "dscp_configuration_id",
				Description:   "A reference to the dscp configuration to which the network interface is linked.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("InterfacePropertiesFormat.DscpConfiguration.ID"),
				IgnoreInTests: true,
			},
			{
				Name:        "enable_accelerated_networking",
				Description: "If the network interface is accelerated networking enabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.EnableAcceleratedNetworking"),
			},
			{
				Name:        "enable_ip_forwarding",
				Description: "Indicates whether IP forwarding is enabled on this network interface.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.EnableIPForwarding"),
			},
			{
				Name:        "hosted_workloads",
				Description: "List of references to linked BareMetal resources.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.HostedWorkloads"),
			},
			{
				Name:        "mac_address",
				Description: "The MAC address of the network interface.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.MacAddress"),
			},
			{
				Name:        "migration_phase",
				Description: "Migration phase of Network Interface resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.MigrationPhase"),
			},
			{
				Name:          "network_security_group",
				Description:   "The reference to the NetworkSecurityGroup resource.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("InterfacePropertiesFormat.NetworkSecurityGroup.ID"),
				IgnoreInTests: true,
			},
			{
				Name:        "nic_type",
				Description: "Type of Network Interface resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.NicType"),
			},
			{
				Name:        "primary",
				Description: "Whether this is a primary network interface on a virtual machine.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.Primary"),
			},
			{
				Name:        "private_endpoint",
				Description: "A reference to the private endpoint to which the network interface is linked.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.PrivateEndpoint.ID"),
			},
			{
				Name:          "private_link_service",
				Description:   "Privatelinkservice of the network interface resource.",
				Type:          schema.TypeJSON,
				Resolver:      resolveNetworkInterfacePrivateLinkService,
				IgnoreInTests: true,
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the network interface resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.ProvisioningState"),
			},
			{
				Name:        "resource_guid",
				Description: "The provisioning state of the network interface resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.ResourceGUID"),
			},
			{
				Name:        "tap_configurations",
				Description: "A list of TapConfigurations of the network interface.",
				Type:        schema.TypeJSON,
				Resolver:    resolveNetworkInterfaceTapConfigurations,
			},
			{
				Name:        "virtual_machine_id",
				Description: "The reference to a virtual machine.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.VirtualMachine.ID"),
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "type",
				Description: "Resource type.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_network_interface_ip_configurations",
				Description: "NetworkInterface IP Configurations. ",
				Resolver:    fetchNetworkInterfaceIPConfigurations,
				Columns: []schema.Column{
					{
						Name:        "interface_cq_id",
						Description: "Unique CloudQuery ID of azure_network_interface table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "Resource Id",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "Resource name",
						Type:        schema.TypeString,
					},
					{
						Name:        "etag",
						Description: "A unique read-only string that changes whenever the resource is updated.",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Resource type",
						Type:        schema.TypeString,
					},
					{
						Name:        "primary",
						Description: "Whether this is a primary network interface on a virtual machine.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.Primary"),
					},
					{
						Name:          "application_gateway_backend_address_pools",
						Description:   "The reference to ApplicationGatewayBackendAddressPool resource.",
						Type:          schema.TypeJSON,
						Resolver:      schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.ApplicationGatewayBackendAddressPools"),
						IgnoreInTests: true,
					},
					{
						Name:          "application_security_groups",
						Description:   "Application security groups in which the IP configuration is included.",
						Type:          schema.TypeJSON,
						Resolver:      schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.ApplicationSecurityGroups"),
						IgnoreInTests: true,
					},
					{
						Name:          "load_balancer_backend_address_pools",
						Description:   "The reference to LoadBalancerBackendAddressPool resource.",
						Type:          schema.TypeJSON,
						Resolver:      schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.LoadBalancerBackendAddressPools"),
						IgnoreInTests: true,
					},
					{
						Name:          "load_balancer_inbound_nat_rules",
						Description:   "A list of references of LoadBalancerInboundNatRules.",
						Type:          schema.TypeJSON,
						Resolver:      schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.LoadBalancerInboundNatRules"),
						IgnoreInTests: true,
					},
					{
						Name:          "private_ip_address",
						Description:   "Private IP address of the IP configuration.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.PrivateIPAddress"),
						IgnoreInTests: true,
					},
					{
						Name:        "private_ip_address_version",
						Description: "Whether the specific IP configuration is IPv4 or IPv6. Default is IPv4. Possible values include: 'IPVersionIPv4', 'IPVersionIPv6",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.PrivateIPAddressVersion"),
					},
					{
						Name:        "private_ip_allocation_method",
						Description: "Private IP address allocation method.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.PrivateIPAllocationMethod"),
					},
					{
						Name:          "private_link_connection_properties",
						Description:   "PrivateLinkConnection properties for the network interface.",
						Type:          schema.TypeJSON,
						Resolver:      resolveInterfaceIPConfigurationPrivateLinkConnectionProperties,
						IgnoreInTests: true,
					},
					{
						Name:        "provisioning_state",
						Description: "The provisioning state of the network interface IP configuration.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.ProvisioningState"),
					},
					{
						Name:          "public_ip_address",
						Description:   "Public IP address bound to the IP configuration.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.PublicIPAddress.ID"),
						IgnoreInTests: true,
					},
					{
						Name:          "subnet_id",
						Description:   "subnet ID of network interface ip configuration",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.Subnet.ID"),
						IgnoreInTests: true,
					},
					{
						Name:          "virtual_network_taps",
						Description:   "The reference to Virtual Network Taps.",
						Type:          schema.TypeJSON,
						Resolver:      schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.VirtualNetworkTaps"),
						IgnoreInTests: true,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchNetworkInterfaces(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.Interfaces
	response, err := svc.ListAll(ctx)
	if err != nil {
		return diag.WrapError(err)
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
func fetchNetworkInterfaceIPConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	ni := parent.Item.(network.Interface)
	if ni.IPConfigurations != nil {
		res <- *ni.IPConfigurations
	}
	return nil
}
func resolveNetworkInterfacePrivateLinkService(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(network.Interface)
	if p.InterfacePropertiesFormat == nil ||
		p.InterfacePropertiesFormat.PrivateLinkService == nil {
		return nil
	}

	out, err := json.Marshal(p.InterfacePropertiesFormat.PrivateLinkService)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, out))
}
func resolveNetworkInterfaceTapConfigurations(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(network.Interface)
	if p.InterfacePropertiesFormat == nil ||
		p.InterfacePropertiesFormat.TapConfigurations == nil {
		return nil
	}

	out, err := json.Marshal(p.InterfacePropertiesFormat.TapConfigurations)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, out))
}

func resolveInterfaceIPConfigurationPrivateLinkConnectionProperties(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(network.InterfaceIPConfiguration)
	if p.PrivateLinkConnectionProperties == nil {
		return nil
	}

	out, err := json.Marshal(map[string]interface{}{
		"fqdns":              p.PrivateLinkConnectionProperties.Fqdns,
		"requiredMemberName": p.PrivateLinkConnectionProperties.RequiredMemberName,
		"groupId":            p.PrivateLinkConnectionProperties.GroupID,
	})
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, out))
}
