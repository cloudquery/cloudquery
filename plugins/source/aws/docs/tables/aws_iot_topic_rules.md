# Table: aws_iot_topic_rules

https://docs.aws.amazon.com/iot/latest/apireference/API_GetTopicRule.html

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
|tags|JSON|
|arn (PK)|String|
|rule|JSON|
|rule_arn|String|
|result_metadata|JSON|