package eventbridge

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EventBuses() *schema.Table {
	tableName := "aws_eventbridge_event_buses"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_EventBus.html`,
		Resolver:    fetchEventBuses,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "events"),
		Transform:   client.TransformWithStruct(&types.EventBus{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEventBusTags,
			},
			{
				Name: "arn",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			eventBusRules(),
		},
	}
}

func fetchEventBuses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input eventbridge.ListEventBusesInput
	c := meta.(*client.Client)
	svc := c.Services().Eventbridge
	for {
		response, err := svc.ListEventBuses(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.EventBuses
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func resolveEventBusTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	eventBusArn := resource.Item.(types.EventBus).Arn
	return resolveTags(ctx, meta, resource, c, *eventBusArn)
}

func resolveTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column, resourceArn string) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Eventbridge
	input := eventbridge.ListTagsForResourceInput{
		ResourceARN: &resourceArn,
	}
	output, err := svc.ListTagsForResource(ctx, &input)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(output.Tags))
}
