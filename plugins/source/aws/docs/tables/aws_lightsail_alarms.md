# Table: aws_lightsail_alarms

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Alarm.html

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
|comparison_operator|String|
|contact_protocols|StringArray|
|created_at|Timestamp|
|datapoints_to_alarm|Int|
|evaluation_periods|Int|
|location|JSON|
|metric_name|String|
|monitored_resource_info|JSON|
|name|String|
|notification_enabled|Bool|
|notification_triggers|StringArray|
|period|Int|
|resource_type|String|
|state|String|
|statistic|String|
|support_code|String|
|threshold|Float|
|treat_missing_data|String|
|unit|String|