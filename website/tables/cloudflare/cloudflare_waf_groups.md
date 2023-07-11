# Table: cloudflare_waf_groups

This table shows data for Cloudflare WAF Groups.

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
|name|`utf8`|
|description|`utf8`|
|rules_count|`int64`|
|modified_rules_count|`int64`|
|package_id|`utf8`|
|mode|`utf8`|
|allowed_modes|`list<item: utf8, nullable>`|