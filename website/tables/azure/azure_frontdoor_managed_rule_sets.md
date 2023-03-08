# Table: azure_frontdoor_managed_rule_sets

https://learn.microsoft.com/en-us/rest/api/frontdoor/webapplicationfirewall/managed-rule-sets/list#managedrulesetdefinition

The composite primary key for this table is (**subscription_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id (PK)|String|
|location|String|
|properties|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|