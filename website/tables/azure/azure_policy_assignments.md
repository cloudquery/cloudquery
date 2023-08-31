# Table: azure_policy_assignments

This table shows data for Azure Policy Assignments.

https://learn.microsoft.com/en-us/rest/api/policy/policy-assignments/list?tabs=HTTP#policyassignment

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|identity|`json`|
|location|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|