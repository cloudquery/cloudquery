package route53recoveryreadiness

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness"
	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ReadinessChecks() *schema.Table {
	tableName := "aws_route53recoveryreadiness_readiness_checks"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/recovery-readiness/latest/api/readinesschecks.html`,
		Resolver:    fetchReadinessChecks,
		Transform:   transformers.TransformWithStruct(&types.ReadinessCheckOutput{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "route53-recovery-control-config"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ReadinessCheckArn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchReadinessChecks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRoute53recoveryreadiness).Route53recoveryreadiness
	paginator := route53recoveryreadiness.NewListReadinessChecksPaginator(svc, &route53recoveryreadiness.ListReadinessChecksInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *route53recoveryreadiness.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.ReadinessChecks
	}
	return nil
}
