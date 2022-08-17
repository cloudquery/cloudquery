package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/mitchellh/mapstructure"
)

//go:generate cq-gen --resource topics --config topics.hcl --output .
func Topics() *schema.Table {
	return &schema.Table{
		Name:         "aws_sns_topics",
		Description:  "Amazon SNS topic",
		Resolver:     fetchSnsTopics,
		Multiplex:    client.ServiceAccountRegionMultiplexer("sns"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveSnsTopicTags,
			},
			{
				Name:        "delivery_policy",
				Description: "The JSON serialization of the topic's delivery policy",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "display_name",
				Description: "The human-readable name used in the From field for notifications to email and email-json endpoints",
				Type:        schema.TypeString,
			},
			{
				Name:        "owner",
				Description: "The AWS account ID of the topic's owner",
				Type:        schema.TypeString,
			},
			{
				Name:        "policy",
				Description: "The JSON serialization of the topic's access control policy",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "subscriptions_confirmed",
				Description: "The number of confirmed subscriptions for the topic",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "subscriptions_deleted",
				Description: "The number of deleted subscriptions for the topic",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "subscriptions_pending",
				Description: "The number of subscriptions pending confirmation for the topic",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the topic",
				Type:        schema.TypeString,
			},
			{
				Name:        "effective_delivery_policy",
				Description: "The JSON serialization of the effective delivery policy, taking system defaults into account",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "kms_master_key_id",
				Description: "The ID of an Amazon Web Services managed customer master key (CMK) for Amazon SNS or a custom CMK",
				Type:        schema.TypeString,
			},
			{
				Name:        "fifo_topic",
				Description: "When this is set to true, a FIFO topic is created",
				Type:        schema.TypeBool,
			},
			{
				Name:        "content_based_deduplication",
				Description: "Enables content-based deduplication for FIFO topics",
				Type:        schema.TypeBool,
			},
			{
				Name:        "unknown_fields",
				Description: "Other subscription attributes",
				Type:        schema.TypeJSON,
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
		output, err := svc.ListTopics(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
		}
		for _, topic := range output.Topics {
			attrs, err := svc.GetTopicAttributes(ctx, &sns.GetTopicAttributesInput{TopicArn: topic.TopicArn})
			if err != nil {
				if c.IsNotFoundError(err) {
					continue
				}
				return diag.WrapError(err)
			}
			t := Topic{Arn: topic.TopicArn}
			dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &t})
			if err != nil {
				return diag.WrapError(err)
			}
			if err := dec.Decode(attrs.Attributes); err != nil {
				return diag.WrapError(err)
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
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, client.TagsToMap(tags.Tags)))
}
