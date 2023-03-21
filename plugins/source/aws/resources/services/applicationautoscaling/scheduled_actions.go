package applicationautoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ScheduledActions() *schema.Table {
	tableName := "aws_applicationautoscaling_scheduled_actions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/autoscaling/application/APIReference/API_ScheduledAction.html`,
		Resolver:    fetchScheduledActions,
		Multiplex:   client.ServiceAccountRegionNamespaceMultiplexer(tableName, "application-autoscaling"),
		Transform:   transformers.TransformWithStruct(&types.ScheduledAction{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ScheduledActionARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchScheduledActions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Applicationautoscaling

	config := applicationautoscaling.DescribeScheduledActionsInput{
		ServiceNamespace: types.ServiceNamespace(c.AutoscalingNamespace),
	}
	for {
		output, err := svc.DescribeScheduledActions(ctx, &config)
		if err != nil {
			return err
		}

		res <- output.ScheduledActions

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return nil
}
