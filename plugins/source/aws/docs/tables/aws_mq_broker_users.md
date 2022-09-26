# Table: aws_mq_broker_users


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_mq_brokers`](aws_mq_brokers.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|broker_arn|String|
|broker_id|String|
|console_access|Bool|
|groups|StringArray|
|pending|JSON|
|username|String|
|result_metadata|JSON|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|