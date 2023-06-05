# Table: gandi_livedns_domain_snapshots

This table shows data for Gandi LiveDNS Domain Snapshots.

The composite primary key for this table is (**fqdn**, **id**).

## Relations

This table depends on [gandi_livedns_domains](gandi_livedns_domains).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|fqdn (PK)|`utf8`|
|automatic|`bool`|
|created_at|`timestamp[us, tz=UTC]`|
|id (PK)|`utf8`|
|name|`utf8`|
|snapshot_href|`utf8`|
|zone_data|`json`|