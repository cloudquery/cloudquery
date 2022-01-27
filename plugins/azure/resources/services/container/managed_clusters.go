package container

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2021-03-01/containerservice"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ContainerManagedClusters() *schema.Table {
	return &schema.Table{
		Name:         "azure_container_managed_clusters",
		Description:  "ManagedCluster managed cluster",
		Resolver:     fetchContainerManagedClusters,
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
				Name:        "provisioning_state",
				Description: "The current deployment or provisioning state.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedClusterProperties.ProvisioningState"),
			},
			{
				Name:        "power_state_code",
				Description: "Tells whether the cluster is Running or Stopped.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedClusterProperties.PowerState.Code"),
			},
			{
				Name:        "max_agent_pools",
				Description: "The max number of agent pools for the managed cluster",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ManagedClusterProperties.MaxAgentPools"),
			},
			{
				Name:        "kubernetes_version",
				Description: "Version of Kubernetes specified when creating the managed cluster",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedClusterProperties.KubernetesVersion"),
			},
			{
				Name:        "dns_prefix",
				Description: "DNS prefix specified when creating the managed cluster",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedClusterProperties.DNSPrefix"),
			},
			{
				Name:          "fqdn_subdomain",
				Description:   "FQDN subdomain specified when creating private cluster with custom private dns zone",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ManagedClusterProperties.FqdnSubdomain"),
				IgnoreInTests: true,
			},
			{
				Name:        "fqdn",
				Description: "FQDN for the master pool",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedClusterProperties.Fqdn"),
			},
			{
				Name:          "private_fqdn",
				Description:   "FQDN of private cluster",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ManagedClusterProperties.PrivateFQDN"),
				IgnoreInTests: true,
			},
			{
				Name:        "azure_portal_fqdn",
				Description: "FQDN for the master pool which used by proxy config",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedClusterProperties.AzurePortalFQDN"),
			},
			{
				Name:          "linux_profile_admin_username",
				Description:   "The administrator username to use for Linux VMs",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ManagedClusterProperties.LinuxProfile.AdminUsername"),
				IgnoreInTests: true,
			},
			{
				Name:        "windows_profile_admin_username",
				Description: "Specifies the name of the administrator account.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedClusterProperties.WindowsProfile.AdminUsername"),
			},
			{
				Name:          "windows_profile_admin_password",
				Description:   "Specifies the password of the administrator account.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ManagedClusterProperties.WindowsProfile.AdminPassword"),
				IgnoreInTests: true,
			},
			{
				Name:        "windows_profile_license_type",
				Description: "The licenseType to use for Windows VMs.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedClusterProperties.WindowsProfile.LicenseType"),
			},
			{
				Name:        "windows_profile_enable_csi_proxy",
				Description: "Whether to enable CSI proxy",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ManagedClusterProperties.WindowsProfile.EnableCSIProxy"),
			},
			{
				Name:        "service_principal_profile_client_id",
				Description: "The ID for the service principal",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedClusterProperties.ServicePrincipalProfile.ClientID"),
			},
			{
				Name:          "service_principal_profile_secret",
				Description:   "The secret password associated with the service principal in plain text",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ManagedClusterProperties.ServicePrincipalProfile.Secret"),
				IgnoreInTests: true,
			},
			{
				Name:        "addon_profiles",
				Description: "Profile of managed cluster add-on",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ManagedClusterProperties.AddonProfiles"),
			},
			{
				Name:          "pod_identity_profile_enabled",
				Description:   "Whether the pod identity addon is enabled",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("ManagedClusterProperties.PodIdentityProfile.Enabled"),
				IgnoreInTests: true,
			},
			{
				Name:          "pod_identity_profile_allow_network_plugin_kubenet",
				Description:   "Customer consent for enabling AAD pod identity addon in cluster using Kubenet network plugin",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("ManagedClusterProperties.PodIdentityProfile.AllowNetworkPluginKubenet"),
				IgnoreInTests: true,
			},
			{
				Name:        "node_resource_group",
				Description: "Name of the resource group containing agent pool nodes",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedClusterProperties.NodeResourceGroup"),
			},
			{
				Name:        "enable_rbac",
				Description: "Whether to enable Kubernetes Role-Based Access Control",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ManagedClusterProperties.EnableRBAC"),
			},
			{
				Name:        "network_profile_network_plugin",
				Description: "Network plugin used for building Kubernetes network.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedClusterProperties.NetworkProfile.NetworkPlugin"),
			},
			{
				Name:        "network_profile_network_policy",
				Description: "Network policy used for building Kubernetes network.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedClusterProperties.NetworkProfile.NetworkPolicy"),
			},
			{
				Name:        "network_profile_network_mode",
				Description: "Network mode used for building Kubernetes network.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedClusterProperties.NetworkProfile.NetworkMode"),
			},
			{
				Name:          "network_profile_pod_cidr",
				Description:   "A CIDR notation IP range from which to assign pod IPs when kubenet is used",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ManagedClusterProperties.NetworkProfile.PodCidr"),
				IgnoreInTests: true,
			},
			{
				Name:        "network_profile_service_cidr",
				Description: "A CIDR notation IP range from which to assign service cluster IPs.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedClusterProperties.NetworkProfile.ServiceCidr"),
			},
			{
				Name:        "network_profile_dns_service_ip",
				Description: "An IP address assigned to the Kubernetes DNS service.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedClusterProperties.NetworkProfile.DNSServiceIP"),
			},
			{
				Name:        "network_profile_docker_bridge_cidr",
				Description: "A CIDR notation IP range assigned to the Docker bridge network.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedClusterProperties.NetworkProfile.DockerBridgeCidr"),
			},
			{
				Name:        "network_profile_outbound_type",
				Description: "The outbound (egress) routing method.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedClusterProperties.NetworkProfile.OutboundType"),
			},
			{
				Name:        "network_profile_load_balancer_sku",
				Description: "The load balancer sku for the managed cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedClusterProperties.NetworkProfile.LoadBalancerSku"),
			},
			{
				Name:        "network_profile_load_balancer_managed_outbound_ips_count",
				Description: "Desired number of outbound IP created/managed by Azure for the cluster load balancer.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ManagedClusterProperties.NetworkProfile.LoadBalancerProfile.ManagedOutboundIPs.Count"),
			},
			{
				Name:          "network_profile_load_balancer_outbound_ip_prefixes",
				Description:   "A list of public IP prefix resources",
				Type:          schema.TypeStringArray,
				Resolver:      resolveContainerManagedClusterNetworkProfileLoadBalancerOutboundIPPrefixes,
				IgnoreInTests: true,
			},
			{
				Name:          "network_profile_load_balancer_outbound_ips",
				Description:   "A list of public IP resources",
				Type:          schema.TypeStringArray,
				Resolver:      resolveContainerManagedClusterNetworkProfileLoadBalancerOutboundIps,
				IgnoreInTests: true,
			},
			{
				Name:        "network_profile_load_balancer_effective_outbound_ips",
				Description: "The effective outbound IP resources of the cluster load balancer",
				Type:        schema.TypeStringArray,
				Resolver:    resolveContainerManagedClusterNetworkProfileLoadBalancerEffectiveOutboundIps,
			},
			{
				Name:          "network_profile_load_balancer_allocated_outbound_ports",
				Description:   "Desired number of allocated SNAT ports per VM.",
				Type:          schema.TypeInt,
				Resolver:      schema.PathResolver("ManagedClusterProperties.NetworkProfile.LoadBalancerProfile.AllocatedOutboundPorts"),
				IgnoreInTests: true,
			},
			{
				Name:          "network_profile_load_balancer_idle_timeout",
				Description:   "Desired outbound flow idle timeout in minutes.",
				Type:          schema.TypeInt,
				Resolver:      schema.PathResolver("ManagedClusterProperties.NetworkProfile.LoadBalancerProfile.IdleTimeoutInMinutes"),
				IgnoreInTests: true,
			},
			{
				Name:          "aad_profile_managed",
				Description:   "Whether to enable managed AAD",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("ManagedClusterProperties.AadProfile.Managed"),
				IgnoreInTests: true,
			},
			{
				Name:          "aad_profile_enable_azure_rbac",
				Description:   "Whether to enable Azure RBAC for Kubernetes authorization",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("ManagedClusterProperties.AadProfile.EnableAzureRBAC"),
				IgnoreInTests: true,
			},
			{
				Name:          "aad_profile_admin_group_object_ids",
				Description:   "AAD group object IDs that will have admin role of the cluster",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("ManagedClusterProperties.AadProfile.AdminGroupObjectIDs"),
				IgnoreInTests: true,
			},
			{
				Name:          "aad_profile_client_app_id",
				Description:   "The client AAD application ID",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ManagedClusterProperties.AadProfile.ClientAppID"),
				IgnoreInTests: true,
			},
			{
				Name:          "aad_profile_server_app_id",
				Description:   "The server AAD application ID",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ManagedClusterProperties.AadProfile.ServerAppID"),
				IgnoreInTests: true,
			},
			{
				Name:          "aad_profile_server_app_secret",
				Description:   "The server AAD application secret",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ManagedClusterProperties.AadProfile.ServerAppSecret"),
				IgnoreInTests: true,
			},
			{
				Name:          "aad_profile_tenant_id",
				Description:   "The AAD tenant ID to use for authentication.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ManagedClusterProperties.AadProfile.TenantID"),
				IgnoreInTests: true,
			},
			{
				Name:        "auto_upgrade_profile_upgrade_channel",
				Description: "upgrade channel for auto upgrade.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedClusterProperties.AutoUpgradeProfile.UpgradeChannel"),
			},
			{
				Name:        "auto_scaler_profile_expander",
				Description: "Possible values include: 'LeastWaste', 'MostPods', 'Priority', 'Random'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedClusterProperties.AutoScalerProfile.Expander"),
			},
			{
				Name:          "api_server_access_profile_authorized_ip_ranges",
				Description:   "Authorized IP Ranges to kubernetes API server",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("ManagedClusterProperties.APIServerAccessProfile.AuthorizedIPRanges"),
				IgnoreInTests: true,
			},
			{
				Name:        "api_server_access_profile_enable_private_cluster",
				Description: "Whether to create the cluster as a private cluster or not",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ManagedClusterProperties.APIServerAccessProfile.EnablePrivateCluster"),
			},
			{
				Name:          "api_server_access_profile_private_dns_zone",
				Description:   "Private dns zone mode for private cluster",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ManagedClusterProperties.APIServerAccessProfile.PrivateDNSZone"),
				IgnoreInTests: true,
			},
			{
				Name:          "disk_encryption_set_id",
				Description:   "ResourceId of the disk encryption set to use for enabling encryption at rest",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ManagedClusterProperties.DiskEncryptionSetID"),
				IgnoreInTests: true,
			},
			{
				Name:        "identity_profile",
				Description: "Identities associated with the cluster",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ManagedClusterProperties.IdentityProfile"),
			},
			{
				Name:        "disable_local_accounts",
				Description: "If set to true, getting static credential will be disabled for this cluster.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ManagedClusterProperties.DisableLocalAccounts"),
			},
			{
				Name:          "http_proxy_config_http_proxy",
				Description:   "HTTP proxy server endpoint to use",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ManagedClusterProperties.HTTPProxyConfig.HTTPProxy"),
				IgnoreInTests: true,
			},
			{
				Name:          "http_proxy_config_https_proxy",
				Description:   "HTTPS proxy server endpoint to use",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ManagedClusterProperties.HTTPProxyConfig.HTTPSProxy"),
				IgnoreInTests: true,
			},
			{
				Name:          "http_proxy_config_no_proxy",
				Description:   "Endpoints that should not go through proxy",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("ManagedClusterProperties.HTTPProxyConfig.NoProxy"),
				IgnoreInTests: true,
			},
			{
				Name:          "http_proxy_config_trusted_ca",
				Description:   "Alternative CA cert to use for connecting to proxy servers",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ManagedClusterProperties.HTTPProxyConfig.TrustedCa"),
				IgnoreInTests: true,
			},
			{
				Name:        "identity_principal_id",
				Description: "The principal id of the system assigned identity which is used by master components",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.PrincipalID"),
			},
			{
				Name:        "identity_tenant_id",
				Description: "The tenant id of the system assigned identity which is used by master components",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.TenantID"),
			},
			{
				Name:        "identity_type",
				Description: "The type of identity used for the managed cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Type"),
			},
			{
				Name:          "identity_user_assigned_identities",
				Description:   "The user identity associated with the managed cluster.",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("Identity.UserAssignedIdentities"),
				IgnoreInTests: true,
			},
			{
				Name:        "sku_name",
				Description: "Name of a managed cluster SKU.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "Tier of a managed cluster SKU.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Tier"),
			},
			{
				Name:          "extended_location_name",
				Description:   "The name of the extended location",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ExtendedLocation.Name"),
				IgnoreInTests: true,
			},
			{
				Name:        "extended_location_type",
				Description: "The type of the extended location.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExtendedLocation.Type"),
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
				Name:          "azure_container_managed_cluster_pip_user_assigned_id_exceptions",
				Description:   "ManagedClusterPodIdentityException",
				Resolver:      fetchContainerManagedClusterPipUserAssignedIdentityExceptions,
				Options:       schema.TableCreationOptions{PrimaryKeys: []string{"managed_cluster_cq_id", "name"}},
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "managed_cluster_cq_id",
						Description: "Unique CloudQuery ID of azure_container_managed_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "Name of the pod identity exception",
						Type:        schema.TypeString,
					},
					{
						Name:        "namespace",
						Description: "Namespace of the pod identity exception",
						Type:        schema.TypeString,
					},
					{
						Name:        "pod_labels",
						Description: "Pod labels to match",
						Type:        schema.TypeJSON,
					},
				},
			},
			{
				Name:          "azure_container_managed_cluster_private_link_resources",
				Description:   "PrivateLinkResource a private link resource",
				Resolver:      fetchContainerManagedClusterPrivateLinkResources,
				Options:       schema.TableCreationOptions{PrimaryKeys: []string{"managed_cluster_cq_id", "id"}},
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "managed_cluster_cq_id",
						Description: "Unique CloudQuery ID of azure_container_managed_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "The ID of the private link resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "The name of the private link resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The resource type",
						Type:        schema.TypeString,
					},
					{
						Name:        "group_id",
						Description: "The group ID of the resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("GroupID"),
					},
					{
						Name:        "required_members",
						Description: "RequiredMembers of the resource",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "private_link_service_id",
						Description: "The private link service ID of the resource, this field is exposed only to NRP internally",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateLinkServiceID"),
					},
				},
			},
			{
				Name:        "azure_container_managed_cluster_agent_pool_profiles",
				Description: "ManagedClusterAgentPoolProfile profile for the container service agent pool",
				Resolver:    fetchContainerManagedClusterAgentPoolProfiles,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"managed_cluster_cq_id", "name"}},
				Columns: []schema.Column{
					{
						Name:        "managed_cluster_cq_id",
						Description: "Unique CloudQuery ID of azure_container_managed_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "Unique name of the agent pool profile in the context of the subscription and resource group",
						Type:        schema.TypeString,
					},
					{
						Name:        "count",
						Description: "Number of agents (VMs) to host docker containers.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "vm_size",
						Description: "Size of agent VMs",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VMSize"),
					},
					{
						Name:        "os_disk_size_gb",
						Description: "OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("OsDiskSizeGB"),
					},
					{
						Name:        "os_disk_type",
						Description: "OS disk type to be used for machines in a given agent pool.",
						Type:        schema.TypeString,
					},
					{
						Name:        "kubelet_disk_type",
						Description: "KubeletDiskType determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.",
						Type:        schema.TypeString,
					},
					{
						Name:        "vnet_subnet_id",
						Description: "VNet SubnetID specifies the VNet's subnet identifier for nodes and maybe pods",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VnetSubnetID"),
					},
					{
						Name:          "pod_subnet_id",
						Description:   "Pod SubnetID specifies the VNet's subnet identifier for pods",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("PodSubnetID"),
						IgnoreInTests: true,
					},
					{
						Name:        "max_pods",
						Description: "Maximum number of pods that can run on a node",
						Type:        schema.TypeInt,
					},
					{
						Name:        "os_type",
						Description: "OsType to be used to specify os type.",
						Type:        schema.TypeString,
					},
					{
						Name:        "os_sku",
						Description: "OsSKU to be used to specify os sku.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("OsSKU"),
					},
					{
						Name:        "max_count",
						Description: "Maximum number of nodes for auto-scaling",
						Type:        schema.TypeInt,
					},
					{
						Name:        "min_count",
						Description: "Minimum number of nodes for auto-scaling",
						Type:        schema.TypeInt,
					},
					{
						Name:        "enable_auto_scaling",
						Description: "Whether to enable auto-scaler",
						Type:        schema.TypeBool,
					},
					{
						Name:        "type",
						Description: "AgentPoolType represents types of an agent pool.",
						Type:        schema.TypeString,
					},
					{
						Name:        "mode",
						Description: "AgentPoolMode represents mode of an agent pool.",
						Type:        schema.TypeString,
					},
					{
						Name:        "orchestrator_version",
						Description: "Version of orchestrator specified when creating the managed cluster",
						Type:        schema.TypeString,
					},
					{
						Name:        "node_image_version",
						Description: "Version of node image",
						Type:        schema.TypeString,
					},
					{
						Name:          "upgrade_settings_max_surge",
						Description:   "Count or percentage of additional nodes to be added during upgrade.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("UpgradeSettings.MaxSurge"),
						IgnoreInTests: true,
					},
					{
						Name:        "provisioning_state",
						Description: "The current deployment or provisioning state, which only appears in the response",
						Type:        schema.TypeString,
					},
					{
						Name:        "power_state_code",
						Description: "Tells whether the cluster is Running or Stopped.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PowerState.Code"),
					},
					{
						Name:          "availability_zones",
						Description:   "Availability zones for nodes.",
						Type:          schema.TypeStringArray,
						IgnoreInTests: true,
					},
					{
						Name:        "enable_node_public_ip",
						Description: "Enable public IP for nodes",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("EnableNodePublicIP"),
					},
					{
						Name:          "node_public_ip_prefix_id",
						Description:   "Public IP Prefix ID VM nodes use IPs assigned from this Public IP Prefix",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("NodePublicIPPrefixID"),
						IgnoreInTests: true,
					},
					{
						Name:        "scale_set_priority",
						Description: "ScaleSetPriority to be used to specify virtual machine scale set priority.",
						Type:        schema.TypeString,
					},
					{
						Name:        "scale_set_eviction_policy",
						Description: "ScaleSetEvictionPolicy to be used to specify eviction policy for Spot virtual machine scale set.",
						Type:        schema.TypeString,
					},
					{
						Name:          "spot_max_price",
						Description:   "SpotMaxPrice to be used to specify the maximum price you are willing to pay in US Dollars.",
						Type:          schema.TypeFloat,
						IgnoreInTests: true,
					},
					{
						Name:        "tags",
						Description: "Agent pool tags to be persisted on the agent pool virtual machine scale set",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "node_labels",
						Description: "Agent pool node labels to be persisted across all nodes in agent pool",
						Type:        schema.TypeJSON,
					},
					{
						Name:          "node_taints",
						Description:   "Taints added to new nodes during node pool create and scale.",
						Type:          schema.TypeStringArray,
						IgnoreInTests: true,
					},
					{
						Name:          "proximity_placement_group_id",
						Description:   "The ID for Proximity Placement Group",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("ProximityPlacementGroupID"),
						IgnoreInTests: true,
					},
					{
						Name:          "kubelet_config_cpu_manager_policy",
						Description:   "CPU Manager policy to use",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("KubeletConfig.CPUManagerPolicy"),
						IgnoreInTests: true,
					},
					{
						Name:          "kubelet_config_cpu_cfs_quota",
						Description:   "Enable CPU CFS quota enforcement for containers that specify CPU limits",
						Type:          schema.TypeBool,
						Resolver:      schema.PathResolver("KubeletConfig.CPUCfsQuota"),
						IgnoreInTests: true,
					},
					{
						Name:          "kubelet_config_cpu_cfs_quota_period",
						Description:   "Sets CPU CFS quota period value",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("KubeletConfig.CPUCfsQuotaPeriod"),
						IgnoreInTests: true,
					},
					{
						Name:          "kubelet_config_image_gc_high_threshold",
						Description:   "The percent of disk usage after which image garbage collection is always run",
						Type:          schema.TypeInt,
						Resolver:      schema.PathResolver("KubeletConfig.ImageGcHighThreshold"),
						IgnoreInTests: true,
					},
					{
						Name:          "kubelet_config_image_gc_low_threshold",
						Description:   "The percent of disk usage before which image garbage collection is never run",
						Type:          schema.TypeInt,
						Resolver:      schema.PathResolver("KubeletConfig.ImageGcLowThreshold"),
						IgnoreInTests: true,
					},
					{
						Name:          "kubelet_config_topology_manager_policy",
						Description:   "Topology Manager policy to use",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("KubeletConfig.TopologyManagerPolicy"),
						IgnoreInTests: true,
					},
					{
						Name:          "kubelet_config_allowed_unsafe_sysctls",
						Description:   "Allowlist of unsafe sysctls or unsafe sysctl patterns (ending in `*`)",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("KubeletConfig.AllowedUnsafeSysctls"),
						IgnoreInTests: true,
					},
					{
						Name:          "kubelet_config_fail_swap_on",
						Description:   "If set to true it will make the Kubelet fail to start if swap is enabled on the node",
						Type:          schema.TypeBool,
						Resolver:      schema.PathResolver("KubeletConfig.FailSwapOn"),
						IgnoreInTests: true,
					},
					{
						Name:          "kubelet_config_container_log_max_size_mb",
						Description:   "The maximum size (eg 10Mi) of container log file before it is rotated",
						Type:          schema.TypeInt,
						Resolver:      schema.PathResolver("KubeletConfig.ContainerLogMaxSizeMB"),
						IgnoreInTests: true,
					},
					{
						Name:          "kubelet_config_container_log_max_files",
						Description:   "The maximum number of container log files that can be present for a container.",
						Type:          schema.TypeInt,
						Resolver:      schema.PathResolver("KubeletConfig.ContainerLogMaxFiles"),
						IgnoreInTests: true,
					},
					{
						Name:          "kubelet_config_pod_max_pids",
						Description:   "The maximum number of processes per pod",
						Type:          schema.TypeInt,
						Resolver:      schema.PathResolver("KubeletConfig.PodMaxPids"),
						IgnoreInTests: true,
					},
					{
						Name:          "linux_os_config",
						Description:   "LinuxOSConfig specifies the OS configuration of linux agent nodes",
						Type:          schema.TypeJSON,
						Resolver:      resolveContainerManagedClusterAgentPoolProfileLinuxOsConfig,
						IgnoreInTests: true,
					},
					{
						Name:        "enable_encryption_at_host",
						Description: "Whether to enable EncryptionAtHost",
						Type:        schema.TypeBool,
					},
					{
						Name:        "enable_fips",
						Description: "Whether to use FIPS enabled OS",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("EnableFIPS"),
					},
					{
						Name:        "gpu_instance_profile",
						Description: "GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "azure_container_managed_cluster_pip_user_assigned_identities",
				Description:   "ManagedClusterPodIdentity",
				Resolver:      fetchContainerManagedClusterPipUserAssignedIdentities,
				Options:       schema.TableCreationOptions{PrimaryKeys: []string{"managed_cluster_cq_id", "name"}},
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "managed_cluster_cq_id",
						Description: "Unique CloudQuery ID of azure_container_managed_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "Name of the pod identity",
						Type:        schema.TypeString,
					},
					{
						Name:        "namespace",
						Description: "Namespace of the pod identity",
						Type:        schema.TypeString,
					},
					{
						Name:        "binding_selector",
						Description: "Binding selector to use for the AzureIdentityBinding resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "identity_resource_id",
						Description: "The resource id of the user assigned identity",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Identity.ResourceID"),
					},
					{
						Name:        "identity_client_id",
						Description: "The client id of the user assigned identity",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Identity.ClientID"),
					},
					{
						Name:        "identity_object_id",
						Description: "The object id of the user assigned identity",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Identity.ObjectID"),
					},
					{
						Name:        "provisioning_state",
						Description: "The current provisioning state of the pod identity.",
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
func fetchContainerManagedClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Container.ManagedClusters
	result, err := svc.List(ctx)
	if err != nil {
		return err
	}
	for result.NotDone() {
		res <- result.Values()
		if err := result.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil
}

func resolveContainerManagedClusterNetworkProfileLoadBalancerOutboundIPPrefixes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	mc, ok := resource.Item.(containerservice.ManagedCluster)
	if !ok {
		return fmt.Errorf("not a containerservice.ManagedCluster instance: %T", resource.Item)
	}
	if mc.NetworkProfile == nil ||
		mc.NetworkProfile.LoadBalancerProfile == nil ||
		mc.NetworkProfile.LoadBalancerProfile.OutboundIPPrefixes == nil ||
		mc.NetworkProfile.LoadBalancerProfile.OutboundIPPrefixes.PublicIPPrefixes == nil {
		return nil
	}
	items := *mc.NetworkProfile.LoadBalancerProfile.OutboundIPPrefixes.PublicIPPrefixes
	ids := make([]string, 0, len(items))
	for _, p := range items {
		if p.ID != nil {
			ids = append(ids, *p.ID)
		}
	}
	return resource.Set(c.Name, ids)
}

