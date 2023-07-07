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

func groupLifecycleHooks() *schema.Table {
	tableName := "aws_autoscaling_group_lifecycle_hooks"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_LifecycleHook.html`,
		Resolver:    fetchAutoscalingGroupLifecycleHooks,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "autoscaling"),
		Transform:   transformers.TransformWithStruct(&types.LifecycleHook{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "group_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchAutoscalingGroupLifecycleHooks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(models.AutoScalingGroupWrapper)
	cl := meta.(*client.Client)
	svc := cl.Services().Autoscaling
	config := autoscaling.DescribeLifecycleHooksInput{AutoScalingGroupName: p.AutoScalingGroupName}

	output, err := svc.DescribeLifecycleHooks(ctx, &config, func(options *autoscaling.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if isAutoScalingGroupNotExistsError(err) {
			return nil
		}
		return err
	}
	res <- output.LifecycleHooks
	return nil
}
