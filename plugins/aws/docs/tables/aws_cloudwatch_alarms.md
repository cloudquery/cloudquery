
# Table: aws_cloudwatch_alarms
The details about a metric alarm.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|actions_enabled|boolean|Indicates whether actions should be executed during any changes to the alarm state.|
|alarm_actions|text[]|The actions to execute when this alarm transitions to the ALARM state from any other state.|
|alarm_arn|text|The Amazon Resource Name (ARN) of the alarm.|
|alarm_configuration_updated_timestamp|timestamp without time zone|The time stamp of the last update to the alarm configuration.|
|alarm_description|text|The description of the alarm.|
|alarm_name|text|The name of the alarm.|
|comparison_operator|text|The arithmetic operation to use when comparing the specified statistic and threshold.|
|datapoints_to_alarm|integer|The number of data points that must be breaching to trigger the alarm.|
|dimensions|jsonb|The dimensions for the metric associated with the alarm.|
|evaluate_low_sample_count_percentile|text|Used only for alarms based on percentiles.|
|evaluation_periods|integer|The number of periods over which data is compared to the specified threshold.|
|extended_statistic|text|The percentile statistic for the metric associated with the alarm.|
|insufficient_data_actions|text[]|The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state.|
|metric_name|text|The name of the metric associated with the alarm, if this is an alarm based on a single metric.|
|namespace|text|The namespace of the metric associated with the alarm.|
|ok_actions|text[]|The actions to execute when this alarm transitions to the OK state from any other state.|
|period|integer|The period, in seconds, over which the statistic is applied.|
|state_reason|text|An explanation for the alarm state, in text format.|
|state_reason_data|text|An explanation for the alarm state, in JSON format.|
|state_updated_timestamp|timestamp without time zone|The time stamp of the last update to the alarm state.|
|state_value|text|The state value for the alarm.|
|statistic|text|The statistic for the metric associated with the alarm, other than percentile.|
|threshold|float|The value to compare with the specified statistic.|
|threshold_metric_id|text|In an alarm based on an anomaly detection model, this is the ID of the ANOMALY_DETECTION_BAND function used as the threshold for the alarm.|
|treat_missing_data|text|Sets how this alarm is to handle missing data points.|
|unit|text|The unit of the metric associated with the alarm.|
