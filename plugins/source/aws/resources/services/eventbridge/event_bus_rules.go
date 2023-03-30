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

func eventBusRules() *schema.Table {
	return &schema.Table{
		Name:        "aws_eventbridge_event_bus_rules",
		Description: `https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Rule.html`,
		Resolver:    fetchEventBusRules,
		Transform:   client.TransformWithStruct(&types.Rule{}, transformers.WithPrimaryKeys("Arn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "event_bus_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEventBusRuleTags,
			},
		},
		Relations: []*schema.Table{
			eventBusTargets(),
		},
	}
}

func fetchEventBusRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(types.EventBus)
	input := eventbridge.ListRulesInput{
		EventBusName: p.Arn,
	}
	c := meta.(*client.Client)
	svc := c.Services().Eventbridge
	for {
		response, err := svc.ListRules(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.Rules
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func resolveEventBusRuleTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	eventBusArn := resource.Item.(types.Rule).Arn
	return resolveTags(ctx, meta, resource, c, *eventBusArn)
}
