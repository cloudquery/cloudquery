package applicationautoscaling

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Policies() *schema.Table {
	tableName := "aws_applicationautoscaling_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/autoscaling/application/APIReference/API_ScalingPolicy.html`,
		Resolver:    fetchPolicies,
		Multiplex:   client.ServiceAccountRegionNamespaceMultiplexer(tableName, "application-autoscaling"),
		Transform:   transformers.TransformWithStruct(&types.ScalingPolicy{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("PolicyARN"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Applicationautoscaling

	config := applicationautoscaling.DescribeScalingPoliciesInput{
		ServiceNamespace: types.ServiceNamespace(cl.AutoscalingNamespace),
	}
	paginator := applicationautoscaling.NewDescribeScalingPoliciesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *applicationautoscaling.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.ScalingPolicies
	}

	return nil
}
