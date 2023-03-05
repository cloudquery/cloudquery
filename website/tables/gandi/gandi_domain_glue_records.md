# Table: gandi_domain_glue_records

The composite primary key for this table is (**name**, **fqdn**).

## Relations

This table depends on [gandi_domains](gandi_domains).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|name (PK)|String|
|fqdn (PK)|String|
|ips|StringArray|
|href|String|
|fqdn_unicode|String|