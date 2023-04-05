package eventbridge

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func eventBusTargets() *schema.Table {
	return &schema.Table{
		Name:        "aws_eventbridge_event_bus_targets",
		Description: `https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Target.html`,
		Resolver:    fetchEventBusTargets,
		Transform:   transformers.TransformWithStruct(&types.Target{}, transformers.WithPrimaryKeys("Id")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "rule_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "event_bus_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("event_bus_arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchEventBusTargets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	rule := parent.Item.(types.Rule)
	bus := parent.Parent.Item.(types.EventBus)

	input := eventbridge.ListTargetsByRuleInput{
		EventBusName: bus.Arn,
		Rule:         rule.Name,
	}
	c := meta.(*client.Client)
	svc := c.Services().Eventbridge
	for {
		response, err := svc.ListTargetsByRule(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.Targets
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
