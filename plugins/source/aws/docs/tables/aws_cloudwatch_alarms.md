# Table: aws_cloudwatch_alarms

https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricAlarm.html

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
|tags|JSON|
|arn (PK)|String|
|dimensions|JSON|
|actions_enabled|Bool|
|alarm_actions|StringArray|
|alarm_configuration_updated_timestamp|Timestamp|
|alarm_description|String|
|alarm_name|String|
|comparison_operator|String|
|datapoints_to_alarm|Int|
|evaluate_low_sample_count_percentile|String|
|evaluation_periods|Int|
|extended_statistic|String|
|insufficient_data_actions|StringArray|
|metric_name|String|
|metrics|JSON|
|namespace|String|
|ok_actions|StringArray|
|period|Int|
|state_reason|String|
|state_reason_data|String|
|state_updated_timestamp|Timestamp|
|state_value|String|
|statistic|String|
|threshold|Float|
|threshold_metric_id|String|
|treat_missing_data|String|
|unit|String|