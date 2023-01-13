# Table: oracle_database_external_container_databases

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
|lifecycle_state|String|
|time_created|Timestamp|
|freeform_tags|JSON|
|defined_tags|JSON|
|lifecycle_details|String|
|db_unique_name|String|
|db_id|String|
|database_version|String|
|database_edition|String|
|time_zone|String|
|character_set|String|
|ncharacter_set|String|
|db_packs|String|
|database_configuration|String|
|database_management_config|JSON|
|stack_monitoring_config|JSON|