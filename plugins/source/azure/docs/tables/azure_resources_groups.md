
# Table: azure_resources_groups
Azure resource group
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|id|text|The ID of the resource group|
|name|text|The name of the resource group|
|type|text|The type of the resource group|
|properties_provisioning_state|text|The provisioning state|
|location|text|The location of the resource group It cannot be changed after the resource group has been created It must be one of the supported Azure locations|
|managed_by|text|The ID of the resource that manages this resource group|
|tags|jsonb|The tags attached to the resource group|
