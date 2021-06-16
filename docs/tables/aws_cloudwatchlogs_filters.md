
# Table: aws_cloudwatchlogs_filters
Metric filters express how CloudWatch Logs would extract metric observations from ingested log events and transform them into metric data in a CloudWatch metric.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|creation_time|bigint|The creation time of the metric filter, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.|
|name|text|The name of the metric filter.|
|pattern|text|A symbolic description of how CloudWatch Logs should interpret the data in each log event.|
|log_group_name|text|The name of the log group.|
