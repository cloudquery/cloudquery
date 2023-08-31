# Table: cloudflare_waf_rules

This table shows data for Cloudflare WAF Rules.

The primary key for this table is **_cq_id**.

## Relations

This table depends on [cloudflare_waf_packages](cloudflare_waf_packages).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|waf_package_id|`utf8`|
|id|`utf8`|
|description|`utf8`|
|priority|`utf8`|
|package_id|`utf8`|
|group|`json`|
|mode|`utf8`|
|default_mode|`utf8`|
|allowed_modes|`list<item: utf8, nullable>`|