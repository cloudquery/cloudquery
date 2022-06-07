package container

import (
	"context"
	"net"

	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ContainerRegistries() *schema.Table {
	return &schema.Table{
		Name:         "azure_container_registries",
		Description:  "Azure compute disk",
		Resolver:     fetchContainerRegistries,
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
				Name:        "sku_name",
				Description: "The SKU name of the container registry",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "The SKU tier based on the SKU name",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Tier"),
			},
			{
				Name:        "login_server",
				Description: "The URL that can be used to log into the container registry",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegistryProperties.LoginServer"),
			},
			{
				Name:        "creation_date",
				Description: "The creation date of the container registry",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("RegistryProperties.CreationDate.Time"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the container registry at the time the operation was called",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegistryProperties.ProvisioningState"),
			},
			{
				Name:          "status",
				Description:   "The short label for the status",
				IgnoreInTests: true,
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("RegistryProperties.Status.DisplayStatus"),
			},
			{
				Name:          "status_message",
				Description:   "The detailed message for the status, including alerts and error messages",
				Type:          schema.TypeString,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("RegistryProperties.Status.Message"),
			},
			{
				Name:        "status_timestamp",
				Description: "The timestamp when the status was changed to the current value",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("RegistryProperties.Status.Timestamp.Time"),
			},
			{
				Name:        "admin_user_enabled",
				Description: "The value that indicates whether the admin user is enabled",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("RegistryProperties.AdminUserEnabled"),
			},
			{
				Name:          "storage_account_id",
				Description:   "The resource ID of the storage account",
				Type:          schema.TypeString,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("RegistryProperties.StorageAccount.ID"),
			},
			{
				Name:        "network_rule_set_default_action",
				Description: "The default action of allow or deny when no other rules match",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegistryProperties.NetworkRuleSet.DefaultAction"),
			},
			{
				Name:        "quarantine_policy_status",
				Description: "The value that indicates whether the policy is enabled or not",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegistryProperties.Policies.QuarantinePolicy.Status"),
			},
			{
				Name:        "trust_policy_type",
				Description: "The type of trust policy",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegistryProperties.Policies.TrustPolicy.Type"),
			},
			{
				Name:        "trust_policy_status",
				Description: "The value that indicates whether the policy is enabled or not",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegistryProperties.Policies.TrustPolicy.Status"),
			},
			{
				Name:        "retention_policy_days",
				Description: "The number of days to retain an untagged manifest after which it gets purged",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("RegistryProperties.Policies.RetentionPolicy.Days"),
			},
			{
				Name:        "retention_policy_last_updated_time",
				Description: "The timestamp when the policy was last updated",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("RegistryProperties.Policies.RetentionPolicy.LastUpdatedTime.Time"),
			},
			{
				Name:        "retention_policy_status",
				Description: "The value that indicates whether the policy is enabled or not",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegistryProperties.Policies.RetentionPolicy.Status"),
			},
			{
				Name:        "id",
				Description: "The resource ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "The location of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The tags of the resource",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "azure_container_registry_network_rule_set_virtual_network_rules",
				Description:   "VirtualNetworkRule virtual network rule",
				Resolver:      fetchContainerRegistryNetworkRuleSetVirtualNetworkRules,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "registry_cq_id",
						Description: "Unique CloudQuery ID of azure_container_registries table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "action",
						Description: "The action of virtual network rule",
						Type:        schema.TypeString,
					},
					{
						Name:        "virtual_network_id",
						Description: "Resource ID of a subnet",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualNetworkResourceID"),
					},
				},
			},
			{
				Name:          "azure_container_registry_network_rule_set_ip_rules",
				Description:   "IPRule IP rule with specific IP or IP range in CIDR format",
				Resolver:      fetchContainerRegistryNetworkRuleSetIpRules,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "registry_cq_id",
						Description: "Unique CloudQuery ID of azure_container_registries table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "action",
						Description: "The action of IP ACL rule",
						Type:        schema.TypeString,
					},
					{
						Name:        "ip_address_or_range",
						Description: "Specifies the IP or IP range in CIDR format",
						Type:        schema.TypeCIDR,
						Resolver:    resolveContainerRegistryNetworkRuleSetIPRulesIpAddressOrRange,
					},
				},
			},
			{
				Name:          "azure_container_registry_replications",
				Description:   "Replication an object that represents a replication for a container registry",
				Resolver:      fetchContainerRegistryReplications,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "registry_cq_id",
						Description: "Unique CloudQuery ID of azure_container_registries table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "provisioning_state",
						Description: "The provisioning state of the replication at the time the operation was called",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ReplicationProperties.ProvisioningState"),
					},
					{
						Name:        "status",
						Description: "The short label for the status",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ReplicationProperties.Status.DisplayStatus"),
					},
					{
						Name:        "status_message",
						Description: "The detailed message for the status, including alerts and error messages",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ReplicationProperties.Status.Message"),
					},
					{
						Name:        "status_timestamp",
						Description: "The timestamp when the status was changed to the current value",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("ReplicationProperties.Status.Timestamp.Time"),
					},
					{
						Name:        "id",
						Description: "The resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "The name of the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The type of the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "location",
						Description: "The location of the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "tags",
						Description: "The tags of the resource",
						Type:        schema.TypeJSON,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchContainerRegistries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ContainerRegistry.Registries
	result, err := svc.List(ctx)
	if err != nil {
		return diag.WrapError(err)
	}
	for result.NotDone() {
		res <- result.Values()
		if err := result.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
func fetchContainerRegistryNetworkRuleSetVirtualNetworkRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(containerregistry.Registry)
	if r.RegistryProperties == nil || r.RegistryProperties.NetworkRuleSet == nil || r.RegistryProperties.NetworkRuleSet.VirtualNetworkRules == nil {
		return nil
	}
	res <- *r.RegistryProperties.NetworkRuleSet.VirtualNetworkRules
	return nil
}
func fetchContainerRegistryNetworkRuleSetIpRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(containerregistry.Registry)
	if r.RegistryProperties == nil || r.RegistryProperties.NetworkRuleSet == nil || r.RegistryProperties.NetworkRuleSet.IPRules == nil {
		return nil
	}
	res <- *r.RegistryProperties.NetworkRuleSet.IPRules
	return nil
}
func resolveContainerRegistryNetworkRuleSetIPRulesIpAddressOrRange(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(containerregistry.IPRule)
	_, cidr, err := net.ParseCIDR(*r.IPAddressOrRange)
	if err != nil {
		ip := net.ParseIP(*r.IPAddressOrRange)
		if ip == nil {
			return diag.WrapError(err)
		}
		if v4 := ip.To4(); v4 != nil {
			return diag.WrapError(resource.Set(c.Name, &net.IPNet{IP: v4, Mask: net.CIDRMask(32, 32)}))
		}
		if v6 := ip.To16(); v6 != nil {
			return diag.WrapError(resource.Set(c.Name, &net.IPNet{IP: v6, Mask: net.CIDRMask(128, 128)}))
		}
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set(c.Name, cidr))
}
func fetchContainerRegistryReplications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(containerregistry.Registry)
	svc := meta.(*client.Client).Services().ContainerRegistry.Replications
	resource, err := client.ParseResourceID(*r.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	result, err := svc.List(ctx, resource.ResourceGroup, *r.Name)
	if err != nil {
		return diag.WrapError(err)
	}
	for result.NotDone() {
		res <- result.Values()
		if err := result.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
