# Table: azure_security_secure_scores

This table shows data for Azure Security Secure Scores.

https://learn.microsoft.com/en-us/rest/api/defenderforcloud/secure-scores/list?tabs=HTTP#securescoreitem

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|properties|`json`|
|type|`utf8`|