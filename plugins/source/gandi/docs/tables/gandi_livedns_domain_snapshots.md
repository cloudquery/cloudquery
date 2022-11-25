# Table: gandi_livedns_domain_snapshots



The composite primary key for this table is (**fqdn**, **id**).

## Relations
This table depends on [gandi_livedns_domains](gandi_livedns_domains.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|fqdn (PK)|String|
|automatic|Bool|
|created_at|Timestamp|
|id (PK)|String|
|name|String|
|snapshot_href|String|
|zone_data|JSON|