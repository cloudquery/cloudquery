# Table: azure_policy_exemptions

This table shows data for Azure Policy Exemptions.

https://learn.microsoft.com/en-us/rest/api/policy/policy-exemptions/list?tabs=HTTP#policyexemption

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|subscription_id|utf8|
|properties|json|
|id (PK)|utf8|
|name|utf8|
|system_data|json|
|type|utf8|