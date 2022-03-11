package network

import (
	"context"
	"encoding/json"
	"fmt"
	"net"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func NetworkPublicIPAddresses() *schema.Table {
	return &schema.Table{
		Name:         "azure_network_public_ip_addresses",
		Description:  "PublicIPAddress public IP address resource.",
		Resolver:     fetchNetworkPublicIpAddresses,
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
				Name:          "extended_location_name",
				Description:   "The name of the extended location.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ExtendedLocation.Name"),
				IgnoreInTests: true,
			},
			{
				Name:          "extended_location_type",
				Description:   "The type of the extended location.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ExtendedLocation.Type"),
				IgnoreInTests: true,
			},
			{
				Name:        "sku_name",
				Description: "Name of a public IP address SKU",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "Tier of a public IP address SKU",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Tier"),
			},
			{
				Name:        "public_ip_allocation_method",
				Description: "The public IP address allocation method",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.PublicIPAllocationMethod"),
			},
			{
				Name:        "public_ip_address_version",
				Description: "The public IP address version",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.PublicIPAddressVersion"),
			},
			{
				Name:          "ip_configuration",
				Description:   "The IP configuration associated with the public IP address.",
				Type:          schema.TypeJSON,
				Resolver:      resolveNetworkPublicIPAddressesIpConfiguration,
				IgnoreInTests: true,
			},
			{
				Name:          "dns_settings_domain_name_label",
				Description:   "The domain name label.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("PublicIPAddressPropertiesFormat.DNSSettings.DomainNameLabel"),
				IgnoreInTests: true,
			},
			{
				Name:          "dns_settings_fqdn",
				Description:   "The Fully Qualified Domain Name of the A DNS record associated with the public IP.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("PublicIPAddressPropertiesFormat.DNSSettings.Fqdn"),
				IgnoreInTests: true,
			},
			{
				Name:          "dns_settings_reverse_fqdn",
				Description:   "The reverse FQDN.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("PublicIPAddressPropertiesFormat.DNSSettings.ReverseFqdn"),
				IgnoreInTests: true,
			},
			{
				Name:          "ddos_settings_ddos_custom_policy_id",
				Description:   "Resource ID.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("PublicIPAddressPropertiesFormat.DdosSettings.DdosCustomPolicy.ID"),
				IgnoreInTests: true,
			},
			{
				Name:        "ddos_settings_protection_coverage",
				Description: "The DDoS protection policy customizability of the public IP",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.DdosSettings.ProtectionCoverage"),
			},
			{
				Name:          "ddos_settings_protected_ip",
				Description:   "Enables DDoS protection on the public IP.",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("PublicIPAddressPropertiesFormat.DdosSettings.ProtectedIP"),
				IgnoreInTests: true,
			},
			{
				Name:        "ip_tags",
				Description: "The list of tags associated with the public IP address.",
				Type:        schema.TypeJSON,
				Resolver:    resolveNetworkPublicIPAddressesIpTags,
			},
			{
				Name:        "ip_address",
				Description: "The IP address associated with the public IP address resource.",
				Type:        schema.TypeInet,
				Resolver:    resolveNetworkPublicIPAddressesIpAddress,
			},
			{
				Name:          "public_ip_prefix_id",
				Description:   "Resource ID.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("PublicIPAddressPropertiesFormat.PublicIPPrefix.ID"),
				IgnoreInTests: true,
			},
			{
				Name:        "idle_timeout_in_minutes",
				Description: "The idle timeout of the public IP address.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.IdleTimeoutInMinutes"),
			},
			{
				Name:        "resource_guid",
				Description: "The resource GUID property of the public IP address resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.ResourceGUID"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the public IP address resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.ProvisioningState"),
			},
			{
				Name:          "service_public_ip_address",
				Description:   "The service public IP address of the public IP address resource.",
				Type:          schema.TypeJSON,
				Resolver:      resolveNetworkPublicIPAddressesServicePublicIpAddress,
				IgnoreInTests: true,
			},
			{
				Name:          "nat_gateway",
				Description:   "The NatGateway for the Public IP address.",
				Type:          schema.TypeJSON,
				Resolver:      resolveNetworkPublicIPAddressesNatGateway,
				IgnoreInTests: true,
			},
			{
				Name:        "migration_phase",
				Description: "Migration phase of Public IP Address",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.MigrationPhase"),
			},
			{
				Name:          "linked_public_ip_address",
				Description:   "The linked public IP address of the public IP address resource.",
				Type:          schema.TypeJSON,
				Resolver:      resolveNetworkPublicIPAddressesLinkedPublicIpAddress,
				IgnoreInTests: true,
			},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated.",
				Type:        schema.TypeString,
			},
			{
				Name:        "zones",
				Description: "A list of availability zones denoting the IP allocated for the resource needs to come from.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "id",
				Description: "Resource ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type.",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "Resource location.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchNetworkPublicIpAddresses(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.PublicIPAddresses
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
func resolveNetworkPublicIPAddressesIpConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(network.PublicIPAddress)
	if !ok {
		return fmt.Errorf("expected to have network.SecurityGroup but got %T", resource.Item)
	}

	if p.PublicIPAddressPropertiesFormat == nil ||
		p.PublicIPAddressPropertiesFormat.IPConfiguration == nil {
		return nil
	}

	out, err := json.Marshal(p.PublicIPAddressPropertiesFormat.IPConfiguration)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, out)
}
func resolveNetworkPublicIPAddressesIpTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(network.PublicIPAddress)
	if !ok {
		return fmt.Errorf("expected to have network.SecurityGroup but got %T", resource.Item)
	}
	if p.IPTags == nil {
		return nil
	}
	j := map[string]interface{}{}
	for _, t := range *p.IPTags {
		j[*t.IPTagType] = *t.Tag
	}
	return resource.Set(c.Name, j)
}
func resolveNetworkPublicIPAddressesIpAddress(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(network.PublicIPAddress)
	if !ok {
		return fmt.Errorf("expected to have network.SecurityGroup but got %T", resource.Item)
	}
	if p.IPAddress == nil {
		return nil
	}
	i := net.ParseIP(*p.IPAddress)
	if i == nil {
		return fmt.Errorf("wrong format of IP: %s", *p.IPAddress)
	}
	return resource.Set(c.Name, i)
}
func resolveNetworkPublicIPAddressesServicePublicIpAddress(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(network.PublicIPAddress)
	if !ok {
		return fmt.Errorf("expected to have network.SecurityGroup but got %T", resource.Item)
	}

	if p.PublicIPAddressPropertiesFormat == nil ||
		p.PublicIPAddressPropertiesFormat.ServicePublicIPAddress == nil {
		return nil
	}

	out, err := json.Marshal(p.PublicIPAddressPropertiesFormat.ServicePublicIPAddress)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, out)
}
func resolveNetworkPublicIPAddressesNatGateway(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(network.PublicIPAddress)
	if !ok {
		return fmt.Errorf("expected to have network.SecurityGroup but got %T", resource.Item)
	}

	if p.PublicIPAddressPropertiesFormat == nil ||
		p.PublicIPAddressPropertiesFormat.NatGateway == nil {
		return nil
	}

	out, err := json.Marshal(p.PublicIPAddressPropertiesFormat.NatGateway)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, out)
}
func resolveNetworkPublicIPAddressesLinkedPublicIpAddress(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(network.PublicIPAddress)
	if !ok {
		return fmt.Errorf("expected to have network.SecurityGroup but got %T", resource.Item)
	}

	if p.PublicIPAddressPropertiesFormat == nil ||
		p.PublicIPAddressPropertiesFormat.LinkedPublicIPAddress == nil {
		return nil
	}

	out, err := json.Marshal(p.PublicIPAddressPropertiesFormat.LinkedPublicIPAddress)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, out)
}
