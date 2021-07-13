
# Table: azure_monitor_activity_log_alerts
ActivityLogAlertResource an activity log alert resource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|scopes|text[]|A list of resourceIds that will be used as prefixes The alert will only apply to activityLogs with resourceIds that fall under one of these prefixes This list must include at least one item|
|enabled|boolean|Indicates whether this activity log alert is enabled If an activity log alert is not enabled, then none of its actions will be activated|
|description|text|A description of this activity log alert|
|id|text|Azure resource Id|
|name|text|Azure resource name|
|type|text|Azure resource type|
|location|text|Resource location|
|tags|jsonb|Resource tags|
