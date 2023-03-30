package applicationautoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Policies() *schema.Table {
	tableName := "aws_applicationautoscaling_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/autoscaling/application/APIReference/API_ScalingPolicy.html`,
		Resolver:    fetchPolicies,
		Multiplex:   client.ServiceAccountRegionNamespaceMultiplexer(tableName, "application-autoscaling"),
		Transform:   client.TransformWithStruct(&types.ScalingPolicy{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Applicationautoscaling

	config := applicationautoscaling.DescribeScalingPoliciesInput{
		ServiceNamespace: types.ServiceNamespace(c.AutoscalingNamespace),
	}
	for {
		output, err := svc.DescribeScalingPolicies(ctx, &config)
		if err != nil {
			return err
		}

		res <- output.ScalingPolicies

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return nil
}
