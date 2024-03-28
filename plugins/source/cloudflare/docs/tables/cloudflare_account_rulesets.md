# Table: cloudflare_account_rulesets

This table shows data for Cloudflare Account Rulesets.

https://developers.cloudflare.com/api/operations/listAccountRulesets

The composite primary key for this table is (**account_id**, **ruleset_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|ruleset_id (PK)|`utf8`|
|id|`utf8`|
|name|`utf8`|
|description|`utf8`|
|kind|`utf8`|
|version|`utf8`|
|last_updated|`timestamp[us, tz=UTC]`|
|phase|`utf8`|
|rules|`json`|
|shareable_entitlement_name|`utf8`|