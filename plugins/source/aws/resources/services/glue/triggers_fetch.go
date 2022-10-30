package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchGlueTriggers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Glue
	input := glue.ListTriggersInput{MaxResults: aws.Int32(200)}
	for {
		response, err := svc.ListTriggers(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.TriggerNames
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func getTrigger(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	name := resource.Item.(string)
	svc := c.Services().Glue
	dc, err := svc.GetTrigger(ctx, &glue.GetTriggerInput{
		Name: &name,
	})
	if err != nil {
		return err
	}
	resource.Item = *dc.Trigger
	return nil
}

func resolveGlueTriggerArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	arn := aws.String(triggerARN(cl, aws.ToString(resource.Item.(types.Trigger).Name)))
	return resource.Set(c.Name, arn)
}

func resolveGlueTriggerTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	result, err := svc.GetTags(ctx, &glue.GetTagsInput{
		ResourceArn: aws.String(triggerARN(cl, aws.ToString(resource.Item.(types.Trigger).Name))),
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, result.Tags)
}

func triggerARN(cl *client.Client, name string) string {
	return cl.ARN(client.GlueService, "trigger", name)
}
