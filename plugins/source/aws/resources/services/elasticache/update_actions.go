package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func UpdateActions() *schema.Table {
	tableName := "aws_elasticache_update_actions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_UpdateAction.html`,
		Resolver:    fetchElasticacheUpdateAction,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticache"),
		Transform: transformers.TransformWithStruct(&types.UpdateAction{},
			transformers.WithPrimaryKeyComponents("CacheClusterId", "ReplicationGroupId", "ServiceUpdateName"),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}

func fetchElasticacheUpdateAction(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input elasticache.DescribeUpdateActionsInput
	cl := meta.(*client.Client)

	paginator := elasticache.NewDescribeUpdateActionsPaginator(meta.(*client.Client).Services(client.AWSServiceElasticache).Elasticache, &input)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx, func(options *elasticache.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		for _, ua := range v.UpdateActions {
			// We can either have cache cluster ID or replication group ID,
			// but we don't want to have nil values in PK columns.
			// See AWS CLI examples:
			// https://docs.aws.amazon.com/cli/latest/reference/elasticache/describe-update-actions.html#examples
			if ua.CacheClusterId == nil {
				ua.CacheClusterId = aws.String("")
			}
			if ua.ReplicationGroupId == nil {
				ua.ReplicationGroupId = aws.String("")
			}
		}
		res <- v.UpdateActions
	}
	return nil
}
