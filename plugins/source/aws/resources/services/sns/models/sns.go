package models

// Amazon SNS subscription.
type Subscription struct {
	// The subscription's endpoint (format depends on the protocol).
	Endpoint *string
	// The subscription's owner.
	Owner *string
	// The subscription's protocol.
	Protocol *string
	// The subscription's ARN.
	SubscriptionArn *string
	// The ARN of the subscription's topic.
	TopicArn *string
	// True if the subscription confirmation request was authenticated.
	ConfirmationWasAuthenticated *bool
	// The JSON serialization of the subscription's delivery policy.
	DeliveryPolicy *string
	// The JSON serialization of the effective delivery policy that takes into account the topic delivery policy and account system defaults.
	EffectiveDeliveryPolicy *string
	// The filter policy JSON that is assigned to the subscription.
	FilterPolicy *string
	// True if the subscription hasn't been confirmed.
	PendingConfirmation *bool
	// True if raw message delivery is enabled for the subscription.
	RawMessageDelivery *bool
	// When specified, sends undeliverable messages to the specified Amazon SQS dead-letter queue.
	RedrivePolicy *string
	// The ARN of the IAM role that has permission to write to the Kinesis Data Firehose delivery stream and has Amazon SNS listed as a trusted entity.
	SubscriptionRoleArn *string
	// Other subscription attributes.
	UnknownFields map[string]interface{} `mapstructure:",remain"`
}

// Amazon SNS topic.
type Topic struct {
	// The JSON serialization of the topic's delivery policy.
	DeliveryPolicy *string
	// The human-readable name used in the From field for notifications to email and email-json endpoints.
	DisplayName *string
	// The AWS account ID of the topic's owner.
	Owner *string
	// The JSON serialization of the topic's access control policy.
	Policy *string
	// The number of confirmed subscriptions for the topic.
	SubscriptionsConfirmed *int
	// The number of deleted subscriptions for the topic.
	SubscriptionsDeleted *int
	// The number of subscriptions pending confirmation for the topic.
	SubscriptionsPending *int
	// The Amazon Resource Name (ARN) of the topic.
	Arn *string `mapstructure:"TopicArn"`
	// The JSON serialization of the effective delivery policy, taking system defaults into account.
	EffectiveDeliveryPolicy *string
	// The ID of an Amazon Web Services managed customer master key (CMK) for Amazon SNS or a custom CMK.
	KmsMasterKeyId *string
	// When this is set to true, a FIFO topic is created.
	FifoTopic *bool
	// Enables content-based deduplication for FIFO topics.
	ContentBasedDeduplication *bool
	// Other subscription attributes.
	UnknownFields map[string]interface{} `mapstructure:",remain"`
}
