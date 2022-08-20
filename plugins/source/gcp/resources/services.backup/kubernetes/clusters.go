package kubernetes

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/container/v1"
)

//go:generate cq-gen --resource clusters --config gen.hcl --output .
func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "gcp_kubernetes_clusters",
		Description: "A Google Kubernetes Engine cluster",
		Resolver:    fetchKubernetesClusters,
		Multiplex:   client.ProjectMultiplex,

		Options: schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "addons_config_cloud_run_config_disabled",
				Description: "Whether Cloud Run addon is enabled for this cluster",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AddonsConfig.CloudRunConfig.Disabled"),
			},
			{
				Name:        "addons_config_cloud_run_config_load_balancer_type",
				Description: "\"LOAD_BALANCER_TYPE_UNSPECIFIED\" - Load balancer type for Cloud Run is unspecified   \"LOAD_BALANCER_TYPE_EXTERNAL\" - Install external load balancer for Cloud Run   \"LOAD_BALANCER_TYPE_INTERNAL\" - Install internal load balancer for Cloud Run",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AddonsConfig.CloudRunConfig.LoadBalancerType"),
			},
			{
				Name:        "addons_config_config_connector_config_enabled",
				Description: "Whether Cloud Connector is enabled for this cluster",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AddonsConfig.ConfigConnectorConfig.Enabled"),
			},
			{
				Name:        "addons_config_dns_cache_config_enabled",
				Description: "Whether NodeLocal DNSCache is enabled for this cluster",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AddonsConfig.DnsCacheConfig.Enabled"),
			},
			{
				Name:        "addons_config_gce_persistent_disk_csi_driver_config_enabled",
				Description: "Whether the Compute Engine PD CSI driver is enabled for this cluster",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AddonsConfig.GcePersistentDiskCsiDriverConfig.Enabled"),
			},
			{
				Name:        "addons_config_horizontal_pod_autoscaling_disabled",
				Description: "Whether the Horizontal Pod Autoscaling feature is enabled in the cluster",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AddonsConfig.HorizontalPodAutoscaling.Disabled"),
			},
			{
				Name:        "addons_config_http_load_balancing_disabled",
				Description: "Whether the HTTP Load Balancing controller is enabled in the cluster",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AddonsConfig.HttpLoadBalancing.Disabled"),
			},
			{
				Name:        "addons_config_kubernetes_dashboard_disabled",
				Description: "Whether the Kubernetes Dashboard is enabled for this cluster",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AddonsConfig.KubernetesDashboard.Disabled"),
			},
			{
				Name:        "addons_config_network_policy_config_disabled",
				Description: "Whether NetworkPolicy is enabled for this cluster",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AddonsConfig.NetworkPolicyConfig.Disabled"),
			},
			{
				Name:        "authenticator_groups_config_enabled",
				Description: "Whether this cluster should return group membership lookups during authentication using a group of security groups",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AuthenticatorGroupsConfig.Enabled"),
			},
			{
				Name:        "authenticator_groups_config_security_group",
				Description: "The name of the security group-of-groups to be used Only relevant if enabled = true",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AuthenticatorGroupsConfig.SecurityGroup"),
			},
			{
				Name:        "autopilot_enabled",
				Description: "Enable Autopilot",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Autopilot.Enabled"),
			},
			{
				Name:          "autoscaling_autoprovisioning_locations",
				Description:   "The list of Google Compute Engine zones (https://cloudgooglecom/compute/docs/zones#available) in which the NodePool's nodes can be created by NAP",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("Autoscaling.AutoprovisioningLocations"),
				IgnoreInTests: true,
			},
			{
				Name:          "autoscaling_autoprovisioning_node_pool_defaults",
				Description:   "AutoprovisioningNodePoolDefaults contains defaults for a node pool created by NAP",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("Autoscaling.AutoprovisioningNodePoolDefaults"),
				IgnoreInTests: true,
			},
			{
				Name:        "autoscaling_profile",
				Description: "\"PROFILE_UNSPECIFIED\" - No change to autoscaling configuration   \"OPTIMIZE_UTILIZATION\" - Prioritize optimizing utilization of resources   \"BALANCED\" - Use default (balanced) autoscaling configuration",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Autoscaling.AutoscalingProfile"),
			},
			{
				Name:        "autoscaling_enable_node_autoprovisioning",
				Description: "Enables automatic node pool creation and deletion",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Autoscaling.EnableNodeAutoprovisioning"),
			},
			{
				Name:          "autoscaling_resource_limits",
				Description:   "Contains global constraints regarding minimum and maximum amount of resources in the cluster",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("Autoscaling.ResourceLimits"),
			},
			{
				Name:        "binary_authorization_enabled",
				Description: "Enable Binary Authorization for this cluster",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("BinaryAuthorization.Enabled"),
			},
			{
				Name:          "cluster_ipv4_cidr",
				Description:   "The IP address range of the container pods in this cluster, in CIDR (http://enwikipediaorg/wiki/Classless_Inter-Domain_Routing) notation (eg",
				Type:          schema.TypeCIDR,
				Resolver:      client.AllowEmptyStringIPNetResolver("ClusterIpv4Cidr"),
				IgnoreInTests: true,
			},
			{
				Name:          "conditions",
				Description:   "Which conditions caused the current cluster state",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:        "confidential_nodes_enabled",
				Description: "Whether Confidential Nodes feature is enabled for all nodes in this cluster",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ConfidentialNodes.Enabled"),
			},
			{
				Name:        "create_time",
				Description: "The time the cluster was created, in RFC3339 (https://wwwietforg/rfc/rfc3339txt) text format",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.DateResolver("CreateTime"),
			},
			{
				Name:        "current_master_version",
				Description: "The current software version of the master endpoint",
				Type:        schema.TypeString,
			},
			{
				Name:        "current_node_count",
				Description: "The number of nodes currently in the cluster",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "current_node_version",
				Description: "Deprecated, use NodePoolsversion (https://cloudgooglecom/kubernetes-engine/docs/reference/rest/v1/projectslocationsclustersnodePools) instead",
				Type:        schema.TypeString,
			},
			{
				Name:        "database_encryption_key_name",
				Description: "Name of CloudKMS key to use for the encryption of secrets in etcd",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseEncryption.KeyName"),
			},
			{
				Name:        "database_encryption_state",
				Description: "\"UNKNOWN\" - Should never be set   \"ENCRYPTED\" - Secrets in etcd are encrypted   \"DECRYPTED\" - Secrets in etcd are stored in plain text (at etcd level) - this is unrelated to Compute Engine level full disk encryption",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseEncryption.State"),
			},
			{
				Name:        "default_max_pods_constraint_max_pods_per_node",
				Description: "Constraint enforced on the max num of pods per node",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DefaultMaxPodsConstraint.MaxPodsPerNode"),
			},
			{
				Name:        "description",
				Description: "An optional description of this cluster",
				Type:        schema.TypeString,
			},
			{
				Name:        "enable_kubernetes_alpha",
				Description: "Kubernetes alpha features are enabled on this cluster",
				Type:        schema.TypeBool,
			},
			{
				Name:        "enable_tpu",
				Description: "Enable the ability to use Cloud TPUs in this cluster",
				Type:        schema.TypeBool,
			},
			{
				Name:        "endpoint",
				Description: "The IP address of this cluster's master endpoint",
				Type:        schema.TypeInet,
				Resolver:    schema.IPAddressResolver("Endpoint"),
			},
			{
				Name:          "expire_time",
				Description:   "The time the cluster will be automatically deleted in RFC3339 (https://wwwietforg/rfc/rfc3339txt) text format",
				Type:          schema.TypeTimestamp,
				Resolver:      schema.DateResolver("ExpireTime"),
				IgnoreInTests: true,
			},
			{
				Name:        "id",
				Description: "Output only",
				Type:        schema.TypeString,
			},
			{
				Name:        "initial_cluster_version",
				Description: "The initial Kubernetes version for this cluster",
				Type:        schema.TypeString,
			},
			{
				Name:        "initial_node_count",
				Description: "The number of nodes to create in this cluster",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "instance_group_urls",
				Description: "Deprecated",
				Type:        schema.TypeStringArray,
			},
			{
				Name:          "ip_allocation_policy_cluster_ipv4_cidr",
				Description:   "This field is deprecated, use cluster_ipv4_cidr_block",
				Type:          schema.TypeCIDR,
				Resolver:      client.AllowEmptyStringIPNetResolver("IpAllocationPolicy.ClusterIpv4Cidr"),
				IgnoreInTests: true,
			},
			{
				Name:          "ip_allocation_policy_cluster_ipv4_cidr_block",
				Description:   "The IP address range for the cluster pod IPs If this field is set, then `clustercluster_ipv4_cidr` must be left blank",
				Type:          schema.TypeCIDR,
				Resolver:      client.AllowEmptyStringIPNetResolver("IpAllocationPolicy.ClusterIpv4CidrBlock"),
				IgnoreInTests: true,
			},
			{
				Name:        "ip_allocation_policy_cluster_secondary_range_name",
				Description: "The name of the secondary range to be used for the cluster CIDR block",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("IpAllocationPolicy.ClusterSecondaryRangeName"),
			},
			{
				Name:        "ip_allocation_policy_create_subnetwork",
				Description: "Whether a new subnetwork will be created automatically for the cluster",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("IpAllocationPolicy.CreateSubnetwork"),
			},
			{
				Name:          "ip_allocation_policy_node_ipv4_cidr",
				Description:   "This field is deprecated, use node_ipv4_cidr_block",
				Type:          schema.TypeCIDR,
				Resolver:      client.AllowEmptyStringIPNetResolver("IpAllocationPolicy.NodeIpv4Cidr"),
				IgnoreInTests: true,
			},
			{
				Name:          "ip_allocation_policy_node_ipv4_cidr_block",
				Description:   "The IP address range of the instance IPs in this cluster",
				Type:          schema.TypeCIDR,
				Resolver:      client.AllowEmptyStringIPNetResolver("IpAllocationPolicy.NodeIpv4CidrBlock"),
				IgnoreInTests: true,
			},
			{
				Name:          "ip_allocation_policy_services_ipv4_cidr",
				Description:   "This field is deprecated, use services_ipv4_cidr_block",
				Type:          schema.TypeCIDR,
				Resolver:      client.AllowEmptyStringIPNetResolver("IpAllocationPolicy.ServicesIpv4Cidr"),
				IgnoreInTests: true,
			},
			{
				Name:          "ip_allocation_policy_services_ipv4_cidr_block",
				Description:   "The IP address range of the services IPs in this cluster",
				Type:          schema.TypeCIDR,
				Resolver:      client.AllowEmptyStringIPNetResolver("IpAllocationPolicy.ServicesIpv4CidrBlock"),
				IgnoreInTests: true,
			},
			{
				Name:        "ip_allocation_policy_services_secondary_range_name",
				Description: "The name of the secondary range to be used as for the services CIDR block",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("IpAllocationPolicy.ServicesSecondaryRangeName"),
			},
			{
				Name:        "ip_allocation_policy_subnetwork_name",
				Description: "A custom subnetwork name to be used if `create_subnetwork` is true",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("IpAllocationPolicy.SubnetworkName"),
			},
			{
				Name:          "ip_allocation_policy_tpu_ipv4_cidr_block",
				Description:   "The IP address range of the Cloud TPUs in this cluster",
				Type:          schema.TypeCIDR,
				Resolver:      client.AllowEmptyStringIPNetResolver("IpAllocationPolicy.TpuIpv4CidrBlock"),
				IgnoreInTests: true,
			},
			{
				Name:        "ip_allocation_policy_use_ip_aliases",
				Description: "Whether alias IPs will be used for pod IPs in the cluster",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("IpAllocationPolicy.UseIpAliases"),
			},
			{
				Name:        "ip_allocation_policy_use_routes",
				Description: "Whether routes will be used for pod IPs in the cluster This is used in conjunction with use_ip_aliases",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("IpAllocationPolicy.UseRoutes"),
			},
			{
				Name:        "label_fingerprint",
				Description: "The fingerprint of the set of labels for this cluster",
				Type:        schema.TypeString,
			},
			{
				Name:        "legacy_abac_enabled",
				Description: "Whether the ABAC authorizer is enabled for this cluster When enabled, identities in the system, including service accounts, nodes, and controllers, will have statically granted permissions beyond those provided by the RBAC configuration or IAM",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("LegacyAbac.Enabled"),
			},
			{
				Name:        "location",
				Description: "The name of the Google Compute Engine zone (https://cloudgooglecom/compute/docs/regions-zones/regions-zones#available) or region (https://cloudgooglecom/compute/docs/regions-zones/regions-zones#available) in which the cluster resides",
				Type:        schema.TypeString,
			},
			{
				Name:        "locations",
				Description: "The list of Google Compute Engine zones (https://cloudgooglecom/compute/docs/zones#available) in which the cluster's nodes should be located",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "logging_config_component_config_enable_components",
				Description: "Select components to collect logs",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("LoggingConfig.ComponentConfig.EnableComponents"),
			},
			{
				Name:        "logging_service",
				Description: "The logging service the cluster should use to write logs",
				Type:        schema.TypeString,
			},
			{
				Name:        "maintenance_policy_resource_version",
				Description: "A hash identifying the version of this policy, so that updates to fields of the policy won't accidentally undo intermediate changes (and so that users of the API unaware of some fields won't accidentally remove other fields)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MaintenancePolicy.ResourceVersion"),
			},
			{
				Name:        "maintenance_policy_window_daily_maintenance_window_duration",
				Description: "Duration of the time window, automatically chosen to be smallest possible in the given scenario",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MaintenancePolicy.Window.DailyMaintenanceWindow.Duration"),
			},
			{
				Name:        "maintenance_policy_window_daily_maintenance_window_start_time",
				Description: "Time within the maintenance window to start the maintenance operations",
				Type:        schema.TypeTimestamp,
				// https://pkg.go.dev/google.golang.org/api/container/v1#DailyMaintenanceWindow
				Resolver: schema.DateResolver("MaintenancePolicy.Window.DailyMaintenanceWindow.StartTime", "15:04"),
			},
			{
				Name:          "maintenance_policy_window_maintenance_exclusions",
				Description:   "Exceptions to maintenance window Non-emergency maintenance should not occur in these windows",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("MaintenancePolicy.Window.MaintenanceExclusions"),
				IgnoreInTests: true,
			},
			{
				Name:        "maintenance_policy_window_recurring_window_recurrence",
				Description: "An RRULE (https://toolsietforg/html/rfc5545#section-3853) for how this window reccurs",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MaintenancePolicy.Window.RecurringWindow.Recurrence"),
			},
			{
				Name:          "maintenance_policy_window_recurring_window_window_end_time",
				Description:   "The time that the window ends",
				Type:          schema.TypeTimestamp,
				Resolver:      schema.DateResolver("MaintenancePolicy.Window.RecurringWindow.Window.EndTime"),
				IgnoreInTests: true,
			},
			{
				Name:          "maintenance_policy_window_recurring_window_window_start_time",
				Description:   "The time that the window first starts",
				Type:          schema.TypeTimestamp,
				Resolver:      schema.DateResolver("MaintenancePolicy.Window.RecurringWindow.Window.StartTime"),
				IgnoreInTests: true,
			},
			{
				Name:        "master_auth_client_certificate",
				Description: "Base64-encoded public certificate used by clients to authenticate to the cluster endpoint",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MasterAuth.ClientCertificate"),
			},
			{
				Name:        "master_auth_client_certificate_config_issue_client_certificate",
				Description: "Issue a client certificate",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("MasterAuth.ClientCertificateConfig.IssueClientCertificate"),
			},
			{
				Name:        "master_auth_client_key",
				Description: "Base64-encoded private key used by clients to authenticate to the cluster endpoint",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MasterAuth.ClientKey"),
			},
			{
				Name:        "master_auth_cluster_ca_certificate",
				Description: "Base64-encoded public certificate that is the root of trust for the cluster",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MasterAuth.ClusterCaCertificate"),
			},
			{
				Name:        "master_auth_password",
				Description: "The password to use for HTTP basic authentication to the master endpoint",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MasterAuth.Password"),
			},
			{
				Name:        "master_auth_username",
				Description: "The username to use for HTTP basic authentication to the master endpoint",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MasterAuth.Username"),
			},
			{
				Name:          "master_authorized_networks_config_cidr_blocks",
				Description:   "cidr_blocks define up to 50 external networks that could access Kubernetes master through HTTPS",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("MasterAuthorizedNetworksConfig.CidrBlocks"),
				IgnoreInTests: true,
			},
			{
				Name:        "master_authorized_networks_config_enabled",
				Description: "Whether or not master authorized networks is enabled",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("MasterAuthorizedNetworksConfig.Enabled"),
			},
			{
				Name:        "monitoring_config_component_config_enable_components",
				Description: "Select components to collect metrics",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("MonitoringConfig.ComponentConfig.EnableComponents"),
			},
			{
				Name:        "monitoring_service",
				Description: "The monitoring service the cluster should use to write metrics",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The name of this cluster",
				Type:        schema.TypeString,
			},
			{
				Name:        "network",
				Description: "The name of the Google Compute Engine network (https://cloudgooglecom/compute/docs/networks-and-firewalls#networks) to which the cluster is connected",
				Type:        schema.TypeString,
			},
			{
				Name:        "network_config_datapath_provider",
				Description: "The desired datapath provider for this cluster",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("NetworkConfig.DatapathProvider"),
			},
			{
				Name:        "network_config_default_snat_status_disabled",
				Description: "Disables cluster default sNAT rules",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("NetworkConfig.DefaultSnatStatus.Disabled"),
			},
			{
				Name:        "network_config_enable_intra_node_visibility",
				Description: "Whether Intra-node visibility is enabled for this cluster",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("NetworkConfig.EnableIntraNodeVisibility"),
			},
			{
				Name:        "network_config_enable_l4ilb_subsetting",
				Description: "Whether L4ILB Subsetting is enabled for this cluster",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("NetworkConfig.EnableL4ilbSubsetting"),
			},
			{
				Name:        "network_config_network",
				Description: "Output only",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("NetworkConfig.Network"),
			},
			{
				Name:        "network_config_private_ipv6_google_access",
				Description: "The desired state of IPv6 connectivity to Google Services",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("NetworkConfig.PrivateIpv6GoogleAccess"),
			},
			{
				Name:        "network_config_subnetwork",
				Description: "Output only",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("NetworkConfig.Subnetwork"),
			},
			{
				Name:        "network_policy_enabled",
				Description: "Whether network policy is enabled on the cluster",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("NetworkPolicy.Enabled"),
			},
			{
				Name:        "network_policy_provider",
				Description: "\"PROVIDER_UNSPECIFIED\" - Not set   \"CALICO\" - Tigera (Calico Felix)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("NetworkPolicy.Provider"),
			},
			{
				Name:        "node_config",
				Description: "Parameters used in creating the cluster's nodes",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "node_ipv4_cidr_size",
				Description: "The size of the address space on each node for hosting containers",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "notification_config_pubsub_enabled",
				Description: "Enable notifications for Pub/Sub",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("NotificationConfig.Pubsub.Enabled"),
			},
			{
				Name:        "notification_config_pubsub_topic",
				Description: "The desired Pub/Sub topic to which notifications will be sent by GKE",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("NotificationConfig.Pubsub.Topic"),
			},
			{
				Name:        "private_cluster_config_enable_private_endpoint",
				Description: "Whether the master's internal IP address is used as the cluster endpoint",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("PrivateClusterConfig.EnablePrivateEndpoint"),
			},
			{
				Name:        "private_cluster_config_enable_private_nodes",
				Description: "Whether nodes have internal IP addresses only",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("PrivateClusterConfig.EnablePrivateNodes"),
			},
			{
				Name:        "private_cluster_config_master_global_access_config_enabled",
				Description: "Whenever master is accessible globally or not",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("PrivateClusterConfig.MasterGlobalAccessConfig.Enabled"),
			},
			{
				Name:          "private_cluster_config_master_ipv4_cidr_block",
				Description:   "The IP range in CIDR notation to use for the hosted master network",
				Type:          schema.TypeCIDR,
				Resolver:      client.AllowEmptyStringIPNetResolver("PrivateClusterConfig.MasterIpv4CidrBlock"),
				IgnoreInTests: true,
			},
			{
				Name:        "private_cluster_config_peering_name",
				Description: "Output only",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PrivateClusterConfig.PeeringName"),
			},
			{
				Name:        "private_cluster_config_private_endpoint",
				Description: "Output only",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PrivateClusterConfig.PrivateEndpoint"),
			},
			{
				Name:        "private_cluster_config_public_endpoint",
				Description: "Output only",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PrivateClusterConfig.PublicEndpoint"),
			},
			{
				Name:        "release_channel",
				Description: "\"UNSPECIFIED\" - No channel specified   \"RAPID\" - RAPID channel is offered on an early access basis for customers who want to test new releases",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ReleaseChannel.Channel"),
			},
			{
				Name:          "resource_labels",
				Description:   "The resource labels for the cluster to use to annotate any related Google Compute Engine resources",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:          "resource_usage_export_config",
				Description:   "Configuration for exporting resource usages",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
			{
				Name:          "services_ipv4_cidr",
				Description:   "The IP address range of the Kubernetes services in this cluster, in CIDR (http://enwikipediaorg/wiki/Classless_Inter-Domain_Routing) notation (eg",
				Type:          schema.TypeCIDR,
				Resolver:      client.AllowEmptyStringIPNetResolver("ServicesIpv4Cidr"),
				IgnoreInTests: true,
			},
			{
				Name:        "shielded_nodes_enabled",
				Description: "Whether Shielded Nodes features are enabled on all nodes in this cluster",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ShieldedNodes.Enabled"),
			},
			{
				Name:        "status",
				Description: "\"STATUS_UNSPECIFIED\" - Not set   \"PROVISIONING\" - The PROVISIONING state indicates the cluster is being created   \"RUNNING\" - The RUNNING state indicates the cluster has been created and is fully usable   \"RECONCILING\" - The RECONCILING state indicates that some work is actively being done on the cluster, such as upgrading the master or node software",
				Type:        schema.TypeString,
			},
			{
				Name:        "status_message",
				Description: "Deprecated",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnetwork",
				Description: "The name of the Google Compute Engine subnetwork (https://cloudgooglecom/compute/docs/subnetworks) to which the cluster is connected",
				Type:        schema.TypeString,
			},
			{
				Name:          "tpu_ipv4_cidr_block",
				Description:   "The IP address range of the Cloud TPUs in this cluster, in CIDR (http://enwikipediaorg/wiki/Classless_Inter-Domain_Routing) notation (eg",
				Type:          schema.TypeCIDR,
				Resolver:      client.AllowEmptyStringIPNetResolver("TpuIpv4CidrBlock"),
				IgnoreInTests: true,
			},
			{
				Name:        "vertical_pod_autoscaling_enabled",
				Description: "Enables vertical pod autoscaling",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VerticalPodAutoscaling.Enabled"),
			},
			{
				Name:        "workload_identity_config_workload_pool",
				Description: "The workload pool to attach all Kubernetes service accounts to",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("WorkloadIdentityConfig.WorkloadPool"),
			},
			{
				Name:        "zone",
				Description: "The name of the Google Compute Engine zone (https://cloudgooglecom/compute/docs/zones#available) in which the cluster resides",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_kubernetes_cluster_node_pools",
				Description: "NodePool contains the name and configuration for a cluster's node pool",
				Resolver:    fetchKubernetesClusterNodePools,
				Columns: []schema.Column{
					{
						Name:        "cluster_cq_id",
						Description: "Unique CloudQuery ID of gcp_kubernetes_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "autoscaling_autoprovisioned",
						Description: "Can this node pool be deleted automatically",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Autoscaling.Autoprovisioned"),
					},
					{
						Name:        "autoscaling_enabled",
						Description: "Is autoscaling enabled for this node pool",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Autoscaling.Enabled"),
					},
					{
						Name:        "autoscaling_max_node_count",
						Description: "Maximum number of nodes in the NodePool",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("Autoscaling.MaxNodeCount"),
					},
					{
						Name:        "autoscaling_min_node_count",
						Description: "Minimum number of nodes in the NodePool",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("Autoscaling.MinNodeCount"),
					},
					{
						Name:          "conditions",
						Description:   "Which conditions caused the current node pool state",
						Type:          schema.TypeJSON,
						IgnoreInTests: true,
					},
					{
						Name:          "config_accelerators",
						Description:   "A list of hardware accelerators to be attached to each node",
						Type:          schema.TypeJSON,
						Resolver:      schema.PathResolver("Config.Accelerators"),
						IgnoreInTests: true,
					},
					{
						Name:        "config_boot_disk_kms_key",
						Description: "The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Config.BootDiskKmsKey"),
					},
					{
						Name:        "config_disk_size_gb",
						Description: "Size of the disk attached to each node, specified in GB The smallest allowed disk size is 10GB",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("Config.DiskSizeGb"),
					},
					{
						Name:        "config_disk_type",
						Description: "Type of the disk attached to each node (eg",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Config.DiskType"),
					},
					{
						Name:        "config_gvnic_enabled",
						Description: "Whether gVNIC features are enabled in the node pool",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Config.Gvnic.Enabled"),
					},
					{
						Name:        "config_image_type",
						Description: "The image type to use for this node",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Config.ImageType"),
					},
					{
						Name:        "config_kubelet_config_cpu_cfs_quota",
						Description: "Enable CPU CFS quota enforcement for containers that specify CPU limits",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Config.KubeletConfig.CpuCfsQuota"),
					},
					{
						Name:        "config_kubelet_config_cpu_cfs_quota_period",
						Description: "Set the CPU CFS quota period value 'cpucfs_period_us'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Config.KubeletConfig.CpuCfsQuotaPeriod"),
					},
					{
						Name:        "config_kubelet_config_cpu_manager_policy",
						Description: "Control the CPU management policy on the node",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Config.KubeletConfig.CpuManagerPolicy"),
					},
					{
						Name:        "config_labels",
						Description: "The map of Kubernetes labels (key/value pairs) to be applied to each node",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Config.Labels"),
					},
					{
						Name:          "config_linux_node_config_sysctls",
						Description:   "The Linux kernel parameters to be applied to the nodes and all pods running on the nodes",
						Type:          schema.TypeJSON,
						Resolver:      schema.PathResolver("Config.LinuxNodeConfig.Sysctls"),
						IgnoreInTests: true,
					},
					{
						Name:        "config_local_ssd_count",
						Description: "The number of local SSD disks to be attached to the node",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("Config.LocalSsdCount"),
					},
					{
						Name:        "config_machine_type",
						Description: "The name of a Google Compute Engine machine type (https://cloudgooglecom/compute/docs/machine-types) If unspecified, the default machine type is `e2-medium`",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Config.MachineType"),
					},
					{
						Name:        "config_metadata",
						Description: "The metadata key/value pairs assigned to instances in the cluster",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Config.Metadata"),
					},
					{
						Name:        "config_min_cpu_platform",
						Description: "Minimum CPU platform to be used by this instance",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Config.MinCpuPlatform"),
					},
					{
						Name:        "config_node_group",
						Description: "Setting this field will assign instances of this pool to run on the specified node group",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Config.NodeGroup"),
					},
					{
						Name:        "config_oauth_scopes",
						Description: "The set of Google API scopes to be made available on all of the node VMs under the \"default\" service account",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Config.OauthScopes"),
					},
					{
						Name:        "config_preemptible",
						Description: "Whether the nodes are created as preemptible VM instances",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Config.Preemptible"),
					},
					{
						Name:        "config_reservation_affinity_consume_reservation_type",
						Description: "\"UNSPECIFIED\" - Default value",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Config.ReservationAffinity.ConsumeReservationType"),
					},
					{
						Name:        "config_reservation_affinity_key",
						Description: "Corresponds to the label key of a reservation resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Config.ReservationAffinity.Key"),
					},
					{
						Name:          "config_reservation_affinity_values",
						Description:   "Corresponds to the label value(s) of reservation resource(s)",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("Config.ReservationAffinity.Values"),
						IgnoreInTests: true,
					},
					{
						Name:        "config_sandbox_config_type",
						Description: "\"UNSPECIFIED\" - Default value",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Config.SandboxConfig.Type"),
					},
					{
						Name:        "config_service_account",
						Description: "The Google Cloud Platform Service Account to be used by the node VMs",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Config.ServiceAccount"),
					},
					{
						Name:        "config_shielded_instance_config_enable_integrity_monitoring",
						Description: "Defines whether the instance has integrity monitoring enabled",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Config.ShieldedInstanceConfig.EnableIntegrityMonitoring"),
					},
					{
						Name:        "config_shielded_instance_config_enable_secure_boot",
						Description: "Defines whether the instance has Secure Boot enabled",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Config.ShieldedInstanceConfig.EnableSecureBoot"),
					},
					{
						Name:        "config_tags",
						Description: "The list of instance tags applied to all nodes",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Config.Tags"),
					},
					{
						Name:          "config_taints",
						Description:   "List of kubernetes taints to be applied to each node",
						Type:          schema.TypeJSON,
						Resolver:      schema.PathResolver("Config.Taints"),
						IgnoreInTests: true,
					},
					{
						Name:        "config_workload_metadata_config_mode",
						Description: "\"MODE_UNSPECIFIED\" - Not set   \"GCE_METADATA\" - Expose all Compute Engine metadata to pods   \"GKE_METADATA\" - Run the GKE Metadata Server on this node",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Config.WorkloadMetadataConfig.Mode"),
					},
					{
						Name:        "initial_node_count",
						Description: "The initial node count for the pool",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "instance_group_urls",
						Description: "The resource URLs of the managed instance groups (https://cloudgooglecom/compute/docs/instance-groups/creating-groups-of-managed-instances) associated with this node pool",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "locations",
						Description: "The list of Google Compute Engine zones (https://cloudgooglecom/compute/docs/zones#available) in which the NodePool's nodes should be located",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "management_auto_repair",
						Description: "A flag that specifies whether the node auto-repair is enabled for the node pool",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Management.AutoRepair"),
					},
					{
						Name:        "management_auto_upgrade",
						Description: "A flag that specifies whether node auto-upgrade is enabled for the node pool",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Management.AutoUpgrade"),
					},
					{
						Name:          "management_upgrade_options_auto_upgrade_start_time",
						Description:   "This field is set when upgrades are about to commence with the approximate start time for the upgrades, in RFC3339 (https://wwwietforg/rfc/rfc3339txt) text format",
						Type:          schema.TypeTimestamp,
						Resolver:      schema.DateResolver("Management.UpgradeOptions.AutoUpgradeStartTime"),
						IgnoreInTests: true,
					},
					{
						Name:        "management_upgrade_options_description",
						Description: "This field is set when upgrades are about to commence with the description of the upgrade",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Management.UpgradeOptions.Description"),
					},
					{
						Name:        "max_pods_constraint_max_pods_per_node",
						Description: "Constraint enforced on the max num of pods per node",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("MaxPodsConstraint.MaxPodsPerNode"),
					},
					{
						Name:        "name",
						Description: "The name of the node pool",
						Type:        schema.TypeString,
					},
					{
						Name:        "network_config_create_pod_range",
						Description: "Input only",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("NetworkConfig.CreatePodRange"),
					},
					{
						Name:          "network_config_pod_ipv4_cidr_block",
						Description:   "The IP address range for pod IPs in this node pool Only applicable if `create_pod_range` is true",
						Type:          schema.TypeCIDR,
						Resolver:      client.AllowEmptyStringIPNetResolver("NetworkConfig.PodIpv4CidrBlock"),
						IgnoreInTests: true,
					},
					{
						Name:        "network_config_pod_range",
						Description: "The ID of the secondary range for pod IPs",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("NetworkConfig.PodRange"),
					},
					{
						Name:        "pod_ipv4_cidr_size",
						Description: "The pod CIDR block size per node in this node pool",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "self_link",
						Description: "Server-defined URL for the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "\"STATUS_UNSPECIFIED\" - Not set   \"PROVISIONING\" - The PROVISIONING state indicates the node pool is being created   \"RUNNING\" - The RUNNING state indicates the node pool has been created and is fully usable   \"RUNNING_WITH_ERROR\" - The RUNNING_WITH_ERROR state indicates the node pool has been created and is partially usable",
						Type:        schema.TypeString,
					},
					{
						Name:        "status_message",
						Description: "Deprecated",
						Type:        schema.TypeString,
					},
					{
						Name:        "upgrade_settings_max_surge",
						Description: "The maximum number of nodes that can be created beyond the current size of the node pool during the upgrade process",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("UpgradeSettings.MaxSurge"),
					},
					{
						Name:        "upgrade_settings_max_unavailable",
						Description: "The maximum number of nodes that can be simultaneously unavailable during the upgrade process",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("UpgradeSettings.MaxUnavailable"),
					},
					{
						Name:        "version",
						Description: "The version of the Kubernetes of this node",
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

func fetchKubernetesClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	output, err := c.Services.Container.Projects.Locations.Clusters.List("projects/" + c.ProjectId + "/locations/-").Do()
	if err != nil {
		return errors.WithStack(err)
	}

	res <- output.Clusters
	return nil
}

func fetchKubernetesClusterNodePools(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*container.Cluster)

	res <- p.NodePools

	return nil
}
