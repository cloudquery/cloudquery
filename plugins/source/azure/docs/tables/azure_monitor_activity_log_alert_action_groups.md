
# Table: azure_monitor_activity_log_alert_action_groups
ActivityLogAlertActionGroup a pointer to an Azure Action Group
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|activity_log_alert_cq_id|uuid|Unique ID of azure_monitor_activity_log_alerts table (FK)|
|activity_log_alert_id|text|ID of azure_monitor_activity_log_alerts table (FK)|
|action_group_id|text|The resourceId of the action group This cannot be null or empty|
|webhook_properties|jsonb|the dictionary of custom properties to include with the post operation These data are appended to the webhook payload|
