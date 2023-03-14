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
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|id (PK)|String|
|name|String|
|description|String|
|zone_id|String|
|detection_mode|String|
|sensitivity|String|
|action_mode|String|