# Table: oracle_virtualnetwork_capture_filters

This table shows data for Oracle Virtual Network Capture Filters.

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
|lifecycle_state|`utf8`|
|defined_tags|`json`|
|display_name|`utf8`|
|freeform_tags|`json`|
|filter_type|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|vtap_capture_filter_rules|`json`|