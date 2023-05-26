# Table: gandi_domain_glue_records

This table shows data for Gandi Domain Glue Records.

The composite primary key for this table is (**name**, **fqdn**).

## Relations

This table depends on [gandi_domains](gandi_domains).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|name (PK)|utf8|
|fqdn (PK)|utf8|
|ips|list<item: utf8, nullable>|
|href|utf8|
|fqdn_unicode|utf8|