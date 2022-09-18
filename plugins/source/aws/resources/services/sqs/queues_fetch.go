package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/mitchellh/mapstructure"
)

func fetchSqsQueues(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().SQS
	var params sqs.ListQueuesInput
	for {
		result, err := svc.ListQueues(ctx, &params)
		if err != nil {
			return err
		}

		for _, url := range result.QueueUrls {
			input := sqs.GetQueueAttributesInput{
				QueueUrl:       aws.String(url),
				AttributeNames: []types.QueueAttributeName{types.QueueAttributeNameAll},
			}
			out, err := svc.GetQueueAttributes(ctx, &input)
			if err != nil {
				if cl.IsNotFoundError(err) {
					continue
				}
				return err
			}

			var q Queue
			d, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &q})
			if err != nil {
				return err
			}
			if err := d.Decode(out.Attributes); err != nil {
				return err
			}
			q.URL = url
			res <- q
		}
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return nil
}
func resolveSqsQueueTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().SQS
	q := resource.Item.(Queue)
	result, err := svc.ListQueueTags(ctx, &sqs.ListQueueTagsInput{QueueUrl: &q.URL})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, result.Tags)
}
