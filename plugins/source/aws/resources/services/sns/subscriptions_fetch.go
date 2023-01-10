package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sns/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/mitchellh/mapstructure"
)

func fetchSnsSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Sns
	config := sns.ListSubscriptionsInput{}
	for {
		output, err := svc.ListSubscriptions(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.Subscriptions

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func getSnsSubscription(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Sns
	item := resource.Item.(types.Subscription)
	s := models.Subscription{
		SubscriptionArn: item.SubscriptionArn,
		Owner:           item.Owner,
		Protocol:        item.Protocol,
		TopicArn:        item.TopicArn,
		Endpoint:        item.Endpoint,
	}
	// Return early if SubscriptionARN is not set because it is still pending
	if aws.ToString(item.SubscriptionArn) == "PendingConfirmation" {
		resource.Item = s
		return nil
	}

	attrs, err := svc.GetSubscriptionAttributes(ctx, &sns.GetSubscriptionAttributesInput{SubscriptionArn: item.SubscriptionArn})
	if err != nil {
		return err
	}
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &s})
	if err != nil {
		return err
	}
	if err := dec.Decode(attrs.Attributes); err != nil {
		return err
	}
	resource.Item = s
	return nil
}
