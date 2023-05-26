# Table: gandi_livedns_domains

This table shows data for Gandi LiveDNS Domains.

The primary key for this table is **fqdn**.

## Relations

The following tables depend on gandi_livedns_domains:
  - [gandi_livedns_domain_snapshots](gandi_livedns_domain_snapshots)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|sharing_id|utf8|
|fqdn (PK)|utf8|
|domain_href|utf8|
|domain_keys_href|utf8|
|domain_records_href|utf8|
|automatic_snapshots|bool|