# Table: aws_elbv2_target_group_target_health_descriptions

https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_TargetHealthDescription.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_elbv2_target_groups](aws_elbv2_target_groups.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|target_group_arn|String|
|health_check_port|String|
|target|JSON|
|target_health|JSON|