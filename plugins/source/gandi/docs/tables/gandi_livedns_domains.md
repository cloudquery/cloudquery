# Table: gandi_livedns_domains



The primary key for this table is **fqdn**.

## Relations

The following tables depend on gandi_livedns_domains:
  - [gandi_livedns_domain_snapshots](gandi_livedns_domain_snapshots.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|sharing_id|String|
|fqdn (PK)|String|
|domain_href|String|
|domain_keys_href|String|
|domain_records_href|String|
|automatic_snapshots|Bool|