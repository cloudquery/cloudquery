# Table: oracle_identity_domains

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|region (PK)|String|
|compartment_id (PK)|String|
|id (PK)|String|
|display_name|String|
|description|String|
|url|String|
|home_region_url|String|
|home_region|String|
|replica_regions|JSON|
|type|String|
|license_type|String|
|is_hidden_on_login|Bool|
|time_created|Timestamp|
|lifecycle_state|String|
|lifecycle_details|String|
|freeform_tags|JSON|
|defined_tags|JSON|