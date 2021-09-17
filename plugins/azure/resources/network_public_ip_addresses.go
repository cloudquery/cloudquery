package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func NetworkPublicIPAddresses() *schema.Table {
	return &schema.Table{
		Name:         "azure_network_public_ip_addresses",
		Description:  "PublicIPAddress public IP address resource",
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
				Name:        "extended_location_name",
				Description: "The name of the extended location",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExtendedLocation.Name"),
			},
			{
				Name:        "extended_location_type",
				Description: "The type of the extended location",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExtendedLocation.Type"),
			},
			{
				Name:        "sku_name",
				Description: "Name of a public IP address SKU Possible values include: 'PublicIPAddressSkuNameBasic', 'PublicIPAddressSkuNameStandard'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "Tier of a public IP address SKU Possible values include: 'PublicIPAddressSkuTierRegional', 'PublicIPAddressSkuTierGlobal'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Tier"),
			},
			{
				Name:        "public_ip_allocation_method",
				Description: "The public IP address allocation method Possible values include: 'Static', 'Dynamic'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.PublicIPAllocationMethod"),
			},
			{
				Name:        "public_ip_address_version",
				Description: "The public IP address version Possible values include: 'IPv4', 'IPv6'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.PublicIPAddressVersion"),
			},
			{
				Name:        "private_ip_address",
				Description: "The private IP address of the IP configuration",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.PrivateIPAddress"),
			},
			{
				Name:        "private_ip_allocation_method",
				Description: "The private IP address allocation method Possible values include: 'Static', 'Dynamic'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.PrivateIPAllocationMethod"),
			},
			{
				Name:        "subnet",
				Description: "The reference to the subnet resource",
				Type:        schema.TypeJSON,
				Resolver:    resolveNetworkPublicIPAddressSubnet,
			},
			{
				Name:        "public_ip_address",
				Description: "The reference to the public IP resource",
				Type:        schema.TypeJSON,
				Resolver:    resolveNetworkPublicIPAddressPublicIPAddress,
			},
			{
				Name:        "dns_settings_domain_name_label",
				Description: "The domain name label The concatenation of the domain name label and the regionalized DNS zone make up the fully qualified domain name associated with the public IP address If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.DNSSettings.DomainNameLabel"),
			},
			{
				Name:        "dns_settings_fqdn",
				Description: "The Fully Qualified Domain Name of the A DNS record associated with the public IP This is the concatenation of the domainNameLabel and the regionalized DNS zone",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.DNSSettings.Fqdn"),
			},
			{
				Name:        "dns_settings_reverse_fqdn",
				Description: "The reverse FQDN A user-visible, fully qualified domain name that resolves to this public IP address If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addrarpa domain to the reverse FQDN",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.DNSSettings.ReverseFqdn"),
			},
			{
				Name:        "ddos_settings_ddos_custom_policy_id",
				Description: "Resource ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.DdosSettings.DdosCustomPolicy.ID"),
			},
			{
				Name:        "ddos_settings_protection_coverage",
				Description: "The DDoS protection policy customizability of the public IP Only standard coverage will have the ability to be customized Possible values include: 'DdosSettingsProtectionCoverageBasic', 'DdosSettingsProtectionCoverageStandard'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.DdosSettings.ProtectionCoverage"),
			},
			{
				Name:        "ddos_settings_protected_ip",
				Description: "Enables DDoS protection on the public IP",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.DdosSettings.ProtectedIP"),
			},
			{
				Name:        "ip_address",
				Description: "The IP address associated with the public IP address resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.IPAddress"),
			},
			{
				Name:        "public_ip_prefix_id",
				Description: "Resource ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.PublicIPPrefix.ID"),
			},
			{
				Name:        "idle_timeout_in_minutes",
				Description: "The idle timeout of the public IP address",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.IdleTimeoutInMinutes"),
			},
			{
				Name:        "resource_guid",
				Description: "The resource GUID property of the public IP address resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.ResourceGUID"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the public IP address resource Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PublicIPAddressPropertiesFormat.ProvisioningState"),
			},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated",
				Type:        schema.TypeString,
			},
			{
				Name:        "zones",
				Description: "A list of availability zones denoting the IP allocated for the resource needs to come from",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "id",
				Description: "Resource ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource name",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "Resource location",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_network_public_ip_address_ip_tags",
				Description: "IPTag contains the IpTag associated with the object",
				Resolver:    fetchNetworkPublicIpAddressIpTags,
				Columns: []schema.Column{
					{
						Name:        "public_ip_address_cq_id",
						Description: "Unique CloudQuery ID of azure_network_public_ip_addresses table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "ip_tag_type",
						Description: "The IP tag type Example: FirstPartyUsage",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IPTagType"),
					},
					{
						Name:        "tag",
						Description: "The value of the IP tag associated with the public IP Example: SQL",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchNetworkPublicIpAddresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Network.PublicIPAddresses
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
func resolveNetworkPublicIPAddressSubnet(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(network.PublicIPAddress)
	if !ok {
		return fmt.Errorf("expected to have network.SecurityGroup but got %T", resource.Item)
	}

	if p.PublicIPAddressPropertiesFormat == nil ||
		p.PublicIPAddressPropertiesFormat.IPConfiguration == nil ||
		p.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat == nil ||
		p.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.Subnet == nil {
		return nil
	}

	out, err := json.Marshal(p.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.Subnet)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out)
}
func resolveNetworkPublicIPAddressPublicIPAddress(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(network.PublicIPAddress)
	if !ok {
		return fmt.Errorf("expected to have network.SecurityGroup but got %T", resource.Item)
	}

	if p.PublicIPAddressPropertiesFormat == nil ||
		p.PublicIPAddressPropertiesFormat.IPConfiguration == nil ||
		p.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat == nil ||
		p.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.PublicIPAddress == nil {
		return nil
	}

	out, err := json.Marshal(p.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.PublicIPAddress)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out)
}
func fetchNetworkPublicIpAddressIpTags(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(network.PublicIPAddress)
	if !ok {
		return fmt.Errorf("expected to have network.SecurityGroup but got %T", parent.Item)
	}

	if p.IPTags == nil {
		return nil
	}
	res <- *p.IPTags
	return nil
}
