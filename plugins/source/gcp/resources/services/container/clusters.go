// Code generated by codegen; DO NOT EDIT.

package container

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "gcp_container_clusters",
		Description: `https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster`,
		Resolver:    fetchClusters,
		Multiplex:   client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "initial_node_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("InitialNodeCount"),
			},
			{
				Name:     "node_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NodeConfig"),
			},
			{
				Name:     "master_auth",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MasterAuth"),
			},
			{
				Name:     "logging_service",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LoggingService"),
			},
			{
				Name:     "monitoring_service",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MonitoringService"),
			},
			{
				Name:     "network",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Network"),
			},
			{
				Name:     "cluster_ipv4_cidr",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterIpv4Cidr"),
			},
			{
				Name:     "addons_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AddonsConfig"),
			},
			{
				Name:     "subnetwork",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Subnetwork"),
			},
			{
				Name:     "node_pools",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NodePools"),
			},
			{
				Name:     "locations",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Locations"),
			},
			{
				Name:     "enable_kubernetes_alpha",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableKubernetesAlpha"),
			},
			{
				Name:     "resource_labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResourceLabels"),
			},
			{
				Name:     "label_fingerprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LabelFingerprint"),
			},
			{
				Name:     "legacy_abac",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LegacyAbac"),
			},
			{
				Name:     "network_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkPolicy"),
			},
			{
				Name:     "ip_allocation_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IpAllocationPolicy"),
			},
			{
				Name:     "master_authorized_networks_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MasterAuthorizedNetworksConfig"),
			},
			{
				Name:     "maintenance_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MaintenancePolicy"),
			},
			{
				Name:     "binary_authorization",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BinaryAuthorization"),
			},
			{
				Name:     "autoscaling",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Autoscaling"),
			},
			{
				Name:     "network_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkConfig"),
			},
			{
				Name:     "default_max_pods_constraint",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DefaultMaxPodsConstraint"),
			},
			{
				Name:     "resource_usage_export_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResourceUsageExportConfig"),
			},
			{
				Name:     "authenticator_groups_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AuthenticatorGroupsConfig"),
			},
			{
				Name:     "private_cluster_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrivateClusterConfig"),
			},
			{
				Name:     "database_encryption",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DatabaseEncryption"),
			},
			{
				Name:     "vertical_pod_autoscaling",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VerticalPodAutoscaling"),
			},
			{
				Name:     "shielded_nodes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ShieldedNodes"),
			},
			{
				Name:     "release_channel",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ReleaseChannel"),
			},
			{
				Name:     "workload_identity_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WorkloadIdentityConfig"),
			},
			{
				Name:     "mesh_certificates",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MeshCertificates"),
			},
			{
				Name:     "cost_management_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CostManagementConfig"),
			},
			{
				Name:     "notification_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NotificationConfig"),
			},
			{
				Name:     "confidential_nodes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ConfidentialNodes"),
			},
			{
				Name:     "identity_service_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IdentityServiceConfig"),
			},
			{
				Name:     "self_link",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SelfLink"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Zone"),
			},
			{
				Name:     "endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Endpoint"),
			},
			{
				Name:     "initial_cluster_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InitialClusterVersion"),
			},
			{
				Name:     "current_master_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CurrentMasterVersion"),
			},
			{
				Name:     "current_node_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CurrentNodeVersion"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreateTime"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("Status"),
			},
			{
				Name:     "status_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StatusMessage"),
			},
			{
				Name:     "node_ipv4_cidr_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NodeIpv4CidrSize"),
			},
			{
				Name:     "services_ipv4_cidr",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServicesIpv4Cidr"),
			},
			{
				Name:     "instance_group_urls",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("InstanceGroupUrls"),
			},
			{
				Name:     "current_node_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CurrentNodeCount"),
			},
			{
				Name:     "expire_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExpireTime"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "enable_tpu",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableTpu"),
			},
			{
				Name:     "tpu_ipv4_cidr_block",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TpuIpv4CidrBlock"),
			},
			{
				Name:     "conditions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Conditions"),
			},
			{
				Name:     "autopilot",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Autopilot"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "node_pool_defaults",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NodePoolDefaults"),
			},
			{
				Name:     "logging_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LoggingConfig"),
			},
			{
				Name:     "monitoring_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MonitoringConfig"),
			},
			{
				Name:     "node_pool_auto_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NodePoolAutoConfig"),
			},
		},
	}
}
