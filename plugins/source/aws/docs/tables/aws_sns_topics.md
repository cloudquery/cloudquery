
# Table: aws_sns_topics
Amazon SNS topic
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb||
|delivery_policy|jsonb|The JSON serialization of the topic's delivery policy|
|display_name|text|The human-readable name used in the From field for notifications to email and email-json endpoints|
|owner|text|The AWS account ID of the topic's owner|
|policy|jsonb|The JSON serialization of the topic's access control policy|
|subscriptions_confirmed|bigint|The number of confirmed subscriptions for the topic|
|subscriptions_deleted|bigint|The number of deleted subscriptions for the topic|
|subscriptions_pending|bigint|The number of subscriptions pending confirmation for the topic|
|arn|text|The Amazon Resource Name (ARN) of the topic|
|effective_delivery_policy|jsonb|The JSON serialization of the effective delivery policy, taking system defaults into account|
|kms_master_key_id|text|The ID of an Amazon Web Services managed customer master key (CMK) for Amazon SNS or a custom CMK|
|fifo_topic|boolean|When this is set to true, a FIFO topic is created|
|content_based_deduplication|boolean|Enables content-based deduplication for FIFO topics|
|unknown_fields|jsonb|Other subscription attributes|
