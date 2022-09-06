package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource replication_groups --config ./gen.hcl --output .
func ReplicationGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_elasticache_replication_groups",
		Description:  "Contains all of the attributes of a specific Redis replication group.",
		Resolver:     fetchElasticacheReplicationGroups,
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
				Description: "The ARN (Amazon Resource Name) of the replication group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ARN"),
			},
			{
				Name:        "at_rest_encryption_enabled",
				Description: "A flag that enables encryption at-rest when set to true",
				Type:        schema.TypeBool,
			},
			{
				Name:        "auth_token_enabled",
				Description: "A flag that enables using an AuthToken (password) when issuing Redis commands. Default: false",
				Type:        schema.TypeBool,
			},
			{
				Name:        "auth_token_last_modified_date",
				Description: "The date the auth token was last modified",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "auto_minor_version_upgrade",
				Description: "Auto minor version upgrade.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "automatic_failover",
				Description: "Indicates the status of automatic failover for this Redis replication group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cache_node_type",
				Description: "The name of the compute and memory capacity node type for each node in the replication group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cluster_enabled",
				Description: "A flag indicating whether or not this replication group is cluster enabled; i.e., whether its data can be partitioned across multiple shards (API/CLI: node groups)",
				Type:        schema.TypeBool,
			},
			{
				Name:        "configuration_endpoint_address",
				Description: "The DNS hostname of the cache node.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ConfigurationEndpoint.Address"),
			},
			{
				Name:        "configuration_endpoint_port",
				Description: "The port number that the cache engine is listening on.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ConfigurationEndpoint.Port"),
			},
			{
				Name:        "data_tiering",
				Description: "Enables data tiering",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "The user supplied description of the replication group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "global_replication_group_id",
				Description: "The name of the Global datastore",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("GlobalReplicationGroupInfo.GlobalReplicationGroupId"),
			},
			{
				Name:        "global_replication_group_member",
				Description: "The role of the replication group in a Global datastore",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("GlobalReplicationGroupInfo.GlobalReplicationGroupMemberRole"),
			},
			{
				Name:        "kms_key_id",
				Description: "The ID of the KMS key used to encrypt the disk in the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "member_clusters",
				Description: "The names of all the cache clusters that are part of this replication group.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "member_clusters_outpost_arns",
				Description: "The outpost ARNs of the replication group's member clusters.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "multi_az",
				Description: "A flag indicating if you have Multi-AZ enabled to enhance fault tolerance",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MultiAZ"),
			},
			{
				Name:        "pending_auth_token_status",
				Description: "Pending modified auth token status",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.AuthTokenStatus"),
			},
			{
				Name:        "pending_automatic_failover_status",
				Description: "pending autmatic failover for this redis replication group",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.AutomaticFailoverStatus"),
			},
			{
				Name:        "pending_primary_cluster_id",
				Description: "The primary cluster ID that is applied immediately (if --apply-immediately was specified), or during the next maintenance window.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.PrimaryClusterId"),
			},
			{
				Name:        "pending_resharding_slot_migration_progress_percentage",
				Description: "The percentage of the slot migration that is complete.",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("PendingModifiedValues.Resharding.SlotMigration.ProgressPercentage"),
			},
			{
				Name:        "pending_user_group_ids_to_add",
				Description: "The ID of the user group to add.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("PendingModifiedValues.UserGroups.UserGroupIdsToAdd"),
			},
			{
				Name:        "pending_user_group_ids_to_remove",
				Description: "The ID of the user group to remove.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("PendingModifiedValues.UserGroups.UserGroupIdsToRemove"),
			},
			{
				Name:        "replication_group_create_time",
				Description: "The date and time when the cluster was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "replication_group_id",
				Description: "The identifier for the replication group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "snapshot_retention_limit",
				Description: "The number of days for which ElastiCache retains automatic cluster snapshots before deleting them",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "snapshot_window",
				Description: "The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard)",
				Type:        schema.TypeString,
			},
			{
				Name:        "snapshotting_cluster_id",
				Description: "The cluster ID that is used as the daily snapshot source for the replication group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The current state of this replication group - creating, available, modifying, deleting, create-failed, snapshotting.",
				Type:        schema.TypeString,
			},
			{
				Name:        "transit_encryption_enabled",
				Description: "A flag that enables in-transit encryption when set to true",
				Type:        schema.TypeBool,
			},
			{
				Name:        "user_group_ids",
				Description: "The ID of the user group associated to the replication group.",
				Type:        schema.TypeStringArray,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_elasticache_replication_group_log_delivery_configurations",
				Description: "Returns the destination, format and type of the logs.",
				Resolver:    schema.PathTableResolver("LogDeliveryConfigurations"),
				Columns: []schema.Column{
					{
						Name:        "replication_group_cq_id",
						Description: "Unique CloudQuery ID of aws_elasticache_replication_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "cloudwatch_destination_log_group",
						Description: "The log group of the CloudWatch Logs destination",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DestinationDetails.CloudWatchLogsDetails.LogGroup"),
					},
					{
						Name:        "kinesis_firehose_destination_delivery_stream",
						Description: "The Kinesis Data Firehose delivery stream of the Kinesis Data Firehose destination",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DestinationDetails.KinesisFirehoseDetails.DeliveryStream"),
					},
					{
						Name:        "destination_type",
						Description: "Returns the destination type, either cloudwatch-logs or kinesis-firehose.",
						Type:        schema.TypeString,
					},
					{
						Name:        "log_format",
						Description: "Returns the log format, either JSON or TEXT.",
						Type:        schema.TypeString,
					},
					{
						Name:        "log_type",
						Description: "Refers to slow-log (https://redis.io/commands/slowlog) or engine-log.",
						Type:        schema.TypeString,
					},
					{
						Name:        "message",
						Description: "Returns an error message for the log delivery configuration.",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "Returns the log delivery configuration status",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_elasticache_replication_group_node_groups",
				Description: "Represents a collection of cache nodes in a replication group",
				Resolver:    schema.PathTableResolver("NodeGroups"),
				Columns: []schema.Column{
					{
						Name:        "replication_group_cq_id",
						Description: "Unique CloudQuery ID of aws_elasticache_replication_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "node_group_id",
						Description: "The identifier for the node group (shard)",
						Type:        schema.TypeString,
					},
					{
						Name:        "primary_endpoint_address",
						Description: "The DNS hostname of the cache node.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrimaryEndpoint.Address"),
					},
					{
						Name:        "primary_endpoint_port",
						Description: "The port number that the cache engine is listening on.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("PrimaryEndpoint.Port"),
					},
					{
						Name:        "reader_endpoint_address",
						Description: "The DNS hostname of the cache node.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ReaderEndpoint.Address"),
					},
					{
						Name:        "reader_endpoint_port",
						Description: "The port number that the cache engine is listening on.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ReaderEndpoint.Port"),
					},
					{
						Name:        "slots",
						Description: "The keyspace for this node group (shard).",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "The current state of this replication group - creating, available, modifying, deleting.",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_elasticache_replication_group_node_group_members",
						Description: "Represents a single node within a node group (shard).",
						Resolver:    schema.PathTableResolver("NodeGroupMembers"),
						Columns: []schema.Column{
							{
								Name:        "replication_group_node_group_cq_id",
								Description: "Unique CloudQuery ID of aws_elasticache_replication_group_node_groups table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "cache_cluster_id",
								Description: "The ID of the cluster to which the node belongs.",
								Type:        schema.TypeString,
							},
							{
								Name:        "cache_node_id",
								Description: "The ID of the node within its cluster",
								Type:        schema.TypeString,
							},
							{
								Name:        "current_role",
								Description: "The role that is currently assigned to the node - primary or replica",
								Type:        schema.TypeString,
							},
							{
								Name:        "preferred_availability_zone",
								Description: "The name of the Availability Zone in which the node is located.",
								Type:        schema.TypeString,
							},
							{
								Name:        "preferred_outpost_arn",
								Description: "The outpost ARN of the node group member.",
								Type:        schema.TypeString,
							},
							{
								Name:        "read_endpoint_address",
								Description: "The DNS hostname of the cache node.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ReadEndpoint.Address"),
							},
							{
								Name:        "read_endpoint_port",
								Description: "The port number that the cache engine is listening on.",
								Type:        schema.TypeBigInt,
								Resolver:    schema.PathResolver("ReadEndpoint.Port"),
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchElasticacheReplicationGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	paginator := elasticache.NewDescribeReplicationGroupsPaginator(meta.(*client.Client).Services().ElastiCache, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- v.ReplicationGroups
	}
	return nil
}
