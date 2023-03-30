# Table: aws_elbv2_listener_rules

This table shows data for Amazon Elastic Load Balancer (ELB) v2 Listener Rules.

https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_Rule.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_elbv2_listeners](aws_elbv2_listeners).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|listener_arn|String|
|arn (PK)|String|
|actions|JSON|
|conditions|JSON|
|is_default|Bool|
|priority|String|
|rule_arn|String|