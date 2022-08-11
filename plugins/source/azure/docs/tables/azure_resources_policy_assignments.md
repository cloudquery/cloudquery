
# Table: azure_resources_policy_assignments
Azure network watcher
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|display_name|text|The display name of the policy assignment|
|policy_definition_id|text|The ID of the policy definition or policy set definition being assigned|
|scope|text|The scope for the policy assignment|
|not_scopes|text[]|The policy's excluded scopes|
|parameters|jsonb|The parameter values for the assigned policy rule|
|description|text|This message will be part of response in case of policy violation|
|metadata|jsonb|The policy assignment metadata|
|enforcement_mode|text|The policy assignment enforcement mode|
|id|text|The ID of the policy assignment|
|type|text|The type of the policy assignment|
|name|text|The name of the policy assignment|
|sku_name|text|The name of the policy sku|
|sku_tier|text|The policy sku tier|
|location|text|The location of the policy assignment|
|identity_principal_id|text|The principal ID of the resource identity|
|identity_tenant_id|text|The tenant ID of the resource identity|
|identity_type|text|The identity type|
