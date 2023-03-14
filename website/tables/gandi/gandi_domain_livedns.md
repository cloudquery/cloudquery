# Table: gandi_domain_livedns

This table shows data for Gandi Domain Livedns.

The primary key for this table is **fqdn**.

## Relations

This table depends on [gandi_domains](gandi_domains).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|fqdn (PK)|String|
|current|String|
|nameservers|StringArray|
|dnssec_available|Bool|
|livednssec_available|Bool|