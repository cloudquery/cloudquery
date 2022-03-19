package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/mitchellh/mapstructure"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SQSQueues() *schema.Table {
	return &schema.Table{
		Name:         "aws_sqs_queues",
		Description:  "Simple Queue Service",
		Resolver:     fetchSQSQueues,
		Multiplex:    client.ServiceAccountRegionMultiplexer("sqs"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
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
				Name:        "url",
				Description: "Queue URL",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("URL"),
			},
			{
				Name:          "policy",
				Description:   "The queue's policy. A valid Amazon Web Services policy.",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:        "visibility_timeout",
				Description: "The visibility timeout for the queue, in seconds.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "maximum_message_size",
				Description: "The limit of how many bytes a message can contain before Amazon SQS rejects it.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "message_retention_period",
				Description: "The length of time, in seconds, for which Amazon SQS retains a message.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "approximate_number_of_messages",
				Description: "The approximate number of messages available for retrieval from the queue.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "approximate_number_of_messages_not_visible",
				Description: "The approximate number of messages that are in flight.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "created_timestamp",
				Description: "UNIX time when the queue was created.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "last_modified_timestamp",
				Description: "UNIX time when the queue was last changed.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "arn",
				Description: "Amazon resource name (ARN) of the queue.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("QueueArn"),
			},
			{
				Name:        "approximate_number_of_messages_delayed",
				Description: "The approximate number of messages in the queue that are delayed and not available for reading immediately.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "delay_seconds",
				Description: "The default delay on the queue in seconds.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "receive_message_wait_time_seconds",
				Description: "the length of time, in seconds, for which the ReceiveMessage action waits for a message to arrive.",
				Type:        schema.TypeInt,
			},
			{
				Name:          "redrive_policy",
				Description:   "The parameters for the dead-letter queue functionality of the source queue as a JSON object.",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:          "fifo_queue",
				Description:   "True if the queue is FIFO queue.",
				Type:          schema.TypeBool,
				IgnoreInTests: true,
			},
			{
				Name:          "content_based_deduplication",
				Description:   "True if content-based deduplication is enabled for the queue.",
				Type:          schema.TypeBool,
				IgnoreInTests: true,
			},
			{
				Name:          "kms_master_key_id",
				Description:   "ID of an Amazon Web Services managed customer master key (CMK) for Amazon SQS or a custom CMK.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "kms_data_key_reuse_period_seconds",
				Description:   "The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling KMS again.",
				Type:          schema.TypeInt,
				IgnoreInTests: true,
			},
			{
				Name:          "deduplication_scope",
				Description:   "Specifies whether message deduplication occurs at the message group or queue level.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "fifo_throughput_limit",
				Description:   "Specifies whether message deduplication occurs at the message group or queue level.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "redrive_allow_policy",
				Description:   "The parameters for the permissions for the dead-letter queue redrive permission.",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:          "tags",
				Description:   "Queue tags.",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:        "unknown_fields",
				Description: "Other queue attributes",
				Type:        schema.TypeJSON,
			},
		},
	}
}

type sqsQueue struct {
	URL                                   string
	Policy                                *string
	VisibilityTimeout                     *int32
	MaximumMessageSize                    *int32
	MessageRetentionPeriod                *int32
	ApproximateNumberOfMessages           *int32
	ApproximateNumberOfMessagesNotVisible *int32
	CreatedTimestamp                      *int32
	LastModifiedTimestamp                 *int32
	QueueArn                              *string
	ApproximateNumberOfMessagesDelayed    *int32
	DelaySeconds                          *int32
	ReceiveMessageWaitTimeSeconds         *int32
	RedrivePolicy                         *string
	FifoQueue                             *bool
	ContentBasedDeduplication             *bool
	KmsMasterKeyId                        *string
	KmsDataKeyReusePeriodSeconds          *int32
	DeduplicationScope                    *string
	FifoThroughputLimit                   *string
	RedriveAllowPolicy                    *string

	UnknownFields map[string]interface{} `mapstructure:",remain"`

	Tags map[string]string
}

func fetchSQSQueues(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client := meta.(*client.Client)
	sqsClient := client.Services().SQS
	optsFn := func(o *sqs.Options) {
		o.Region = client.Region
	}
	var params sqs.ListQueuesInput
	for {
		result, err := sqsClient.ListQueues(ctx, &params, optsFn)
		if err != nil {
			return diag.WrapError(err)
		}

		for _, url := range result.QueueUrls {
			input := sqs.GetQueueAttributesInput{
				QueueUrl:       &url,
				AttributeNames: []types.QueueAttributeName{types.QueueAttributeNameAll},
			}
			out, err := sqsClient.GetQueueAttributes(ctx, &input, optsFn)
			if err != nil {
				return diag.WrapError(err)
			}

			var q sqsQueue
			d, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &q})
			if err != nil {
				return diag.WrapError(err)
			}
			if err := d.Decode(out.Attributes); err != nil {
				return err
			}
			q.URL = url

			tagsOut, err := sqsClient.ListQueueTags(ctx, &sqs.ListQueueTagsInput{QueueUrl: &url}, optsFn)
			if err != nil {
				return diag.WrapError(err)
			}
			q.Tags = tagsOut.Tags
			res <- q
		}
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return nil
}
