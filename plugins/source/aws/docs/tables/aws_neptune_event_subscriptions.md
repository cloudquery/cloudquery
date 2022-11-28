# Table: aws_neptune_event_subscriptions

https://docs.aws.amazon.com/neptune/latest/userguide/api-events.html#DescribeEventSubscriptions

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|cust_subscription_id|String|
|customer_aws_id|String|
|enabled|Bool|
|event_categories_list|StringArray|
|sns_topic_arn|String|
|source_ids_list|StringArray|
|source_type|String|
|status|String|
|subscription_creation_time|String|