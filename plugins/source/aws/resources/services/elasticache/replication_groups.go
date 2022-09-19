// Code generated by codegen; DO NOT EDIT.

package elasticache

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ReplicationGroups() *schema.Table {
	return &schema.Table{
		Name:      "aws_elasticache_replication_groups",
		Resolver:  fetchElasticacheReplicationGroups,
		Multiplex: client.ServiceAccountRegionMultiplexer("elasticache"),
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
				Name:     "automatic_failover",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutomaticFailover"),
			},
			{
				Name:     "cache_node_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CacheNodeType"),
			},
			{
				Name:     "cluster_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ClusterEnabled"),
			},
			{
				Name:     "configuration_endpoint",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ConfigurationEndpoint"),
			},
			{
				Name:     "data_tiering",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DataTiering"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "global_replication_group_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("GlobalReplicationGroupInfo"),
			},
			{
				Name:     "kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KmsKeyId"),
			},
			{
				Name:     "log_delivery_configurations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LogDeliveryConfigurations"),
			},
			{
				Name:     "member_clusters",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("MemberClusters"),
			},
			{
				Name:     "member_clusters_outpost_arns",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("MemberClustersOutpostArns"),
			},
			{
				Name:     "multi_az",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MultiAZ"),
			},
			{
				Name:     "node_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NodeGroups"),
			},
			{
				Name:     "pending_modified_values",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PendingModifiedValues"),
			},
			{
				Name:     "replication_group_create_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ReplicationGroupCreateTime"),
			},
			{
				Name:     "replication_group_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicationGroupId"),
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
				Name:     "snapshotting_cluster_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SnapshottingClusterId"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "transit_encryption_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("TransitEncryptionEnabled"),
			},
			{
				Name:     "user_group_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("UserGroupIds"),
			},
		},
	}
}
