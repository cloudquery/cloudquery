
# Table: azure_postgresql_server_configurations
Azure postgresql server configuration
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|server_id|uuid|Unique ID of azure_postgresql_servers table (FK)|
|value|text|Value of the configuration|
|description|text|Description of the configuration|
|default_value|text|Default value of the configuration|
|data_type|text|Data type of the configuration|
|allowed_values|text|Allowed values of the configuration|
|source|text|Source of the configuration|
|resource_id|text|Fully qualified resource ID for the resource Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}|
|name|text|The name of the resource|
|type|text|The type of the resource Eg "MicrosoftCompute/virtualMachines" or "MicrosoftStorage/storageAccounts"|
