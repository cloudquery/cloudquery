# Table: oracle_identity_domains

This table shows data for Oracle Identity Domains.

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
|display_name|`utf8`|
|description|`utf8`|
|url|`utf8`|
|home_region_url|`utf8`|
|home_region|`utf8`|
|replica_regions|`json`|
|type|`utf8`|
|license_type|`utf8`|
|is_hidden_on_login|`bool`|
|time_created|`timestamp[us, tz=UTC]`|
|lifecycle_state|`utf8`|
|lifecycle_details|`utf8`|
|freeform_tags|`json`|
|defined_tags|`json`|