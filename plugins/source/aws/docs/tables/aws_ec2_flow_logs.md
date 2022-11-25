# Table: aws_ec2_flow_logs

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_FlowLog.html

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
|creation_time|Timestamp|
|deliver_cross_account_role|String|
|deliver_logs_error_message|String|
|deliver_logs_permission_arn|String|
|deliver_logs_status|String|
|destination_options|JSON|
|flow_log_id|String|
|flow_log_status|String|
|log_destination|String|
|log_destination_type|String|
|log_format|String|
|log_group_name|String|
|max_aggregation_interval|Int|
|resource_id|String|
|tags|JSON|
|traffic_type|String|