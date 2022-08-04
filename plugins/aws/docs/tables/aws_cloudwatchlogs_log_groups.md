
# Table: aws_cloudwatchlogs_log_groups
Represents a log group.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb|The tags for the log group.|
|arn|text|The Amazon Resource Name (ARN) of the log group.|
|creation_time|bigint|The creation time of the log group, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.|
|kms_key_id|text|The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.|
|log_group_name|text|The name of the log group.|
|metric_filter_count|bigint|The number of metric filters.|
|retention_in_days|bigint|The number of days to retain the log events in the specified log group|
|stored_bytes|bigint|The number of bytes stored.|
