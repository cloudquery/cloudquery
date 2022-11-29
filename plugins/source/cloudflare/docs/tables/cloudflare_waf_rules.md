# Table: cloudflare_waf_rules



The primary key for this table is **_cq_id**.

## Relations
This table depends on [cloudflare_waf_packages](cloudflare_waf_packages.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|waf_package_id|String|
|id|String|
|description|String|
|priority|String|
|package_id|String|
|group|JSON|
|mode|String|
|default_mode|String|
|allowed_modes|StringArray|