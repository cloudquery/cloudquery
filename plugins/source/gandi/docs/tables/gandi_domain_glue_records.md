# Table: gandi_domain_glue_records



The composite primary key for this table is (**domain_fqdn**, **name**).

## Relations
This table depends on [gandi_domains](gandi_domains.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|domain_fqdn (PK)|String|
|name (PK)|String|
|ips|StringArray|
|fqdn|String|
|href|String|
|fqdn_unicode|String|