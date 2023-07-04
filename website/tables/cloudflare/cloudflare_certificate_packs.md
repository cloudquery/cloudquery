# Table: cloudflare_certificate_packs

This table shows data for Cloudflare Certificate Packs.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|zone_id|`utf8`|
|id (PK)|`utf8`|
|type|`utf8`|
|hosts|`list<item: utf8, nullable>`|
|certificates|`json`|
|primary_certificate|`utf8`|
|validation_records|`json`|
|validation_errors|`json`|
|validation_method|`utf8`|
|validity_days|`int64`|
|certificate_authority|`utf8`|
|cloudflare_branding|`bool`|