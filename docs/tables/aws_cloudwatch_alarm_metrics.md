
# Table: aws_cloudwatch_alarm_metrics
This structure is used in both GetMetricData and PutMetricAlarm.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|alarm_id|uuid|Unique ID of aws_cloudwatch_alarms table (FK)|
|metric_id|text|A short name used to tie this object to the results in the response.|
|expression|text|The math expression to be performed on the returned data, if this object is performing a math expression.|
|label|text|A human-readable label for this metric or expression.|
|metric_stat_metric_dimensions|jsonb|The dimensions for the metric.|
|metric_stat_metric_name|text|The name of the metric.|
|metric_stat_metric_namespace|text|The namespace of the metric.|
|metric_stat_period|integer|The granularity, in seconds, of the returned data points.|
|metric_stat|text|The statistic to return.|
|metric_stat_unit|text|When you are using a Put operation, this defines what unit you want to use when storing the metric.|
|period|integer|The granularity, in seconds, of the returned data points.|
|return_data|boolean|When used in GetMetricData, this option indicates whether to return the timestamps and raw data values of this metric.|