func resolveContainerManagedClusterNetworkProfileLoadBalancerOutboundIps(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	mc, ok := resource.Item.(containerservice.ManagedCluster)
	if !ok {
		return fmt.Errorf("not a containerservice.ManagedCluster instance: %T", resource.Item)
	}
	if mc.NetworkProfile == nil ||
		mc.NetworkProfile.LoadBalancerProfile == nil ||
		mc.NetworkProfile.LoadBalancerProfile.OutboundIPs == nil ||
		mc.NetworkProfile.LoadBalancerProfile.OutboundIPs.PublicIPs == nil {
		return nil
	}
	items := *mc.NetworkProfile.LoadBalancerProfile.OutboundIPs.PublicIPs
	ids := make([]string, 0, len(items))
	for _, p := range items {
		if p.ID != nil {
			ids = append(ids, *p.ID)
		}
	}
	return resource.Set(c.Name, ids)
}

func resolveContainerManagedClusterNetworkProfileLoadBalancerEffectiveOutboundIps(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	mc, ok := resource.Item.(containerservice.ManagedCluster)
	if !ok {
		return fmt.Errorf("not a containerservice.ManagedCluster instance: %T", resource.Item)
	}
	if mc.NetworkProfile == nil ||
		mc.NetworkProfile.LoadBalancerProfile == nil ||
		mc.NetworkProfile.LoadBalancerProfile.EffectiveOutboundIPs == nil {
		return nil
	}
	items := *mc.NetworkProfile.LoadBalancerProfile.EffectiveOutboundIPs
	ids := make([]string, 0, len(items))
	for _, p := range items {
		if p.ID != nil {
			ids = append(ids, *p.ID)
		}
	}
	return resource.Set(c.Name, ids)
}

