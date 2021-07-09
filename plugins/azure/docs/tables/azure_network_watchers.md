
# Table: azure_network_watchers
Azure network watcher
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|etag|text|A unique read-only string that changes whenever the resource is updated|
|provisioning_state|text|The provisioning state of the network watcher resource Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
|location|text|Resource location|
|tags|jsonb|Resource tags|
|flow_log_storage_id|text|ID of the storage account which is used to store the flow log|
|flow_log_enabled|boolean|Flag to enable/disable flow logging|
|flow_log_retention_policy_days|integer|Number of days to retain flow log records|
|flow_log_retention_policy_enabled|boolean|Flag to enable/disable retention|
|flow_log_format_type|text|The file type of flow log Possible values include: 'JSON'|
|flow_log_format_version|integer|The version (revision) of the flow log|
