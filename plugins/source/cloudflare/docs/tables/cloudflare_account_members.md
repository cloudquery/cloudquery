# Table: cloudflare_account_members



The primary key for this table is **_cq_id**.

## Relations
This table depends on [cloudflare_accounts](cloudflare_accounts.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|id|String|
|code|String|
|user|JSON|
|status|String|
|roles|JSON|