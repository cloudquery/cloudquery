# Table: azure_security_secure_score_control_definitions

This table shows data for Azure Security Secure Score Control Definitions.

https://learn.microsoft.com/en-us/rest/api/defenderforcloud/secure-score-control-definitions/list?tabs=HTTP#securescorecontroldefinitionitem

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
|type|`utf8`|