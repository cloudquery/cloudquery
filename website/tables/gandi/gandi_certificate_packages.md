# Table: gandi_certificate_packages

This table shows data for Gandi Certificate Packages.

The primary key for this table is **name**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|sharing_id|`utf8`|
|name (PK)|`utf8`|
|name_label|`utf8`|
|href|`utf8`|
|max_domains|`int64`|
|type|`utf8`|
|type_label|`utf8`|
|wildcard|`bool`|