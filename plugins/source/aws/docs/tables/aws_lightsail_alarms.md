# Table: aws_lightsail_alarms

This table shows data for Lightsail Alarms.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Alarm.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|comparison_operator|`utf8`|
|contact_protocols|`list<item: utf8, nullable>`|
|created_at|`timestamp[us, tz=UTC]`|
|datapoints_to_alarm|`int64`|
|evaluation_periods|`int64`|
|location|`json`|
|metric_name|`utf8`|
|monitored_resource_info|`json`|
|name|`utf8`|
|notification_enabled|`bool`|
|notification_triggers|`list<item: utf8, nullable>`|
|period|`int64`|
|resource_type|`utf8`|
|state|`utf8`|
|statistic|`utf8`|
|support_code|`utf8`|
|threshold|`float64`|
|treat_missing_data|`utf8`|
|unit|`utf8`|