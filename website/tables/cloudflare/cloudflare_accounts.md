# Table: cloudflare_accounts

This table shows data for Cloudflare Accounts.

The primary key for this table is **id**.

## Relations

The following tables depend on cloudflare_accounts:
  - [cloudflare_account_members](cloudflare_account_members)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|
|created_on|`timestamp[us, tz=UTC]`|
|settings|`json`|