# Table: cloudflare_accounts


The primary key for this table is **id**.

## Relations
The following tables depend on `cloudflare_accounts`:
  - [`cloudflare_account_members`](cloudflare_account_members.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|id (PK)|String|
|name|String|
|type|String|
|created_on|Timestamp|
|settings|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|