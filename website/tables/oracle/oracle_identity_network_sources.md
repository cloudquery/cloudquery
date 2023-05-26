# Table: oracle_identity_network_sources

This table shows data for Oracle Identity Network Sources.

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|region (PK)|`utf8`|
|compartment_id (PK)|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|description|`utf8`|
|public_source_list|`list<item: utf8, nullable>`|
|virtual_source_list|`json`|
|services|`list<item: utf8, nullable>`|
|time_created|`timestamp[us, tz=UTC]`|
|freeform_tags|`json`|
|defined_tags|`json`|