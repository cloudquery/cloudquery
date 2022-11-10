// Auto generated code - DO NOT EDIT.

package container

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ManagedClusters() *schema.Table {
	return &schema.Table{
		Name:        "azure_container_managed_clusters",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2021-03-01/containerservice#ManagedCluster`,
		Resolver:    fetchContainerManagedClusters,
		Multiplex:   client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "power_state",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PowerState"),
			},
			{
				Name:     "max_agent_pools",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxAgentPools"),
			},
			{
				Name:     "kubernetes_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KubernetesVersion"),
			},
			{
				Name:     "dns_prefix",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DNSPrefix"),
			},
			{
				Name:     "fqdn_subdomain",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FqdnSubdomain"),
			},
			{
				Name:     "fqdn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Fqdn"),
			},
			{
				Name:     "private_fqdn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrivateFQDN"),
			},
			{
				Name:     "azure_portal_fqdn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AzurePortalFQDN"),
			},
			{
				Name:     "agent_pool_profiles",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AgentPoolProfiles"),
			},
			{
				Name:     "linux_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LinuxProfile"),
			},
			{
				Name:     "windows_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WindowsProfile"),
			},
			{
				Name:     "service_principal_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ServicePrincipalProfile"),
			},
			{
				Name:     "addon_profiles",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AddonProfiles"),
			},
			{
				Name:     "pod_identity_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PodIdentityProfile"),
			},
			{
				Name:     "node_resource_group",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NodeResourceGroup"),
			},
			{
				Name:     "enable_rbac",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableRBAC"),
			},
			{
				Name:     "enable_pod_security_policy",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnablePodSecurityPolicy"),
			},
			{
				Name:     "network_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkProfile"),
			},
			{
				Name:     "aad_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AadProfile"),
			},
			{
				Name:     "auto_upgrade_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AutoUpgradeProfile"),
			},
			{
				Name:     "auto_scaler_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AutoScalerProfile"),
			},
			{
				Name:     "api_server_access_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("APIServerAccessProfile"),
			},
			{
				Name:     "disk_encryption_set_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskEncryptionSetID"),
			},
			{
				Name:     "identity_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IdentityProfile"),
			},
			{
				Name:     "private_link_resources",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrivateLinkResources"),
			},
			{
				Name:     "disable_local_accounts",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DisableLocalAccounts"),
			},
			{
				Name:     "http_proxy_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HTTPProxyConfig"),
			},
			{
				Name:     "identity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Identity"),
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sku"),
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

func fetchContainerManagedClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Container.ManagedClusters

	response, err := svc.List(ctx)

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
