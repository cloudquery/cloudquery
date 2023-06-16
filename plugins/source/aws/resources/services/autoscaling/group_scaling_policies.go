package autoscaling

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/autoscaling/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func groupScalingPolicies() *schema.Table {
	tableName := "aws_autoscaling_group_scaling_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_ScalingPolicy.html`,
		Resolver:    fetchAutoscalingGroupScalingPolicies,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "autoscaling"),
		Transform:   transformers.TransformWithStruct(&types.ScalingPolicy{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "group_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("PolicyARN"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchAutoscalingGroupScalingPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(models.AutoScalingGroupWrapper)
	cl := meta.(*client.Client)
	svc := cl.Services().Autoscaling
	paginator := autoscaling.NewDescribePoliciesPaginator(svc, &autoscaling.DescribePoliciesInput{AutoScalingGroupName: p.AutoScalingGroupName})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *autoscaling.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if isAutoScalingGroupNotExistsError(err) {
				return nil
			}
			return err
		}
		res <- page.ScalingPolicies
	}
	return nil
}
