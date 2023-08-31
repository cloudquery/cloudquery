# Table: cloudflare_waf_packages

This table shows data for Cloudflare WAF Packages.

The primary key for this table is **id**.

## Relations

The following tables depend on cloudflare_waf_packages:
  - [cloudflare_waf_groups](cloudflare_waf_groups)
  - [cloudflare_waf_rules](cloudflare_waf_rules)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|description|`utf8`|
|zone_id|`utf8`|
|detection_mode|`utf8`|
|sensitivity|`utf8`|
|action_mode|`utf8`|