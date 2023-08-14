# Table: azure_policy_set_definitions

This table shows data for Azure Policy Set Definitions.

https://learn.microsoft.com/en-us/rest/api/policy/policy-set-definitions/list?tabs=HTTP#policysetdefinition

The composite primary key for this table is (**subscription_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id (PK)|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|