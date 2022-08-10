
# Table: aws_rds_event_subscriptions
Contains the results of a successful invocation of the DescribeEventSubscriptions action.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|cust_subscription_id|text|The RDS event notification subscription Id.|
|customer_aws_id|text|The AWS customer account associated with the RDS event notification subscription.|
|enabled|boolean|A Boolean value indicating if the subscription is enabled|
|event_categories_list|text[]|A list of event categories for the RDS event notification subscription.|
|arn|text|The Amazon Resource Name (ARN) for the event subscription.|
|sns_topic_arn|text|The topic ARN of the RDS event notification subscription.|
|source_ids_list|text[]|A list of source IDs for the RDS event notification subscription.|
|source_type|text|The source type for the RDS event notification subscription.|
|status|text|The status of the RDS event notification subscription|
|subscription_creation_time|text|The time the RDS event notification subscription was created.|
|tags|jsonb|List of tags|
