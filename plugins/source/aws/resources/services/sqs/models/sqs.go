package models

// Amazon Simple Queue Service.
type Queue struct {
	// The URL of the Amazon SQS queue.
	URL string
	// The approximate number of messages available for retrieval from the queue.
	ApproximateNumberOfMessages *int32
	// The approximate number of messages in the queue that are delayed and not available for reading immediately.
	ApproximateNumberOfMessagesDelayed *int32
	// The approximate number of messages that are in flight.
	ApproximateNumberOfMessagesNotVisible *int32
	// The time when the queue was created in seconds (epoch time).
	CreatedTimestamp *int32
	// The default delay on the queue in seconds.
	DelaySeconds *int32
	// The time when the queue was last changed in seconds (epoch time).
	LastModifiedTimestamp *int32
	// The limit of how many bytes a message can contain before Amazon SQS rejects it.
	MaximumMessageSize *int32
	// The length of time, in seconds, for which Amazon SQS retains a message.
	MessageRetentionPeriod *int32
	// The policy of the queue.
	Policy *string
	// The Amazon resource name (ARN) of the queue.
	Arn *string `mapstructure:"QueueArn"`
	// The length of time, in seconds, for which the ReceiveMessage action waits for a message to arrive.
	ReceiveMessageWaitTimeSeconds *int32
	// The parameters for the dead-letter queue functionality of the source queue as a JSON object.
	RedrivePolicy *string
	// The visibility timeout for the queue.
	VisibilityTimeout *int32
	// The ID of an Amazon Web Services managed customer master key (CMK) for Amazon SQS or a custom CMK.
	KmsMasterKeyId *string
	// The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling KMS again.
	KmsDataKeyReusePeriodSeconds *int32
	// True if the queue is using SSE-SQS encryption using SQS owned encryption keys.
	SqsManagedSseEnabled *bool
	// True if the queue is FIFO queue.
	FifoQueue *bool
	// True if content-based deduplication is enabled for the queue.
	ContentBasedDeduplication *bool
	// Specifies whether message deduplication occurs at the message group or queue level.
	DeduplicationScope *string
	// Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group.
	FifoThroughputLimit *string
	// The parameters for the dead-letter queue functionality of the source queue as a JSON object.
	RedriveAllowPolicy *string

	UnknownFields map[string]any `mapstructure:",remain"`
}
