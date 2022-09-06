package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource clusters --config ./gen.hcl --output .
func Clusters() *schema.Table {
	return &schema.Table{
		Name:         "aws_elasticache_clusters",
		Description:  "Contains all of the attributes of a specific cluster.",
		Resolver:     fetchElasticacheClusters,
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
				Description: "The ARN (Amazon Resource Name) of the cache cluster.",
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
				Description: "Auto minor version upgrade",
				Type:        schema.TypeBool,
			},
			{
				Name:        "create_time",
				Description: "The date and time when the cluster was created.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("CacheClusterCreateTime"),
			},
			{
				Name:        "id",
				Description: "The user-supplied identifier of the cluster",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CacheClusterId"),
			},
			{
				Name:        "status",
				Description: "The current state of this cluster, one of the following values: available, creating, deleted, deleting, incompatible-network, modifying, rebooting cluster nodes, restore-failed, or snapshotting.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CacheClusterStatus"),
			},
			{
				Name:        "cache_node_type",
				Description: "The name of the compute and memory capacity node type for the cluster",
				Type:        schema.TypeString,
			},
			{
				Name:        "cache_parameter_group_cache_node_ids_to_reboot",
				Description: "A list of the cache node IDs which need to be rebooted for parameter changes to be applied",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("CacheParameterGroup.CacheNodeIdsToReboot"),
			},
			{
				Name:        "cache_parameter_group_name",
				Description: "The name of the cache parameter group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CacheParameterGroup.CacheParameterGroupName"),
			},
			{
				Name:        "cache_parameter_group_parameter_apply_status",
				Description: "The status of parameter updates.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CacheParameterGroup.ParameterApplyStatus"),
			},
			{
				Name:        "cache_subnet_group_name",
				Description: "The name of the cache subnet group associated with the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "client_download_landing_page",
				Description: "The URL of the web page where you can download the latest ElastiCache client library.",
				Type:        schema.TypeString,
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
				Name:        "engine",
				Description: "The name of the cache engine (memcached or redis) to be used for this cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "engine_version",
				Description: "The version of the cache engine that is used in this cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "notification_configuration_topic_arn",
				Description: "The arn of a notification topic used for publishing ElastiCache events to subscribers using Amazon Simple Notification Service (SNS)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("NotificationConfiguration.TopicArn"),
			},
			{
				Name:        "notification_configuration_topic_status",
				Description: "The current state of a notification topic used for publishing ElastiCache events to subscribers using Amazon Simple Notification Service (SNS)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("NotificationConfiguration.TopicStatus"),
			},
			{
				Name:        "num_cache_nodes",
				Description: "The number of cache nodes in the cluster",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "pending_auth_token_status",
				Description: "Auth token status that is applied to the cluster in the future or is currently being applied",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.AuthTokenStatus"),
			},
			{
				Name:        "pending_cache_node_ids_to_remove",
				Description: "A list of cache node IDs that are being removed (or will be removed) from the cluster",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("PendingModifiedValues.CacheNodeIdsToRemove"),
			},
			{
				Name:        "pending_cache_node_type",
				Description: "The cache node type that this cluster or replication group is scaled to.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.CacheNodeType"),
			},
			{
				Name:        "pending_engine_version",
				Description: "Cache engine version that is being applied to the cluster (or will be applied)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.EngineVersion"),
			},
			{
				Name:        "pending_num_cache_nodes",
				Description: "The new number of cache nodes for the cluster",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("PendingModifiedValues.NumCacheNodes"),
			},
			{
				Name:        "preferred_availability_zone",
				Description: "The name of the Availability Zone in which the cluster is located or \"Multiple\" if the cache nodes are located in different Availability Zones.",
				Type:        schema.TypeString,
			},
			{
				Name:        "preferred_maintenance_window",
				Description: "Specifies the weekly time range during which maintenance on the cluster is performed",
				Type:        schema.TypeString,
			},
			{
				Name:        "preferred_outpost_arn",
				Description: "The outpost ARN in which the cache cluster is created.",
				Type:        schema.TypeString,
			},
			{
				Name:        "replication_group_id",
				Description: "The replication group to which this cluster belongs",
				Type:        schema.TypeString,
			},
			{
				Name:        "replication_group_log_delivery_enabled",
				Description: "A boolean value indicating whether log delivery is enabled for the replication group.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "snapshot_retention_limit",
				Description: "The number of days for which ElastiCache retains automatic cluster snapshots before deleting them",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "snapshot_window",
				Description: "The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your cluster",
				Type:        schema.TypeString,
			},
			{
				Name:        "transit_encryption_enabled",
				Description: "A flag that enables in-transit encryption when set to true",
				Type:        schema.TypeBool,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_elasticache_cluster_cache_nodes",
				Description: "Represents an individual cache node within a cluster",
				Resolver:    schema.PathTableResolver("CacheNodes"),
				Columns: []schema.Column{
					{
						Name:        "cluster_cq_id",
						Description: "Unique CloudQuery ID of aws_elasticache_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "create_time",
						Description: "The date and time when the cache node was created.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("CacheNodeCreateTime"),
					},
					{
						Name:        "id",
						Description: "The cache node identifier",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CacheNodeId"),
					},
					{
						Name:        "status",
						Description: "The current state of this cache node, one of the following values: available, creating, rebooting, or deleting.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CacheNodeStatus"),
					},
					{
						Name:        "customer_availability_zone",
						Description: "The Availability Zone where this node was created and now resides.",
						Type:        schema.TypeString,
					},
					{
						Name:        "customer_outpost_arn",
						Description: "The customer outpost ARN of the cache node.",
						Type:        schema.TypeString,
					},
					{
						Name:        "endpoint_address",
						Description: "The DNS hostname of the cache node.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Endpoint.Address"),
					},
					{
						Name:        "endpoint_port",
						Description: "The port number that the cache engine is listening on.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("Endpoint.Port"),
					},
					{
						Name:        "parameter_group_status",
						Description: "The status of the parameter group applied to this cache node.",
						Type:        schema.TypeString,
					},
					{
						Name:        "source_cache_node_id",
						Description: "The ID of the primary node to which this read replica node is synchronized",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_elasticache_cluster_cache_security_groups",
				Description: "Represents a cluster's status within a particular cache security group.",
				Resolver:    schema.PathTableResolver("CacheSecurityGroups"),
				Columns: []schema.Column{
					{
						Name:        "cluster_cq_id",
						Description: "Unique CloudQuery ID of aws_elasticache_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "The name of the cache security group.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CacheSecurityGroupName"),
					},
					{
						Name:        "status",
						Description: "The membership status in the cache security group",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_elasticache_cluster_log_delivery_configurations",
				Description: "Returns the destination, format and type of the logs.",
				Resolver:    schema.PathTableResolver("LogDeliveryConfigurations"),
				Columns: []schema.Column{
					{
						Name:        "cluster_cq_id",
						Description: "Unique CloudQuery ID of aws_elasticache_clusters table (FK)",
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
				Name:        "aws_elasticache_cluster_security_groups",
				Description: "Represents a single cache security group and its status.",
				Resolver:    schema.PathTableResolver("SecurityGroups"),
				Columns: []schema.Column{
					{
						Name:        "cluster_cq_id",
						Description: "Unique CloudQuery ID of aws_elasticache_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "security_group_id",
						Description: "The identifier of the cache security group.",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "The status of the cache security group membership",
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

func fetchElasticacheClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input elasticache.DescribeCacheClustersInput
	input.ShowCacheNodeInfo = aws.Bool(true)

	paginator := elasticache.NewDescribeCacheClustersPaginator(meta.(*client.Client).Services().ElastiCache, &input)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- v.CacheClusters
	}
	return nil
}
