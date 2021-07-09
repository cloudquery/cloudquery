
# Table: azure_monitor_log_profiles
LogProfileResource the log profile resource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|storage_account_id|text|the resource id of the storage account to which you would like to send the Activity Log|
|service_bus_rule_id|text|The service bus rule ID of the service bus namespace in which you would like to have Event Hubs created for streaming the Activity Log The rule ID is of the format: '{service bus resource ID}/authorizationrules/{key name}'|
|locations|text[]|List of regions for which Activity Log events should be stored or streamed It is a comma separated list of valid ARM locations including the 'global' location|
|categories|text[]|the categories of the logs These categories are created as is convenient to the user Some values are: 'Write', 'Delete', and/or 'Action'|
|retention_policy_enabled|boolean|a value indicating whether the retention policy is enabled|
|retention_policy_days|integer|the number of days for the retention in days A value of 0 will retain the events indefinitely|
|id|text|Azure resource Id|
|name|text|Azure resource name|
|type|text|Azure resource type|
|location|text|Resource location|
|tags|jsonb|Resource tags|
