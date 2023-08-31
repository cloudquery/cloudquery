# Table: aws_iot_topic_rules

This table shows data for AWS IoT Topic Rules.

https://docs.aws.amazon.com/iot/latest/apireference/API_GetTopicRule.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|rule|`json`|
|rule_arn|`utf8`|
|result_metadata|`json`|