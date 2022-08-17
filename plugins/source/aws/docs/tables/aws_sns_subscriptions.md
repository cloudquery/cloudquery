
# Table: aws_sns_subscriptions
Amazon SNS subscription
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|endpoint|text|The subscription's endpoint (format depends on the protocol)|
|owner|text|The subscription's owner|
|protocol|text|The subscription's protocol|
|arn|text|The subscription's ARN|
|topic_arn|text|The ARN of the subscription's topic|
|confirmation_was_authenticated|boolean|True if the subscription confirmation request was authenticated|
|delivery_policy|jsonb|The JSON serialization of the subscription's delivery policy|
|effective_delivery_policy|jsonb|The JSON serialization of the effective delivery policy that takes into account the topic delivery policy and account system defaults|
|filter_policy|jsonb|The filter policy JSON that is assigned to the subscription|
|pending_confirmation|boolean|True if the subscription hasn't been confirmed|
|raw_message_delivery|boolean|True if raw message delivery is enabled for the subscription|
|redrive_policy|text|When specified, sends undeliverable messages to the specified Amazon SQS dead-letter queue|
|subscription_role_arn|text|The ARN of the IAM role that has permission to write to the Kinesis Data Firehose delivery stream and has Amazon SNS listed as a trusted entity|
|unknown_fields|jsonb|Other subscription attributes|
