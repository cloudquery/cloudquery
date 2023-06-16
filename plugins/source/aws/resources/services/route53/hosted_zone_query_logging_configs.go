package route53

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/route53/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func hostedZoneQueryLoggingConfigs() *schema.Table {
	tableName := "aws_route53_hosted_zone_query_logging_configs"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/Route53/latest/APIReference/API_QueryLoggingConfig.html`,
		Resolver:    fetchRoute53HostedZoneQueryLoggingConfigs,
		Transform:   transformers.TransformWithStruct(&types.QueryLoggingConfig{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "route53"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveRoute53HostedZoneQueryLoggingConfigsArn,
				PrimaryKey: true,
			},
			{
				Name:     "hosted_zone_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchRoute53HostedZoneQueryLoggingConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*models.Route53HostedZoneWrapper)
	cl := meta.(*client.Client)
	svc := cl.Services().Route53
	config := route53.ListQueryLoggingConfigsInput{HostedZoneId: r.Id}
	paginator := route53.NewListQueryLoggingConfigsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *route53.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.QueryLoggingConfigs
	}
	return nil
}

func resolveRoute53HostedZoneQueryLoggingConfigsArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	ql := resource.Item.(types.QueryLoggingConfig)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.Route53Service),
		Region:    "",
		AccountID: "",
		Resource:  fmt.Sprintf("queryloggingconfig/%s", aws.ToString(ql.Id)),
	}.String())
}
