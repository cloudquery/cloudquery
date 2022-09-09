package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ReplicationGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticache_replication_groups",
		Description: "Contains all of the attributes of a specific Redis replication group.",
		Resolver:    fetchElasticacheReplicationGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticache"),
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
				Name:            "arn",
				Description:     "The ARN (Amazon Resource Name) of the replication group.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
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
				Name:     "configuration_endpoint",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ConfigurationEndpoint"),
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
				Name:     "global_replication_group_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("GlobalReplicationGroupInfo"),
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
				Name:     "pending_modified_values",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PendingModifiedValues"),
			},
			{
				Name:        "pending_primary_cluster_id",
				Description: "The primary cluster ID that is applied immediately (if --apply-immediately was specified), or during the next maintenance window.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.PrimaryClusterId"),
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
				Type:        schema.TypeInt,
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
			{
				Name:        "log_delivery_configurations",
				Description: "Returns the destination, format and type of the logs.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("LogDeliveryConfigurations"),
			},
			{
				Name:        "node_groups",
				Description: "Represents a collection of cache nodes in a replication group",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("NodeGroups"),
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
			return err
		}
		res <- v.ReplicationGroups
	}
	return nil
}
