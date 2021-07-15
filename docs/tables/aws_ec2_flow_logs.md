
# Table: aws_ec2_flow_logs
Describes a flow log.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|id|text|The flow log ID.|
|creation_time|timestamp without time zone|The date and time the flow log was created.|
|deliver_logs_error_message|text|Information about the error that occurred.|
|deliver_logs_permission_arn|text|The ARN of the IAM role that posts logs to CloudWatch Logs.|
|deliver_logs_status|text|The status of the logs delivery (SUCCESS | FAILED).|
|flow_log_id|text|The flow log ID.|
|flow_log_status|text|The status of the flow log (ACTIVE).|
|log_destination|text|Specifies the destination to which the flow log data is published.|
|log_destination_type|text|Specifies the type of destination to which the flow log data is published.|
|log_format|text|The format of the flow log record.|
|log_group_name|text|The name of the flow log group.|
|max_aggregation_interval|integer|The maximum interval of time, in seconds, during which a flow of packets is captured and aggregated into a flow log record.|
|resource_id|text|The ID of the resource on which the flow log was created.|
|tags|jsonb|The tags for the flow log.|
|traffic_type|text|The type of traffic captured for the flow log.|
