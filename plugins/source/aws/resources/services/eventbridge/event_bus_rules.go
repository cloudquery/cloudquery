package eventbridge

import (
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func EventBusRules() *schema.Table {
	return &schema.Table{
		Name:        "aws_eventbridge_event_bus_rules",
		Description: `https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Rule.html`,
		Resolver:    fetchEventbridgeEventBusRules,
		Multiplex:   client.ServiceAccountRegionMultiplexer("events"),
		Transform:   transformers.TransformWithStruct(&types.Rule{}, transformers.WithPrimaryKeys("Arn")),
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
				Resolver: resolveEventbridgeEventBusRuleTags,
			},
		},
	}
}
