# Table: aws_cloudwatch_alarms

This table shows data for Cloudwatch Alarms.

https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricAlarm.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|dimensions|`json`|
|actions_enabled|`bool`|
|alarm_actions|`list<item: utf8, nullable>`|
|alarm_arn|`utf8`|
|alarm_configuration_updated_timestamp|`timestamp[us, tz=UTC]`|
|alarm_description|`utf8`|
|alarm_name|`utf8`|
|comparison_operator|`utf8`|
|datapoints_to_alarm|`int64`|
|evaluate_low_sample_count_percentile|`utf8`|
|evaluation_periods|`int64`|
|evaluation_state|`utf8`|
|extended_statistic|`utf8`|
|insufficient_data_actions|`list<item: utf8, nullable>`|
|metric_name|`utf8`|
|metrics|`json`|
|namespace|`utf8`|
|ok_actions|`list<item: utf8, nullable>`|
|period|`int64`|
|state_reason|`utf8`|
|state_reason_data|`utf8`|
|state_transitioned_timestamp|`timestamp[us, tz=UTC]`|
|state_updated_timestamp|`timestamp[us, tz=UTC]`|
|state_value|`utf8`|
|statistic|`utf8`|
|threshold|`float64`|
|threshold_metric_id|`utf8`|
|treat_missing_data|`utf8`|
|unit|`utf8`|