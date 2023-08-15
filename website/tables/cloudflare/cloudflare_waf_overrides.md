# Table: cloudflare_waf_overrides

This table shows data for Cloudflare WAF Overrides.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|zone_id|`utf8`|
|id (PK)|`utf8`|
|description|`utf8`|
|urls|`list<item: utf8, nullable>`|
|priority|`int64`|
|groups|`json`|
|rewrite_action|`json`|
|rules|`json`|
|paused|`bool`|