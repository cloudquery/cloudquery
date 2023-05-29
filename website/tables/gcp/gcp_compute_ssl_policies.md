# Table: gcp_compute_ssl_policies

This table shows data for GCP Compute SSL Policies.

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|creation_timestamp|`utf8`|
|custom_features|`list<item: utf8, nullable>`|
|description|`utf8`|
|enabled_features|`list<item: utf8, nullable>`|
|fingerprint|`utf8`|
|id|`int64`|
|kind|`utf8`|
|min_tls_version|`utf8`|
|name|`utf8`|
|profile|`utf8`|
|region|`utf8`|
|self_link (PK)|`utf8`|
|warnings|`json`|