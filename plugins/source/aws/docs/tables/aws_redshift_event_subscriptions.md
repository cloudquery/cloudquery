
# Table: aws_redshift_event_subscriptions
Describes event subscriptions.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|arn|text|ARN of the event subscription.|
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|id|text|The name of the Amazon Redshift event notification subscription.|
|customer_aws_id|text|The AWS customer account associated with the Amazon Redshift event notification subscription.|
|enabled|boolean|A boolean value indicating whether the subscription is enabled; true indicates that the subscription is enabled.|
|event_categories_list|text[]|The list of Amazon Redshift event categories specified in the event notification subscription|
|severity|text|The event severity specified in the Amazon Redshift event notification subscription|
|sns_topic_arn|text|The Amazon Resource Name (ARN) of the Amazon SNS topic used by the event notification subscription.|
|source_ids_list|text[]|A list of the sources that publish events to the Amazon Redshift event notification subscription.|
|source_type|text|The source type of the events returned by the Amazon Redshift event notification.|
|status|text|The status of the Amazon Redshift event notification subscription.|
|subscription_creation_time|timestamp without time zone|The date and time the Amazon Redshift event notification subscription was created.|
|tags|jsonb|Tags|
