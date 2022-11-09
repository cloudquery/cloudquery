// Code generated by codegen; DO NOT EDIT.

package elasticache

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticache_clusters",
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CacheCluster.html`,
		Resolver:    fetchElasticacheClusters,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticache"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "at_rest_encryption_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AtRestEncryptionEnabled"),
			},
			{
				Name:     "auth_token_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AuthTokenEnabled"),
			},
			{
				Name:     "auth_token_last_modified_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("AuthTokenLastModifiedDate"),
			},
			{
				Name:     "auto_minor_version_upgrade",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AutoMinorVersionUpgrade"),
			},
			{
				Name:     "cache_cluster_create_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CacheClusterCreateTime"),
			},
			{
				Name:     "cache_cluster_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CacheClusterId"),
			},
			{
				Name:     "cache_cluster_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CacheClusterStatus"),
			},
			{
				Name:     "cache_node_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CacheNodeType"),
			},
			{
				Name:     "cache_nodes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CacheNodes"),
			},
			{
				Name:     "cache_parameter_group",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CacheParameterGroup"),
			},
			{
				Name:     "cache_security_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CacheSecurityGroups"),
			},
			{
				Name:     "cache_subnet_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CacheSubnetGroupName"),
			},
			{
				Name:     "client_download_landing_page",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClientDownloadLandingPage"),
			},
			{
				Name:     "configuration_endpoint",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ConfigurationEndpoint"),
			},
			{
				Name:     "engine",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Engine"),
			},
			{
				Name:     "engine_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EngineVersion"),
			},
			{
				Name:     "ip_discovery",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IpDiscovery"),
			},
			{
				Name:     "log_delivery_configurations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LogDeliveryConfigurations"),
			},
			{
				Name:     "network_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NetworkType"),
			},
			{
				Name:     "notification_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NotificationConfiguration"),
			},
			{
				Name:     "num_cache_nodes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumCacheNodes"),
			},
			{
				Name:     "pending_modified_values",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PendingModifiedValues"),
			},
			{
				Name:     "preferred_availability_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PreferredAvailabilityZone"),
			},
			{
				Name:     "preferred_maintenance_window",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PreferredMaintenanceWindow"),
			},
			{
				Name:     "preferred_outpost_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PreferredOutpostArn"),
			},
			{
				Name:     "replication_group_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicationGroupId"),
			},
			{
				Name:     "replication_group_log_delivery_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ReplicationGroupLogDeliveryEnabled"),
			},
			{
				Name:     "security_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SecurityGroups"),
			},
			{
				Name:     "snapshot_retention_limit",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("SnapshotRetentionLimit"),
			},
			{
				Name:     "snapshot_window",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SnapshotWindow"),
			},
			{
				Name:     "transit_encryption_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("TransitEncryptionEnabled"),
			},
		},
	}
}
