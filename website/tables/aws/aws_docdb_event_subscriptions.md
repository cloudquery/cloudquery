# Table: aws_docdb_event_subscriptions

This table shows data for Amazon DocumentDB Event Subscriptions.

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_EventSubscription.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
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