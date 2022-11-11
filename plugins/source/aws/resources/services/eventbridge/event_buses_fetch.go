package eventbridge

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchEventbridgeEventBuses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
func resolveEventbridgeEventBusTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	eventBusArn := resource.Item.(types.EventBus).Arn
	return resolveEventBridgeTags(ctx, meta, resource, c, *eventBusArn)
}
func fetchEventbridgeEventBusRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
func resolveEventbridgeEventBusRuleTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	eventBusArn := resource.Item.(types.Rule).Arn
	return resolveEventBridgeTags(ctx, meta, resource, c, *eventBusArn)
}

func resolveEventBridgeTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column, resourceArn string) error {
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
