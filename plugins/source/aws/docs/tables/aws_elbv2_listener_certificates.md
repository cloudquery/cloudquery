# Table: aws_elbv2_listener_certificates

https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_Certificate.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_elbv2_listeners](aws_elbv2_listeners.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|listener_arn|String|
|certificate_arn|String|
|is_default|Bool|