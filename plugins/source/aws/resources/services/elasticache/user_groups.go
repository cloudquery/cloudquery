package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource user_groups --config ./gen.hcl --output .
func UserGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_elasticache_user_groups",
		Description:  "Describes Elasticache user groups",
		Resolver:     fetchElasticacheUserGroups,
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
				Description: "The Amazon Resource Name (ARN) of the user group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ARN"),
			},
			{
				Name:        "engine",
				Description: "The current supported value is Redis.",
				Type:        schema.TypeString,
			},
			{
				Name:        "minimum_engine_version",
				Description: "The minimum engine version required, which is Redis 6.0",
				Type:        schema.TypeString,
			},
			{
				Name:        "pending_user_ids_to_add",
				Description: "The list of user IDs to add.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("PendingChanges.UserIdsToAdd"),
			},
			{
				Name:        "pending_user_ids_to_remove",
				Description: "The list of user IDs to remove.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("PendingChanges.UserIdsToRemove"),
			},
			{
				Name:        "replication_groups",
				Description: "A list of replication groups that the user group can access.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "status",
				Description: "Indicates user group status",
				Type:        schema.TypeString,
			},
			{
				Name:        "user_group_id",
				Description: "The ID of the user group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "user_ids",
				Description: "The list of user IDs that belong to the user group.",
				Type:        schema.TypeStringArray,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchElasticacheUserGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	paginator := elasticache.NewDescribeUserGroupsPaginator(meta.(*client.Client).Services().ElastiCache, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- v.UserGroups
	}
	return nil
}
