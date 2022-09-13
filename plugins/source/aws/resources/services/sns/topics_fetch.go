package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/mitchellh/mapstructure"
)

func fetchSnsTopics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().SNS
	config := sns.ListTopicsInput{}
	for {
		output, err := svc.ListTopics(ctx, &config)
		if err != nil {
			return err
		}
		for _, topic := range output.Topics {
			attrs, err := svc.GetTopicAttributes(ctx, &sns.GetTopicAttributesInput{TopicArn: topic.TopicArn})
			if err != nil {
				if c.IsNotFoundError(err) {
					continue
				}
				return err
			}
			t := Topic{Arn: topic.TopicArn}
			dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &t})
			if err != nil {
				return err
			}
			if err := dec.Decode(attrs.Attributes); err != nil {
				return err
			}
			res <- t
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func resolveSnsTopicTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	topic := resource.Item.(Topic)
	cl := meta.(*client.Client)
	svc := cl.Services().SNS
	tagParams := sns.ListTagsForResourceInput{
		ResourceArn: topic.Arn,
	}
	tags, err := svc.ListTagsForResource(ctx, &tagParams)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(tags.Tags))
}
