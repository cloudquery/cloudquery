package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticache_clusters",
		Description: "Contains all of the attributes of a specific cluster.",
		Resolver:    fetchElasticacheClusters,
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
				Description:     "The ARN (Amazon Resource Name) of the cache cluster.",
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
				Name:     "cache_parameter_group",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CacheParameterGroup"),
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
				Name:     "configuration_endpoint",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ConfigurationEndpoint"),
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
				Name:        "notification_configuration",
				Type: 			schema.TypeJSON,
			},
			{
				Name:        "num_cache_nodes",
				Description: "The number of cache nodes in the cluster",
				Type:        schema.TypeInt,
			},
			{
				Name:     "pending_modified_values",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PendingModifiedValues"),
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
				Type:        schema.TypeInt,
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
			{
				Name:        "cache_nodes",
				Description: "Represents an individual cache node within a cluster",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("CacheNodes"),
			},
			{
				Name:        "cache_security_groups",
				Description: "Represents a cluster's status within a particular cache security group.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("CacheSecurityGroups"),
			},
			{
				Name:        "log_delivery_configurations",
				Description: "The destination, format and type of the logs.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("LogDeliveryConfigurations"),
			},
			{
				Name:        "security_groups",
				Description: "Represents a single cache security group and its status.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("SecurityGroups"),
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
			return err
		}
		res <- v.CacheClusters
	}
	return nil
}
