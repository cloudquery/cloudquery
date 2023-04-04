package route53

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/route53/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveRoute53HostedZoneQueryLoggingConfigsArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "hosted_zone_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchRoute53HostedZoneQueryLoggingConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*models.Route53HostedZoneWrapper)
	svc := meta.(*client.Client).Services().Route53
	config := route53.ListQueryLoggingConfigsInput{HostedZoneId: r.Id}
	for {
		response, err := svc.ListQueryLoggingConfigs(ctx, &config, func(options *route53.Options) {})
		if err != nil {
			return err
		}
		res <- response.QueryLoggingConfigs
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
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
