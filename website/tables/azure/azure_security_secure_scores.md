# Table: azure_security_secure_scores

This table shows data for Azure Security Secure Scores.

https://learn.microsoft.com/en-us/rest/api/defenderforcloud/secure-scores/list?tabs=HTTP#securescoreitem

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|id (PK)|String|
|name|String|
|properties|JSON|
|type|String|