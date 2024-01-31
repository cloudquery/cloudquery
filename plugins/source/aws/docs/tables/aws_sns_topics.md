# Table: aws_sns_topics

This table shows data for Sns Topics.

https://docs.aws.amazon.com/sns/latest/api/API_GetTopicAttributes.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|delivery_policy|`json`|
|policy|`json`|
|effective_delivery_policy|`json`|
|display_name|`utf8`|
|owner|`utf8`|
|subscriptions_confirmed|`int64`|
|subscriptions_deleted|`int64`|
|subscriptions_pending|`int64`|
|kms_master_key_id|`utf8`|
|fifo_topic|`bool`|
|content_based_deduplication|`bool`|
|unknown_fields|`json`|