
# Table: aws_lightsail_alarms
Describes an alarm
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the alarm|
|comparison_operator|text|The arithmetic operation used when comparing the specified statistic and threshold|
|contact_protocols|text[]|The contact protocols for the alarm, such as Email, SMS (text messaging), or both|
|created_at|timestamp without time zone|The timestamp when the alarm was created|
|datapoints_to_alarm|integer|The number of data points that must not within the specified threshold to trigger the alarm|
|evaluation_periods|integer|The number of periods over which data is compared to the specified threshold|
|location_availability_zone|text|The Availability Zone|
|location_region_name|text|The AWS Region name|
|metric_name|text|The name of the metric associated with the alarm|
|monitored_resource_info_arn|text|The Amazon Resource Name (ARN) of the resource being monitored|
|monitored_resource_info_name|text|The name of the Lightsail resource being monitored|
|monitored_resource_info_resource_type|text|The Lightsail resource type of the resource being monitored|
|name|text|The name of the alarm|
|notification_enabled|boolean|Indicates whether the alarm is enabled|
|notification_triggers|text[]|The alarm states that trigger a notification|
|period|integer|The period, in seconds, over which the statistic is applied|
|resource_type|text|The Lightsail resource type (eg, Alarm)|
|state|text|The current state of the alarm|
|statistic|text|The statistic for the metric associated with the alarm|
|support_code|text|The support code|
|threshold|float|The value against which the specified statistic is compared|
|treat_missing_data|text|Specifies how the alarm handles missing data points|
|unit|text|The unit of the metric associated with the alarm|
