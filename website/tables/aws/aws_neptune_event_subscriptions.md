# Table: aws_neptune_event_subscriptions

This table shows data for Neptune Event Subscriptions.

https://docs.aws.amazon.com/neptune/latest/userguide/api-events.html#DescribeEventSubscriptions

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|cust_subscription_id|`utf8`|
|customer_aws_id|`utf8`|
|enabled|`bool`|
|event_categories_list|`list<item: utf8, nullable>`|
|event_subscription_arn|`utf8`|
|sns_topic_arn|`utf8`|
|source_ids_list|`list<item: utf8, nullable>`|
|source_type|`utf8`|
|status|`utf8`|
|subscription_creation_time|`utf8`|