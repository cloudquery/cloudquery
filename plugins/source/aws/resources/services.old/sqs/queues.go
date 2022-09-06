package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/mitchellh/mapstructure"
)

//go:generate cq-gen --resource queues --config queues.hcl --output .
func Queues() *schema.Table {
	return &schema.Table{
		Name:         "aws_sqs_queues",
		Description:  "Amazon Simple Queue Service",
		Resolver:     fetchSqsQueues,
		Multiplex:    client.ServiceAccountRegionMultiplexer("sqs"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveSqsQueueTags,
			},
			{
				Name:        "url",
				Description: "The URL of the Amazon SQS queue",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("URL"),
			},
			{
				Name:        "approximate_number_of_messages",
				Description: "The approximate number of messages available for retrieval from the queue",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "approximate_number_of_messages_delayed",
				Description: "The approximate number of messages in the queue that are delayed and not available for reading immediately",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "approximate_number_of_messages_not_visible",
				Description: "The approximate number of messages that are in flight",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "created_timestamp",
				Description: "The time when the queue was created in seconds (epoch time)",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "delay_seconds",
				Description: "The default delay on the queue in seconds",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "last_modified_timestamp",
				Description: "The time when the queue was last changed in seconds (epoch time)",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "maximum_message_size",
				Description: "The limit of how many bytes a message can contain before Amazon SQS rejects it",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "message_retention_period",
				Description: "The length of time, in seconds, for which Amazon SQS retains a message",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "policy",
				Description: "The policy of the queue",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "arn",
				Description: "The Amazon resource name (ARN) of the queue",
				Type:        schema.TypeString,
			},
			{
				Name:        "receive_message_wait_time_seconds",
				Description: "The length of time, in seconds, for which the ReceiveMessage action waits for a message to arrive",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "redrive_policy",
				Description: "The parameters for the dead-letter queue functionality of the source queue as a JSON object",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "visibility_timeout",
				Description: "The visibility timeout for the queue",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "kms_master_key_id",
				Description: "The ID of an Amazon Web Services managed customer master key (CMK) for Amazon SQS or a custom CMK",
				Type:        schema.TypeString,
			},
			{
				Name:        "kms_data_key_reuse_period_seconds",
				Description: "The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling KMS again",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "sqs_managed_sse_enabled",
				Description: "True if the queue is using SSE-SQS encryption using SQS owned encryption keys",
				Type:        schema.TypeBool,
			},
			{
				Name:        "fifo_queue",
				Description: "True if the queue is FIFO queue",
				Type:        schema.TypeBool,
			},
			{
				Name:        "content_based_deduplication",
				Description: "True if content-based deduplication is enabled for the queue",
				Type:        schema.TypeBool,
			},
			{
				Name:        "deduplication_scope",
				Description: "Specifies whether message deduplication occurs at the message group or queue level",
				Type:        schema.TypeString,
			},
			{
				Name:        "fifo_throughput_limit",
				Description: "Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group",
				Type:        schema.TypeString,
			},
			{
				Name:        "redrive_allow_policy",
				Description: "The parameters for the dead-letter queue functionality of the source queue as a JSON object",
				Type:        schema.TypeJSON,
			},
			{
				Name: "unknown_fields",
				Type: schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchSqsQueues(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().SQS
	var params sqs.ListQueuesInput
	for {
		result, err := svc.ListQueues(ctx, &params)
		if err != nil {
			return diag.WrapError(err)
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
				return diag.WrapError(err)
			}

			var q Queue
			d, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &q})
			if err != nil {
				return diag.WrapError(err)
			}
			if err := d.Decode(out.Attributes); err != nil {
				return diag.WrapError(err)
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
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, result.Tags))
}
