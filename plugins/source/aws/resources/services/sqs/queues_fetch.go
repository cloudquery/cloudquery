package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sqs/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/mitchellh/mapstructure"
)

func fetchSqsQueues(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Sqs
	var params sqs.ListQueuesInput
	for {
		result, err := svc.ListQueues(ctx, &params)
		if err != nil {
			return err
		}
		res <- result.QueueUrls

		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return nil
}

func getQueue(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Sqs
	qURL := resource.Item.(string)

	input := sqs.GetQueueAttributesInput{
		QueueUrl:       aws.String(qURL),
		AttributeNames: []types.QueueAttributeName{types.QueueAttributeNameAll},
	}
	out, err := svc.GetQueueAttributes(ctx, &input)
	if err != nil {
		return err
	}

	q := &models.Queue{URL: qURL}
	d, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: q})
	if err != nil {
		return err
	}
	if err := d.Decode(out.Attributes); err != nil {
		return err
	}

	resource.Item = q
	return nil
}

func resolveSqsQueueTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sqs
	q := resource.Item.(*models.Queue)
	result, err := svc.ListQueueTags(ctx, &sqs.ListQueueTagsInput{QueueUrl: &q.URL})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, result.Tags)
}
