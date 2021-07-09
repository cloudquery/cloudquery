
# Table: azure_network_security_group_flow_logs
FlowLog a flow log resource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|security_group_cq_id|uuid|Unique ID of azure_network_security_groups table (FK)|
|security_group_id|uuid|ID of azure_network_security_groups table (FK)|
|target_resource_id|text|ID of network security group to which flow log will be applied|
|target_resource_guid|text|Guid of network security group to which flow log will be applied|
|storage_id|text|ID of the storage account which is used to store the flow log|
|enabled|boolean|Flag to enable/disable flow logging|
|retention_policy_days|integer|Number of days to retain flow log records|
|retention_policy_enabled|boolean|Flag to enable/disable retention|
|format_type|text|The file type of flow log Possible values include: 'JSON'|
|format_version|integer|The version (revision) of the flow log|
|flow_analytics_configuration_enabled|boolean|Flag to enable/disable traffic analytics for network watcher|
|flow_analytics_configuration_workspace_id|text|The resource guid of the attached workspace for network watcher|
|flow_analytics_configuration_workspace_region|text|The location of the attached workspace for network watcher|
|flow_analytics_configuration_workspace_resource_id|text|Resource Id of the attached workspace for network watcher|
|flow_analytics_configuration_traffic_analytics_interval|integer|The interval in minutes which would decide how frequently TA service should do flow analytics for network watcher|
|provisioning_state|text|The provisioning state of the flow log Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'|
|etag|text|A unique read-only string that changes whenever the resource is updated|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
|location|text|Resource location|
|tags|jsonb|Resource tags|
