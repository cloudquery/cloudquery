package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchAutoscalingScheduledActions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Autoscaling
	params := &autoscaling.DescribeScheduledActionsInput{
		MaxRecords: aws.Int32(100),
	}
	for {
		output, err := svc.DescribeScheduledActions(ctx, params)
		if err != nil {
			return err
		}
		for _, scheduledUpdateGroupAction := range output.ScheduledUpdateGroupActions {
			res <- scheduledUpdateGroupAction
		}
		if aws.ToString(output.NextToken) == "" {
			break
		}
		params.NextToken = output.NextToken
	}
	return nil
}
