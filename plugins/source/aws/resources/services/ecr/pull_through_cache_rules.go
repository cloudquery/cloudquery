package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func PullThroughCacheRules() *schema.Table {
	tableName := "aws_ecr_pull_through_cache_rules"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_DescribePullThroughCacheRules.html`,
		Resolver:    fetchPullThroughCacheRules,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "api.ecr"),
		Transform:   transformers.TransformWithStruct(&types.PullThroughCacheRule{}, transformers.WithPrimaryKeys("RegistryId", "UpstreamRegistryUrl", "EcrRepositoryPrefix")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}
func fetchPullThroughCacheRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ecr
	paginator := ecr.NewDescribePullThroughCacheRulesPaginator(svc, nil)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *ecr.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.PullThroughCacheRules
	}
	return nil
}
