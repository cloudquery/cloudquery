# Table: azure_security_allowed_connections

This table shows data for Azure Security Allowed Connections.

https://learn.microsoft.com/en-us/rest/api/defenderforcloud/allowed-connections/list?tabs=HTTP#allowedconnectionsresource

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|id (PK)|`utf8`|
|location|`utf8`|
|name|`utf8`|
|properties|`json`|
|type|`utf8`|