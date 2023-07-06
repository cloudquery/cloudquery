package eventbridge

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func eventBusRules() *schema.Table {
	return &schema.Table{
		Name:        "aws_eventbridge_event_bus_rules",
		Description: `https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Rule.html`,
		Resolver:    fetchEventBusRules,
		Transform:   transformers.TransformWithStruct(&types.Rule{}, transformers.WithPrimaryKeys("Arn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "event_bus_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
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
	cl := meta.(*client.Client)
	svc := cl.Services().Eventbridge
	// No paginator available
	for {
		response, err := svc.ListRules(ctx, &input, func(options *eventbridge.Options) {
			options.Region = cl.Region
		})
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