func fetchContainerManagedClusterPipUserAssignedIdentityExceptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	mc, ok := parent.Item.(containerservice.ManagedCluster)
	if !ok {
		return fmt.Errorf("not a containerservice.ManagedCluster instance: %T", parent.Item)
	}
	if mc.PodIdentityProfile == nil ||
		mc.PodIdentityProfile.UserAssignedIdentityExceptions == nil {
		return nil
	}
	res <- *mc.PodIdentityProfile.UserAssignedIdentityExceptions
	return nil
}

func fetchContainerManagedClusterPrivateLinkResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	mc, ok := parent.Item.(containerservice.ManagedCluster)
	if !ok {
		return fmt.Errorf("not a containerservice.ManagedCluster instance: %T", parent.Item)
	}
	if mc.PrivateLinkResources == nil {
		return nil
	}
	res <- *mc.PrivateLinkResources
	return nil
}

func fetchContainerManagedClusterAgentPoolProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	mc, ok := parent.Item.(containerservice.ManagedCluster)
	if !ok {
		return fmt.Errorf("not a containerservice.ManagedCluster instance: %T", parent.Item)
	}
	if mc.AgentPoolProfiles == nil {
		return nil
	}
	res <- *mc.AgentPoolProfiles
	return nil
}

func resolveContainerManagedClusterAgentPoolProfileLinuxOsConfig(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	a, ok := resource.Item.(containerservice.ManagedClusterAgentPoolProfile)
	if !ok {
		return fmt.Errorf("not an containerservice.ManagedClusterAgentPoolProfile instance: %T", resource.Item)
	}
	out, err := json.Marshal(a.LinuxOSConfig)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out)
}

func fetchContainerManagedClusterPipUserAssignedIdentities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	mc, ok := parent.Item.(containerservice.ManagedCluster)
	if !ok {
		return fmt.Errorf("not a containerservice.ManagedCluster instance: %T", parent.Item)
	}
	if mc.PodIdentityProfile == nil ||
		mc.PodIdentityProfile.UserAssignedIdentities == nil {
		return nil
	}
	res <- *mc.PodIdentityProfile.UserAssignedIdentities
	return nil
}
