package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/spf13/cast"
)

func SnsTopics() *schema.Table {
	return &schema.Table{
		Name:                 "aws_sns_topics",
		Resolver:             fetchSnsTopics,
		Multiplex:            client.AccountRegionMultiplex,
		IgnoreError:          client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:         client.DeleteAccountRegionFilter,
		PostResourceResolver: resolveTopicAttributes,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "owner",
				Type: schema.TypeString,
			},
			{
				Name: "policy",
				Type: schema.TypeJSON,
			},
			{
				Name: "delivery_policy",
				Type: schema.TypeJSON,
			},
			{
				Name: "display_name",
				Type: schema.TypeString,
			},
			{
				Name: "subscriptions_confirmed",
				Type: schema.TypeBigInt,
			},
			{
				Name: "subscriptions_deleted",
				Type: schema.TypeBigInt,
			},
			{
				Name: "subscriptions_pending",
				Type: schema.TypeBigInt,
			},
			{
				Name: "effective_delivery_policy",
				Type: schema.TypeJSON,
			},
			{
				Name: "fifo_topic",
				Type: schema.TypeBool,
			},
			{
				Name: "content_based_deduplication",
				Type: schema.TypeBool,
			},
			{
				Name: "topic_arn",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchSnsTopics(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().SNS
	config := sns.ListTopicsInput{}
	for {
		output, err := svc.ListTopics(ctx, &config, func(o *sns.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output.Topics

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func resolveTopicAttributes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	topic, ok := resource.Item.(types.Topic)
	if !ok {
		return fmt.Errorf("%T is not topic", resource.Item)
	}
	c := meta.(*client.Client)
	svc := c.Services().SNS
	// All topic attributes are returned as a string; we have to handle type conversion
	params := sns.GetTopicAttributesInput{
		TopicArn: topic.TopicArn,
	}
	output, err := svc.GetTopicAttributes(ctx, &params, func(o *sns.Options) {
		o.Region = c.Region
	})
	if err != nil {
		return err
	}
	// Set all attributes
	if err := resource.Set("subscriptions_confirmed", cast.ToInt(output.Attributes["SubscriptionsConfirmed"])); err != nil {
		return err
	}
	if err := resource.Set("subscriptions_deleted", cast.ToInt(output.Attributes["SubscriptionsDeleted"])); err != nil {
		return err
	}
	if err := resource.Set("subscriptions_pending", cast.ToInt(output.Attributes["SubscriptionsPending"])); err != nil {
		return err
	}
	if err := resource.Set("fifo_topic", cast.ToBool(output.Attributes["FifoTopic"])); err != nil {
		return err
	}
	if err := resource.Set("content_based_deduplication", cast.ToBool(output.Attributes["ContentBasedDeduplication"])); err != nil {
		return err
	}
	if p, ok := output.Attributes["Policy"]; ok && p != "" {
		if err := resource.Set("policy", p); err != nil {
			return err
		}
	}
	if p, ok := output.Attributes["DeliveryPolicy"]; ok && p != "" {
		if err := resource.Set("delivery_policy", p); err != nil {
			return err
		}
	}
	if err := resource.Set("display_name", output.Attributes["DisplayName"]); err != nil {
		return err
	}
	if err := resource.Set("owner", output.Attributes["Owner"]); err != nil {
		return err
	}
	if p, ok := output.Attributes["EffectiveDeliveryPolicy"]; ok && p != "" {
		if err := resource.Set("effective_delivery_policy", p); err != nil {
			return err
		}
	}

	return nil
}
