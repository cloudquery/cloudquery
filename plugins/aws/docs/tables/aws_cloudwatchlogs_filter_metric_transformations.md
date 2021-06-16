
# Table: aws_cloudwatchlogs_filter_metric_transformations
Indicates how to transform ingested log events to metric data in a CloudWatch metric.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|filter_id|uuid|Unique ID of aws_cloudwatchlogs_filters table (FK)|
|metric_name|text|The name of the CloudWatch metric.|
|metric_namespace|text|A custom namespace to contain your metric in CloudWatch.|
|metric_value|text|The value to publish to the CloudWatch metric when a filter pattern matches a log event.|
|default_value|float|(Optional) The value to emit when a filter pattern does not match a log event.|
