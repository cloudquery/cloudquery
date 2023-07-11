package route53resolver

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ResolverQueryLogConfigs() *schema.Table {
	tableName := "aws_route53resolver_resolver_query_log_configs"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ResolverQueryLogConfig.html`,
		Resolver:    fetchQueryLogConfigs,
		Transform:   transformers.TransformWithStruct(&types.ResolverQueryLogConfig{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "route53resolver"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}

func fetchQueryLogConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Route53resolver
	var input route53resolver.ListResolverQueryLogConfigsInput
	paginator := route53resolver.NewListResolverQueryLogConfigsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *route53resolver.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.ResolverQueryLogConfigs
	}
	return nil
}
