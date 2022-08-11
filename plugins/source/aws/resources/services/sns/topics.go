package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/spf13/cast"
)

func SnsTopics() *schema.Table {
	return &schema.Table{
		Name:                 "aws_sns_topics",
		Description:          "AWS SNS topic",
		Resolver:             fetchSnsTopics,
		Multiplex:            client.ServiceAccountRegionMultiplexer("sns"),
		IgnoreError:          client.IgnoreCommonErrors,
		DeleteFilter:         client.DeleteAccountRegionFilter,
		PostResourceResolver: resolveTopicAttributes,
		Options:              schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "owner",
				Description: "The AWS account ID of the topic's owner.",
				Type:        schema.TypeString,
			},
			{
				Name:        "policy",
				Description: "The JSON serialization of the topic's access control policy.",
				Type:        schema.TypeJSON,
			},
			{
				Name:          "delivery_policy",
				Description:   "The JSON serialization of the topic's delivery policy.",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:        "display_name",
				Description: "The human-readable name used in the From field for notifications to email and email-json endpoints.",
				Type:        schema.TypeString,
			},
			{
				Name:        "subscriptions_confirmed",
				Description: "The number of confirmed subscriptions for the topic.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "subscriptions_deleted",
				Description: "The number of deleted subscriptions for the topic.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "subscriptions_pending",
				Description: "The number of subscriptions pending confirmation for the topic.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "effective_delivery_policy",
				Description: "The JSON serialization of the effective delivery policy, taking system defaults into account.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "fifo_topic",
				Description: "When this is set to true, a FIFO topic is created.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "content_based_deduplication",
				Description: "Enables content-based deduplication for FIFO topics.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "kms_master_key_id",
				Description: "The ID of an AWS managed customer master key (CMK) for Amazon SNS or a custom CMK",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The topic's ARN.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TopicArn"),
			},
			{
				Name:        "tags",
				Description: "Topic tags.",
				Type:        schema.TypeJSON,
				Resolver:    resolveTopicTags,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchSnsTopics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().SNS
	config := sns.ListTopicsInput{}
	for {
		output, err := svc.ListTopics(ctx, &config, func(o *sns.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
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
	topic := resource.Item.(types.Topic)
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
		return diag.WrapError(err)
	}
	// Set all attributes
	if err := resource.Set("subscriptions_confirmed", cast.ToInt(output.Attributes["SubscriptionsConfirmed"])); err != nil {
		return diag.WrapError(err)
	}
	if err := resource.Set("subscriptions_deleted", cast.ToInt(output.Attributes["SubscriptionsDeleted"])); err != nil {
		return diag.WrapError(err)
	}
	if err := resource.Set("subscriptions_pending", cast.ToInt(output.Attributes["SubscriptionsPending"])); err != nil {
		return diag.WrapError(err)
	}
	if err := resource.Set("fifo_topic", cast.ToBool(output.Attributes["FifoTopic"])); err != nil {
		return diag.WrapError(err)
	}
	if err := resource.Set("content_based_deduplication", cast.ToBool(output.Attributes["ContentBasedDeduplication"])); err != nil {
		return diag.WrapError(err)
	}
	if p, ok := output.Attributes["Policy"]; ok && p != "" {
		if err := resource.Set("policy", p); err != nil {
			return diag.WrapError(err)
		}
	}
	if p, ok := output.Attributes["DeliveryPolicy"]; ok && p != "" {
		if err := resource.Set("delivery_policy", p); err != nil {
			return diag.WrapError(err)
		}
	}
	if err := resource.Set("display_name", output.Attributes["DisplayName"]); err != nil {
		return diag.WrapError(err)
	}
	if err := resource.Set("owner", output.Attributes["Owner"]); err != nil {
		return diag.WrapError(err)
	}
	if p, ok := output.Attributes["EffectiveDeliveryPolicy"]; ok && p != "" {
		if err := resource.Set("effective_delivery_policy", p); err != nil {
			return diag.WrapError(err)
		}
	}
	if p, ok := output.Attributes["KmsMasterKeyId"]; ok && p != "" {
		if err := resource.Set("kms_master_key_id", p); err != nil {
			return diag.WrapError(err)
		}
	}

	return nil
}

func resolveTopicTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	topic := resource.Item.(types.Topic)
	c := meta.(*client.Client)
	svc := c.Services().SNS
	tagParams := sns.ListTagsForResourceInput{
		ResourceArn: topic.TopicArn,
	}
	tags, err := svc.ListTagsForResource(ctx, &tagParams, func(o *sns.Options) {
		o.Region = c.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(col.Name, client.TagsToMap(tags.Tags)))
}
