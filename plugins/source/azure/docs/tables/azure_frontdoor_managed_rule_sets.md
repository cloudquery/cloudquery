# Table: azure_frontdoor_managed_rule_sets

This table shows data for Azure Frontdoor Managed Rule Sets.

https://learn.microsoft.com/en-us/rest/api/frontdoor/webapplicationfirewall/managed-rule-sets/list#managedrulesetdefinition

The composite primary key for this table is (**subscription_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id (PK)|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|