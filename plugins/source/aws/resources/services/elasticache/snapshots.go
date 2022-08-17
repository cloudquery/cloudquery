package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource snapshots --config ./gen.hcl --output .
func Snapshots() *schema.Table {
	return &schema.Table{
		Name:         "aws_elasticache_snapshots",
		Description:  "Represents a copy of an entire Redis cluster as of the time when the snapshot was taken.",
		Resolver:     fetchElasticacheSnapshots,
		Multiplex:    client.ServiceAccountRegionMultiplexer("elasticache"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The ARN (Amazon Resource Name) of the snapshot.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ARN"),
			},
			{
				Name:        "auto_minor_version_upgrade",
				Description: "Auto minor version upgrade",
				Type:        schema.TypeBool,
			},
			{
				Name:        "automatic_failover",
				Description: "Indicates the status of automatic failover for the source Redis replication group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cache_cluster_create_time",
				Description: "The date and time when the source cluster was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "cache_cluster_id",
				Description: "The user-supplied identifier of the source cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cache_node_type",
				Description: "The name of the compute and memory capacity node type for the source cluster. The following node types are supported by ElastiCache",
				Type:        schema.TypeString,
			},
			{
				Name:        "cache_parameter_group_name",
				Description: "The cache parameter group that is associated with the source cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cache_subnet_group_name",
				Description: "The name of the cache subnet group associated with the source cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "data_tiering",
				Description: "Data tiering",
				Type:        schema.TypeString,
			},
			{
				Name:        "engine",
				Description: "The name of the cache engine (memcached or redis) used by the source cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "engine_version",
				Description: "The version of the cache engine version that is used by the source cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "kms_key_id",
				Description: "The ID of the KMS key used to encrypt the snapshot.",
				Type:        schema.TypeString,
			},
			{
				Name:        "num_cache_nodes",
				Description: "The number of cache nodes in the source cluster",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "num_node_groups",
				Description: "The number of node groups (shards) in this snapshot",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "port",
				Description: "The port number used by each cache nodes in the source cluster.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "preferred_availability_zone",
				Description: "The name of the Availability Zone in which the source cluster is located.",
				Type:        schema.TypeString,
			},
			{
				Name:        "preferred_maintenance_window",
				Description: "Specifies the weekly time range during which maintenance on the cluster is performed",
				Type:        schema.TypeString,
			},
			{
				Name:        "preferred_outpost_arn",
				Description: "The ARN (Amazon Resource Name) of the preferred outpost.",
				Type:        schema.TypeString,
			},
			{
				Name:        "replication_group_description",
				Description: "A description of the source replication group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "replication_group_id",
				Description: "The unique identifier of the source replication group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "snapshot_name",
				Description: "The name of a snapshot",
				Type:        schema.TypeString,
			},
			{
				Name:        "snapshot_retention_limit",
				Description: "For an automatic snapshot, the number of days for which ElastiCache retains the snapshot before deleting it",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "snapshot_source",
				Description: "Indicates whether the snapshot is from an automatic backup (automated) or was created manually (manual).",
				Type:        schema.TypeString,
			},
			{
				Name:        "snapshot_status",
				Description: "The status of the snapshot",
				Type:        schema.TypeString,
			},
			{
				Name:        "snapshot_window",
				Description: "The daily time range during which ElastiCache takes daily snapshots of the source cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "topic_arn",
				Description: "The Amazon Resource Name (ARN) for the topic used by the source cluster for publishing notifications.",
				Type:        schema.TypeString,
			},
			{
				Name:        "vpc_id",
				Description: "The Amazon Virtual Private Cloud identifier (VPC ID) of the cache subnet group for the source cluster.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_elasticache_snapshot_node_snapshots",
				Description: "Represents an individual cache node in a snapshot of a cluster.",
				Resolver:    schema.PathTableResolver("NodeSnapshots"),
				Columns: []schema.Column{
					{
						Name:        "snapshot_cq_id",
						Description: "Unique CloudQuery ID of aws_elasticache_snapshots table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "cache_cluster_id",
						Description: "A unique identifier for the source cluster.",
						Type:        schema.TypeString,
					},
					{
						Name:        "cache_node_create_time",
						Description: "The date and time when the cache node was created in the source cluster.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "cache_node_id",
						Description: "The cache node identifier for the node in the source cluster.",
						Type:        schema.TypeString,
					},
					{
						Name:        "cache_size",
						Description: "The size of the cache on the source cache node.",
						Type:        schema.TypeString,
					},
					{
						Name:        "node_group_configuration_node_group_id",
						Description: "Either the ElastiCache for Redis supplied 4-digit id or a user supplied id for the node group these configuration values apply to.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("NodeGroupConfiguration.NodeGroupId"),
					},
					{
						Name:        "node_group_configuration_primary_availability_zone",
						Description: "The Availability Zone where the primary node of this node group (shard) is launched.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("NodeGroupConfiguration.PrimaryAvailabilityZone"),
					},
					{
						Name:        "node_group_configuration_primary_outpost_arn",
						Description: "The outpost ARN of the primary node.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("NodeGroupConfiguration.PrimaryOutpostArn"),
					},
					{
						Name:        "node_group_configuration_replica_availability_zones",
						Description: "A list of Availability Zones to be used for the read replicas",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("NodeGroupConfiguration.ReplicaAvailabilityZones"),
					},
					{
						Name:        "node_group_configuration_replica_count",
						Description: "The number of read replica nodes in this node group (shard).",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("NodeGroupConfiguration.ReplicaCount"),
					},
					{
						Name:        "node_group_configuration_replica_outpost_arns",
						Description: "The outpost ARN of the node replicas.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("NodeGroupConfiguration.ReplicaOutpostArns"),
					},
					{
						Name:        "node_group_configuration_slots",
						Description: "A string that specifies the keyspace for a particular node group",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("NodeGroupConfiguration.Slots"),
					},
					{
						Name:        "node_group_id",
						Description: "A unique identifier for the source node group (shard).",
						Type:        schema.TypeString,
					},
					{
						Name:        "snapshot_create_time",
						Description: "The date and time when the source node's metadata and cache data set was obtained for the snapshot.",
						Type:        schema.TypeTimestamp,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchElasticacheSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	paginator := elasticache.NewDescribeSnapshotsPaginator(meta.(*client.Client).Services().ElastiCache, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- v.Snapshots
	}
	return nil
}
