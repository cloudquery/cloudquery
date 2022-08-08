
# Table: azure_monitor_activity_log_alert_conditions
ActivityLogAlertLeafCondition an Activity Log alert condition that is met by comparing an activity log field and value
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|activity_log_alert_cq_id|uuid|Unique ID of azure_monitor_activity_log_alerts table (FK)|
|activity_log_alert_id|text|ID of azure_monitor_activity_log_alerts table (FK)|
|field|text|The name of the field that this condition will examine The possible values for this field are (case-insensitive): 'resourceId', 'category', 'caller', 'level', 'operationName', 'resourceGroup', 'resourceProvider', 'status', 'subStatus', 'resourceType', or anything beginning with 'properties'|
|equals|text|The field value will be compared to this value (case-insensitive) to determine if the condition is met|
