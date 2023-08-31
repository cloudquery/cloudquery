# Table: cloudflare_account_members

This table shows data for Cloudflare Account Members.

The primary key for this table is **_cq_id**.

## Relations

This table depends on [cloudflare_accounts](cloudflare_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|id|`utf8`|
|code|`utf8`|
|user|`json`|
|status|`utf8`|
|roles|`json`|
|policies|`json`|