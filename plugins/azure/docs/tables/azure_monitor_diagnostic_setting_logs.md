
# Table: azure_monitor_diagnostic_setting_logs
LogSettings part of MultiTenantDiagnosticSettings Specifies the settings for a particular log
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|diagnostic_setting_cq_id|uuid|Unique ID of azure_monitor_diagnostic_settings table (FK)|
|diagnostic_setting_id|text|Unique ID of azure_monitor_diagnostic_settings table (FK)|
|category|text|Name of a Diagnostic Log category for a resource type this setting is applied to To obtain the list of Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation|
|enabled|boolean|a value indicating whether this log is enabled|
|retention_policy_enabled|boolean|a value indicating whether the retention policy is enabled|
|retention_policy_days|integer|the number of days for the retention in days A value of 0 will retain the events indefinitely|
