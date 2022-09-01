package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/mitchellh/mapstructure"
)

//go:generate cq-gen --resource subscriptions --config subscriptions.hcl --output .
func Subscriptions() *schema.Table {
	return &schema.Table{
		Name:         "aws_sns_subscriptions",
		Description:  "Amazon SNS subscription",
		Resolver:     fetchSnsSubscriptions,
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
				Name:        "endpoint",
				Description: "The subscription's endpoint (format depends on the protocol)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Subscription.Endpoint"),
			},
			{
				Name:        "owner",
				Description: "The subscription's owner",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Subscription.Owner"),
			},
			{
				Name:        "protocol",
				Description: "The subscription's protocol",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Subscription.Protocol"),
			},
			{
				Name:        "arn",
				Description: "The subscription's ARN",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Subscription.SubscriptionArn"),
			},
			{
				Name:        "topic_arn",
				Description: "The ARN of the subscription's topic",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Subscription.TopicArn"),
			},
			{
				Name:        "confirmation_was_authenticated",
				Description: "True if the subscription confirmation request was authenticated",
				Type:        schema.TypeBool,
			},
			{
				Name:        "delivery_policy",
				Description: "The JSON serialization of the subscription's delivery policy",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "effective_delivery_policy",
				Description: "The JSON serialization of the effective delivery policy that takes into account the topic delivery policy and account system defaults",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "filter_policy",
				Description: "The filter policy JSON that is assigned to the subscription",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "pending_confirmation",
				Description: "True if the subscription hasn't been confirmed",
				Type:        schema.TypeBool,
			},
			{
				Name:        "raw_message_delivery",
				Description: "True if raw message delivery is enabled for the subscription",
				Type:        schema.TypeBool,
			},
			{
				Name:        "redrive_policy",
				Description: "When specified, sends undeliverable messages to the specified Amazon SQS dead-letter queue",
				Type:        schema.TypeString,
			},
			{
				Name:        "subscription_role_arn",
				Description: "The ARN of the IAM role that has permission to write to the Kinesis Data Firehose delivery stream and has Amazon SNS listed as a trusted entity",
				Type:        schema.TypeString,
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

func fetchSnsSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return diag.WrapError(client.ListAndDetailResolver(ctx, meta, res, listSubscriptions, subscriptionDetail))
}

func listSubscriptions(ctx context.Context, meta schema.ClientMeta, detailChan chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().SNS
	config := sns.ListSubscriptionsInput{}
	for {
		output, err := svc.ListSubscriptions(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
		}
		for _, item := range output.Subscriptions {
			detailChan <- item
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func subscriptionDetail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, summary interface{}) {
	c := meta.(*client.Client)
	svc := c.Services().SNS
	item := summary.(types.Subscription)
	s := Subscription{Subscription: item}
	// Return early if SubscriptionARN is not set because it is still pending
	if aws.ToString(item.SubscriptionArn) == "PendingConfirmation" {
		resultsChan <- s
		return
	}

	attrs, err := svc.GetSubscriptionAttributes(ctx, &sns.GetSubscriptionAttributesInput{SubscriptionArn: item.SubscriptionArn})
	if err != nil {
		if c.IsNotFoundError(err) {
			return
		}
		errorChan <- diag.WrapError(err)
		return
	}
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &s})
	if err != nil {
		errorChan <- diag.WrapError(err)
		return
	}
	if err := dec.Decode(attrs.Attributes); err != nil {
		errorChan <- diag.WrapError(err)
		return
	}
	resultsChan <- s
}
