package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource global_replication_groups --config ./gen.hcl --output .
func GlobalReplicationGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_elasticache_global_replication_groups",
		Description:  "Consists of a primary cluster that accepts writes and an associated secondary cluster that resides in a different Amazon region",
		Resolver:     fetchElasticacheGlobalReplicationGroups,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "arn",
				Description: "The ARN (Amazon Resource Name) of the global replication group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ARN"),
			},
			{
				Name:        "at_rest_encryption_enabled",
				Description: "A flag that enables encryption at rest when set to true",
				Type:        schema.TypeBool,
			},
			{
				Name:        "auth_token_enabled",
				Description: "A flag that enables using an AuthToken (password) when issuing Redis commands. Default: false",
				Type:        schema.TypeBool,
			},
			{
				Name:        "cache_node_type",
				Description: "The cache node type of the Global datastore",
				Type:        schema.TypeString,
			},
			{
				Name:        "cluster_enabled",
				Description: "A flag that indicates whether the Global datastore is cluster enabled.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "engine",
				Description: "The Elasticache engine",
				Type:        schema.TypeString,
			},
			{
				Name:        "engine_version",
				Description: "The Elasticache Redis engine version.",
				Type:        schema.TypeString,
			},
			{
				Name:        "global_replication_group_description",
				Description: "The optional description of the Global datastore",
				Type:        schema.TypeString,
			},
			{
				Name:        "global_replication_group_id",
				Description: "The name of the Global datastore",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The status of the Global datastore",
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
				Name:        "aws_elasticache_global_replication_group_global_node_groups",
				Description: "Indicates the slot configuration and global identifier for a slice group.",
				Resolver:    schema.PathTableResolver("GlobalNodeGroups"),
				Columns: []schema.Column{
					{
						Name:        "global_replication_group_cq_id",
						Description: "Unique CloudQuery ID of aws_elasticache_global_replication_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "global_node_group_id",
						Description: "The name of the global node group",
						Type:        schema.TypeString,
					},
					{
						Name:        "slots",
						Description: "The keyspace for this node group",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_elasticache_global_replication_group_members",
				Description: "A member of a Global datastore",
				Resolver:    schema.PathTableResolver("Members"),
				Columns: []schema.Column{
					{
						Name:        "global_replication_group_cq_id",
						Description: "Unique CloudQuery ID of aws_elasticache_global_replication_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "automatic_failover",
						Description: "Indicates whether automatic failover is enabled for the replication group.",
						Type:        schema.TypeString,
					},
					{
						Name:        "replication_group_id",
						Description: "The replication group id of the Global datastore member.",
						Type:        schema.TypeString,
					},
					{
						Name:        "replication_group_region",
						Description: "The Amazon region of the Global datastore member.",
						Type:        schema.TypeString,
					},
					{
						Name:        "role",
						Description: "Indicates the role of the replication group, primary or secondary.",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "The status of the membership of the replication group.",
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

func fetchElasticacheGlobalReplicationGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	paginator := elasticache.NewDescribeGlobalReplicationGroupsPaginator(meta.(*client.Client).Services().ElastiCache, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- v.GlobalReplicationGroups
	}
	return nil
}
